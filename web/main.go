package LevelFuncs

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Runner() {
	fmt.Println("Page is started")
	var mod = gin.Default()
	mod.GET("", Main)

	mod.Run("127.0.0.1:8080")
}

func Main(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"dd": "values"})
}
