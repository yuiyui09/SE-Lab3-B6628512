package models

import "gorm.io/gorm"

type Reviewvote struct {
    gorm.Model
    Title   string `json:"title"`
    Content string `json:"content"`
    Rating  int    `json:"rating"`
}