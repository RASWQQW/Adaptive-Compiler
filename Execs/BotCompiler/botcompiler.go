// bot actually sends the question
// in one language then it waits till it translates
// then this code keeps as global file or text
// and other check cases are just copies global code and changes value

// CHANGES WHCIH REQUIRED
// MAKE TIME OF CODE AND MEMBERY USAGE IMORTANT AS IT  WILL  LEVERAGE USER STAT ON EACH RAN CODE/TOTAL SCORE

// COMPILER WEBSITE
// https://www.tutorialspoint.com/compile_cpp_online.php
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
	"sync"
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
	RunFunc func(map[string]any, chan []string, sync.Mutex) []string,
	lc sync.Mutex) {
	if RunFunc == nil {
		//It runs func of request for default if there is no given func
		RunFunc = CompilerRequester
	}
	go RunFunc(params, ReturnValue, lc) // Runs in background until code it ends or limit time comes to an end
	time.Sleep(time.Second * time.Duration(sleepTime))
	//if the given time ends fun will set value  to chan by itself
	ReturnValue <- []string{"Time limit exceeded", "TIMELIMIT", "1"}
}

// The idea is making possible of running custom codes with their id of file
// whether is user or proper code on one compile request
// and getting all results by specification
func CodeBatcher(gatherer *LevelFuncs.BatchGathererList, stat string) { //(addCode string, headers string) {
	//var out_CompileCode string = stat + ` `
	//var out_FuncRunnerInj string = `` //  I have sitten code of thread running of given funcs

	var RepresenterMethod []string = gatherer.CllRepresentString
	var out_BuiltInLibs string = `
					using namespace std;
					#include <iostream> 
					#include <thread>
					#include<chrono>
					#include <vector>
					#include <string>
					#include <sstream> `

	var out_ParamsDefStruct string = "struct params {\n"
	var ParamsToTransmitFuncTemp string = ""
	var ParamsToPassTemp string = ""
	var TypesToPassTemp string = ""
	var ParamsToMainFuncTemp string = ""
	var concator string = ""

	for _, valAdder := range gatherer.FuncParamNamesAndTypes {
		TypesToPassTemp = concator + TypesToPassTemp + valAdder[0]                                       // temp  "string, int, ..."
		ParamsToPassTemp = concator + ParamsToPassTemp + valAdder[1]                                     // temp  "val1, val2, ..."
		ParamsToTransmitFuncTemp = concator + ParamsToTransmitFuncTemp + valAdder[0] + " " + valAdder[1] // temp  "string  val1, int val2, ..."
		ParamsToMainFuncTemp = concator + ParamsToMainFuncTemp + "values__lister[v]." + valAdder[1]      // temp  "values__lister[v].val1, values__lister[v].val2, ..." // main values lister in the point

		concator = ", " // after all used strings it goes to pass

		//Add values to the struct
		out_ParamsDefStruct = out_ParamsDefStruct + valAdder[0] + " " + valAdder[1] + ";\n" // temp  "string  val1; int val2;..."
	}

	out_ParamsDefStruct = out_ParamsDefStruct + "};" // closing code of struct

	var out_FuncReturnType string = `typedef ` + gatherer.FuncReturnType + ` to_run_ret;`
	var out_TypeDefs string = `
			typedef ` + gatherer.FuncReturnType + ` to_run_ret;
			typedef to_run_ret (*run_code_temp)(` + TypesToPassTemp /*types only no need of names*/ + `); //  there we need just pass parameters of func when preparing
			`

	// CONSTRUCT TRANSMIT FUNC
	var out_transmitFuncChanger string = `
			void transmitFunc(run_code_temp to_run, 
						void(*Representer)(string[], to_run_ret), 
						vector<vector<string>> *arr, 
						string FuncId,
						int place,
						float time_limit_val,
						` + ParamsToTransmitFuncTemp + `/*THERE FUNC WILL BE PREPARED AS TEMPLATE TO OWN PARAMS*/){

			std::chrono::steady_clock::time_point begin_t = std::chrono::steady_clock::now();
			std::stringstream s;

			try{
				//std::cout << "FUNC EXCEEDED CHECKbefore" << endl;
				string pointer[2] = {"0", "0"};
				///throw runtime_error("CHECK_TEXT!");
				auto runner = [Representer, to_run, ` + ParamsToPassTemp + `, &pointer]() {
					Representer(pointer, to_run(` + ParamsToPassTemp + `));};

				std::thread ddx(runner);
				ddx.detach();
				// ddx.join();
				// float checker = ((place + rand() % 3)) * ((float) 1 / 21);
				//cout << "FUNC EXCEEDED CHECK1 pointer: " << pointer << endl;
				int st = 0;  
				while(st < (time_limit_val*1000` + /*THERE JUST LAYS FEATURE /15.6436/ - NEEDED ONLY WHEN IN LOCAL*/ `)){
					std::this_thread::sleep_for(std::chrono::milliseconds(1));
					st = st + 1;
					if(pointer[1] == "1"){
						break;
					}        
				}
				std::chrono::time_point end_t = std::chrono::steady_clock::now();
				string diff = to_string((float)(std::chrono::duration_cast<std::chrono::milliseconds>(end_t-begin_t).count()) / 1000);
				//cout <<  "END SEC: " << st << " get pointer: " << pointer[0]  << "Time diff: " << diff<< endl;

				if(pointer[1] != "1" && pointer[1] == "0"){
					(*arr).at(place) = {FuncId, diff};/// )};
				} else { 
					(*arr).at(place) = {FuncId, diff, pointer[0]}; /// )};
				}
				//cout << pointer[1] << " " << pointer[0] << "out_res: " + (*arr).at(place)[0] << " " << (*arr).at(place)[1] << (*arr).at(place)[2];

			}catch(const exception &e){
				// make sense to notify it gone from outer codes not from templates
				//cout << "come to error" << endl;
				s << e.what();
				//cout << "errorcode: " << s.str() << endl;
				(*arr).at(place) = {FuncId, "<code_exec_error>" + s.str() + "</code_exec_error>"}; /// )};
			}
		}
	`

	//There Have to bes passed list of
	//var FuncNamesLists string = `list<` + returnType + `(*)(` + passingParamsTypes + `)> RUNNFUNCS {`

	var FuncsDeclarings string = ``

	var ValuesListerRealValues string = "{"
	var FuncIds string = ""
	var FuncNamesInLIst string = "{"

	var FuncsRunList []string = []string{} // there i have to show funcs list syntax

	var funcCreate = func(name string, funcCode string) {
		FuncsDeclarings = FuncsDeclarings + "\n\n" + funcCode[:strings.Index(funcCode, " ")-1] + name + funcCode[strings.Index(funcCode, "("):]
		FuncsRunList = append(FuncsRunList, name) //fmt.Sprintf(gatherer.CllRepresentString, name, gatherer.CllTypePassingString)

		// Add all funcIds to the list
		if len(FuncIds) > 0 {
			FuncIds = FuncIds + ", "
		}
		FuncIds = FuncIds + (name + "FuncId")
	}

	for vv := 0; vv < len(gatherer.Collection); vv = vv + 1 {
		var nname string = gatherer.Collection[vv].CllProfile.Name
		funcCreate(nname+"ProperBased", gatherer.Collection[vv].CllProperCode)
		funcCreate(nname+"UserBased", gatherer.Collection[vv].CllUserCode)
	}

	FuncNamesInLIst = strings.Join(FuncsRunList, ", ")

	// time rewriter
	var Time_limits_list string = "{"
	var Values_temp string = ""
	var SIZE_COUNT int = len(gatherer.Collection)
	concator = ""
	for col := range gatherer.Collection {
		var ttime float32 = gatherer.Collection[col].ParamTimeLimitToRun
		if ttime < 3.40 {
			gatherer.Collection[col].ParamTimeLimitToRun = ttime + 0.035 // there would little interval to get any trepsaces avoided
		}
		Time_limits_list = concator + Time_limits_list + LevelFuncs.ToString(ttime)

		// ALL TYPES OF GENERATIVE IN CODE MUST BE STAND IN ORDER WHERE PROPERCODE THEN USERCODE COMES ONE AND AFTER
		// Values lister process starts
		Values_temp = "{"

		for ii := range gatherer.FuncParamNamesAndTypes { // need to write some relevant search
			for innerB := range gatherer.Collection[col].ParamBuncher {
				if gatherer.FuncParamNamesAndTypes[ii][1] == gatherer.Collection[col].ParamBuncher[innerB][0] {
					if len(Values_temp) > 1 {
						Values_temp = Values_temp + ", "
					}
					Values_temp = Values_temp + gatherer.Collection[col].ParamBuncher[innerB][2]
					break
				}
			}
		}
		if len(ValuesListerRealValues) > 1 {
			ValuesListerRealValues = ValuesListerRealValues + ", "
		}
		ValuesListerRealValues = ValuesListerRealValues + (Values_temp + "}")

	}

	var out_MainFuncChanger string = `
	// I HAVE TO CREATE TRANSMIT FUNCTION TO PERFORM A REPRESENT FUNCTION
	int main(){
		
		// THERE GOES CHECKER OF TIME LIITER LOGIC IF THERE COMES AT LEAST UNMATCHED TIME
		// IT WILL JUST STOP AND KILL ANY OTHER PROCESESS
		//std::thread(CheckerDD, Printer, "CheckerTest")

		// pre-forming values
		params values__lister[` + strconv.Itoa(SIZE_COUNT/2) + `] = ` + ValuesListerRealValues + `;
		string FuncIds[` + strconv.Itoa(SIZE_COUNT) + `] = {` + FuncIds + `};
		float given_time_limits[` + strconv.Itoa(SIZE_COUNT) + `] = ` + Time_limits_list + `//{3.12, 3.12};
		vector<run_code_temp> objectsd = ` + FuncNamesInLIst + `//{Printer, Printer};
		
		// after forming values
		vector<vector<string>> catcher = {};
		vector<std::thread> threads;

		int Objects_size = objectsd.size();
		int valSent = 0;

		// promise<vector<string>> promiser;
		// std::future<vector<string>> ff = promiser.get_future();
		for(int v = 0; v < Objects_size; v ++){
			// THERE IT GOES VIA RUNTIME AND 
			// EACH PASSING ARGUMENT WILL HAVE PARAM NAME AS WELL WHEN PASSING
			catcher.push_back({FuncIds[v], "-1", "-1"});
			/*  v - is place of func to  locate its value on a list
				given_time_limits - is unique time limit collection of each func
			*/
			threads.push_back(std::thread(transmitFunc, Printer,  ` + RepresenterMethod[0] + `, &catcher, FuncIds[v], v, given_time_limits[v], /**//* further goes changed values */` + ParamsToMainFuncTemp /*values__lister[v].val1, values__lister[v].val2*/ + `));//
			if (v % 2 == 1) {
				valSent = valSent + 1;
			}
		}
		
		for(auto&& isd : threads) {
			isd.join();
		}
		
		std::cout << "FUNC GETTIME" << endl;
		vector<vector<string>> FinalCatcher;
		vector<int> OnPerformFuncs;
		std::vector<int>::iterator it;
	
		// fetching the results
	
		int conc = 0; //to handle finsihed funcs
		int conc2 = conc; // to make cycle on func check proceess
		bool istostop = false;
		string funcId = "";
		int reverser = 1;
		while (true) {
			
			if(istostop) {
				break; // stops when there is no need
			}
	
			if(OnPerformFuncs.size() >= Objects_size){
				break;
			}
	
			if(conc >= Objects_size){ // if there is loop to look out is ended
				if(reverser == 1 || int(OnPerformFuncs.size() >= Objects_size / 2)){
					conc = int(Objects_size / 2);
					reverser = 0;
				}else{
					conc = 0;
					reverser = 1;
				}
				// int min_to_start = 999; //it may change
				// OnPerformFuncs.clear();
	
				// for(int in = 0; in < catcher.size(); in++){
				//     it = std::find(OfPerformFuncs.begin(), OfPerformFuncs.end(), in);
				//     if (it != OfPerformFuncs.end()){    
				//         OnPerformFuncs.push_back(conc);
				//         if(min_to_start > conc){
				//             min_to_start = conc;
				//         }
				//     }
				// }        
				// conc = min_to_start;
			}
			
			it = std::find(OnPerformFuncs.begin(), OnPerformFuncs.end(), conc);
			if (it == OnPerformFuncs.end()){
				// vector<string> catcher = ff.get();
				funcId = catcher[conc][0];
				cout << "<rrn::"<< funcId << ">";
				//cout << "Size: " << sizeof((catcher[conc][0])) << "Div: " << sizeof(string) << sizeof(vector<string>);
				if(catcher[conc].size() >= 2 
					// && (catcher[conc][0] != "-1" 
					&& catcher[conc][1] != "-1" 
					&& catcher[conc][2] != "-1"
					){
				
					//cout << "Run time: " << catcher[conc][0] << endl << catcher[conc][1] << sizeof(catcher) / sizeof(vector<string>);
					//cout << "Error is true: " << catcher[conc][1].find("code_exec_error");
					if(catcher[conc][1].find("code_exec_error") != std::string::npos){
						istostop = true;
						cout << catcher[conc][1]; // to print the error
						//break; // no more will come any compilation
					}else if(stod(catcher[conc][1]) > TIME_LIMIT){
						istostop = true;
						// THERE HAVE TO BE LOGIC OF STOPPING THE PROCESS
						cout << "<time_exceeded_err" << " time::" << catcher[conc][1] << " func::" << funcId << "/>";
						//break; // no more will come any compilation
					}else{
						if (catcher[conc].size() >= 3 || catcher[conc][2] == "-1"){
							cout << "<res>" << catcher[conc][2] << "</res>"; // finally printing the fun result
							cout  << endl << "<stat::run_time:" << catcher[conc][2] <<  "/>"; // there i should add value of mempor usage
						}
					}
					// THERE IS JUST NO NEED OF PRINTING TWICE ALL WE NEED
					// IS GETTING ITS ERROR OR TIME LIMIT EXEED ERROR
					// cout << "<" + catcher.at(0) + ">" + catcher.at(1) + "<>";
					// FinalCatcher.push_back(catcher);
					conc2 = conc2 + 1;
					OnPerformFuncs.push_back(conc);
				}
				conc = conc + 1;
				cout << "</rrn::"<< funcId << ">\n";
			}
		}
	
		// std::thread object_t(ffer, 3);
		// std::thread object_t2(ffer, 1);
		// auto  funcs[]  = {{ffer, }, ffer};  
		// object_t.join();
		// object_t2.join();
		// object_t2.detach();
	
		return 0;
	}`

	// HAVE TO ADD TWO FUNCS OF REPRESENTER
	// FITTING OF FUNC RETURN TYPE

	//  I NEED TO TEST USING ALREADY CREATED OBJECTS SENDING INTO THIS METHOD
	// IF I WANT TO TEST IT TO FULL CORRECTNESS

	var RetCode string = ``
	RetCode = stat + "\n\n" + out_BuiltInLibs +
		"\n\n" + out_ParamsDefStruct +

		"\n" + RepresenterMethod[1] +

		"\n" + out_FuncReturnType +
		"\n" + out_TypeDefs +
		"\n" + out_transmitFuncChanger +
		"\n" + out_MainFuncChanger

	res_getter := make(chan []string)
	var lsc sync.Mutex
	WebSocketRunner(map[string]any{"lang": "cpp", "code": RetCode}, res_getter, lsc)

	// THERE HAVE TO STAY EXTRACTING OF PROPER RESULT FROM PRINTED CODE

}

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

func CompilerRequester(Values map[string]any, RetVals chan []string, lsc sync.Mutex) []string {
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
		RetVals <- []string{"no token", "-105"}
		return []string{"no token", "-105"}
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

	RetVals <- []string{rVal.Output, fmt.Sprintf("Error: %s", "-108")} //rVal.Output
	return []string{rVal.Output, fmt.Sprintf("Error: %s", "-108")}     //rVal.Output
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

func WebSocketRunner(params_c map[string]any, returnValue chan []string, lsc sync.Mutex) []string /*get 0 result, 1 status */ {
	var lang string = LevelFuncs.ToString(params_c["lang"])
	var code string = LevelFuncs.ToString(params_c["code"])

	// lsc.Lock()
	// defer lsc.Unlock()

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
	// STRING LOOKING POINTa

	// getstr, _ := strconv.Unquote(fmt.Sprintf(`"%s"`, code))
	inv_ <- code
	InitWebsocketClient(cc /*lang need to change */, "cpp", outv_, inv_, false, false)
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
					returnValue <- []string{strdd, ""}
					fmt.Println("Settled vals:")
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
