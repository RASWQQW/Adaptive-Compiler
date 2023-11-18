package LevelFuncs

import (
	"fmt"
	"reflect"
)

func main() {

	//There happens all top env useable funcs imp
}

func ToString(val interface{}) string {
	var kinds map[string]string = map[string]string{reflect.Int.String(): "%s",
		reflect.String.String(): "%s", reflect.Float64.String(): "%f", reflect.Bool.String(): "%"}
	return fmt.Sprintf(kinds[reflect.ValueOf(val).Kind().String()], val)
}
func rept(vals ...any) {
	var valss []interface{} = []interface{}{}
	for v, _ := range vals {
		valss = append(valss, ToString(vals[v]))
	}
	fmt.Println(valss...)
}
