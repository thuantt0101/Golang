package helper

import (
	"github.com/Golang/Proj/week2-exercise2/crawler"
	"github.com/Golang/Proj/week2-exercise2/model"
)

func FillDataToArticle(article *model.Article, data crawler.Data) {
	article.Title = data.Title
	article.PublishedAt = data.PublishedDate
	article.Content = data.Content
	article.Author = data.Author
}
