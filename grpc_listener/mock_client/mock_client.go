package main

import (
	"context"
	"flag"
	"io"
	"log"
	"time"

	ecpb "unified-observability/grpc_listener/proto"

	"google.golang.org/grpc"
)

var addr = flag.String("addr", "localhost:50001", "the address to connect to")

func StreamRequest(client ecpb.ALSServiceClient, query string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := client.StreamMetric(ctx, &ecpb.MetricRequest{Query: query})
	if err != nil {
		log.Fatalf("client.StreamRequest(_) = _, %v: ", err)
	}

	for {
		resp, err := res.Recv()

		if err == io.EOF {
			return
		}
		if err != nil {
			log.Fatalln("Receiving", err)
			return
		}
		log.Fatalf("client.StreamRequest(_) = _, %v: ", resp)

	}

}

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	rgc := ecpb.NewALSServiceClient(conn)

	StreamRequest(rgc, "Send Stream Metric")
}
