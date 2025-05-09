package infrastructure

import (
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"xyz_golang/internal/config"
	"xyz_golang/internal/consumer"
	"xyz_golang/internal/domain"
	"xyz_golang/internal/limit"
	"xyz_golang/pkg/xlogger"
)

var (
	cfg config.Config

	consumerRepository domain.ConsumerRepository
	limitRepository    domain.LimitRepository

	consumerService domain.ConsumerService
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

	consumerRepository = consumer.NewPgsqlUserRepository(db)
	limitRepository = limit.NewPgsqlLimitRepository(db)

	consumerService = consumer.NewConsumerService(consumerRepository, limitRepository)
}
