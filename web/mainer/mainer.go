package Mainer

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"

	inputing "ep/Execs"
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
