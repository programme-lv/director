/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a server for Greeter service.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/joho/godotenv"
	pb "github.com/programme-lv/director/msg"
	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedDirectorServer
	rmqConnStr string
}

const ResponseQueueName = "res_q"

func (s *server) EvaluateSubmission(req *pb.EvaluationRequest, stream pb.Director_EvaluateSubmissionServer) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var rmq *amqp.Connection
	rmq, err := amqp.Dial(s.rmqConnStr)
	if err != nil {
		return err
	}

	ch, err := rmq.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"eval_q",     // name
		true,         // durable
		false,        // autoDelete
		false,        // exclusive
		false,        // no-wait
		amqp.Table{}, // arguments
	)
	if err != nil {
		return err
	}

	err = ch.PublishWithContext(ctx, "", q.Name, false, false, amqp.Publishing{
		ContentType:   "application/json",
		Body:          bodyJson,
		ReplyTo:       ResponseQueueName,
		CorrelationId: string(correlationJson),
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
	pb.RegisterDirectorServer(grpcServer, server)
	log.Printf("server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
