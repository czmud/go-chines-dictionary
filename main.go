package main

import (
	"github.com/czmud/go-chinese-dictionary/src/controllers"
	"github.com/czmud/go-chinese-dictionary/src/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.POST("/words/create", controllers.CreateNewWord)
	r.PUT("/words/update/:id", controllers.UpdateWordByID)
	r.GET("/words/get_all", controllers.GetAllWords)
	r.GET("/words/get/:id", controllers.GetWordByID)
	r.DELETE("/words/delete/:id", controllers.DeletePostByID)

	r.Run()
}
