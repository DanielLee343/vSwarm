package main

import (
	"context"
	"sync"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	log "github.com/sirupsen/logrus"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/vhive-serverless/vSwarm/tools/benchmarking_eventing/proto"

	"github.com/vhive-serverless/vSwarm/tools/endpoint"
)

var (
	tsdbConn   *grpc.ClientConn
	tsdbClient proto.TimeseriesClient
	lock       sync.Mutex
)

func Start(tdbAddr string, endpoints []*endpoint.Endpoint, workflowIDs map[*endpoint.Endpoint]string) {
	lock.Lock()
	defer lock.Unlock()

	// Start the TimeseriesDB only if there exist at least one endpoint
	// that uses eventing
	enable := false
	for _, endpoint := range endpoints {
		if endpoint.Eventing {
			enable = true
			break
		}
	}
	if !enable {
		return
	}

	workflowDefinitions := make(map[string]*proto.WorkflowDefinition)

	for _, ep := range endpoints {
		workflowID := workflowIDs[ep]
		workflowDefinitions[workflowID] = &proto.WorkflowDefinition{
			Id: workflowID,
			CompletionEventDescriptors: []*proto.CompletionEventDescriptor{
				{
					AttrMatchers: ep.Matchers,
				},
			},
		}
	}

	dialOptions := make([]grpc.DialOption, 0)
	dialOptions = append(dialOptions, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if *withTracing {
		dialOptions = append(dialOptions, grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()))
	}
	var err error
	tsdbConn, err = grpc.Dial(tdbAddr, dialOptions...)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	tsdbClient = proto.NewTimeseriesClient(tsdbConn)
	ctx, cancel := context.WithTimeout(context.Background(), grpcTimeout)
	defer cancel()

	if _, err := tsdbClient.StartExperiment(ctx, &proto.ExperimentDefinition{WorkflowDefinitions: workflowDefinitions}); err != nil {
		log.Fatalln("failed to start experiment", err)
	}
}

func End() (durations []time.Duration) {
	lock.Lock()
	defer lock.Unlock()

	// TimeseriesDB is started only if there existed at least one endpoint
	// that used eventing; tsdbConn is nil if not started.
	if tsdbConn == nil {
		return
	}

	defer tsdbConn.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	res, err := tsdbClient.EndExperiment(ctx, &empty.Empty{})
	if err != nil {
		log.Fatalln("failed to end experiment", err)
	}

	for _, wrk := range res.WorkflowResults {
		for _, inv := range wrk.Invocations {
			// Skip incomplete invocations
			if inv.Status != proto.InvocationStatus_COMPLETED {
				continue
			}
			durations = append(durations, inv.Duration.AsDuration())
		}
	}
	return
}
