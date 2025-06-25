package main

import (
	"github.com/xiaofan193/xifancloud193/internal/framework"
	"github.com/xiaofan193/xifancloud193/internal/framework/kernel"
	"github.com/xiaofan193/xifancloud193/internal/framework/provider/app"
	"github.com/xiaofan193/xifancloud193/pkg/app/grpc"
	"github.com/xiaofan193/xifancloud193/pkg/app/http"
)

func main() {
	// initiallization server container
	container := framework.NewXfContainer()
	//bind App service provider
	_ = container.Bind(&app.XfAppProvider{})
	//  The servie providers that need to be bound for subsequent initialization ...

	// Initialize HTTP and GRPC engines and bind the as service provider to the service container
	kernelProvider := &kernel.XfKernelProvider{}

	if engine, err := http.NewHttpEngine(container); err != nil {
		kernelProvider.HttpEngine = engine
	}
	if engine, err := grpc.NewGrpcEngine(container); err != nil {
		kernelProvider.GrpcEngine = engine
	}

	// todo some method mission
	//_ = container.Bind(kernelProvider)

	// todo run root

	//_ = console.Runconmand(container)
}
