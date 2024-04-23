package generator

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func Checker() {
	var regStr string = "https://www.blackbox.ai/api/chat"
	var usa string = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36"
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
								"content": "write for loop in python to count 100",
								"role": "user"}],
							"previewToken": null,
							"trendingAgentMode" : {},
							"userId" : "9ae48abf-5aae-48ec-940b-c044dfc04298"}`

	req, _ := http.NewRequest("POST", regStr, bytes.NewBuffer([]byte(regString)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", usa)
	var cl *http.Client = &http.Client{}
	res, _ := cl.Do(req)
	red, _ := io.ReadAll(res.Body)
	fmt.Println("gotten value << " + string(red))
}
