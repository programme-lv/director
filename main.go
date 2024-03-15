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
	rmqConnStr string
}

func (s *server) EvaluateSubmission(req *pb.EvaluationRequest, stream pb.Director_EvaluateSubmissionServer) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	body = snappy.Encode(nil, body)

	var rmq *amqp.Connection
	rmq, err = amqp.Dial(s.rmqConnStr)
	if err != nil {
		return err
	}

	ch, err := rmq.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	evalQ, err := ch.QueueDeclare(
		"eval_q",     // name
		true,         // durable (messages will survive broker restarts)
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
			ContentType: "application/protobuf",
			Body:        body,
			ReplyTo:     respQ.Name,
			// Expiration:  "10000", // 10 seconds
			Expiration: fmt.Sprintf("%d", 1000*60*30), // 30 minutes
		})

	if err != nil {
		return err
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
	grpcServer := grpc.NewServer()
	server := &server{}
	server.rmqConnStr = os.Getenv("RMQ_CONN_STR")
	log.Println("RMQ_CONN_STR: ", server.rmqConnStr)
	pb.RegisterDirectorServer(grpcServer, server)
	log.Printf("server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
