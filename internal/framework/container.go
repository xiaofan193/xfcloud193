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
	return x.findServiceProvider(key) != nil
}
func (x *XfContainer) findServiceProvider(key string) ServiceProvider {

	x.Lock.RLock()
	defer x.Lock.Unlock()
	if sp, ok := x.Providers[key]; ok {
		return sp
	}
	return nil
}

func (x *XfContainer) Make(key string) (interface{}, error) {
	return x.make(key, nil, false)
}

func (x *XfContainer) MustMake(key string) interface{} {
	serv, err := x.make(key, nil, false)
	if err != nil {
		panic("container not contain key" + key)
	}
	return serv
}
func (x *XfContainer) MakeNew(key string, params []interface{}) (interface{}, error) {
	return x.make(key, params, true)
}

func (x *XfContainer) newInstance(sp ServiceProvider, params []interface{}) (interface{}, error) {
	if err := sp.Boot(x); err != nil {
		return nil, err
	}
	if params == nil {
		params = sp.Params(x)
	}
	method := sp.Register(x)
	ins, err := method(params...)
	if err != nil {
		return nil, err
	}
	return ins, err
}

// the real instantiate a service
func (x *XfContainer) make(key string, params []interface{}, forceNew bool) (interface{}, error) {
	x.Lock.RLock()
	defer x.Lock.RUnlock()
	sp := x.findServiceProvider(key)
	if sp == nil {
		return nil, fmt.Errorf("container %s have not register", key)
	}

	if forceNew {
		return x.newInstance(sp, params)
	}
	// use the instance in container
	if ins, ok := x.Instances[key]; ok {
		return ins, nil
	}
	// if container is not instantiate, instantiate onece again
	inst, err := x.newInstance(sp, nil)
	if err != nil {
		return nil, err
	}

	x.Instances[key] = inst
	return inst, nil
}

// List the string credentials of all service providers int the container
func (x *XfContainer) NameList() []string {
	ret := []string{}
	for _, provider := range x.Providers {
		name := provider.Name()
		ret = append(ret, name)
	}
	return ret
}
