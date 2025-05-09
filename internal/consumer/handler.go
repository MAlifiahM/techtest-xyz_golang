package consumer

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"net/http"
	"time"
	"xyz_golang/internal/domain"
	"xyz_golang/internal/middleware/validation"
	"xyz_golang/internal/utilities"
)

type HttpHandler struct {
	consumerSvc domain.ConsumerService
}

func NewHttpHandler(r fiber.Router, consumerSvc domain.ConsumerService) {
	handler := &HttpHandler{
		consumerSvc: consumerSvc,
	}

	r.Post("/", validation.New[domain.ConsumerRequest](), handler.Store)
	r.Get("/:id", handler.GetByID)
	r.Get("/:id/limit", handler.GetLimit)
}

func (h *HttpHandler) GetByID(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.Response{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	consumer, err := h.consumerSvc.GetByID(id)
	if err != nil {
		status := fiber.StatusInternalServerError
		if err == fiber.ErrNotFound {
			status = fiber.StatusNotFound
		}
		c.Status(status)
		return c.JSON(domain.Response{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.Status(http.StatusOK).JSON(domain.Response{
		Status:  true,
		Message: "success",
		Data:    consumer,
	})
}

func (h *HttpHandler) GetLimit(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.Response{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	limit, err := h.consumerSvc.GetLimit(id)
	if err != nil {
		status := fiber.StatusInternalServerError
		if err == fiber.ErrNotFound {
			status = fiber.StatusNotFound
		}
		c.Status(status)
		return c.JSON(domain.Response{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.Status(http.StatusOK).JSON(domain.Response{
		Status:  true,
		Message: "success",
		Data:    limit,
	})
}

func (h *HttpHandler) Store(c *fiber.Ctx) error {
	consumerReq := utilities.ExtractStructFromValidator[domain.ConsumerRequest](c)

	dateOfBirth, err := time.Parse("2006-01-02", consumerReq.DateOfBirth)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.Response{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	consumer := &domain.Consumer{
		NIK:          consumerReq.NIK,
		FullName:     consumerReq.FullName,
		LegalName:    consumerReq.LegalName,
		PlaceOfBirth: consumerReq.PlaceOfBirth,
		DateOfBirth:  dateOfBirth,
		Salary:       consumerReq.Salary,
		PhotoKTP:     consumerReq.PhotoKTP,
		PhotoSelfie:  consumerReq.PhotoSelfie,
	}

	if err := h.consumerSvc.Store(consumer); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(domain.Response{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.Status(http.StatusCreated).JSON(domain.Response{
		Status:  true,
		Message: "success",
		Data:    consumer,
	})
}
