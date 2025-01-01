package api

import (
	"HalykProject/app/models"
	"HalykProject/app/service"
	"HalykProject/util/exception"
	"context"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type OrderController struct {
	service *service.OrderService
}

func NewUserHandler(service *service.OrderService) *OrderController {
	return &OrderController{service: service}
}

func (h *OrderController) CreateOrder(c *fiber.Ctx) error {
	var user models.Order

	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"CreateOrder: Ошибка при парсинге ордера \n": err.Error()})
	}

	if err := h.service.CreateOrder(context.Background(), user); err != nil {
		return exception.Handle(err, c)
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "User created"})
}

func (h *OrderController) GetOrderByID(c *fiber.Ctx) error {
	id := c.Params("id")

	order, err := h.service.GetOrderByID(context.Background(), id)
	if err != nil {
		return exception.Handle(err, c)
	}

	return c.Status(http.StatusOK).JSON(order)
}

func (h *OrderController) GetAllOrder(c *fiber.Ctx) error {
	orders, err := h.service.GetAllOrders(context.Background())
	if err != nil {
		return exception.Handle(err, c)
	}
	return c.Status(http.StatusOK).JSON(orders)
}

func (h *OrderController) GetOrderByUserId(c *fiber.Ctx) error {
	id := c.Params("userId")
	orders, err := h.service.GetOrdersByUserId(context.Background(), id)
	if err != nil {
		return exception.Handle(err, c)
	}
	return c.Status(http.StatusOK).JSON(orders)
}

func (h *OrderController) GetOrderByBoxId(c *fiber.Ctx) error {
	id := c.Params("boxId")
	order, err := h.service.GetOrderByBoxId(context.Background(), id)
	if err != nil {
		return exception.Handle(err, c)
	}
	return c.Status(http.StatusOK).JSON(order)
}

func (h *OrderController) GetOrderByStatus(c *fiber.Ctx) error {
	status := c.Params("status")
	orders, err := h.service.GetOrderByStatus(context.Background(), status)
	if err != nil {
		return exception.Handle(err, c)
	}
	return c.Status(http.StatusOK).JSON(orders)
}

func (h *OrderController) Complete(c *fiber.Ctx) error {
	id := c.Params("id")

	err := h.service.Complete(context.Background(), id)

	if err != nil {
		return exception.Handle(err, c)
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Success"})

}

func (h *OrderController) UpdateOrder(c *fiber.Ctx) error {
	var user models.Order
	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Не валидный json"})
	}

	if err := h.service.Update(context.Background(), user); err != nil {
		return exception.Handle(err, c)
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Success"})
}

func (h *OrderController) Extend(c *fiber.Ctx) error {
	type rp struct {
		RentalPeriod int `json:"rentalPeriod"`
	}
	var rentalP rp

	if err := c.BodyParser(&rentalP); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	id := c.Params("id")

	err := h.service.Extend(context.Background(), id, rentalP.RentalPeriod)

	if err != nil {
		return exception.Handle(err, c)
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Success"})
}
