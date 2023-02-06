package routers

import (
	"fmt"

	"sp-interview/internal/env"

	"github.com/gofiber/fiber/v2"

	"github.com/rs/zerolog"

	fiber_recover "github.com/gofiber/fiber/v2/middleware/recover"

	fiber_logger "github.com/gofiber/fiber/v2/middleware/logger"

	v1 "sp-interview/internal/fiber/routers/v1"

	util_error "sp-interview/internal/fiber/utils/errors"

	redis_handler "sp-interview/internal/redis"
)

// expose an interface with its functions
type Router interface {
	Init()
}

// wrap fiber into route struct and add necessary dependencies
type route struct {
	app *fiber.App

	redis redis_handler.RedisHandler

	logger *zerolog.Logger

	config *env.Config
}

// NewRouter create an structure that implemented `Router` interface
func NewRouter(r redis_handler.RedisHandler, l *zerolog.Logger, cfg *env.Config) Router {
	return &route{
		redis:  r,
		logger: l,
		config: cfg,
	}
}

// initSpaceHandler register all api's that related to order
func (r *route) initOrderHandler(router fiber.Router) {
	order_handler := v1.NewOrderHandler(r.redis, r.logger)

	router.Post("/", order_handler.AddOrder) // POST /api/order
}

func (r *route) Init() {
	cfg := fiber.Config{
		EnablePrintRoutes:     true,
		DisableStartupMessage: true,
		ErrorHandler:          util_error.ErrorHandler,
	}

	r.app = fiber.New(cfg)

	// register recover middleware for catching error
	r.app.Use(fiber_recover.New())

	// register logger middleware for logging requests
	r.app.Use(fiber_logger.New())

	// register all v1 api's
	v1Group := r.app.Group("/api/v1")
	r.initOrderHandler(v1Group.Group("/order"))

	// run fiber
	addr := fmt.Sprintf("0.0.0.0:%v", r.config.Port)
	r.logger.Info().Msgf("Running fiber with address: %v", addr)
	err := r.app.Listen(addr)

	if err != nil {
		r.logger.Error().Msgf(err.Error())
	}
}
