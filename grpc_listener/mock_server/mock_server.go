package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"time"
	pb "unified-observability/grpc_listener/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var port = flag.Int("port", 50001, "the port to serve on")

type alsServer struct {
	pb.UnimplementedALSServiceServer
}

func main() {
	flag.Parse()
	fmt.Printf("server starting on port %d...\n", *port)
	// ctx := context.Background()

	s := grpc.NewServer()
	pb.RegisterALSServiceServer(s, &alsServer{})
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *alsServer) StreamMetric(req *pb.MetricRequest, stream pb.ALSService_StreamMetricServer) error {
	for {
		select {
		case <-stream.Context().Done():
			return status.Error(codes.Canceled, "Stream has ended")
		default:
			time.Sleep(1 * time.Second)

			if err := stream.SendMsg(&pb.Entry{
				AriaProvider:               "AriaProvider",
				AriaService:                "AriaService",
				CloudAccountID:             "CloudAccountID",
				KubernetesClusterName:      "KubernetesClusterName",
				KubernetesWorkloadName:     "KubernetesWorkloadName",
				KubernetesNamespaceName:    "KubernetesNamespaceName",
				KubernetesPodFullname:      "KubernetesPodFullname",
				KubernetesPodName:          "KubernetesPodName",
				KubernetesServiceName:      "KubernetesServiceName",
				Location:                   "Location",
				PCCloudAddress:             "PCCloudAddress",
				Region:                     "Region",
				KubernetesServiceRequestor: "KubernetesServiceRequestor",
				KubernetesServiceResponder: "KubernetesServiceResponder",
				KubernetesSrcClusterName:   "KubernetesSrcClusterName",
				KubernetesSrcNamespace:     "KubernetesSrcNamespace",
				EntityUID:                  "EntityUID",
				DestSpaceName:              "DestSpaceName",
				SourceSpaceName:            "SourceSpaceName",
				ResponseCode:               "ResponseCode",
			}); err != nil {
				return status.Error(codes.Canceled, "Stream has cancelled")
			}
		}

	}
}
