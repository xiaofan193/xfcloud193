package apiserver

import cliflag "github.com/xiaofan193/xifancloud193/internal/pkg/flag"

type CliOptions interface {
	Flags() (fss cliflag.NamedFlagSets)
	Validate() []error
}

type ConfigurableOption interface {
	ApplyFlags() []error
}

type PrintableOption interface {
	String() string
}
