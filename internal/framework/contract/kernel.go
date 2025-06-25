package contract

import (
	"google.golang.org/grpc"
	"net/http"
)

const KernelKey = "xf:kernel"

// Kernel interface provider the core of framework
type Kernel interface {
	HttpEngine() http.Handler
	GrpcEngine() *grpc.Server
}
