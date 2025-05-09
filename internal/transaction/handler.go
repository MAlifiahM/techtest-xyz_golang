package transaction

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"net/http"
	"xyz_golang/internal/domain"
	"xyz_golang/internal/middleware/validation"
	"xyz_golang/internal/utilities"
)

type HttpHandler struct {
	transactionSvc domain.TransactionService
}

func NewHttpHandler(r fiber.Router, transactionSvc domain.TransactionService) {
	handler := &HttpHandler{
		transactionSvc: transactionSvc,
	}

	r.Post("/", validation.New[domain.TransactionRequest](), handler.Store)
	r.Get("/:id", handler.GetByID)
}

func (h *HttpHandler) Store(c *fiber.Ctx) error {
	transactionReq := utilities.ExtractStructFromValidator[domain.TransactionRequest](c)

	transaction := &domain.Transaction{
		ConsumerID:     transactionReq.ConsumerID,
		ContractNumber: transactionReq.ContractNumber,
		Tenor:          transactionReq.Tenor,
		OTR:            transactionReq.OTR,
		AdminFee:       transactionReq.AdminFee,
		Installment:    transactionReq.Installment,
		Interest:       transactionReq.Interest,
		AssetName:      transactionReq.AssetName,
	}

	if err := h.transactionSvc.Store(transaction); err != nil {
		status := fiber.StatusInternalServerError
		if err == fiber.ErrNotFound {
			status = fiber.StatusNotFound
		}
		return c.Status(status).JSON(domain.Response{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.Status(http.StatusCreated).JSON(domain.Response{
		Status:  true,
		Message: "success",
		Data:    transaction,
	})
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

	transaction, err := h.transactionSvc.GetByID(id)
	if err != nil {
		status := fiber.StatusInternalServerError
		if err == fiber.ErrNotFound {
			status = fiber.StatusNotFound
		}
		return c.Status(status).JSON(domain.Response{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.Status(http.StatusOK).JSON(domain.Response{
		Status:  true,
		Message: "success",
		Data:    transaction,
	})
}
