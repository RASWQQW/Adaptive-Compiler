package obj

import (
	"fmt"
	"reflect"
	"strings"
	"time"
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
func SetValsOfContainer(code string, topic string) *Container {
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

// Have to check with struct rather than just list or pointer
// carries values of user new call and gets and sets values as reference type
// TODO WORK WITH **TYPE** VALUES THAT CREATES OBJECTS LIKE TYPE [ANY D] APPLE(D, D)D
//
//	type MapFunc[A any, D any] func(A) D
type Career struct {
	OUTS   map[string]interface{}
	INOUTS map[string]interface{}
}

func (cd *Career) Finder() {
	return
}

func (cd *Career) ValFinder(needVal string, itype string, wTime int) any { // needVal is string to get from map wTime is to wait until given time
	if wTime == -1 {
		wTime = 600
	}
	for wTime != 0 {
		var sVal map[string]interface{}
		if strings.ToLower(itype) == "inout" {
			sVal = cd.INOUTS
		} else if strings.ToLower(itype) == "out" {
			sVal = cd.OUTS
		} else {
			panic("Given string name is not found amid struct values")
		}
		for kk, _ := range sVal {
			if kk == needVal {
				return sVal[kk]
			}
		}
		time.Sleep(time.Second)
		wTime = wTime - 1
	}
	panic("Value Finding time exeeded")
}

func Converter[C any](fromVal any) (ok C) {
	ok, err := fromVal.(C)
	if !err {
		panic("types are not equal")
	}
	fmt.Println(ok, reflect.ValueOf(ok).Kind())
	return
}
