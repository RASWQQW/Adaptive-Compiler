package inputing

import (
	bs "example/base/abst"
	"example/comps/cpp"
	obj "example/inputing/obj"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func Compiler(cmp obj.Container) string {
	return All(cmp)
}

// have to break down
func All(cmp obj.Container) string {
	const LANG = "cpp"

	// AND THERE GOES LITTLE REPORT
	var rep = func(message string, val interface{}) {
		fmt.Println(message, val)
	}

	// and passing its values from aboveds
	// this is a func of current topic TOPIC uses
	var gb bs.BaseConnection = bs.GetBase()
	var loCode = cmp.GetVal("code")
	var TopicName = cmp.GetVal("topic")
	

	GetTaskId, taskType, _ := gb.GetTaskByName(TopicName)
	rep("Function name: ", TopicName)
	rep("User Code: ", loCode)
	rep("Function id: ", GetTaskId)
	rep("Func Type: ", taskType)

	var getFunc = gb.GetFunction(GetTaskId)
	var propCode = gb.GetProperCode(GetTaskId, LANG)

	var CurrentFuncName string = getFunc[1]
	rep("Func name: ", getFunc[0])

	var CurrentFuncReturnType string = getFunc[0]
	rep("Func return: ", getFunc[1])

	var ProperBaseCode string = strings.ReplaceAll(propCode, "FUNC", CurrentFuncName) //changes for all proper code funcs
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

	var StaticStartCode string = "#include <iostream>\nusing namespace std;\nint main()"
	var path string = GlobalPath + "\\cpp\\cppsep"
	var SaveCode = func(path string, res string) {
		cc, err := os.Create(path)
		if err != nil {
			defer cc.Close()
			cc.WriteString(res)
		}
	}

	var codeSaving = func(staticHeader string, Inputs string) string {
		// here goes saving a two proper and user func code each other
		// its a code from user
		var ppcodeUser string = path + "//UserCode//compiler.cpp"
		SaveCode(ppcodeUser, staticHeader+"\n"+loCode)
		var userCompRes []string = cpp.Runner(ppcodeUser)

		// its a proper one
		var ppcodeProper string = path + "//ProperCode//compiler.cpp"
		SaveCode(ppcodeProper, staticHeader+"\n"+ProperBaseCode)
		var properCompRes []string = cpp.Runner(ppcodeProper)

		if len(userCompRes[1]) > 0 || len(properCompRes[1]) > 0 {
			return "Error: " + userCompRes[1] + properCompRes[1]
		} else {
			if strings.ReplaceAll(properCompRes[0], " ", "") != strings.ReplaceAll(userCompRes[0], " ", "") {
				// WrongCode = true
				return "Input: " + Inputs + "\n Proper Output: " + properCompRes[0] + "\nYour output: " + userCompRes[0]
			} else {
				return "False"
			}
		}
	}

	// there have to check that value is
	// PRINT type or I/O type

	// here users have to choose wether ther will be compilling by method or only with a input
	// but it have to output so concrete
	/// THERE STARTS PRINT TESTING------------
	if taskType == "PT" {
		if strings.Contains(loCode, "#include iostream") || strings.Contains(cmp.code, "using namespace std;") {
			// here goes code only whcih sent with custom and untemplated input and with other values
			// executes with regex

			if strings.Contains(strings.ReplaceAll(loCode, " ", ""), "cin>>") {
				// var inputs []string = make([]string, 10)
				// for vals := range inputs {
				// }
				// var staticHeader = ""
			}
		} else {
			if strings.Contains(loCode, CurrentFuncName) && cmp.code[:5] == "void" {
				var printUserCode = fmt.Sprintf("void %s(){\n%s\n}", CurrentFuncName, cmp.code)
				var SaveCode = StaticStartCode + fmt.Sprint("{%s}", printUserCode)

				// This have to changed
				var Inputs string = loCode[:len("Arg Count")]
				if dd := codeSaving(SaveCode, Inputs); dd == "False" {
					return "Your code is correct"
				} else {
					return dd
				}
			}
		}

	} else if taskType == "RT" {
		// THERE STARTS METHOD TESTING -----------
		// checks name and return type of value
		if strings.Contains(loCode, strings.ReplaceAll(CurrentFuncName, " ", "")) {
			// var types []string = []string{"int", "string", "double"}
			// var cppcusttypes []string = []string{"list", "map"}
			// if !(len(regexp.MustCompile(fmt.Sprintf("(%s)+", strings.Join(types, "|"))).FindString(cmp.code[0:(len(max(types)))])) > 0) {
			// }

			// HERE ALSO HAVE TO STAY THE THING THAT KEEPS
			// CHECKING PARAM TYPES OF USER CODE

			// checks access method type and tis existance
			coType := strings.ReplaceAll(CurrentFuncReturnType, " ", "")
			var substr = strings.ReplaceAll(loCode[0:strings.Index(cmp., coType)], " ", "")
			if substr != "static" && substr != "" {
				return "access type isn't relevant"
			}

			cmp.code = Replacer([]string{"static"}, cmp.code)
			if coType == string(cmp.code[0:len(coType)]) {
				var namefullIndex int = len(coType) + 1 + len(CurrentFuncName)
				if !(cmp.code[len(coType)+2:namefullIndex] == CurrentFuncName) {
					return "Func Name doesn't match"
				} else {
					// ARG NAME AND TYPE CHECKING
					var mainer = MainParamPassing
					if strings.Index(cmp.code, "(") == namefullIndex+1 {
						var arguments string = cmp.code[namefullIndex+2 : strings.Index(cmp.code, ")")]
						var argLists []string = strings.Split(strings.Join(strings.Split(arguments, ", "), ""), ",")
						for arg := range argLists {
							var sect []string = strings.Split(argLists[arg], " ")

							// if ther remained values in argLists
							if len(mainer) == 0 {
								return fmt.Sprintf("%s, value and other next are exessive", sect[1])
							}

							for ind, param := range mainer {

								if sect[1] == param[1] {
									if strings.Contains(param[0], "[") {
										var insq string = param[0][strings.Index(param[0], "["):strings.Index(param[0], "]")]
										sect[0] = strings.ReplaceAll(param[0], insq, "") + insq
									}
									if sect[0] == param[0] {
										// THERE HAVE TO BE SOLUTION TO ADDING TO THIS
										mainer = append(mainer[0:ind], mainer[ind+1:]...)
										break
									} else {
										// strings.Contains(sect[0], "*")There goes type is not contains that lasted values
										return fmt.Sprintf("The parameter %s had unappropriate type, given %s need to %s", sect[1], sect[0], param[1])
									}
								}
							}
						}
						// if there remained values in mainer
						if len(mainer) > 0 {
							return fmt.Sprintf("Some types are not defined {%s}", mainer)
						}
					}

				}

			} else {
				return "return type doesn't match"
			}

		} else {
			return "Name or return type doesn't match"
		}

		//var TypesKeys map[string]interface{} = map[string]interface{}{"dict[int]": map[int]string, "list[float]": []interface{}, "list[int]": []int, "int": int, "double": ""}
		//DEF CONST VALS AND METHODS TO SAVE
		var ParamCheckingTime int = rand.Intn(300) + 100

		for t := 0; t < ParamCheckingTime; t++ {
			// two of them fils ahead all of and gets place when they are ready
			var DataTypes, FuncNameCopy string = "", ""
			// var WrongCode bool = false

			for vvs := range MainParamPassing {
				var ParamType string = MainParamPassing[vvs][0]
				var name string = MainParamPassing[vvs][1]

				var RandVal string

				if strings.Contains(ParamType, "[") {
					var ListComped = FindList(ParamType)
					var matrixComped = FindMatrix(ParamType)
					// here starts checking that val is list or matrix
					if len(ListComped) > 0 {
						// there starts whole process of checking that type and giving value
						var glen string = GetLen(ListComped)[0]
						var gtypename string = GetParamType(ListComped)
						RandVal = GetFullType(name, gtypename, []string{glen}) + RandValSetting(gtypename, glen, "list") + ";"

					} else if len(matrixComped) > 0 {
						var glen []string = GetLen(matrixComped)
						var gtypename string = GetParamType(matrixComped)
						var FullType, valsStringsformat string
						FullType = GetFullType(name, gtypename, glen)
						for conts := range glen {
							valsStringsformat = valsStringsformat + ", " + RandValSetting(gtypename, glen[conts], "matrix")
						}

						RandVal = FullType + "{" + valsStringsformat + "};"
					}

				} else {
					RandVal = fmt.Sprintf("%s %s %s", ParamType, name, TypeGuesser(string(ParamType)))
				}
				// here goes adding all inner parameters to put in a row line
				DataTypes = DataTypes + RandVal + "\n"
				// it is just FuncName and its parameters adding on sequence
				FuncNameCopy = FuncNameCopy + fmt.Sprintf("%s=%s,", name, name)
			}

			var staticHeader string = StaticStartCode +
				"{\n\t" + fmt.Sprintf("%s \n\n cout << %s(%s);", DataTypes, CurrentFuncName, FuncNameCopy) + "\n}"

			// and finally goes checking by compiling and waiting it
			var res string = codeSaving(staticHeader, DataTypes)
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
