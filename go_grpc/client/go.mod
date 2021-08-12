module client

go 1.16

replace hello_grpc/service => ../hello_grpc/service

require (
	google.golang.org/grpc v1.40.0
	hello_grpc/service v0.0.0-00010101000000-000000000000
)
