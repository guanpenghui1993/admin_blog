package services

type ArticleService struct {
}

var ArticleService = newArticleService()

func newArticleService() *ArticleService {
	return &ArticleService{}
}

func (a *ArticleService) GetArticleList() []models.Article {
	
}