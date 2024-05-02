package methods

import (
	"ep/Execs/obj"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"
)

var envMainPath, _ = os.Getwd()

func TypeComplexityCalc(Params []string, is_right_numb bool) float32 {
	var regComp = regexp.MustCompile("(\\[\\d*\\])+")
	var rgs *regexp.Regexp = regexp.MustCompile("\\d+")
	var typeComplixity float32 = 0
	for _, paramType := range Params {
		var counts []string = regComp.FindAllString(paramType, 1)

		// For counting amount thickness of type if it's array
		var matchs = rgs.FindAllString(paramType, -1)
		var SizeThick float32 = 1 // it at least returns 1 value cause it have its type even if it not defiend
		if len(matchs) > 0 {
			gotterd, _ := strconv.Atoi(matchs[0])
			SizeThick = float32(gotterd)
			for tm := 1; tm < len(matchs); tm++ {
				conv, _ := strconv.Atoi(matchs[tm])
				if is_right_numb {
					SizeThick = SizeThick * float32(float32(conv/100)*25)
				} else {
					SizeThick = SizeThick * float32(float32(conv)*0.0125)
				}
			}
		}

		var lls map[string]float32 = map[string]float32{"string": 1.12, "int": 1.12, "double": 1.25, "float": 1.15, "char": 0.5}
		for llsd, ls := range lls {
			if strings.Contains(paramType, llsd) {
				SizeThick = SizeThick * float32(ls)
			}
		}
		typeComplixity = typeComplixity + float32(len(counts)) + SizeThick // + 1 there is no way of adding just one [number]
	}
	return typeComplixity
}

func StepGiving(task_id int, MainParams []string, MainParamsTypes []string) int {
	// Step giving determines whether func have to checked long or short

	if 1 == 1 {
		return 1
	}

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

	var typeComplixity int = len(MainParams) * int(TypeComplexityCalc(getCodeParamTypes, true)) // it sort of multipling to own count one time to increase type's complixity by one time
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

// KEY WORDS TO REFACTOR: DICT
// THE CALCULER NOT CONSIDERS CALCULATING A DICT VALUES BUT TO TIME ARRAY AND OTHER VALUES
func ExecTimeComp(GenTypes [][3]string, RetType []string) float32 { //[0] param type [1] value
	//return 20 // default value

	if 1 == 1 {
		return 21
	}

	var TypesExt []string = []string{}
	var ValuesExt []string = []string{}
	var ValuesComplexity float32

	for s := 0; s < len(GenTypes); s++ {
		TypesExt = append(TypesExt, GenTypes[s][0])
		ValuesExt = append(ValuesExt, GenTypes[s][1])
	}
	for point, val := range ValuesExt {
		var OverAllCom float32 = 0
		val = strings.Replace(
			strings.Replace(
				strings.Replace(
					strings.Replace(val, " ", "", -1),
					"[", "", -1),
				"]", "", -1),
			"'", "", -1) // for strings
		// removes off blanks list of things: [, ]

		for _, valCut := range strings.Split(val, ",") {
			if len(valCut) > 0 {
				if unicode.IsDigit(rune(valCut[0])) && !strings.Contains(TypesExt[point], "string") { //further have to elobrate when there will be DICT // sort of speaking a char like value
					// the points each tenth of number have 8 mult time affect
					OverAllCom = float32(float32(float32(len(strings.Replace(valCut, ".", "", -1))*8)/1000) * 30) // 8 byte is minimal size of int // to get rid of point values
				} else {
					OverAllCom = float32(float32(float32(len(valCut)*4)/1000) * 40) // 4 byte is minimal size of string
				}
			}
			ValuesComplexity = ValuesComplexity + OverAllCom
		}
	}
	var TotalTime float32 = 0.335 * float32(ValuesComplexity) * float32(TypeComplexityCalc(TypesExt, false)+((TypeComplexityCalc(RetType, false)/100)*33))
	return TotalTime
	// make sure that i could generate enough relevant code
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
