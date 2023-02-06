package v1

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"

	"github.com/rs/zerolog"

	redis_handler "sp-interview/internal/redis"

	repositories "sp-interview/internal/fiber/repositories"

	order_dto "sp-interview/internal/fiber/models/dtos/order"

	util_error "sp-interview/internal/fiber/utils/errors"

	route_validator "sp-interview/internal/fiber/utils/validator"
)

type OrderHandler interface {
	AddOrder(ctx *fiber.Ctx) error
}

type orderHandler struct {
	repo   repositories.OrderRepository
	redis  redis_handler.RedisHandler
	logger *zerolog.Logger
}

func NewOrderHandler(redis redis_handler.RedisHandler, l *zerolog.Logger) OrderHandler {
	return &orderHandler{
		repo:   repositories.NewOrderRepository(redis),
		redis:  redis,
		logger: l,
	}
}

func (o *orderHandler) AddOrder(ctx *fiber.Ctx) error {
	o_dto := new(order_dto.AddOrderDto)

	// map body to order_dto instance
	if err := ctx.BodyParser(o_dto); err != nil {
		o.logger.Warn().Msgf("unable to map body to add order dto:%s", err.Error())
		return util_error.NewBadRequestError(err.Error())
	}

	// validate input body data
	errors := route_validator.Validate(o_dto)

	if errors != nil {
		return util_error.NewBadRequestError(errors)
	}

	order := order_dto.ConvertToModel(o_dto)
	message, err := json.Marshal(order)

	if err != nil {
		o.logger.Warn().Msgf("an error occured in marshaling json into string: %s", err.Error())
		return util_error.NewInternalServerError(err.Error())
	}

	// push message to queue
	err = o.redis.Push(string(message), ctx.Context())

	if err != nil {
		o.logger.Warn().Msgf("error in sending message: %s", err.Error())
		return util_error.NewInternalServerError(err.Error())
	}

	return ctx.Status(200).JSON(
		map[string]any{
			"message": "ok",
		},
	)

}
