package main

import (
	"github.com/xiaofan193/xifancloud193/internal/framework"
	"github.com/xiaofan193/xifancloud193/internal/framework/provider/app"
)

func main() {
	// initiallization server container
	container := framework.NewXfContainer()
	//bind App service provider
	_ = container.Bind(&app.XfAppProvider{})
	//  The servie providers that need to be bound for subsequent initialization

}
