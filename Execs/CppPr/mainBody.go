package inputing

import (
	"ep/Execs/BotCompiler"
	compil "ep/Execs/BotCompiler"
	"ep/Execs/methods"
	"ep/LevelFuncs"
	_ "ep/comps/cpp"

	mt "ep/Execs/methods"
	obj "ep/Execs/obj"
	lv "ep/LevelFuncs"
	bs "ep/base/abst"
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/exp/slices"
)

// for gathering all for delete files
type Gatherer struct {
	Lister []string
}

var gbPath, _ = os.Getwd()

// var getFolders []string = strings.Split(gbPath, "\\")
// var GlobalPath = strings.Join(getFolders[:len(getFolders)-3], "\\")
var path string = gbPath + "\\comps\\cpp" // NEED TO MAKE SURE THERE CODE COMES RIGHT WITH  EXACT PATH

func CodeMatching(LoCode string, CurrentFuncName string, CurrentFuncReturnType string, MainParamPassing [][]string) string {
	// THERE STARTS METHOD TESTING -----------
	// checks name and return type of value
	if strings.Contains(LoCode, strings.ReplaceAll(CurrentFuncName, " ", "")) {
		// var types []string = []string{"intypet", "string", "double"}
		// var cppcusttypes []string = []string{"list", "map"}
		// if !(len(regexp.MustCompile(fmt.Sprintf("(%s)+", strings.Join(types, "|"))).FindString(LoCode[0:(len(max(types)))])) > 0) {
		// }

		// HERE ALSO HAVE TO STAY THE THING THAT KEEPS
		// CHECKING PARAM TYPES OF USER CODE

		// checks access method type and tis existance

		coType := strings.ReplaceAll(CurrentFuncReturnType, " ", "")
		stLen := strings.Index(LoCode, coType)
		if stLen != -1 || stLen <= len("static")+1 {
			if strings.ReplaceAll(LoCode[stLen:strings.Index(LoCode, CurrentFuncName)], " ", "") != coType {
				return "return type doesn't match"
			}
		} else {
			return "return type doesn't match"
		}

		var substr = strings.ReplaceAll(LoCode[0:stLen], " ", "") // strings.Index(LoCode, coType)
		if substr != "static" && substr != "" || slices.Contains([]string{"public", "private"}, substr) {
			return "access type isn't relevant"
		}

		// LoCode = obj.Replacer([]string{"static"}, LoCode)
		var namefullIndex int = stLen + len(coType) + 1 + len(CurrentFuncName)
		fmt.Println(LoCode[stLen+len(coType)+1:namefullIndex], CurrentFuncName)
		if !(LoCode[stLen+len(coType)+1:namefullIndex] == CurrentFuncName) {
			return "Func Name doesn't match"
		} else {
			// ARG NAME AND TYPE CHECKING
			var mainer = slices.Clone(MainParamPassing)
			if strings.Index(LoCode, "(") == namefullIndex {
				var arguments string = LoCode[namefullIndex+1 : strings.Index(LoCode, ")")]
				// var argLists []string = strings.Split(strings.Join(strings.Split(arguments, ", "), ""), ",")
				var markarg = strings.ReplaceAll(strings.ReplaceAll(arguments, ",", "+"), "+ ", "+")
				var argLists []string = strings.Split(markarg, "+")
				for arg := range argLists {
					var sect []string = strings.Split(argLists[arg], " ")

					// if there remained values in argLists
					if len(mainer) == 0 {
						return fmt.Sprintf("%s, value and other next are exessive", sect[1])
					}
					for ind, param := range mainer {
						// if current given values aren't matrix nor list
						if sect[1] == param[1] {
							if sect[0] == param[0] {
								// THERE HAVE TO BE SOLUTION TO ADDING TO THIS
								mainer = append(mainer[0:ind], mainer[ind+1:]...)
								break
							} else {
								// strings.Contains(sect[0], "*")There goes type is not contains that lasted values
								return fmt.Sprintf("The parameter %s had unappropriate type, given %s need to %s", sect[1], sect[0], param[1])
							}
						} else if strings.Contains(param[0], "[") {
							var insq string = param[0][strings.Index(param[0], "["):strings.LastIndex(param[0], "]")]

							// REFACTORING
							// by logic there have to change either userCode type or type from base
							if param[1] == strings.Trim(strings.ReplaceAll(sect[1], insq, ""), " ") {
								if param[0] == sect[0]+insq {
									mainer = append(mainer[0:ind], mainer[ind+1:]...)
									break
								} else {
									return fmt.Sprintf("The current %s value's type(matrix) are not matching, given %s need %s", param[1], sect[0]+insq, param[0])
								}
							}
						} else {
							return "Types of given value aren't matching"
						}
					}
				}
				// if there remained values in mainer
				if len(mainer) > 0 {
					return fmt.Sprintf("Some types are not defined {%s}", mainer)
				}
			} else {
				return "There no matching of func types"
			}
		}
		return "true"
	} else {
		return "Func Name even located inside code"
	}
}

func SaveCode(path string, res string) {
	//ar dd sync.Mutex too long to create by sequence
	//dd.Lock() // all of process stops
	cc, err := os.Create(path)
	defer cc.Close()
	if err != nil {
		//panic(err)
		fmt.Println("Error code shown (not serious): ", err)
	}
	cc.WriteString(string(res))
	//dd.Unlock() // and realeses
}

// GLOBAL CHANNEL VALUE
var LANG_G string = ""
var Ncounter chan int = make(chan int, 1)
var DELETEFOLDERNAMES chan []string = make(chan []string)

// SINGLE INS GLOBAL VALUES
var GRetFunc []string = []string{}

func codeSaving(
	LoCode string,
	ProperBaseCode string,
	staticHeader string,
	callMain string,
	Inputs string,
	chStep int,
	ProfGat *Gatherer,
	CodeVals [][]string) string {
	// here goes saving a two proper and user func code each other
	// its a code from user
	var Profile = compil.SaveProfile(LoCode, ProperBaseCode, LANG_G, Ncounter)
	var GetTimeCalced float64 = mt.ExecTimeComp(CodeVals, GRetType)
	var CommonPath string = path + "\\ParalelVaries"

	// IT MAKES MORE SENSE ADDING BEFORE SO ALL OFF CREATED OR DOES NOT, CAN BE COUNTED
	// adding to the pointer list the value of current Profile
	ProfGat.Lister = append(ProfGat.Lister, Profile.Name)
	fmt.Println("CUR VAL: ", ProfGat)

	//User to Compile
	SaveCode(CommonPath+"\\"+Profile.Name+"\\"+Profile.UserCName, staticHeader+"\n"+LoCode+"\n"+callMain)
	vdd1 := make(chan []string)

	// its a proper one
	SaveCode(CommonPath+"\\"+Profile.Name+"\\"+Profile.ProperCName, staticHeader+"\n"+ProperBaseCode+"\n"+callMain)
	vvd2 := make(chan []string)
	// because its a method  and can gain res via chan
	//go TimeLimitCompiler(GetTimeCalced, vvd2, map[string]any{"CommonPath": CommonPath, "ProfileObj": Profile, "filename": Profile.ProperCName}, cpp.Runner) //Profile

	//RUNNING VALUES INP TIMELIMITED AND API COMPILERS
	for _, vals := range []LevelFuncs.StrListStrChan{{Profile.UserCName, vdd1}, {Profile.ProperCName, vvd2}} {
		go BotCompiler.TimeLimitCompiler(
			GetTimeCalced,
			vals.Val2,
			map[string]any{"CommonPath": CommonPath, "ProfileObj": Profile, "filename": vals.Val2},
			BotCompiler.CompilerRequester) //Profile
	}
	properCompRes := <-vvd2
	userCompRes := <-vdd1

	fmt.Println("Checking Step: ", chStep)
	if len(userCompRes[1]) > 0 || len(properCompRes[1]) > 0 {
		if len(userCompRes) > 2 && userCompRes[2] == "1" {
			return userCompRes[1]
		}
		return "Error: " + userCompRes[1] + properCompRes[1]
	} else {
		if strings.ReplaceAll(properCompRes[0], " ", "") != strings.ReplaceAll(userCompRes[0], " ", "") {
			// WrongCode = true
			return "Input: " + Inputs + "\n Proper Output: " + properCompRes[0] + "\nYour output: " + userCompRes[0]
		} else {
			// Where all of code is proper
			fmt.Printf("User result: %s \nProper Code result: %s\n\n", userCompRes[0], properCompRes[0])
			return "False"
		}
	}
}

// have to break down
func CompilingResult(
	LoCode string,
	ProperBaseCode string,
	CurrentFuncName string,
	StaticStartCode string,
	taskType string,
	MainParamPassing [][]string,
	ParamCheckingTime int) string {

	// there have to check that value is
	// PRINT type or I/O type

	// here users have to choose wether ther will be compilling by method or only with a input
	// but it have to output so concrete
	/// THERE STARTS PRINT TESTING------------
	if taskType == "PT" {
		if strings.Contains(LoCode, "#include iostream") || strings.Contains(LoCode, "using namespace std;") {
			// here goes code only whcih sent with custom and untemplated input and with other values
			// executes with regex
			if strings.Contains(strings.ReplaceAll(LoCode, " ", ""), "cin>>") {
			}
		} else {
			if strings.Contains(LoCode, CurrentFuncName) && LoCode[:5] == "void" {
				var printUserCode = fmt.Sprintf("void %s(){\n%s\n}", CurrentFuncName, LoCode)
				var SaveCode = StaticStartCode + fmt.Sprint("{%s}", printUserCode)

				// This have to changed
				var Inputs string = LoCode[:len("Arg Count")]
				if dd := codeSaving(
					LoCode, ProperBaseCode, SaveCode, Inputs, "", 1, &Gatherer{[]string{}},
					[][]string{} /* optional value need to optemize */); dd == "False" {
					return "Your code is correct"
				} else {
					return dd
				}
			}
		}

	} else if taskType == "RT" {
		//var TypesKeys map[string]interface{} = map[string]interface{}{"dict[int]": map[int]string, "list[float]": []interface{}, "list[int]": []int, "int": int, "double": ""}
		//DEF CONST VALS AND mt TO SAVE

		fmt.Println("Code checking time: ", lv.ToString(ParamCheckingTime))
		var results []chan string = []chan string{}
		var ids []int = []int{}
		var NGatherer *Gatherer = &Gatherer{[]string{}} // gathering all profiles around goroutins

		for t := 0; t < ParamCheckingTime; t++ {
			// Func for run in goroutine and set result into channel
			var TypeGiver = func(t int, giver chan string) {
				// two of them fils ahead all of and gets place when they are ready
				var DataTypes string = ""
				var FuncParamsPass []string = []string{}

				//VALUES FOR GETTING INFOS ABOUT FUNC PARAMS
				var ParamBunch [][]string = [][]string{}
				// var WrongCode bool = false

				for vvs := range MainParamPassing {
					var ParamType string = MainParamPassing[vvs][0]
					var name string = MainParamPassing[vvs][1]
					var RandVal string

					if strings.Contains(ParamType, "[") {
						var ListComped = mt.FindList(ParamType)
						var matrixComped = mt.FindMatrix(ParamType)
						// here starts checking that val is list or matrix

						//THERE MUST DESCRIBED WHAT SHOULD BE DONE IF RETURNS ENUMERABLE VALUE
						if len(ListComped) > 0 {
							// there starts whole process of checking that type and giving value
							var glen string = mt.GetLen(ListComped)[0]
							var gtypename string = mt.GetParamType(ListComped)
							var RndValType = mt.GetFullType(name, gtypename, []string{glen})
							var RndValue = mt.RandValSetting(t, gtypename, glen, "list")
							RandVal = RndValType + RndValue + ";"

							ParamBunch = append(ParamBunch, []string{RndValType, RndValue}) // adding to value collector

						} else if len(matrixComped) > 0 {
							var glen []string = mt.GetLen(matrixComped)
							var gtypename string = mt.GetParamType(matrixComped)
							var FullType, valsStringsformat string
							FullType = mt.GetFullType(name, gtypename, glen)
							for conts := range glen {
								valsStringsformat = valsStringsformat + ", " + mt.RandValSetting(t, gtypename, glen[conts], "matrix")
							}
							//THERE I HAVE TO LOOK AT SOME CODE BY TEST RUNNIN WITH LITTLE DEBUG
							//Real generated values and full type is in randval word
							RandVal = FullType + "{" + valsStringsformat + "};"
							ParamBunch = append(ParamBunch, []string{FullType, "{" + valsStringsformat + "};"})
						}

					} else {
						var GeneratedSingleValue string = LevelFuncs.ToString(mt.TypeGuesser(t, string(ParamType)))
						ParamBunch = append(ParamBunch, []string{ParamType, GeneratedSingleValue}) // adding to value collector
						RandVal = fmt.Sprintf("%s %s = %d;", ParamType, name, GeneratedSingleValue)
					}
					// here goes adding all inner parameters to put in a row line
					DataTypes = DataTypes + RandVal + "\n"
					// it is just FuncName and its parameters adding on sequence
					FuncParamsPass = append(FuncParamsPass, fmt.Sprintf("%s=%s", name, name))
				}

				var outputerFunc string = ""
				if len(GRetFunc) == 2 {
					outputerFunc = GRetFunc[0] + "(%s(%s));" // by print func check
				} else {
					outputerFunc = "cout << %s(%s);" //by mere cout
				}
				var MainFunc string = "\n\nint main(){\n\t" + fmt.Sprintf("%s \n\n"+outputerFunc, DataTypes, CurrentFuncName, strings.Join(FuncParamsPass, ", ")) + "\n}"

				// THEE I GOTTA WRITE LITTLE FILE MANAGEMENT TO WRITE AND GET OUTPUT OF EXACT PROCESS RUNNING
				// and finally goes checking by compiling and waiting it
				var res string = codeSaving(LoCode, ProperBaseCode, StaticStartCode, MainFunc, DataTypes, t, NGatherer, ParamBunch)
				if res != "False" {
					giver <- res // There goes All results as Time limit or Error etc, beside Proper ones
				} else {
					giver <- "Good result:" + res //to get proper result
				}
			}

			var resHandler = make(chan string)
			go TypeGiver(t, resHandler)
			ids = append(ids, t)
			results = append(results, resHandler)
		}

		// var checked []int = []int{}
		for v := 0; v < len(results); v++ {
			//if !slices.Contains(checked, v) {}
			select {
			case copier := <-results[v]:
				if !strings.Contains(copier, "Good") {
					// if one of them already gives error it returns it
					// code must be checked sequantially due to compelexity rate of parameters
					go DeleteAllWhenDone(v+1, results, NGatherer)
					return copier

				}
				continue
			}
		}
		//IF ITS  WITH ONLY HAVING A ALL PROPER OUT
		go DeleteAllWhenDone(-1, results, NGatherer)

		return "Your code is proper as well"
	}
	// and here have to be func creator which creates func code to save in file
	// var creator = func() {

	// }
	//Lately goes other stuff like looping concerning inner types and giving it like random params
	// and relatively save and check its result
	return "Noting were checked"
}

// DELETING ALWAYS ALL VALUES NO MATTER WRONG OR CORRECT
func DeleteAllWhenDone(fromId int, Callers []chan string, Gatherers *Gatherer) {
	if fromId > -1 {
		for fromC := fromId; fromC < len(Callers); fromC++ {
			select {
			case <-Callers[fromC]:
				continue // anyway it continues to go
			}
		}
	}
	fmt.Println("CODE COUNT: ", len(Gatherers.Lister))
	//CLEANING ALL FOLDERS THAT FOR CHECK
	for _, value := range Gatherers.Lister {
		var pathCur string = path + "\\ParalelVaries\\" + value
		err := os.RemoveAll(pathCur)
		if err == nil {
			fmt.Println("Path was gone removed " + pathCur)
			log.Println("Path was gone removed " + pathCur)
		} else {
			log.Println(err)
		}
	}
}

func Compiler(cmp *obj.Container) string {
	// return All(cmp)
	const LANG = "cpp"
	LANG_G = LANG
	Ncounter <- 0
	defer func() { <-Ncounter }()

	// AND THERE GOES LITTLE REPORT
	var rep = func(message string, val interface{}) {
		fmt.Println(message, val)
	}

	var _, LoCode = mt.Aligner(cmp.GetVal("Code"))
	var TopicName = cmp.GetVal("Topic")

	// and passing its values from aboveds
	// this is a func of current topic TOPIC uses
	var gb bs.BaseConnection = bs.GetBase()

	// full collector of objects that constists Career object which contains value of each async running func
	//var collector []*obj.Career = []*obj.Career{}

	// here goes all saved func values that gives access locallys
	var collects *obj.Career = &obj.Career{INOUTS: map[string]any{"task_name_id": TopicName, "lang": LANG}, OUTS: map[string]any{}}
	//collector = append(collector, taskIdCar)

	var funcs = []func(*obj.Career){
		gb.GetTaskByName, gb.GetFuncParams, gb.GetFunction, gb.GetProperCode}

	methods.RoutineRunner(funcs, collects)

	var CurrentFuncName string = lv.ToString(collects.ValFinder("func_name", "out", -1))
	var TaskType string = lv.ToString(collects.ValFinder("task_type", "out", -1))
	var PropCode string = lv.ToString(collects.ValFinder("prop_code", "out", -1))
	var CurrentFuncReturnType string = lv.ToString(collects.ValFinder("return_type", "out", -1))
	var Params = obj.Converter[[][]string](collects.ValFinder("args", "out", -1))
	var ParamsNames []string = obj.Converter[[]string](collects.ValFinder("arg_names", "out", -1))
	var ParamsTypes []string = obj.Converter[[]string](collects.ValFinder("arg_types", "out", -1))

	if res := cpp_control(CurrentFuncReturnType); res[0] != "-1" {
		GRetFunc = res
	}

	var ProperBaseCode string = strings.ReplaceAll(
		strings.ReplaceAll(PropCode, "FUNC", CurrentFuncName),
		"RET_TYPE", CurrentFuncReturnType) //changes for all proper code funcs

	rep("Proper code: ", ProperBaseCode)
	var ParamCheckingTime int = mt.StepGiving(-1, ParamsNames, ParamsTypes)

	// There i have to ready current topic giving vals and current code
	// and it happens once but rndomly saves it many time

	// MainParamPassing = gb.GetFuncParams(GetTaskId)
	var StaticStartCode string = "using namespace std;\n#include <iostream>\n\n\n"
	if d := CodeMatching(LoCode, CurrentFuncName, CurrentFuncReturnType, Params); d != "true" {
		return d
	}
	return CompilingResult(LoCode, ProperBaseCode, CurrentFuncName, StaticStartCode, TaskType, Params, ParamCheckingTime)

	//return "The last val of checking"

	// var GetTaskId = make(chan []string)
	// var taskType = make(chan []string)
	// var getFunc = make(chan []string)
	// var propCode = make(chan []string)
	// var prPassType = make(chan []string)
	// var prPassName = make(chan []string)
	// var MainParamPassingV2 [][]string = [][]string{}

	// go gb.GetTaskByName(map[string]interface{}{"task_name_id": TopicName},
	// 	map[string]chan []string{"taskType": taskType, "GetTaskId": GetTaskId})

	// var TaskId = <-GetTaskId
	// var funcParams []map[string]interface{} = []map[string]interface{}{
	// 	{"TaskId": TaskId},
	// 	{"TaskId": TaskId, "LANG": LANG}}
	// // {"TaskId": TaskId}}

	// var funcChans = []map[string]chan []string{
	// 	{"getFunc": getFunc},
	// 	{"propCode": propCode}}
	// {"prPassName": prPassName, "prPassType": prPassType}}

	// TESTING BY SINGLE RUN METHOD IN MAIN PARAM PASSINGS
	// gb.GetFuncParams(map[string]interface{}{"task_id": TaskId}, map[string]chan []string{}, map[string]interface{}{"Mainpr": MainParamPassingV2})

	// var gottenFunc []string = <-getFunc

	// var CurrentFuncName string = gottenFunc[1]
	// rep("Func name: ", CurrentFuncName)

	// var CurrentFuncReturnType string = gottenFunc[0]
	// rep("Func return: ", CurrentFuncReturnType)

	// var prNames []string = <-prPassName
	// var prTypes []string = <-prPassType
	// var MainParamPassing [][]string = [][]string{}

	// if len(prNames) == len(prTypes) {
	// 	for v := range prNames {
	// 		MainParamPassing = append(MainParamPassing, []string{prNames[v], prTypes[v]})
	// 	}
	// }

	// var propCodeStr []string = <-propCode
	// var taskTypeStr []string = <-taskType

	// GetTaskId, taskType, _ := gb.GetTaskByName(TopicName)
	// rep("Function name: ", TopicName)
	// rep("User Code: ", LoCode)
	// rep("Function id: ", TaskId)
	// rep("Func Type: ", <-taskType)

	// var getFunc = gb.GetFunction(GetTaskId)
	// var propCode = gb.GetProperCode(GetTaskId, LANG)

}
