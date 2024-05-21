package generator

import (
	"bytes"
	baseSS "ep/base/abst"
	"fmt"
	"io"
	"net/http"
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

func BasePart(properCode string, FuncMadeName string, ret_val_typ string) {

	var crBase = baseSS.GetBase()

	// first had to generate task id and task
	var task_id int64 = crBase.CreateFuncDec(FuncMadeName, ret_val_typ)
	crBase.CreateProperCode("cpp", properCode, task_id)
}

func Checker(prompt string, lang string, to_translate bool) map[string]string {
	if to_translate {
		prompt = Translator(prompt)
	}

	prompt = `write 'Suffix ' function to find 'S' and returning amount of it in given 'word' parameter in cpp sent only code not any other words`

	var regStr string = "https://www.blackbox.ai/api/chat"
	var usa string = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36"
	var userId string = `9ae48abf-5aae-48ec-940b-c044dfc04298` // one may to scarcly change

	var regString string = `{"agentMode": {},
							"clickedAnswer2": false,
							"clickedAnswer3": false,
							"codeModelMode": true,
							"githubToken": null,
							"id": "tMckt0d",
							"isChromeExt": false,
							"isMicMode": false,
							"messages": [{
								"id": "tMckt1d", 
								"content": "` + prompt + /*"write for loop in python to count 100"*/ `", 
								"role": "user"}],
							"previewToken": null,
							"trendingAgentMode" : {},
							"userId" : "` + userId + `"}` // userId , message.id might need dynamic change of value?

	req, _ := http.NewRequest("POST", regStr, bytes.NewBuffer([]byte(regString)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", usa)
	var cl *http.Client = &http.Client{}
	res, _ := cl.Do(req)
	red, _ := io.ReadAll(res.Body)

	fmt.Println("gotten value << " + string(red))
	var properCodeFetch string = strings.Split(string(red), "```")[1] // just fetching by cutting  to the three
	properCodeFetch = properCodeFetch[len(lang):]
	fmt.Println("gotten code << " + string(properCodeFetch))

	var ProperCode string = properCodeFetch

	// next goes saving into base
	return map[string]string{} // there might be need to look at generated code
}
