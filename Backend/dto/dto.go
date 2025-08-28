package dto

type CreateProductRequest struct {
	Product_name string `json:"product_name" validate:"required,min=3"`
	SKU          string `json:"sku" validate:"required,min=3"`
	Quantity     int    `json:"quantity" validate:"required,gte=0"`
	Location     string `json:"location" validate:"required"`
	Status       string `json:"status" validate:"omitempty,oneof=in_stock low_stock out_of_stock"`
}
type UpdateProductRequest struct {
	Product_name string `json:"product_name" validate:"required,min=3"`
	SKU          string `json:"sku" validate:"required,min=3"`
	Quantity     int    `json:"quantity" validate:"required,gte=0"`
	Location     string `json:"location" validate:"required"`
	Status       string `json:"status" validate:"omitempty,oneof=in_stock low_stock out_of_stock"`
}

type ProductResponse struct {
	ID           int    `json:"id"`
	Product_name string `json:"product_name"`
	SKU          string `json:"sku"`
	Quantity     int    `json:"quantity"`
	Status       string `json:"status"`
	Location     string `json:"location"`
}

type RegisterRequest struct {
	Fullname string `json:"fullname" validate:"required,min=3"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}
type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type RegisterResponse struct {
	Message string `json:"message"`
	User    struct {
		ID       int    `json:"id"`
		Fullname string `json:"fullname"`
		Email    string `json:"email"`
	} `json:"user"`
}

type AuthResponse struct {
	Token    string `json:"token"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	ID       int    `json:"id"`
}
