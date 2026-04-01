package Mainer

import (
	"fmt"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	inputing "ep/Execs"
	gn "ep/Execs/GENERATE_TASKS"

	"ep/Execs/obj"
	lvl "ep/LevelFuncs"
	bs "ep/base/abst"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var messagesd = make(chan string, 1)

func Runnersss() {

	fmt.Println("Page is started")
	var mod = gin.Default()
	var ups = websocket.Upgrader{} //ReadBufferSize: 1024, WriteBufferSize: 1024

	mod.LoadHTMLGlob("web/mainer/temps/html/*")
	mod.GET("/htesting/:topicname", func(c *gin.Context) {

		var bss = bs.GetBase()

		var topicName string = c.Param("topicname")
		var crs = &obj.Career{INOUTS: map[string]any{"lang": "cpp", "task_name_id": topicName /*"AddTwoNumbers"*/}, OUTS: map[string]any{}}
		bss.GetTaskByName(crs)
		bss.GetProperCode(crs)
		bss.GetFunction(crs)

		var FindCode string = regexp.MustCompile("(.)+{").FindString(lvl.ToString(crs.OUTS["prop_code"]))
		FindCode = strings.ReplaceAll(FindCode, "RET_TYPE", lvl.ToString(crs.OUTS["return_type"]))
		FindCode = strings.ReplaceAll(FindCode, "FUNC", topicName)
		c.HTML(http.StatusOK, "index.tmpl",
			gin.H{
				"TOPICS":        topicName,
				"FUNCSTARTCODE": FindCode + "\n/*start writing code right here*/\n\n}", //double AddTwoNumbers (int a, int b){ return a + b; }
				"TOPICNAME":     topicName,
				"GIVENEXP":      crs.OUTS["task_exp"],
				"OUTVAL":        "Click on RUN button to see the output"})
	})

	// BOTH TO GENERATED CODE
	// IT IS THE INTERFACE TO SEE WHAT WAS GENERATED FROM PROMPT
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
	// IT IS TO SAVE THE LAST RESULTED GENERATION
	mod.POST("/cdgen", func(c *gin.Context) {

		//GETTING VALUES PARSED ONE AND AFTER
		fmt.Println("Current request keys: ", c.Keys)
		// var ProperCode string = ""
		// var FuncName string = ""
		// var ret_val_type string = ""
		// var func_params [][]string = [][]string{}

		for v := 0; v < len(c.Params); v++ {
			fmt.Println("Gotten value: ", c.Params[v].Key, "  ", c.Params[v].Key)
			if c.Params[v].Key == "FuncName" {

			}
		}

		// gn.BasePart()
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
