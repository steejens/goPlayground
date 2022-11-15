package tronics

import (
	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"strconv"
)

type ProductValidator struct {
	validator *validator.Validate
}

func (p *ProductValidator) Validate(i interface{}) error {
	return p.validator.Struct(i)
}

var products = []map[int]string{{1: "Phones"}, {2: "Tablets"}}

func getProducts(c echo.Context) error {
	return c.JSON(http.StatusOK, products)
}

func getProduct(c echo.Context) error {

	var product map[int]string
	pId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	for i, p := range products {
		if pId == i+1 {
			product = p
		}
	}
	if product == nil {
		return c.JSON(http.StatusNotFound, "product not found")
	}
	return c.JSON(http.StatusOK, product)
}
func deleteProduct(c echo.Context) error {

	var product map[int]string
	pId, err := strconv.Atoi(c.Param("id"))
	var arrIndex int
	if err != nil {
		return err
	}
	for i, p := range products {
		if pId == i+1 {
			product = p
			arrIndex = i
		}
	}
	if product == nil {
		return c.JSON(http.StatusNotFound, "product not found")
	}
	splice := func(s []map[int]string, index int) []map[int]string {
		return append(s[:index], s[index+1:]...)
	}
	products = splice(products, arrIndex)
	return c.JSON(http.StatusOK, product)
}
func updateProduct(c echo.Context) error {

	var product map[int]string
	pId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	for i, p := range products {
		if pId == i+1 {
			product = p
		}
	}
	if product == nil {
		return c.JSON(http.StatusNotFound, "product not found")
	}
	type body struct {
		Name string `json:"product_name" validate:"required,min=4"`
	}
	var reqBody body
	e.Validator = &ProductValidator{validator: v}
	if err := c.Bind(&reqBody); err != nil {
		return err
	}
	if err := c.Validate(reqBody); err != nil {
		return err
	}
	product[pId] = reqBody.Name
	return c.JSON(http.StatusOK, product)
}
func createProduct(c echo.Context) error {
	type body struct {
		Name string `json:"product_name" validate:"required,min=4"`
	}
	var reqBody body
	e.Validator = &ProductValidator{validator: v}
	if err := c.Bind(&reqBody); err != nil {
		return err
	}
	if err := c.Validate(reqBody); err != nil {
		return err
	}
	product := map[int]string{
		len(products) + 1: reqBody.Name,
	}
	products = append(products, product)
	return c.JSON(http.StatusOK, product)
}
