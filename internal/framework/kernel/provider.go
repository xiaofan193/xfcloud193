package kernel

import (
	"github.com/xiaofan193/xifancloud193/internal/framework"
	"github.com/xiaofan193/xifancloud193/internal/framework/contract"
	"google.golang.org/grpc"
)

// provider web engine
type XfKernelProvider struct {
	HttpEngine *framework.GinEngine
	GrpcEngine *grpc.Server
}

// todo
func (provider *XfKernelProvider) Register(c framework.Container) framework.Container {
	//	return NewXfKernelService
	return nil
}

func (provider *XfKernelProvider) Boot(c framework.Container) error {
	ginr := &framework.GinEngine{}
	if provider.HttpEngine == nil {
		provider.HttpEngine = ginr.Defalut()

	}
	if provider.GrpcEngine == nil {
		provider.GrpcEngine = grpc.NewServer()
	}
	// todo
	//provider.HttpEngine.SetContainer(c)
	return nil
}

func (provider *XfKernelProvider) IsDefer() bool { return false }

func (provider *XfKernelProvider) Params(c framework.Container) []interface{} {
	return []interface{}{
		provider.HttpEngine, provider.GrpcEngine,
	}
}

func (provider *XfKernelProvider) Name() string {
	return contract.KernelKey
}
