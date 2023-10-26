package obj

import (
	"fmt"
	"reflect"
)

type Compillers struct {
	Code  string `"Apple: "`
	Topic string
}

type Container struct {
	Compiler Compillers
}

var container *Container = nil

// Its main fucntion because it gets FIRST struct to use further
func SetVals(code string, topic string) *Container {
	container = &Container{Compiler: Compillers{code, topic}}
	return container
}

func (ct Container) GetVal(valName string) string {
	var obj = reflect.ValueOf(ct.Compiler)
	fmt.Println("New name: ", obj, "obj Number: ", obj.NumField())
	for v := 0; v < obj.NumField(); v++ {
		fmt.Println(obj.Type().Field(v).Name, "Current Val")
		if obj.Type().Field(v).Name == valName {
			return obj.Field(v).String()
		}
	}

	return "There is no ready object or value"
}

func (ct *Container) SetVal(key string, value string) {
	fmt.Println(ct.Compiler)
	var isval = reflect.ValueOf(&ct.Compiler)

	elems := isval.Elem()
	fmt.Println("elem Kind: ", elems.Kind())
	fmt.Println("Current key value: ", key, elems.FieldByName(key).String())
	if isval.Kind() == reflect.Ptr {
		name := elems.FieldByName(key)
		fmt.Println(name.CanSet(), name.CanAddr(), name.Addr())
		// name.Addr().SetString(value)
		if name.CanSet() {
			name.Set(reflect.ValueOf(value))
		}
	}
}
