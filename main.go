package main

import (
	// comping "ep/comps/jv/comping"
	// connecting "ep/base"

	"encoding/json"
	inputing "ep/Execs"
	baseCalls "ep/base/abst"

	// obj "ep/inputing/obj"
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
	// var userCode = "double AddTwoNumbers(int a, int b){return a + b;}"
	// RunCode("cpp", userCode, "AddTwoNumbers")

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

	// Runner()

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

	// web.Runner()
	Checking()
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
