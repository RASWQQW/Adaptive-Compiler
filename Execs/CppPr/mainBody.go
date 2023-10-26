package inputing

import (
	mt "example/Execs/methods"
	obj "example/Execs/obj"
	bs "example/base/abst"
	"example/comps/cpp"
	"fmt"
	"os"
	"strings"

	"golang.org/x/exp/slices"
)

var gbPath, _ = os.Getwd()

// var getFolders []string = strings.Split(gbPath, "\\")
// var GlobalPath = strings.Join(getFolders[:len(getFolders)-3], "\\")
var path string = gbPath + "\\comps\\cpp"

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
									return fmt.Sprintf("The current %s values type(matrix) are not matching, given %s need %s", param[1], sect[0]+insq, param[0])
								}
							}
						} else {
							return "Types of given value aren't mathcing"
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
	cc, err := os.Create(path)
	defer cc.Close()
	if err != nil {
		panic(err)
	}
	cc.WriteString(string(res))

}

func codeSaving(LoCode string, ProperBaseCode string, staticHeader string, callMain string, Inputs string, chStep int) string {
	// here goes saving a two proper and user func code each other
	// its a code from user
	var ppcodeUser string = path + "\\UserCode"
	SaveCode(ppcodeUser+"\\compiler.cpp", staticHeader+"\n"+LoCode+"\n"+callMain)
	var userCompRes []string = cpp.Runner(ppcodeUser)

	// its a proper one
	var ppcodeProper string = path + "\\ProperCode"
	SaveCode(ppcodeProper+"\\compiler.cpp", staticHeader+"\n"+ProperBaseCode+"\n"+callMain)
	var properCompRes []string = cpp.Runner(ppcodeProper)

	fmt.Println("Checking Step: ", chStep)
	if len(userCompRes[1]) > 0 || len(properCompRes[1]) > 0 {
		return "Error: " + userCompRes[1] + properCompRes[1]
	} else {
		if strings.ReplaceAll(properCompRes[0], " ", "") != strings.ReplaceAll(userCompRes[0], " ", "") {
			// WrongCode = true
			return "Input: " + Inputs + "\n Proper Output: " + properCompRes[0] + "\nYour output: " + userCompRes[0]
		} else {
			// Where are all of code is proper
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
	MainParamPassing [][]string) string {

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
				if dd := codeSaving(LoCode, ProperBaseCode, SaveCode, Inputs, "", 1); dd == "False" {
					return "Your code is correct"
				} else {
					return dd
				}
			}
		}

	} else if taskType == "RT" {
		//var TypesKeys map[string]interface{} = map[string]interface{}{"dict[int]": map[int]string, "list[float]": []interface{}, "list[int]": []int, "int": int, "double": ""}
		//DEF CONST VALS AND mt TO SAVE

		var ParamCheckingTime int = mt.StepGiving(LoCode)
		fmt.Println("Code checking time: ", string(ParamCheckingTime))

		for t := 0; t < ParamCheckingTime; t++ {
			// two of them fils ahead all of and gets place when they are ready
			var DataTypes string = ""
			var FuncNameCopy []string = []string{}
			// var WrongCode bool = false

			for vvs := range MainParamPassing {
				var ParamType string = MainParamPassing[vvs][0]
				var name string = MainParamPassing[vvs][1]
				var RandVal string

				if strings.Contains(ParamType, "[") {
					var ListComped = mt.FindList(ParamType)
					var matrixComped = mt.FindMatrix(ParamType)
					// here starts checking that val is list or matrix
					if len(ListComped) > 0 {
						// there starts whole process of checking that type and giving value
						var glen string = mt.GetLen(ListComped)[0]
						var gtypename string = mt.GetParamType(ListComped)
						RandVal = mt.GetFullType(name, gtypename, []string{glen}) + mt.RandValSetting(gtypename, glen, "list") + ";"

					} else if len(matrixComped) > 0 {
						var glen []string = mt.GetLen(matrixComped)
						var gtypename string = mt.GetParamType(matrixComped)
						var FullType, valsStringsformat string
						FullType = mt.GetFullType(name, gtypename, glen)
						for conts := range glen {
							valsStringsformat = valsStringsformat + ", " + mt.RandValSetting(gtypename, glen[conts], "matrix")
						}
						RandVal = FullType + "{" + valsStringsformat + "};"
					}

				} else {
					RandVal = fmt.Sprintf("%s %s = %d;", ParamType, name, mt.TypeGuesser(string(ParamType)))
				}
				// here goes adding all inner parameters to put in a row line
				DataTypes = DataTypes + RandVal + "\n"
				// it is just FuncName and its parameters adding on sequence
				FuncNameCopy = append(FuncNameCopy, fmt.Sprintf("%s=%s", name, name))
			}

			var MainFunc string = "\n\nint main(){\n\t" + fmt.Sprintf("%s \n\n cout << %s(%s);", DataTypes, CurrentFuncName, strings.Join(FuncNameCopy, ", ")) + "\n}"

			// and finally goes checking by compiling and waiting it
			var res string = codeSaving(LoCode, ProperBaseCode, StaticStartCode, MainFunc, DataTypes, t)
			if res != "False" {
				return res
			}
		}
		return "Your code is proper as well"
	}
	// and here have to be func creator which creates func code to save in file
	// var creator = func() {

	// }
	//Lately goes other stuff like looping concerning inner types and giving it like random params
	// and relatively save and check its result
	return "Noting were checked"
}

func Compiler(cmp *obj.Container) string {
	// return All(cmp)
	const LANG = "cpp"

	// AND THERE GOES LITTLE REPORT
	var rep = func(message string, val interface{}) {
		fmt.Println(message, val)
	}

	// and passing its values from aboveds
	// this is a func of current topic TOPIC uses
	var gb bs.BaseConnection = bs.GetBase()
	var LoCode, _ = mt.Aligner(cmp.GetVal("Code"))
	var TopicName = cmp.GetVal("Topic")

	GetTaskId, taskType, _ := gb.GetTaskByName(TopicName)
	rep("Function name: ", TopicName)
	rep("User Code: ", LoCode)
	rep("Function id: ", GetTaskId)
	rep("Func Type: ", taskType)

	var getFunc = gb.GetFunction(GetTaskId)
	var propCode = gb.GetProperCode(GetTaskId, LANG)

	var CurrentFuncName string = getFunc[1]
	rep("Func name: ", getFunc[1])

	var CurrentFuncReturnType string = getFunc[0]
	rep("Func return: ", getFunc[0])

	var ProperBaseCode string = strings.ReplaceAll(
		strings.ReplaceAll(propCode, "FUNC", CurrentFuncName),
		"RET_TYPE", CurrentFuncReturnType) //changes for all proper code funcs

	rep("Proper code: ", ProperBaseCode)

	// There i have to ready current topic giving vals and current code
	// and it happens once but rndomly saves it many time

	// mainPassing params gets usally as int Apple2 and int Apple2
	// And There could be matrix, mere integer or double or random list
	var MainParamPassing [][]string = [][]string{{"int", "Setting"},
		{"double", "Setting2"},
		{"list", "JustMember"},
		{"matrix", "JustMember2"},
		{"dict", "JustMember3"}}

	MainParamPassing = gb.GetFuncParams(GetTaskId)
	var StaticStartCode string = "using namespace std;\n#include <iostream>\n\n\n"
	if d := CodeMatching(LoCode, CurrentFuncName, CurrentFuncReturnType, MainParamPassing); d != "true" {
		return d
	}
	return CompilingResult(LoCode, ProperBaseCode, CurrentFuncName, StaticStartCode, taskType, MainParamPassing)
	// return "The last val of checking"
}
