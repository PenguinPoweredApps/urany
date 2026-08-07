package main

import (
  "context"
  "log"
  "time"

  "google.golang.org/grpc"
  "google.golang.org/grpc/credentials/insecure"
  
  // Import the generated pb package
  pb "example.com/helloworld/pb" 
)

func main() {
  // 1. Set up a connection to the server. 
  // We use insecure credentials here since we don't have TLS set up for this example.
  conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
  if err != nil {
    log.Fatalf("did not connect: %v", err)
  }
  defer conn.Close()

  // 2. Create a new client stub
  c := pb.NewGreeterClient(conn)

  // 3. Set a timeout for our request
  ctx, cancel := context.WithTimeout(context.Background(), time.Second)
  defer cancel()
  
  // 4. Call the Unary method
  r, err := c.SayHello(ctx, &pb.HelloRequest{Name: "Gopher"})
  if err != nil {
    log.Fatalf("could not greet: %v", err)
  }
  
  log.Printf("Greeting from server: %s", r.GetMessage())
}