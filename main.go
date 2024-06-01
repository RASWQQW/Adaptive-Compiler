package main

import (
	// comping "ep/comps/jv/comping"
	// connecting "ep/base"

	"encoding/json"
	inputing "ep/Execs"
	inputed "ep/Execs/CppPr"
	"ep/Execs/obj"
	"strconv"

	"ep/web/Mainer"

	//"math"
	chtest "ep/Execs/BotCompiler"

	lvl "ep/LevelFuncs"
	"net/http"
	"strings"

	// Mainer "ep/web/Mainer"

	CodeGenerator "ep/Execs/GENERATE_TASKS"
	baseCalls "ep/base/abst"
	bs "ep/base/abst"

	// obj "ep/inputing/obj"

	// dds "ep/web/gorTest"
	// execs "ep/Execs/BotCompiler"
	// rn "ep/comps/cpp"
	"fmt"
	util "io/ioutil"
	"os"
	"os/exec"
)

func BaseCheck() {
	//baseCalls.BaseConnection
	isid := baseCalls.GetBase()
	fmt.Println("id ", isid.GetTaskById(2))
}

func main() {
	// base.Connecting()
	// var args map[interface{}]string = ReadJson()
	// RunCode(string(args["code"]), string(args["task_name_id"]))
	// var bf = `  int    AddTwoBB(int a,   int s)  {
	// 	int apple = 12  ;
	// 	int dos = 13  ;}  `
	// one, two := obj.Aligner(bf)
	// fmt.Println(
	// 	"Before: ", bf,
	// 	"After: ", one, two)
	// TestingSlice()

	// Runner1()

	// var pointerCheck []string = []string{"Apple"}
	// Turner(pointerCheck)
	// fmt.Println(pointerCheck[0])

	//Checking()
	//web.Runnersss()

	// comp.FileCodeGenerator(1212)
	// dd.Runnersss()

	// restulter := make(chan []string)

	// getName := "cpp-19781285399816371991TP"
	// var pr *execs.Profile = &execs.Profile{Name: getName, UserCName: getName + "_User.cpp", ProperCName: getName + "_Proper.cpp", Lang: "cpp"}
	// rn.Runner("", restulter, pr, pr.UserCName)

	//fmt.Println(math.Abs(float64(1 - 1)))

	// var res = make(map[interface{}]interface{}, 2)
	// var byterd bytes.Buffer

	//io.ReadAll(&byterd, )
	//byterd.ReadString()
	// inputing2.Requester2()
	// json.Unmarshal([]byte("{\"checker\": \"check\"}"), &res)
	// fmt.Println(res["checker"])

	// var ast interface{} = 12
	// fmt.Println(runCode.ToString(ast), runCode.ToString(false))
	// web.Runnersss()
	//var Checker = mme.StepGiving(-1, []string{"a", "b"}, []string{"int", "int"})
	//fmt.Println(Checker)

	// fmt.Println(os.Args[0])
	//gorTest.Mainerd()

	// for {
	// 	var userCode = "double AddTwoNumbers(int a, int b){return a + b;}"
	// 	RunCode("cpp", userCode, "AddTwoNumbers")
	// }
	// WebRunner()

	//BotCompiler.CompilerRequester()
	// BotCompiler.TimeLimitCompiler(1.1)

	// var userCode = "double AddTwoNumbers(int a, int b){return a + b;}"
	// RunCode("cpp", userCode, "AddTwoNumbers")
	// fmt.Println(mt.FindMatrix("appler[12][4545]"))
	// fmt.Println(mt.GetLen("appler[12][12]"))
	// WebRunner()

	// fmt.Println("Write a code: ")
	// reader := bufio.NewReader(os.Stdin)
	// line, _ := reader.ReadString('\n')
	// var sum int = 0
	// str := make(chan int, 3)
	// go func() {
	// 	for v := 0; v < 3; v = v + 1 {
	// 		fmt.Println("cc::")
	// 		fmt.Println(len(str))
	// 		sum = sum + len(str)
	// 		fmt.Println(<-str)
	// 	}
	// }()
	// str <- 1
	// time.Sleep(time.Second)
	// str <- 2
	// time.Sleep(time.Second * 2)
	// str <- 3
	// time.Sleep(time.Second * 1)
	// fmt.Println(sum)

	// fmt.Println("Given code", line)
	// WebSocketRunner()
	// Checkerss()
	// fmt.Println(unicode.IsDigit(rune('1')))
	// var regComp = regexp.MustCompile("(\\[\\d*\\])+")
	// fmt.Println(regComp.FindAllString("[43532][324]", -1))
	//fmt.Println(LevelFuncs.ToString(mt.TypeGuesser(1, string("int"))))
	// fmt.Println(chtest.GetSessionId("cpp"))

	// WebRunner()
	// var userCode = "double AddTwoNumbers(int a, int b){return a + b;}"
	// RunCode("cpp", userCode, "AddTwoNumbers")

	// fmt.Println(
	// 	strconv.Quote(`int main(){std::cout << "CODEER"

	// 	}`), 'b'>>1, byte('c'))
	// var gb baseCalls.BaseConnection = baseCalls.GetBase()
	// var collects *obj.Career = &obj.Career{INOUTS: map[string]any{"task_name_id": "AddTwoNumbers", "lang": "cpp"}, OUTS: map[string]any{}}
	// //collector = append(collector, taskIdCar)
	// gb.GetTaskByName(collects)
	// gb.GetFuncParams(collects)
	// fmt.Println("Vals ", collects.OUTS["args"])
	//fmt.Println("Checker: ")
	//{"string", `'5436546436trdhgfhgfhd'`}, {"string", `'checker'`}, {"int", "12"}
	//fmt.Println("Checker2: ", methods.ExecTimeComp([][]string{{"string[3][11]", `{{"2","2", "2", "2", "2", "2","2","2","2","1", "2", "2", "2", "2",}}`}}, []string{"string"}))
	// var CheckerText string = "%s%s"
	// var CheckerTwo string = fmt.Sprintf("Apple%s", CheckerText, "Txt2", "Txt3")
	// fmt.Println(CheckerTwo)
	// chtest.GetSessionId("cpp")
	WebRunner()
	// CodeRunner()
	// BaseCheck()
	// var bsva bs.BaseConnection = bs.GetBase()
	// // fmt.Println(bsva.CreateFuncDec("CHECKERNAME", "int"))
	// bsva.CreateFuncArgs(int64(42), [][]string{{"int", "checker"}})
}

func OnlineRunner() {

	var cc int = 5 // amount of socket outputs varies
	var outv_ = make(chan string, cc)
	var inv_ = make(chan string, 1)

	var cdedd string = "using namespace std;\n #include <iostream> \n#include <string> \n#include <array>\n\n\n\ndouble AddTwoNumbers(int a, int b){\n\treturn a + b;\n\t}\n\n\ntemplate<typename T>\n\t\tvoid Printer(T PrintType){\n\t\t\tcout << std::to_string(PrintType);\n\t\t}\n\nint main(){\n\tint a = 0;\nint b = 0;\n \n\nPrinter<double>(AddTwoNumbers(a=a, b=b));\n\n\n}"
	gotChan, err := strconv.Unquote(fmt.Sprintf(`"%s"`, cdedd))

	fmt.Println(err)
	fmt.Println(gotChan)
	inv_ <- cdedd
	chtest.InitWebsocketClient(cc /*lang need to change */, "cpp", outv_, inv_, false, false)
}

func CodeRunner() {
	CodeGenerator.GeneratorOfPrompt("", "cpp", false)
}

func CppCodeCreator() {

	var gb bs.BaseConnection = bs.GetBase()
	var checker *lvl.BatchGathererList = new(lvl.BatchGathererList)

	var collects *obj.Career = &obj.Career{INOUTS: map[string]any{"task_name_id": "AddTwoNumbers", "lang": "cpp"}, OUTS: map[string]any{}}
	//collector = append(collector, taskIdCar)
	gb.GetTaskByName(collects)
	gb.GetFuncParams(collects)

	checker.CllRepresentString = inputed.Cpp_control("", 1)
	checker.CllTypePassingString = ""
	checker.FuncReturnType = "int"
	checker.FuncParamNamesAndTypes = obj.Converter[[][]string](collects.ValFinder("args", "out", -1))

	var prod *lvl.Profile = &lvl.Profile{"CHECKERNAME", "CHECKERNAMEUSER", "CHECKERNAMEPROPER", "cpp"}
	getter := make(chan string)

	var BatchGathererIns *lvl.BatchGatherer = &lvl.BatchGatherer{
		ParamBuncher:        [][3]string{{"a", "int", "12"}, {"b", "int", "13"}},
		ParamTimeLimitToRun: 15.1,
		CllCodeParams:       "",
		CllUserCode:         "double AddTwoNumbers(int a, int b){return a + b;}",
		CllProperCode:       "double AddTwoNumbers(int a, int b){return a + b;}",
		CllProfile:          prod,
		CllReturns:          getter}

	checker.Collection = append(checker.Collection, BatchGathererIns)
	chtest.CodeBatcher(checker, "using namespace std;\n#include <iostream>\n\n\n")

}

func WebSocketRunners() {
	vad := chtest.WebSocketRunner //("", "")
	fmt.Println(vad)
}
func WebRunner() {
	Mainer.Runnersss()
}

func RequestTesting() {
	var sendTmp string = `
					curl -H "Content-Type: application/json" 
					-X POST -d '{{JSON}}'
					{{URL}}`

	//EXAMPLE
	//curl.exe -X POST https://api.jdoodle.com/v1/auth-token/ -H 'Content-Type: application/json' -H 'User-Agent: Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:47.0) Gecko/20100101 Firefox/47.0' --data-binary '{"clientId": "b86eafbab040cfd06f729b4f7d233f2d", "clientSecret": "85abad0125e7efdfeb9db2deb1b4155b919c65963deb43a0391ee6209cc71ba9"}'

	var regOut strings.Builder
	var regLink string = "https://api.jdoodle.com/v1/auth-token"
	var clId string = "b86eafbab040cfd06f729b4f7d233f2d"
	var cl_sec string = "35d1146666e6034587c38e86a81aa7f614696a7d172c5ec588a4aaa6788b6405s"

	regJs := map[string]string{"clientId": clId, "clientSecret": cl_sec}

	regJsString, _ := json.Marshal(regJs)

	reg := exec.Command("cmd.exe", "-c", strings.ReplaceAll(strings.ReplaceAll(sendTmp, "{{JSON}}", string(regJsString)), "{{URL}}", regLink))
	reg.Stdout = &regOut
	exec.Command("cmd.exe", "-c")

	//WITH QUERY LIBRARY
	query, _ := http.Post(regLink, "application-json", strings.NewReader(string(regJsString)))
	defer query.Body.Close()
	var res map[any]any
	json.NewDecoder(query.Body).Decode(&res)

}

func RunCode(lang string, code string, topic string) {
	fmt.Println(inputing.Main(lang, code, topic))
}

func ReadJson() map[string]string {
	var givenPath string = PathGiving()
	var jsonSource, err = os.Open(givenPath)
	defer jsonSource.Close()

	if err != nil {
		panic(err)
	}

	var ReadValues map[string]string
	var read, _ = util.ReadAll(jsonSource)
	json.Unmarshal(read, &ReadValues)
	fmt.Printf("Is topic(task_name_id):\n%s\nIs code: \n%s", ReadValues["task_name_id"], ReadValues["code"])

	return ReadValues
	//string{"d1": ReadValues["task_name_id"], "c":  ReadValues["code"]}
}

func PathGiving() string {
	// IT GIVES EXACT PATH OF USER CODE WHICH PREVENTS SYNC MISTAKING OF FILES
	// AND IT HAVE TO BE HAPPEN BY CONSOLE IN API SIDE BUT BETTER DO IT IN REQUEST BY CREATING
	// DIRECT ENDPOINT
	var path string
	fmt.Scan(&path)
	fmt.Println("Given path", path)
	return path

}

func Testing() {
	fmt.Println("Hello, Bro How are You? Are you ok.")
	// var istring string = "Current Command"
	// Checker()
	// cpp.Runner()
	// py.Runner(istring)
	// pointerTesting.Testing()
	// comping.PathCheck()
	// var connectValue = abstImport.GetBase()
	// var result []string = comping.Runner(istring)
	// fmt.Println(result, " ")
	// fmt.Println("Task result as template:", result[0])
	// fmt.Println("Inner Runtime error:", result[1])
}
