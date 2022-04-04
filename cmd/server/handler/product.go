package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/vincentconace/concurrencia-product/internal"
)

type Product struct {
	service internal.Service
}

func NewHandler(p internal.Service) *Product {
	return &Product{p}
}

func (p *Product) GetAll(c *fiber.Ctx) error {
	products, err := p.service.GetAll(c.Context())
	if err != nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": true,
			"msg":   "Products were not found",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":     false,
		"productos": products,
	})
}

func (p *Product) GetById(c *fiber.Ctx) error {
	id := c.Params("id")
	idInt, err := strconv.Atoi(id)
	products, err := p.service.GetByID(c.Context(), idInt)
	if err != nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": true,
			"msg":   "Product was not found",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":     false,
		"productos": products,
	})
}

func (p *Product) Create(c *fiber.Ctx) error {
	var r internal.Product
	if err := c.BodyParser(&r); err != nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": true,
			"msg":   "Not bodyParser",
		})
	}
	product, err := p.service.Create(c.Context(), r)
	if err != nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": true,
			"msg":   "Product was not created",
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"error":     false,
		"productos": product,
	})
}

func (p *Product) Update(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": true,
			"msg":   "Not parse Query",
		})
	}

	var r internal.Product
	if err := c.BodyParser(&r); err != nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": true,
			"msg":   "Not bodyParser",
		})
	}

	err = p.service.Update(c.Context(), r, id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "Product not exist",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"error": false,
		"msg":   "Product was updated",
	})
}

func (p *Product) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": true,
			"msg":   "Not bodyParser",
		})
	}
	err = p.service.Delete(c.Context(), idInt)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "Product not exist",
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"error": false,
		"msg":   "Product was deleted",
	})
}
