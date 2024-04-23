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
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"reflect"
	"strconv"
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
	sleepTime float32,
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

// The idea is making possible of running custom codes with their id of file
// whether is user or proper code on one compile request
// and getting all results by specification
// func CodeBatcher(gatherer *LevelFuncs.BatchGathererList, stat string) { //(addCode string, headers string) {
// 	var CompileCode string = stat + ` `
// 	var FuncRunnerInj string = `` //  I have sitten code of thread running of given funcs

// 	//There Have to bes passed list of

// 	var returnType string = ``
// 	var passingParamsTypes string = ``
// 	var FuncNamesLists string = `list<` + returnType + `(*)(` + passingParamsTypes + `)> RUNNFUNCS {`

// 	var FuncsDeclaring string = ``
// 	var FuncsRunList []string = []string{} // there i have to show funcs list syntax
// 	var funcCreate = func(name string, funcCode string) {
// 		var FuncsDeclaringFunc = func() {}
// 		FuncsDeclaring = FuncsDeclaring + "\n\n" + funcCode[:strings.Index(funcCode, " ")-1] + name + funcCode[strings.Index(funcCode, "("):]
// 		FuncsRunList = append(FuncsRunList, fmt.Sprintf(gatherer.CllRepresentString, name, gatherer.CllTypePassingString))
// 	}

// 	FuncNamesLists + strings.Join(FuncsRunList, ", ") + "};"
// 	for vv := 0; vv < len(gatherer.Collection); vv = vv + 1 {
// 		var nname string = gatherer.Collection[vv].CllProfile.Name
// 		funcCreate(nname+"ProperBased", gatherer.Collection[vv].CllProperCode)
// 		funcCreate(nname+"UserBased", gatherer.Collection[vv].CllUserCode)
// 	}

// }

func GetBoardData(username string) int {
	var regStr string = "https://api.twitch.tv/helix/users?login=" + username
	var clId string = "b86eafbab040cfd06f729b4f7d233f2d"
	var Result string = ""
	req, _ := http.NewRequest("POST", regStr, nil)
	req.Header.Set("Client-ID", clId)
	req.Header.Set("Content-Type", "application/json")

	var cl *http.Client = &http.Client{}
	res, _ := cl.Do(req)
	red, _ := io.ReadAll(res.Body)

	//getting headers
	for key, val := range res.Header {
		fmt.Println("Key " + string(key) + "Val: " + string(val[0]))
		if key == "X-RateLimit-Remaining" {
			Result = val[0]
		}
	}
	// printing body
	fmt.Println(string(red))
	resd, _ := strconv.Atoi(Result)
	return resd
}

func CompilerRequester(Values map[string]any, RetVals chan []string) []string {
	var regLink string = "https://api.jdoodle.com/v1/execute"
	var clId string = ""   //"b86eafbab040cfd06f729b4f7d233f2d"
	var cl_sec string = "" //"85abad0125e7efdfeb9db2deb1b4155b919c65963deb43a0391ee6209cc71ba9"

	read__c, errs := os.ReadFile("C:\\Users\\rasul\\Documents\\PROJECTS\\GOPROJECT\\Execs\\BotCompiler\\conf.json") //(MAINPATH + "\\Exces\\BotCompiler\\conf.json")
	if errs != nil {
		panic(errs)
	}
	var Creds []map[string]string
	json.Unmarshal(read__c, &Creds)

	clId = Creds[1]["ClientId"]
	cl_sec = Creds[1]["ClientSecret"]

	// for _, contains := range Creds {
	// 	if GetBoardData(contains["UserName"]) < 200 {
	// 		clId = Creds[1]["ClientId"]
	// 		cl_sec = Creds[1]["ClientSecret"]
	// 		break
	// 	}
	// }
	if len(clId) < 1 {
		RetVals <- []string{"-105"}
		return []string{"-105"}
	}

	getProf := reflect.ValueOf(Values["ProfileObj"])
	if reflect.TypeOf(Values["ProfileObj"]).Name() != "Profile" ||
		getProf.Type().Name() != "Profile" { // CHECKS IN TWO WAY TO MAKE SURE OF HAVING PROFILE TYPE AND VALUE
		panic("Given name of object is wrong typed which must be Profile instance")
	}

	read, _ := os.ReadFile(LevelFuncs.ToString(Values["CommonPath"]) + "\\" + getProf.FieldByName("Name").String() + "\\" + LevelFuncs.ToString(Values["filename"])) //MAINPATH + "\\comps\\cpp\\UserCode\\compiler.cpp")
	var ReadAndPassString string = string(read)

	regJs := map[string]string{
		"clientId":     clId,
		"clientSecret": cl_sec,
		"script":       ReadAndPassString, //"using namespace std; \n#include <iostream> \n\nint main() {std::cout << \"Apple is sweet!\" << endl; return 0; }", //"std::cout<< 'Apple is sweet';",
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
	fmt.Println("Compl Stat: '", rVal.CompilationStatus+`'`)
	fmt.Println("Body: ", string(resbod))
	//fmt.Println("res after: ", res)
	//fmt.Println("Compiled Code: ", res["output"])
	//fmt.Println("Stat: ", query.Status)
	//fmt.Println("Response:", query.Request.Response)

	if rVal.StatusCode == 200 {
		//need of proper status indication and defining of it
		if rVal.CompilationStatus == "" || len(rVal.CompilationStatus) == 0 || rVal.CompilationStatus == "1" || rVal.CompilationStatus == "null" || rVal.CompilationStatus == "0.00" {
			fmt.Println("Vals inside:  " + rVal.Output)
			RetVals <- []string{rVal.Output, ""}
			return []string{rVal.Output, ""}
		}
	}

	RetVals <- []string{rVal.Output, fmt.Sprintf("Error: %s", "1")} //rVal.Output
	return []string{rVal.Output, fmt.Sprintf("Error: %s", "1")}     //rVal.Output
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

func WebSocketRunner(params map[string]any, returnValue chan []string) []string /*get 0 resul, 1 status */ {
	var lang string = LevelFuncs.ToString(params["lang"])
	var code string = LevelFuncs.ToString(params["code"])
	if len(lang) < 1 {
		lang = "cpp"
	}
	if len(code) < 1 {
		code = `#include <iostream>\n\nusing namespace std;\n\nint main()\n{\n   std::cout << "ssHello World" << endl << \"Checker Part \tTab check\"; \n   \n   return 0;\n}`
	}

	var cc int = 5 // amount of socket outputs varies
	var string_gatherer_ []string = []string{}
	var outv_ = make(chan string, cc)
	var inv_ = make(chan string, 1)
	// THERE HAVE TO CHECKED THE FORMAT OF THE CODE ESPESSILY HOW IT LOOKS FROM
	// STRING LOOKING POINT

	inv_ <- strconv.Quote(code)
	InitWebsocketClient(cc, lang, outv_, inv_, false, false)
	for v := 0; v < cc; v = v + 1 {
		select {
		case out_tuck := <-outv_:
			string_gatherer_ = append(string_gatherer_, out_tuck)
			fmt.Println("OutPut:" + string_gatherer_[v])
			if strings.Contains(string_gatherer_[v], "ERROR_POINT") {
				returnValue <- []string{string_gatherer_[v][len("ERROR_POINT>"):], "-107"}
				return []string{}
			} else if strings.Contains(string_gatherer_[v], `"output"`) {
				var strdd string = strings.Split(string_gatherer_[v], "[")[1]          // removes all [] list thing
				strdd = strings.Split(strings.Trim(strdd[:len(strdd)-1], " "), ",")[1] // and catches exact code return part from list
				strdd, _ = strconv.Unquote(strdd)
				if len(strdd) > 0 {
					fmt.Println("Out code(CONVERTED PROPER TO CHECK VERSION)>>")
					fmt.Println(strdd)
					returnValue <- []string{strdd, "RES"}
					return []string{}

				}
			}
		default:
			break
		}
	}
	returnValue <- []string{"malfunctioning	", "-106"} // the online compiler unsucces code
	return []string{}

}
