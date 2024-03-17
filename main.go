package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/golang/snappy"
	"github.com/joho/godotenv"
	pb "github.com/programme-lv/director/msg"
	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedDirectorServer
	conn *amqp.Connection
}

func (s *server) EvaluateSubmission(req *pb.EvaluationRequest, stream pb.Director_EvaluateSubmissionServer) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	body = snappy.Encode(nil, body)

	ch, err := s.conn.Channel()
	if err != nil {
		return err
	}

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

	err = ch.PublishWithContext(ctx,
		"",         // exchange
		evalQ.Name, // routing key
		true,       // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "application/octet-stream",
			Body:        body,
			ReplyTo:     respQ.Name,
			// Expiration:  "10000", // 10 seconds
			Expiration: fmt.Sprintf("%d", 1000*60*30), // 30 minutes
		})

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
		log.Printf("Received a message: %+v", msg.Body)
		decompressed, err := snappy.Decode(nil, msg.Body)
		if err != nil {
			return fmt.Errorf("failed to decompress message with snappy: %v", err)
		}

		feedback := pb.EvaluationFeedback{}
		if err := proto.Unmarshal(decompressed, &feedback); err != nil {
			return err
		}

		switch feedback.FeedbackTypes.(type) {
		case *pb.EvaluationFeedback_Start:
			log.Printf("Received start message")
		}

		err = stream.Send(&feedback)
		if err != nil {
			return err
		}

		msg.Ack(false)
	}

	return nil
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
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

	grpcServer := grpc.NewServer()
	server := &server{}
	server.conn = conn
	pb.RegisterDirectorServer(grpcServer, server)
	log.Printf("server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
