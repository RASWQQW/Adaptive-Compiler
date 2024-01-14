// bot actually sends the question
// in one language then it waits till it translates
// then this code keeps as global file or text
// and other check cases are just copies global code and changes value
package BotCompiler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
	"strings"
)

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

func Requester2() {

	var regLink string = "https://api.jdoodle.com/v1/execute"
	var clId string = "b86eafbab040cfd06f729b4f7d233f2d"
	var cl_sec string = "85abad0125e7efdfeb9db2deb1b4155b919c65963deb43a0391ee6209cc71ba9"

	regJs := map[string]string{
		"clientId":     clId,
		"clientSecret": cl_sec,
		"script":       "print('Apple is sweet')",
		"language":     "python3"}

	regJsString, _ := json.Marshal(regJs)

	//WITH QUERY LIBRARY
	query, _ := http.Post(regLink, "application/json", strings.NewReader(string(regJsString)))
	defer query.Body.Close()

	res := make(map[any]any)
	rVal := ReadVals{}

	resbod, _ := ioutil.ReadAll(query.Body)
	json.Unmarshal(resbod, &res)

	fmt.Println("res before: ", rVal)
	fmt.Println("res after: ", res)
	fmt.Println("StatCode: ", query.StatusCode)
	fmt.Println("Body: ", string(resbod))
	fmt.Println("Compiled Code: ", res["output"])

	//fmt.Println("Stat: ", query.Status)
	//fmt.Println("Response:", query.Request.Response)
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
