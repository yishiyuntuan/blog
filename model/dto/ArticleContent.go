package dto

import "blog/model/entity"

type ArticleContent struct {
	entity.Article
	Content string
}
