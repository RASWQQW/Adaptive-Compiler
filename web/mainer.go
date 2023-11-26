package LevelFuncs

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Runnersss() {
	fmt.Println("Page is started")
	var mod = gin.Default()
	mod.GET("", Main)
	mod.GET("/form", Form)
	mod.GET("/cpprun", CppInter)
	mod.GET("/checker/:name/:val?", func(c *gin.Context) {
		c.JSON(200, gin.H{"c": c.Keys, "d": c.Params /*"pp": c.Request.URL*/})
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
