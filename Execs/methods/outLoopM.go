package methods

import (
	"example/base/abst"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"regexp"
	"time"
)

var envMainPath, _ = os.Getwd()

func StepGiving(task_id int) int {
	// Step giving detemines whether func have to checked long or short

	cnh := make(chan []string)
	var mainFunc abst.BaseConnection = abst.GetBase()
	mainFunc.GetFuncParams(task_id, cnh)

	var getCodeParamTypes []string = <-cnh
	var getCodeParamsAmount []string = []string{}
	var regComp = regexp.MustCompile("(\\[\\d*\\])+")

	var typeComplixity int = 0
	for _, paramType := range getCodeParamTypes {
		var counts []string = regComp.FindAllString(paramType, -1)
		typeComplixity = typeComplixity + (len(counts) + 1)
	}

	typeComplixity = len(getCodeParamsAmount) * typeComplixity
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
			thre[0] = typeComplixity + 1
		}
		if typeComplixity < thre[0] {
			from = colls[anchor][1][0]
			to = colls[anchor][1][1]
		}
	}

	var time int = rand.Intn(to) + from
	return time
}

func RoutineRunner(funcs ...func(vals ...any)) {
	//There would be incr level
	// to increase level of variable complixity on checking when passing

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

func ExecTimeComp() int {
	return 1
}

func ExecTimeLimier(path string, execer func(string, chan []string) []string) []string {
	var getReturns = make(chan []string, 1)
	getReturns <- []string{}
	go execer(path, getReturns)

	for v := 0; v < ExecTimeComp(); v++ {
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
