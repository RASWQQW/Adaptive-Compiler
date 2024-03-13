package methods

import (
	"ep/Execs/obj"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"time"
)

var envMainPath, _ = os.Getwd()

func StepGiving(task_id int, MainParams []string, MainParamsTypes []string) int {
	// Step giving determines whether func have to checked long or short

	var getCodeParamTypes []string
	if task_id >= 0 {
		// var mainFunc abst.BaseConnection = abst.GetBase()
		// cnh := make(chan []string)
		// mainFunc.GetFuncParams(task_id, cnh)
		// getCodeParamTypes = <-cnh
	} else {
		//There have to be func that extracts types from [][]object
		getCodeParamTypes = MainParamsTypes
	}

	var regComp = regexp.MustCompile("(\\[\\d*\\])+")

	var typeComplixity int = 0
	for _, paramType := range getCodeParamTypes {
		var counts []string = regComp.FindAllString(paramType, -1)
		typeComplixity = typeComplixity + (len(counts)) // + 1 there is no way of adding just one [number]
	}

	typeComplixity = len(MainParams) * typeComplixity
	var from int = 25
	var to int = 95 // in reality it goes to "to" + "from" `400`

	var colls [][][]int = [][][]int{
		{{2}, {15, 10}},
		{{6}, {30, 30}},
		{{12}, {75, 35}},
		{{}, {100, 30}}}

	for anchor := range colls {
		var thre []int = colls[anchor][0]
		if len(thre) < 1 {
			thre[0] = typeComplixity + 1 // to make next value been chosen so the anchor's 0 val is big at typecomp
		}
		// it checks by ascending of anc values so there is no need of more than
		if typeComplixity < thre[0] {
			from = colls[anchor][1][0]
			to = colls[anchor][1][1]
			break
		}
	}

	var time int = rand.Intn(to) + from
	return time
}

func ValIncr(cat string, val int) int {
	// there have to by some logic
	if cat == "default" {

	} else if cat == "matrix" {
		return 8
	} else if cat == "list" {
		return 15
	}
	return 12
}

func ExecTimeComp(GenTypes [][]string, RetType [][]string) float64 { //[0] param type [1] value
	return 1
}

func ExecTimeLimiter(path string, execer func(string, chan []string) []string) []string {
	var getReturns = make(chan []string, 1)
	getReturns <- []string{}
	go execer(path, getReturns)

	for v := 0; v < /*ExecTimeComp()*/ 12; v++ {
		time.Sleep(time.Millisecond)
	}
	if len(<-getReturns) < 1 {
		panic("Time limit exeeded")
	}
	return <-getReturns
}

func FileCreateFunc(code string, CodeStatus string, lang string) string {

	var LangPath map[string]string = map[string]string{"java": "comps/jv", "cpp": "comps/cpp", "py": "comps/py", "csharp": "comps/chsarp"}
	var path string = LangPath[lang] + "\\" + CodeStatus

	op, _ := os.Open(envMainPath + path)
	defer op.Close()
	if st, _ := op.ReadDir(0); len(st) > 201 {
		dd, _ := os.Executable()
		panic(fmt.Sprintf("The file save limit exeeded, loc: %s", filepath.Dir(dd)))
	}

	// while creating goroutine in parallel there must be saver that makes parralel saves
	// in files

	// when just calling it generates name and creates it
	// parrralel version of codeSaving

	var genName = fmt.Sprintf("%sccd%f", CodeStatus, rand.Float64())
	if stat, err := os.Stat(genName); os.IsNotExist(err) {
	} else {
		fmt.Println(stat.IsDir())
		panic(fmt.Sprintf("%q given path already exists", genName))
	}

	var savePath string = envMainPath + path + genName
	saving, _ := os.Create(savePath)
	defer saving.Close()
	saving.WriteString(code)
	return savePath
}

func FileDeleteFunc(delFilePath string) {
	os.Remove(envMainPath + delFilePath)
}

func AnySetter(ids []*interface{}, anys []interface{}) {
	if len(ids) == len(anys) {
		var valsd []interface{} = append([]interface{}{}, anys...)

		for dd, _ := range valsd {
			typeN := reflect.ValueOf(valsd[dd])
			if typeN.Kind() == reflect.String {
				*ids[dd] = typeN.String()
			} else if typeN.Kind() == reflect.Chan {
				if reflect.ValueOf(*ids[dd]).Kind() == reflect.Int {
					*ids[dd] = typeN.Int()
				}
			}
		}
	}
}

func RoutineRunner(funcs []func(*obj.Career), params *obj.Career) {
	//if len(funcs) == len(funcs) len(funcs) == len(params) {
	for funcdd, _ := range funcs {
		go funcs[funcdd](params)
	}
	//}

	//There would be incr level
	// to increase level of variable complixity on checking when passing
}
