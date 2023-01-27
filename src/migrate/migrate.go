package main

import (
	"github.com/czmud/go-chinese-dictionary/src/initializers"
	"github.com/czmud/go-chinese-dictionary/src/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Word{})
}
