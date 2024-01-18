package LevelFuncs

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {

	//There happens all top env useable funcs imp
}

func ToString(val interface{}) string {
	var kinds map[string]string = map[string]string{reflect.Int64.String(): "%d", reflect.Int.String(): "%d",
		reflect.String.String(): "%s", reflect.Float64.String(): "%f", reflect.Bool.String(): "%t", reflect.Interface.String(): "%v"}
	// fmt.Println(kinds[reflect.ValueOf(val).Kind().String()], reflect.ValueOf(val).Kind().String())
	fmt.Println(reflect.ValueOf(val).Kind().String())
	for key, vals := range kinds {
		if strings.Contains(reflect.ValueOf(val).Kind().String(), key) {
			return fmt.Sprintf(vals, val)
		}
	}
	return fmt.Sprintf("%v", val)
}

func R(vals ...any) {
	var valss []interface{} = []interface{}{}
	for v, _ := range vals {
		valss = append(valss, ToString(vals[v]))
	}
	fmt.Println(valss...)
}

type DD struct {
	apple  string
	apples string `stringer`
}

// CREATED TO DEFAULT VALUES OF STRUCT OR SORT OF OBJECTS
func Defaulder(objectName reflect.Type) {
	nVal := reflect.New(reflect.TypeOf(objectName)).Elem()
	fmt.Println(nVal) // Checking its presence by printing it
}

func maidn() {
	Defaulder(reflect.TypeOf(DD{}))
}
