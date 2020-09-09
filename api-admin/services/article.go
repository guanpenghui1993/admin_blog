package services

type ArticleService struct {
}

var ArticleService = newArticleService()

func newArticleService() *ArticleService {
	return &ArticleService{}
}
