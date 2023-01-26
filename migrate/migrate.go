package main

import (
	"github.com/czmud/go-chinese-dictionary/initializers"
	"github.com/czmud/go-chinese-dictionary/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Word{})
}
