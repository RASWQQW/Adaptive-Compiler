package main

import (
	// comping "example/comps/jv/comping"
	// connecting "example/base"

	"encoding/json"
	baseCalls "example/base/abst"
	"example/inputing"
	"fmt"
	util "io/ioutil"
	"os"
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
	// var userCode = "int AddTwoNumbers(int a, int b){return a + b;}"
	// RunCode("cpp", userCode, "AddTwoNumbers")

	var compVal = inputing.SetVals("apple", "apple2")
	// fmt.Printf(inputing.container)
	fmt.Printf(compVal.GetVal("code"))
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
