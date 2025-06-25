package app

import (
	"github.com/xiaofan193/xifancloud193/internal/framework"
	"github.com/xiaofan193/xifancloud193/internal/framework/contract"
)

type XfAppProvider struct {
	BaseFolder string
}

// Register XfApp method
func (x *XfAppProvider) Register(container framework.Container) framework.NewInstance {
	return NewXfApp
}

// calling boot
func (x *XfAppProvider) Boot(container framework.Container) error {
	return nil
}

// IsDefer if defer initialization
func (x XfAppProvider) IsDefer() bool {
	return false
}

// Param 获取初始化参数
func (x *XfAppProvider) Params(container framework.Container) []interface{} {
	return []interface{}{container, x.BaseFolder}
}

// Name get str credentials
func (x *XfAppProvider) Name() string {
	return contract.AppKey
}
