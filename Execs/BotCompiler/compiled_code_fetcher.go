package BotCompiler

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"

	"github.com/gorilla/websocket"
)

// var address string = "wss://cpp.repl-web.programiz.com/socket.io/?sessionId=YrNOQVqNZv&lang=cpp&EIO=3&transport=websocket"
var UserAgent string = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36"

// THERE I MIGHT USE ONE ID IN MANY QUERY
// AND ONLY DROP IT WHEN RANS OUT OF USE
func GetSessionId(lang string) string {
	var add_name_string string = ""
	if lang == "cpp" {
		add_name_string = "compile_cpp_online"
	}
	var Glink_ string = "https://www.tutorialspoint.com/" + add_name_string + ".php"
	req, err := http.NewRequest("GET", Glink_, nil)

	if err != nil {
		panic(err)
	}

	req.Header.Set("User-Agent", UserAgent)
	var cl *http.Client = &http.Client{}
	res, _ := cl.Do(req)
	red, _ := io.ReadAll(res.Body)

	templ := regexp.MustCompile(`var sessionId = "[A-Za-z0-9]+";`)
	var StringFound string = templ.FindString(string(red))
	fmt.Println("Code snippet: " + StringFound)
	var SessionFinder string = strings.Replace(strings.Replace(StringFound, `var sessionId = "`, "", -1), `";`, "", -1)
	return strings.Trim(SessionFinder, " ")
}

func InitWebsocketClient(getLimit int, lang string, OutValue chan string, CodeToSolve chan string, HandTest bool, Cycling bool) {
	// give value to the variables
	var SessionId string = GetSessionId(lang)
	fmt.Println("Starting Client")
	var address string = "wss://tpcg2.tutorialspoint.com/socket.io/?sessionId=" + SessionId + "&lang=csharp&EIO=4&transport=websocket" //65fc92c52c21c
	var GivenHeader http.Header = http.Header{}
	GivenHeader.Add("User-Agent", UserAgent)

	//GivenHeader.Add("Host", "cpp.repl-web.programiz.com")
	//GivenHeader.Add("Sec-WebSocket-Key", "IdZCJgyc19fy9soaeuiIWA==")
	conn, resp, err := websocket.DefaultDialer.Dial(address, GivenHeader)
	if err != nil {
		fmt.Println("dial Error:", err)
	}
	defer conn.Close()

	conn.WriteMessage(1, []byte("40"))
	//var SentArray = `42["code",{"code":"/* Online C++ Compiler and Editor */\n#include <iostream>\n\nusing namespace std;\n\nint main()\n{\n   std::cout << \"ssHello World\" << endl; \n   \n   return 0;\n}","language":"cpp","sessionId":"65fca1a2239b7"}]`
	var SentArray string = `42["code", {"code": "` + strings.TrimSuffix(<-CodeToSolve, "\n") + `", "language":"cpp", "sessionId": "` + SessionId + `"}]`
	conn.WriteMessage(websocket.TextMessage, []byte(SentArray))

	if resp != nil && resp.Header != nil {
		fmt.Println("Resonse>>")
		for key, val := range resp.Header {
			fmt.Println("Key: " + string(key) + " Val: " + string(val[0]))
		}
	}

	var done = make(chan string)
	//	var sett sync.WaitGroup
	go func(OutMsg chan string) {
		var queryAmount int = 0
		defer close(done)
		for {
			mt, message, err := conn.ReadMessage()
			if err != nil {
				fmt.Println("Gotten Error:", err)
				OutMsg <- fmt.Sprintf("ERROR_POINT>%v", err)
				return
			}
			queryAmount = queryAmount + 1
			OutMsg <- string(message)
			fmt.Println("size: %s, message: %s", mt, string(message))
			if queryAmount >= getLimit { // process ends for one query socket when there been handled moer than 3 response
				return
			}
		}
	}(OutValue)
	var val any = <-done
	fmt.Printf("End >> %s", val)

	if HandTest && Cycling {
		for {
			select {
			case <-done:
				conn.Close() // when end intiated from runner or cause of error
				return
			case getCode := <-CodeToSolve:
				// fmt.Println("Write a code: ")
				// fmt.Println("Stdin val: " + string(os.Stdin.Name()))

				// reader := bufio.NewReader(os.Stdin)
				// code, _ := reader.ReadString('\n') //code = `#include <iostream> using namespace std; int main() { std::cout << "Hello World" << endl; return 0; }`
				// fmt.Println("Given code", code)

				var SentArray string = `42["code", {"code": "` + strings.TrimSuffix(getCode, "\n") + `", "language":"cpp", "sessionId": "65fc93e27ba2f"}]`
				err := conn.WriteMessage(websocket.TextMessage, []byte(SentArray))

				if err != nil {
					fmt.Println("Error when giving value :", err)
					return
				}
			}

		}
	}
}
