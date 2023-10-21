package obj

import "reflect"

type Compillers struct {
	code  string
	topic string
}

type Container struct {
	compiler Compillers
}

var container *Container = nil

// Its main fucntion because it gets FIRST struct to use further
func SetVals(code string, topic string) *Container {
	container = &Container{compiler: Compillers{code, topic}}
	return container
}

func (ct *Container) GetVal(valName string) string {
	if ct != nil {
		if valName == "code" {
			return ct.compiler.code
		} else if valName == "topic" {
			return ct.compiler.topic
		}
	}
	return "There is no ready object or value"
}

func (ct *Container) SetVal(key string, value string) {
	var isval = reflect.ValueOf(ct)
	// var values = make([]interface{}, isval.NumField())
	for i := 0; i < isval.NumField(); i++ {
		var cfield = isval.Field(i)
		if cfield.Elem().Type().Name() == key {
			isval.Field(i).SetString(value)
		}
	}
}
