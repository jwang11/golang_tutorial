module server

go 1.16

require (
	google.golang.org/grpc v1.53.0
	hello_grpc/service v0.0.0
)

replace hello_grpc/service => ../hello_grpc/service
