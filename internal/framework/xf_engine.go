package framework

import "github.com/gin-gonic/gin"

type GinEngine struct {
	gin.Engine
	container Container // container
}

func (engine *GinEngine) New() *GinEngine {
	return &GinEngine{
		container: NewXfContainer(),
		Engine:    *gin.New(),
	}
}
func (engine *GinEngine) Defalut() *GinEngine {
	return &GinEngine{
		container: NewXfContainer(),
		Engine:    *gin.Default(),
	}
}

// SetContainer for Engine set container
func (engine *GinEngine) SetContainer(container Container) {
	engine.container = container
}

// GetContainer get contaienr from Engine
func (engine *GinEngine) GetContaier() Container {
	return engine.container
}

// Implement container binding and encapsulation useing engine
func (engine *GinEngine) Bind(provider ServiceProvider) error {
	return engine.container.Bind(provider)
}

// IsBind Has the keyword certificate been bound for the service provider
func (engine *GinEngine) IsBind(key string) bool {
	return engine.container.IsBind(key)
}
