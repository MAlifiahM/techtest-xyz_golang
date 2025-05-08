package infrastructure

import (
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"xyz_golang/internal/config"
	"xyz_golang/pkg/xlogger"
)

var (
	cfg config.Config

	//authorRepository  domain.AuthorRepository
	//articleRepository domain.ArticleRepository
	//
	//articleService domain.ArticleService
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	if err := env.Parse(&cfg); err != nil {
		panic(err)
	}
	xlogger.Setup(cfg)
	dbSetup()

	//authorRepository = author.NewMysqlAuthorRepository(db)
	//articleRepository = article.NewMysqlArticleRepository(db)
	//
	//articleService = article.NewArticleService(articleRepository, authorRepository)
}
