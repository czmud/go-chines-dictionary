package controllers

import (
	"fmt"

	"github.com/czmud/go-chinese-dictionary/initializers"
	"github.com/czmud/go-chinese-dictionary/models"
	"github.com/gin-gonic/gin"
)

func CreateNewWord(c *gin.Context) {

	var reqBody struct {
		Definition     string
		Pinyin         string
		Unicode        string
		CharacterCount int
	}

	c.Bind(&reqBody)

	newWord := models.Word{Definition: reqBody.Definition, Pinyin: reqBody.Pinyin, Unicode: reqBody.Unicode, CharacterCount: reqBody.CharacterCount}

	result := initializers.DB.Create(&newWord)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"word": newWord,
	})
}

func GetAllWords(c *gin.Context) {
	var allWords []models.Word
	result := initializers.DB.Find(&allWords)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"words": allWords,
	})
}

func GetWordByID(c *gin.Context) {
	wordID := c.Param("id")

	var oneWord models.Word
	initializers.DB.First(&oneWord, wordID)

	c.JSON(200, gin.H{
		"words": oneWord,
	})
}

func UpdateWordByID(c *gin.Context) {
	wordID := c.Param("id")

	var reqBody struct {
		Definition     string
		Pinyin         string
		Unicode        string
		CharacterCount int
	}

	c.Bind(&reqBody)

	var oneWord models.Word
	initializers.DB.First(&oneWord, wordID)

	result := initializers.DB.Model(&oneWord).Updates(models.Word{
		Definition:     reqBody.Definition,
		Pinyin:         reqBody.Pinyin,
		Unicode:        reqBody.Unicode,
		CharacterCount: reqBody.CharacterCount,
	})

	fmt.Printf("object: %v", *result)

	c.JSON(200, gin.H{
		"word": oneWord,
	})
}

func DeletePostByID(c *gin.Context) {
	wordID := c.Param("id")

	result := initializers.DB.Delete(&models.Word{}, wordID)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.Status(200)
}
