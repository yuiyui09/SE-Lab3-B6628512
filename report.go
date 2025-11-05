package models

import "gorm.io/gorm"

type RevieReport struct {
    gorm.Model
    Title   string `json:"title"`
    Content string `json:"content"`
    Rating  int    `json:"rating"`
}