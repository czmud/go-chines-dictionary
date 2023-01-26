package models

import "gorm.io/gorm"

type Word struct {
	gorm.Model
	Definition     string
	Pinyin         string
	Unicode        string
	CharacterCount int
}
