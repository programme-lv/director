package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	"github.com/golang/snappy"
	"github.com/joho/godotenv"
	pb "github.com/programme-lv/director/msg"

	// amqp "github.com/rabbitmq/amqp091-go"
	amqp "github.com/peake100/rogerRabbit-go/pkg/amqp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// https://github.com/peake100/rogerRabbit-go#readme

var (
	errMissingMetadata = status.Errorf(codes.InvalidArgument, "missing metadata")
	errInvalidToken    = status.Errorf(codes.Unauthenticated, "invalid token")
)
var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedDirectorServer
	conn    *amqp.Connection
	infoLog *log.Logger
}

func (s *server) EvaluateSubmission(req *pb.EvaluationRequest, stream pb.Director_EvaluateSubmissionServer) error {
	body, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	body = snappy.Encode(nil, body)

	ch, err := s.conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	evalQ, err := ch.QueueDeclare(
		"eval_q",     // name
		false,        // durable (messages will survive broker restarts)
		false,        // autoDelete (queue will be deleted when no consumers)
		false,        // exclusive (exclusive use by only this connection)
		false,        // no-wait (don't wait for a confirmation from the server)
		amqp.Table{}, // arguments
	)
	if err != nil {
		return err
	}

	respQ, err := ch.QueueDeclare(
		"",    // name - empty to use server-generated queue name
		false, // durable
		false, // autoDelete
		false, // exclusive
		false, // no-wait
		amqp.Table{
			"x-expires": int32(1000 * 60 * 35), // 35 minutes
		}, // arguments
	)
	if err != nil {
		return err
	}

	err = ch.Publish(
		"",         // exchange
		evalQ.Name, // routing key
		true,       // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "application/octet-stream",
			Body:        body,
			ReplyTo:     respQ.Name,
			Expiration:  fmt.Sprintf("%d", 1000*60*30), // 30 minutes
		},
	)

	if err != nil {
		return err
	}

	// start consuming from the response queue
	msgs, err := ch.Consume(
		respQ.Name, // queue
		"",         // consumer
		false,      // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	if err != nil {
		return err
	}

	for msg := range msgs {
		// unmarshal the message
		decompressed, err := snappy.Decode(nil, msg.Body)
		if err != nil {
			return fmt.Errorf("failed to decompress message with snappy: %v", err)
		}

		feedback := pb.EvaluationFeedback{}
		if err := proto.Unmarshal(decompressed, &feedback); err != nil {
			return err
		}

		err = stream.Send(&feedback)
		if err != nil {
			return err
		}

		finished := false
		switch feedback.FeedbackTypes.(type) {
		case *pb.EvaluationFeedback_StartEvaluation:
			s.infoLog.Printf("StartEvaluation: %+v", feedback.GetStartEvaluation())
		case *pb.EvaluationFeedback_FinishEvaluation:
			s.infoLog.Printf("FinishEvaluation: %+v", feedback.GetFinishEvaluation())
			finished = true
		case *pb.EvaluationFeedback_FinishWithInernalServerError:
			s.infoLog.Printf("FinishWithInernalServerError: %+v", feedback.GetFinishWithInernalServerError())
			finished = true
		case *pb.EvaluationFeedback_StartCompilation:
			s.infoLog.Printf("StartCompilation: %+v", feedback.GetStartCompilation())
		case *pb.EvaluationFeedback_FinishCompilation:
			s.infoLog.Printf("FinishCompilation: %+v", feedback.GetFinishCompilation())
		case *pb.EvaluationFeedback_FinishWithCompilationError:
			s.infoLog.Printf("FinishWithCompilationError: %+v", feedback.GetFinishWithCompilationError())
			finished = true
		case *pb.EvaluationFeedback_StartTesting:
			s.infoLog.Printf("StartTesting: %+v", feedback.GetStartTesting())
		case *pb.EvaluationFeedback_IgnoreTest:
			s.infoLog.Printf("IgnoreTest: %+v", feedback.GetIgnoreTest())
		case *pb.EvaluationFeedback_StartTest:
			s.infoLog.Printf("StartTest: %+v", feedback.GetStartTest())
		case *pb.EvaluationFeedback_ReportTestSubmissionRuntimeData:
			s.infoLog.Printf("ReportTestSubmissionRuntimeData: %+v", feedback.GetReportTestSubmissionRuntimeData())
		case *pb.EvaluationFeedback_FinishTestWithLimitExceeded:
			s.infoLog.Printf("FinishTestWithLimitExceeded: %+v", feedback.GetFinishTestWithLimitExceeded())
		case *pb.EvaluationFeedback_FinishTestWithRuntimeError:
			s.infoLog.Printf("FinishTestWithRuntimeError: %+v", feedback.GetFinishTestWithRuntimeError())
		case *pb.EvaluationFeedback_ReportTestCheckerRuntimeData:
			s.infoLog.Printf("ReportTestCheckerRuntimeData: %+v", feedback.GetReportTestCheckerRuntimeData())
		case *pb.EvaluationFeedback_FinishTestWithVerdictAccepted:
			s.infoLog.Printf("FinishTestWithVerdictAccepted: %+v", feedback.GetFinishTestWithVerdictAccepted())
		case *pb.EvaluationFeedback_FinishTestWithVerdictWrongAnswer:
			s.infoLog.Printf("FinishTestWithVerdictWrongAnswer: %+v", feedback.GetFinishTestWithVerdictWrongAnswer())
		case *pb.EvaluationFeedback_IncrementScore:
			s.infoLog.Printf("IncrementScore: %+v", feedback.GetIncrementScore())
		}
		msg.Ack(false)
		if finished {
			break
		}
	}

	return nil
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	conn, err := amqp.Dial(os.Getenv("RMQ_CONN_STR"))
	if err != nil {
		log.Fatalf("failed to connect to RabbitMQ: %v", err)
	}

	// notify := conn.NotifyClose(make(chan *amqp.Error))
	// go func() {
	// 	for err := range notify {
	// 		log.Printf("RabbitMQ connection closed: %v", err)
	// 	}
	// }()

	apiKey := os.Getenv("GRPC_API_KEY")
	if len(apiKey) < 3 {
		log.Fatalf("GRPC_API_KEY is not set or invalid")
	}

	opts := []grpc.ServerOption{
		grpc.StreamInterceptor(getEnsureValidTokenFunc(apiKey)),
	}
	grpcServer := grpc.NewServer(opts...)
	server := &server{}
	server.conn = conn
	server.infoLog = log.New(os.Stdout, "INFO: ", log.LstdFlags)
	pb.RegisterDirectorServer(grpcServer, server)
	log.Printf("server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// valid validates the authorization.
func valid(authorization []string, apiKey string) bool {
	if len(authorization) < 1 {
		return false
	}
	token := strings.TrimPrefix(authorization[0], "Bearer ")
	// Perform the token validation here. For the sake of this example, the code
	// here forgoes any of the usual OAuth2 token validation and instead checks
	// for a token matching an arbitrary string.
	return token == apiKey
}

func getEnsureValidTokenFunc(apiKey string) func(srv any, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	return func(srv any, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		md, ok := metadata.FromIncomingContext(ss.Context())
		if !ok {
			return errMissingMetadata
		}
		// // The keys within metadata.MD are normalized to lowercase.
		// // See: https://godoc.org/google.golang.org/grpc/metadata#New
		if !valid(md["authorization"], apiKey) {
			return errInvalidToken
		}
		return handler(srv, ss)
	}
}
