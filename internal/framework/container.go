package framework

import (
	"fmt"
	"sync"
)

// Container is service container,provide bind and get function of services
type Container interface {
	// Bind  a service provider,if the credential of keyword  exist,it will  replace operation,return err
	Bind(provider ServiceProvider) error
	// IsBind wether the credetial of keyword is bind the service provider
	IsBind(key string) bool
	// Make Obtain  a service based on keyword credentials
	Make(key string) (interface{}, error)
	// obtain a service based on keyword credentials ,if this keyword credential is not
	//bound to a service provider,it will panic
	MustMake(key string) interface{}
	// botain a service based on keyword credentials,but the service is not single instance mode
	// this function is very useful when that for different parameters boostrap different instance
	MakeNew(key string, params []interface{}) (interface{}, error)
}

type XfContainer struct {
	Container
	// store register service provider，key as a credential
	Providers map[string]ServiceProvider
	// store specific instances
	Instances map[string]interface{}
	// lock
	Lock sync.RWMutex
}

// create a service container
func NewXfContainer() *XfContainer {
	return &XfContainer{
		Providers: map[string]ServiceProvider{},
		Instances: map[string]interface{}{},
		Lock:      sync.RWMutex{},
	}
}

// output keywords registered in the service container
func (x *XfContainer) PrintProviders() []string {
	ret := []string{}
	for _, provider := range x.Providers {
		name := provider.Name()
		line := fmt.Sprint(name)
		ret = append(ret, line)
	}
	return ret
}

// Bind service container with keyword
func (x *XfContainer) Bind(provider ServiceProvider) error {
	x.Lock.Lock()
	key := provider.Name()
	x.Providers[key] = provider
	x.Lock.Unlock()
	// if provider is not defer
	if provider.IsDefer() == false {
		if err := provider.Boot(x); err != nil {
			return err
		}
		// instatiation method
		params := provider.Params(x)
		method := provider.Register(x)
		instance, err := method(params...)
		if err != nil {
			return err
		}
		x.Instances[key] = instance
	}
	return nil
}

func (x *XfContainer) IsBind(key string) bool {
	return true
}

func (x *XfContainer) Make(key string) (interface{}, error) {
	return nil, nil
}
