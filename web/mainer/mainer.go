package Mainer

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	inputing "ep/Execs"
	gn "ep/Execs/GENERATE_TASKS"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var messagesd = make(chan string, 1)

func Runnersss() {

	fmt.Println("Page is started")
	var mod = gin.Default()
	var ups = websocket.Upgrader{} //ReadBufferSize: 1024, WriteBufferSize: 1024

	mod.LoadHTMLGlob("web/mainer/temps/html/*")
	mod.GET("/htesting", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{"OUTVAL": "Click on RUN button to see the output"})
	})

	// BOTH TO GENERATED CODE
	mod.GET("/cdgen", func(c *gin.Context) {
		c.HTML(http.StatusOK, "taskCreateFrom.html", gin.H{})
	})

	// EACH GENERATED TASK FIRSTLY COMES TO HERE AND IT RETURNS BACK TO  /CDGEN
	mod.GET("/code_generate_attempt", func(c *gin.Context) {
		ups.CheckOrigin = func(r *http.Request) bool { return true }
		con, err := ups.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			panic(err)
		}
		fmt.Println("Connected to js: ", con.RemoteAddr())
		defer con.Close()
		for {
			fmt.Println("waiting sent message...")
			vv, message_get_from, err := con.ReadMessage()

			if err != nil {
				panic(err)
			}

			fmt.Println("message type: ", vv)
			fmt.Println("sent array type message: ", strconv.Itoa(vv))

			if len(message_get_from) > 3 {

				var PromptText string = ""
				var mess_splits []string = strings.Split(string(message_get_from), "|")
				for v := 0; v < len(mess_splits); v++ {
					if len(mess_splits[v]) > len("PROMPT_STRING") {
						if mess_splits[v][:len("PROMPT_STRING")] == "PROMPT_STRING" {
							PromptText = mess_splits[v][len("PROMPT_STRING")+1 : len(mess_splits[v])-1]
						}
					}
				}

				properCodeFetch, func_name, ret_val, params := gn.GeneratorOfPrompt(PromptText, "cpp", false)
				var isRetString string = ""
				isRetString = isRetString + "|genCode:" + properCodeFetch
				isRetString = isRetString + "|func_name:" + func_name
				isRetString = isRetString + "|return_val:" + ret_val

				var StringedParams []string = []string{}
				for _, vad := range params {
					StringedParams = append(StringedParams, "{"+vad[0]+"|"+vad[1]+"}")
				}

				isRetString = isRetString + "|params:[" + strings.Join(StringedParams, ", ") + "]"
				con.WriteMessage(vv, []byte(isRetString))
			}
		}
	})

	// IT TO SAVE THE LAST GEN RESULT
	mod.POST("/cdgen", func(c *gin.Context) {
		c.HTML(http.StatusOK, "taskCreateFrom.tmpl", gin.H{})
	})

	mod.GET("", Main)
	mod.GET("/form", Form)
	mod.GET("/cpprun", CppInter)

	mod.GET("/getr/:roomid", func(c *gin.Context) {
		c.SSEvent("Start", "At first step")
	})

	mod.GET("/msw/:ms", func(c *gin.Context) {
		var msg = c.Params.ByName("ms")
		messagesd <- msg
		c.JSON(200, msg+"is sent")
	})

	mod.GET("/checker/:name/:val?", func(c *gin.Context) {
		c.JSON(200, gin.H{"c": c.Keys, "d": c.Params /*"pp": c.Request.URL*/})
	})

	// mod.GET("/checkcode/", func(con *gin.Context) {
	// 	currentTask := con.Query("taskId")

	// 	cons, err := ups.Upgrade(con.Writer, con.Request, nil)
	// 	defer cons.Close()

	// 	// LANG IS DETERMINES BY TIME
	// 	var lang string = ""
	// })

	mod.GET("/checkcode", func(w *gin.Context) {
		ups.CheckOrigin = func(r *http.Request) bool { return true }
		con, err := ups.Upgrade(w.Writer, w.Request, nil)
		if err != nil {
			panic(err)
		}
		fmt.Println("Connection from: ", con.RemoteAddr())
		II := 0
		defer con.Close()
		for {
			II = II + 1
			fmt.Println("web s Is running", II)

			vv, message_get_from, err := con.ReadMessage()

			if err != nil {
				panic(err)
			}

			fmt.Println(vv, reflect.ValueOf(vv).Kind(), "Client message", message_get_from)
			//time.Sleep(time.Second * 1)
			fmt.Println("Current text ", strconv.Itoa(vv))
			if len(message_get_from) > 3 {
				var message = inputing.Main("cpp", string(message_get_from), "AddTwoNumbers")
				con.WriteMessage(vv, []byte(message))
			} else {
				con.WriteMessage(vv, []byte("Input code is too small"))
			}
			// if 1 != 1 {

			// } else {
			// 	con.WriteMessage(vv, []byte(inputing.Main("cpp", string(message_get_from), "AddTwoNumbers")))
			// }
		}
	})
	mod.Run("127.0.0.1:8080")
}

func Main(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"dd": "values"})
}

func Form(c *gin.Context) {
	var b Nstructs
	c.Bind(&b)
	c.JSON(http.StatusOK, gin.H{
		"a": b.Aval,
		"b": b.Bval,
	})
}

func CppInter(c *gin.Context) {

}

type Nstructs struct {
	Aval string `form:"fa"`
	Bval string `form:"fb"`
}

type Nstruct struct {
	FieldX string `form:"field_x"`
	FieldD string `form:"field_d"`
}
