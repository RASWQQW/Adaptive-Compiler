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

type StrListStrChan struct {
	Val1 string
	Val2 chan []string
	Code string
}

type ValType struct {
	matrix bool "false" // any type with two dimensinal cover
	array  bool "false" // any type with one array struct
	simple bool "false" // any simple
}

type Profile struct {
	Name        string
	UserCName   string
	ProperCName string
	Lang        string
}

type BatchGatherer struct {
	ParamTimeLimitToRun float32
	ParamBuncher        [][3]string
	CllCodeParams       string
	CllUserCode         string
	CllProperCode       string
	CllReturns          chan string
	CllProfile          *Profile
	CllParamType        int // type of i/o based type 1 -> simple, 2 -> array,  3 ->  matrix
}

type BatchGathererList struct {
	Collection             []*BatchGatherer
	CllRepresentString     []string // rep string is given and i can make some funcs to be gotten in order
	CllTypePassingString   string
	FuncParamNamesAndTypes [][]string
	FuncReturnType         string
}
