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

// Package main implements a client for Greeter service.
package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "client-server/protoBuf/logMessage"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", ":50051", "the address to connect to")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewLoggerClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	r, err := c.SendMessage(ctx, &pb.LogRequest{
		Timestamp: int64(123454555),
		LogLine:   "HOLA COMO ESTAS ....",
		Tags: []*pb.LogRequest_Tags{
			{Value: "la app", Type: pb.LogRequest_APPLICATION},
			{Value: "el scope", Type: pb.LogRequest_SCOPE},
			{Value: "1.0.0", Type: pb.LogRequest_VERSION},
			{Value: "10.0.20.12", Type: pb.LogRequest_HOST},
		},
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetResponse())
}
