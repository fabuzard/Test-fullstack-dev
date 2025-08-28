package handler

import (
	"encoding/csv"
	"fmt"
	"inventory_backend/dto"
	"inventory_backend/model"
	"inventory_backend/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	Service service.ProductService
}

func NewProductHandler(service service.ProductService) *ProductHandler {
	return &ProductHandler{Service: service}
}

// Create product
func (h *ProductHandler) Create(c echo.Context) error {
	var productRequest dto.CreateProductRequest
	if err := c.Bind(&productRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	if err := c.Validate(&productRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	// Check if status is valid
	if productRequest.Status != "in_stock" && productRequest.Status != "out_of_stock" && productRequest.Status != "low_stock" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid status value"})
	}

	// Set status based on qty
	if productRequest.Quantity <= 0 {
		productRequest.Status = "out_of_stock"
	} else if productRequest.Quantity < 5 {
		productRequest.Status = "low_stock"
	} else {
		productRequest.Status = "in_stock"
	}

	product := model.Product{
		Product_name: productRequest.Product_name,
		SKU:          productRequest.SKU,
		Quantity:     productRequest.Quantity,
		Status:       productRequest.Status,
		Location:     productRequest.Location,
	}
	// Api call to Create
	createdProduct, err := h.Service.Create(product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create product"})
	}
	response := dto.ProductResponse{
		ID:           createdProduct.ID,
		Product_name: createdProduct.Product_name,
		SKU:          createdProduct.SKU,
		Quantity:     createdProduct.Quantity,
		Status:       createdProduct.Status,
		Location:     createdProduct.Location,
	}
	return c.JSON(http.StatusCreated, response)

}

// Find all product
func (h *ProductHandler) FindAll(c echo.Context) error {
	// Filtering by status
	status := c.QueryParam("status")
	// Api call to FindAll
	products, err := h.Service.FindAll(status)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch products"})
	}
	var response []dto.ProductResponse
	for _, p := range products {
		response = append(response, dto.ProductResponse{
			ID:           p.ID,
			Product_name: p.Product_name,
			SKU:          p.SKU,
			Quantity:     p.Quantity,
			Status:       p.Status,
			Location:     p.Location,
		})
	}
	return c.JSON(http.StatusOK, response)
}

func (h *ProductHandler) FindByID(c echo.Context) error {
	idParam := c.Param("id")
	var id int
	_, err := fmt.Sscanf(idParam, "%d", &id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid product ID"})
	}
	// Api call to FindByID
	product, err := h.Service.FindByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch product"})
	}
	if product == nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Product not found"})
	}
	response := dto.ProductResponse{
		ID:           product.ID,
		Product_name: product.Product_name,
		SKU:          product.SKU,
		Quantity:     product.Quantity,
		Status:       product.Status,
		Location:     product.Location,
	}
	return c.JSON(http.StatusOK, response)
}

func (h *ProductHandler) Update(c echo.Context) error {
	idParam := c.Param("id")
	var id int
	_, err := fmt.Sscanf(idParam, "%d", &id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid product ID"})
	}
	var productRequest dto.UpdateProductRequest
	if err := c.Bind(&productRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	// Input validation
	if err := c.Validate(&productRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	existingProduct, err := h.Service.FindByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch product"})
	}
	if existingProduct == nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Product not found"})
	}
	existingProduct.Product_name = productRequest.Product_name
	existingProduct.SKU = productRequest.SKU
	existingProduct.Quantity = productRequest.Quantity
	existingProduct.Status = productRequest.Status
	existingProduct.Location = productRequest.Location

	// Set status based on qty
	if productRequest.Quantity <= 0 {
		existingProduct.Status = "out_of_stock"
	} else if productRequest.Quantity < 5 {
		existingProduct.Status = "low_stock"
	} else {
		existingProduct.Status = "in_stock"
	}
	// Api call to Update
	updatedProduct, err := h.Service.Update(existingProduct)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update product"})
	}
	return c.JSON(http.StatusOK, updatedProduct)
}

func (h *ProductHandler) Delete(c echo.Context) error {
	idParam := c.Param("id")
	var id int
	_, err := fmt.Sscanf(idParam, "%d", &id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid product ID"})
	}
	// Check if product exists
	existingProduct, err := h.Service.FindByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch product"})
	}
	if existingProduct == nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Product not found"})
	}
	// Api call to Delete
	deletedProduct, err := h.Service.Delete(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete product"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": fmt.Sprintf("Product with ID %d deleted successfully", deletedProduct.ID)})
}

func (h *ProductHandler) ExportCSV(c echo.Context) error {
	products, err := h.Service.FindAll("")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch products"})
	}

	// Set CSV headers
	c.Response().Header().Set("Content-Type", "text/csv")
	c.Response().Header().Set("Content-Disposition", `attachment; filename="products.csv"`)

	writer := csv.NewWriter(c.Response())
	defer writer.Flush()

	// Write header row
	writer.Write([]string{"ID", "Product Name", "SKU", "Quantity", "Status", "Location"})

	// Write data rows
	for _, p := range products {
		row := []string{
			strconv.Itoa(p.ID),
			p.Product_name,
			p.SKU,
			strconv.Itoa(p.Quantity),
			p.Status,
			p.Location,
		}
		writer.Write(row)
	}

	return nil
}
