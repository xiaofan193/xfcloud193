package framework

// NewInstance define how to build  a new instance,create service for all service containers
type NewInstance func(...interface{}) (interface{}, error)

// ServiceProvider define the interfaces that a service provider needs to implement
type ServiceProvider interface {
	// A method for instaniating a service has been registed in the service container.
	// Wether to instantiate this service during registration needs refer to the IsDefer interface.
	Register(Container) NewInstance
	// Boot when calling the instantiated service,some preparation work such as basic configuration and initialzation
	// parameter operations can be included here
	// if Boot return error,all the service instantiated failed,return err
	Boot(Container) error
	//IsDefer decides whether to instantiate the service during registration,if it is not instantiated during registration,
	// it will be instantialed during the first make operation
	IsDefer() bool
	//  define the parameters passed NewInstance，which can be customized multiple times,it is recommemed to
	// use container as the firt parameter
	Params(Container) []interface{}
	// Name represents the credentials of this service provider
	Name() string
}
