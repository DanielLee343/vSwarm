module github.com/vhive-serverless/vSwarm/tools/invoker

go 1.18

replace (
	github.com/vhive-serverless/vSwarm/tools/benchmarking_eventing => ../benchmarking_eventing
	github.com/vhive-serverless/vSwarm/tools/endpoint => ../endpoint
	github.com/vhive-serverless/vSwarm/tools/invoker/proto => ./proto
	github.com/vhive-serverless/vSwarm/utils/tracing/go => ../../utils/tracing/go
)

require (
	github.com/containerd/containerd v1.6.12
	github.com/golang/protobuf v1.5.2
	github.com/google/uuid v1.3.0
	github.com/sirupsen/logrus v1.9.0
	github.com/vhive-serverless/vSwarm/tools/benchmarking_eventing v0.0.0-00010101000000-000000000000
	github.com/vhive-serverless/vSwarm/tools/endpoint v0.0.0-00010101000000-000000000000
	github.com/vhive-serverless/vSwarm/tools/invoker/proto v0.0.0-00010101000000-000000000000
	github.com/vhive-serverless/vSwarm/utils/tracing/go v0.0.0-00010101000000-000000000000
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.36.1
	google.golang.org/grpc v1.49.0
)

require (
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/openzipkin/zipkin-go v0.4.0 // indirect
	go.opentelemetry.io/otel v1.11.1 // indirect
	go.opentelemetry.io/otel/exporters/zipkin v1.8.0 // indirect
	go.opentelemetry.io/otel/sdk v1.8.0 // indirect
	go.opentelemetry.io/otel/trace v1.11.1 // indirect
	golang.org/x/net v0.0.0-20220722155237-a158d28d115b // indirect
	golang.org/x/sys v0.0.0-20220722155257-8c9f86f7a55f // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20220722212130-b98a9ff5e252 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
)
