package kafka

import (
	"github.com/eolinker/eosc"
	"github.com/eolinker/eosc/utils/schema"
	"reflect"
)

const name = "kafka_output"

//Register 注册kafka_output驱动工厂
func Register(register eosc.IExtenderDriverRegister) {
	register.RegisterExtenderDriver(name, NewFactory())
}

type Factory struct {
}

func (f *Factory) Render() *schema.Schema {
	render, err := schema.Generate(reflect.TypeOf((*Config)(nil)), nil)
	if err != nil {
		return nil
	}
	return render
}
func NewFactory() *Factory {
	return &Factory{}
}

func (f *Factory) Create(profession string, name string, label string, desc string, params map[string]interface{}) (eosc.IExtenderDriver, error) {
	return &Driver{
		configType: reflect.TypeOf((*Config)(nil)),
	}, nil
}