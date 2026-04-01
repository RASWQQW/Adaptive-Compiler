package generator

import (
	"bytes"
	mts "ep/Execs/methods"
	baseSS "ep/base/abst"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
)

func HeaderExrtacter() {
}

func Translator(promnt_word string) string {
	return ""
}

// to get from task creator
// in order to check functions are clearly works and mathces the already given results of task creator
// WILL COME AS OPTIONAL
func TestCases(inVals map[string]string, OutVals map[string]string) {

}

func CoverTaskId(func_name string, func_params [][]string, returnType string) {
	var prompt_Example string = `write 'Suffix ' function to find 'S' and returning amount of it in given 'word' parameter in cpp sent only code not any other words`
	fmt.Println(prompt_Example)
}

// IT HAS TO BIND WITH POST OF WEB WHERE GENERATED CODE AND DETAILES ARE CONFIRMED TO SAVE IN BASE
func BasePart(properCode string, FuncMadeName string, ret_val_typ string, funArgs [][]string) {

	/*HERE ALSO COMES VALUES PICKED BY USER HERE TO CHANGE VALUES OF WHEN SAVING */

	var crBase = baseSS.GetBase()
	// first had to generate task id and task
	var task_id int64 = crBase.CreateFuncDec(FuncMadeName, ret_val_typ)
	// FUNCTION ARGS, RETURN VALUE AND NAME ARE ALL WILL PARSED FROM PROPER CODE SO THE CREATOR IS ALREADY CHECKED

	crBase.CreateFuncArgs(task_id, funArgs)
	crBase.CreateProperCode("cpp", properCode, task_id)

}

// IT HAD TO GENERATE ANY PROMPT FROM FORM WHEN USER CLICKED GENERATE
func GeneratorOfPrompt(prompt string,
	lang string,
	to_translate bool) (string, string, string, [][]string) /*first name and ret type and second val is for params */ {
	if to_translate {
		prompt = Translator(prompt)
	}

	var comment_text string = "/*PROPER_CODE_STARTS*/"
	prompt = strings.TrimSpace(prompt) + "\nGive a very simple and clear example of code. Only code no text" // to make sure it only sends code
	prompt = prompt + ". Always put only " + comment_text + " comment in the first space of the code"
	// prompt = prompt + ". Do not add main() function and do not run given function"
	// prompt = `write 'Suffix ' function to find 'S' and returning amount of it in given 'word' parameter in cpp sent only code not any other words`

	var regStr string = "https://www.blackbox.ai/api/chat"
	var usa string = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36"
	var userId string = `9ae48abf-5aae-48ec-940b-c044dfc04298` // one may to scarcly change
	userId = "666a9519b0673"

	var regString string = `{"agentMode": {},
							"clickedAnswer2": false,
							"clickedAnswer3": false,
							"codeModelMode": true,
							"githubToken": null,
							"id": "tMD4rk0",
							"isChromeExt": false,
							"isMicMode": false,
							"messages": [{
								"id": "uYSIVzK", 
								"content": "` + prompt + /*"write for loop in python to count 100"*/ `", 
								"role": "user"}],
							"previewToken": null,
							"trendingAgentMode" : {},
							"userId" : "` + userId + `"}` // userId , message.id might need dynamic change of value?

	req, _ := http.NewRequest("POST", regStr, bytes.NewBuffer([]byte(regString)))
	// req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", usa)
	var cl *http.Client = &http.Client{}
	res, _ := cl.Do(req)
	//red11, _ := io.ReadAll(res.Request.Body)
	//var red2 string = string(red11)
	//fmt.Println(red2)
	red, _ := io.ReadAll(res.Body)

	fmt.Println("gotten value << " + string(red))
	var properCodeFetch []string = strings.Split(string(red), "```") // just fetching by cutting  to the three

	var properCodeFetched string = ""
	//var matchString string = "\n#include<iostream>\nusing namespace std;" // for other template of retrun from compiler
	for vv := len(properCodeFetch) - 1; vv >= 0; vv-- {
		for _, langs := range []string{"cpp", "java", "csharp", "python"} {
			if len(properCodeFetch[vv]) > len(langs) &&
				properCodeFetch[vv][:len(langs)] == langs &&
				langs == "cpp" {
				properCodeFetched = properCodeFetch[vv][len(lang):]
				break
			} else if len(properCodeFetch[vv]) > len(comment_text) &&
				properCodeFetch[vv][:len(comment_text)] == comment_text {
				properCodeFetched = properCodeFetch[vv]
				fmt.Println(properCodeFetched)
				break
			}
		}
	}
	fmt.Println("gotten code << " + properCodeFetched)

	// TO FIND EXACT GINVEN FUNC NAME FROM PROMPT
	var FunctionName string = ""
	if isindex := strings.Index(prompt, "function name:"); isindex != -1 {
		var FuncName = strings.Split(prompt[isindex:], " ")[0][len("function name:"):]
		fmt.Println("Function's name: " + FuncName)
	}

	regexd := regexp.MustCompile("^(double|int|string|float|char)[*]{0,2} " + FunctionName + "( )*\\((.)+\\).+{(\r\n|\r|\n||.)+}$")
	var FuncBody string = regexd.FindString(properCodeFetched)

	// var ProperCode string = properCodeFetch

	FuncBody = mts.Replacer([]string{"static", "public", "private"}, FuncBody)
	var blankSep []string = strings.Split(FuncBody, " ")

	// find return type
	var retType string = blankSep[0]       // [`string`]
	var afterFuncName string = blankSep[1] // [`FuncName` ...]

	var bracketsSep []string = strings.Split(afterFuncName, "(")

	// find a name
	var ItsFuncName string = bracketsSep[0]
	var ParamsJoined []string = strings.Split(strings.Split(bracketsSep[1], ")")[0], ",")

	var Cutter = func(befcut []string) []string {
		for dd, _ := range befcut {
			befcut[dd] = strings.ReplaceAll(befcut[dd], " ", "")
		}
		return befcut
	}
	// find a params
	// THIS FUNC ESPESSIALLY WRITTEN FOR CPP TEMPLATE
	//  TO: int val[1] = {}
	var Params [][]string = [][]string{}
	for dd := 0; dd < len(ParamsJoined); dd++ {
		var ParamVs []string = Cutter(strings.Split(ParamsJoined[dd], " "))
		if stp := strings.Split(ParamVs[1], "["); len(stp) > 1 {
			ParamVs[1] = stp[0]
			ParamVs[0] = ParamVs[0] + "[" + stp[1] // its to put into full for type
		}
		Params = append(Params, ParamVs)
	}

	// mts.FindMatrix()

	// next goes saving into base
	/// THIS CODE JUST RETURNS GENERATED VALUE
	return properCodeFetched, ItsFuncName, retType, Params // there might be need to look at generated code
}
