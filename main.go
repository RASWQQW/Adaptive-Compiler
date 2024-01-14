package main

import (
	// comping "ep/comps/jv/comping"
	// connecting "ep/base"

	"encoding/json"
	inputing "ep/Execs"

	//"math"
	"net/http"
	"strings"

	_ "ep/LevelFuncs"
	baseCalls "ep/base/abst"

	// obj "ep/inputing/obj"

	_ "ep/web/gorTest"

	"fmt"
	util "io/ioutil"
	"os"
	"os/exec"
)

func BaseCheck() {
	//baseCalls.BaseConnection
	isid := baseCalls.GetBase()
	fmt.Println(isid.GetTaskById(2))
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

	// SaveFileTest()
	// fmt.Println(cpp.Runner(""))

	// vdss := make(chan string, 1)
	// vdss <- "CCD"
	// fmt.Println("Bottom val")
	// var dds chan string = vdss
	// var dds2 chan string = vdss

	// kk := <-dds
	// fmt.Println(dds, dds2, kk)

	// // fmt.Println(vdss, <-vdss)
	// var dd *string = &kk
	// fmt.Println(dd, *dd)
	// *dd = "Oppa"
	// fmt.Println(dd, *dd, kk)
	// fmt.Println(<-dds2)

	// Runner1()

	// var pointerCheck []string = []string{"Apple"}
	// Turner(pointerCheck)
	// fmt.Println(pointerCheck[0])

	//Checking()
	//web.Runnersss()

	var userCode = "double AddTwoNumbers(int a, int b){return a + b;}"
	RunCode("cpp", userCode, "AddTwoNumbers")

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
	// wb.Runnersss()
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
