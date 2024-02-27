// bot actually sends the question
// in one language then it waits till it translates
// then this code keeps as global file or text
// and other check cases are just copies global code and changes value

// CHANGES WHCIH REQUIRED
// MAKE TIME OF CODE AND MEMBERY USAGE IMORTANT AS IT  WILL  LEVERAGE USER STAT ON EACH RAN CODE/TOTAL SCORE
package BotCompiler

import (
	"encoding/json"
	"ep/LevelFuncs"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"reflect"
	"strings"
	"time"
)

var MAINPATH, _ = os.Getwd()

func Requester() {
	fmt.Println("Start of requesting")
	var bap string = "6059490532:AAHk7sYVhcRhBg7Opfn3ieae6_wo7hfWtNU"
	var link string = "https://api.telegram.org/bot" + bap + "/sendMessage"
	var vals string = "{\"chat_id\": -1001862349713, \"text\": \"MESSAGE\"}"
	query, _ := http.Post(link, "application/json", strings.NewReader(vals))
	defer query.Body.Close()
	var res map[any]any
	json.NewDecoder(query.Body).Decode(&res)
	fmt.Println(query.StatusCode, query.Status, res)

}

type ReadVals struct {
	Output            string `form:"output"`
	StatusCode        int    `form:"statusCode"`
	Memory            string `form:"memory"`
	CpuTime           string `form:"сpuTime"`
	CompilationStatus string `form:"compilationStatus"`
}

// complete this code to run vals
func TimeLimitCompiler(
	sleepTime float64,
	ReturnValue chan []string,
	params map[string]any,
	RunFunc func(map[string]any, chan []string) []string) {
	if RunFunc == nil {
		//It runs func of request for default if there is no given func
		RunFunc = CompilerRequester
	}
	go RunFunc(params, ReturnValue) // Runs in background until code it ends or limit time comes to an end
	time.Sleep(time.Second * time.Duration(sleepTime))
	//if the given time ends fun will set value  to chan by itself
	ReturnValue <- []string{"Time limit exceeded", "TIMELIMIT", "1"}
}

func CompilerRequester(Values map[string]any, RetVals chan []string) []string {

	var regLink string = "https://api.jdoodle.com/v1/execute"
	var clId string = "b86eafbab040cfd06f729b4f7d233f2d"
	var cl_sec string = "85abad0125e7efdfeb9db2deb1b4155b919c65963deb43a0391ee6209cc71ba9"
	getProf := reflect.ValueOf(Values["ProfileObj"])

	if getProf.Kind().String() != "Profile" {
		panic("Given name of object is wrong typed which must be Profile instance")
	}

	read, _ := os.ReadFile(LevelFuncs.ToString(Values["CommonPath"]) + "\\" + getProf.FieldByName("Name").String() + "\\" + LevelFuncs.ToString(Values["filename"])) //MAINPATH + "\\comps\\cpp\\UserCode\\compiler.cpp")
	var ReadAndPassString string = string(read)

	regJs := map[string]string{
		"clientId":     clId,
		"clientSecret": cl_sec,
		"script":       ReadAndPassString, //"using namespace std; \n#include <iostream> \n\nint main() {cout << \"Apple is sweet!\" << endl; return 0; }", //"std::cout<< 'Apple is sweet';",
		"language":     "cpp"}

	regJsString, _ := json.Marshal(regJs)

	//WITH QUERY LIBRARY
	query, _ := http.Post(regLink, "application/json", strings.NewReader(string(regJsString)))
	defer query.Body.Close()

	//res := make(map[any]any)
	rVal := ReadVals{}

	resbod, _ := ioutil.ReadAll(query.Body)
	json.Unmarshal(resbod, &rVal)

	fmt.Println("res before: ", rVal)
	fmt.Println("StatCode: ", query.StatusCode)
	fmt.Println("Body: ", string(resbod))
	//fmt.Println("res after: ", res)
	//fmt.Println("Compiled Code: ", res["output"])
	//fmt.Println("Stat: ", query.Status)
	//fmt.Println("Response:", query.Request.Response)

	if rVal.StatusCode == 200 {
		if rVal.CompilationStatus == "1" {
			RetVals <- []string{rVal.Output, ""}
			return []string{rVal.Output, ""}
		}
	}

	RetVals <- []string{rVal.Output, fmt.Sprintf("Error: %s", rVal.Output)}
	return []string{rVal.Output, fmt.Sprintf("Error: %s", rVal.Output)}
	// panic("Code is not proper")
}

func Junk() {
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
