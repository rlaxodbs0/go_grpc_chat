package repository

import (
	"sync"
)

type ArticleRepositoryImpl interface  {
	GetArticleSliceByUserID(uint64, int, int) ([]model.Article, int, error)
	CreateArticle(a *model.Article) error
	UpdateArticle(*model.Article) error
	DeleteArticle(*model.Article) error
}

var (
	once sync.Once
	articleRepository *ArticleRepositoryImpl
)


func InitArticleRepository(ar *ArticleRepositoryImpl) {
	once.Do(func() {
		articleRepository = ar
	})
}


func GetGrpcListener() *ArticleRepositoryImpl {
	return articleRepository
}
