package infrastructure

import (
	"fmt"
	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/etag"
	recover2 "github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"time"
	"xyz_golang/internal/consumer"
	"xyz_golang/internal/middleware/limiter"
	"xyz_golang/internal/transaction"
	"xyz_golang/pkg/xlogger"
)

func Run() {
	logger := xlogger.Logger

	app := fiber.New(fiber.Config{
		ProxyHeader:           cfg.ProxyHeader,
		DisableStartupMessage: true,
		ErrorHandler:          defaultErrorHandler,
	})

	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: logger,
		Fields: cfg.LogFields,
	}))
	app.Use(recover2.New())
	app.Use(etag.New())
	app.Use(requestid.New())
	app.Use(limiter.RateLimiter())

	limiterRequest := limiter.NewIPRateLimiter()

	api := app.Group("/api")
	//docs.NewHttpHandler(api.Group("/docs"))
	consumerHandler := api.Group("/consumer")
	consumerHandler.Use(limiterRequest.EndpointLimiter(10, 1*time.Minute))
	consumer.NewHttpHandler(consumerHandler, consumerService)

	transactionHandler := api.Group("/transaction")
	transactionHandler.Use(limiterRequest.EndpointLimiter(5, 1*time.Minute))
	transaction.NewHttpHandler(transactionHandler, transactionService)

	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	logger.Info().Msgf("Server is running on address: %s", addr)
	if err := app.Listen(addr); err != nil {
		logger.Fatal().Err(err).Msg("Server failed to start")
	}
}
