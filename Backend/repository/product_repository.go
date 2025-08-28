package repository

import (
	"database/sql"
	"inventory_backend/model"
)

type ProductRepository interface {
	Create(product model.Product) (model.Product, error)
	FindAll(status string) ([]model.Product, error)
	FindByID(id int) (*model.Product, error)
	Update(product *model.Product) (model.Product, error)
	Delete(id int) (model.Product, error)
}

type productRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *productRepository {
	return &productRepository{db}
}
func (r *productRepository) Create(product model.Product) (model.Product, error) {
	result, err := r.db.Exec(
		"INSERT INTO Products (Product_name, SKU, Quantity, Status, Location) VALUES (?, ?, ?, ?, ?)",
		product.Product_name, product.SKU, product.Quantity, product.Status, product.Location)
	if err != nil {
		return model.Product{}, err
	}
	id, _ := result.LastInsertId()
	product.ID = int(id)
	return product, nil
}

func (r *productRepository) FindAll(status string) ([]model.Product, error) {
	query := "SELECT ID, Product_name, SKU, Quantity, Location, Status, Created_at, Updated_at FROM Products"
	args := []interface{}{}
	if status != "" {
		query += " WHERE Status = ?"
		args = append(args, status)
	}

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []model.Product
	for rows.Next() {
		var p model.Product
		if err := rows.Scan(
			&p.ID, &p.Product_name, &p.SKU, &p.Quantity, &p.Location,
			&p.Status, &p.Created_at, &p.Updated_at,
		); err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}

func (r *productRepository) FindByID(id int) (*model.Product, error) {
	row := r.db.QueryRow(
		"SELECT ID, Product_name, SKU, Quantity, Location, Status, Created_at, Updated_at FROM Products WHERE ID = ?", id)
	var p model.Product
	if err := row.Scan(
		&p.ID, &p.Product_name, &p.SKU, &p.Quantity, &p.Location,
		&p.Status, &p.Created_at, &p.Updated_at,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &p, nil
}

func (r *productRepository) Update(product *model.Product) (model.Product, error) {
	_, err := r.db.Exec(
		"UPDATE Products SET Product_name = ?, SKU = ?, Quantity = ?, Location = ?, Status = ?, Updated_at = CURRENT_TIMESTAMP WHERE ID = ?",
		product.Product_name, product.SKU, product.Quantity, product.Location, product.Status, product.ID)
	if err != nil {
		return model.Product{}, err
	}

	var updated model.Product
	err = r.db.QueryRow(
		"SELECT ID, Product_name, SKU, Quantity, Location, Status, Created_at, Updated_at FROM Products WHERE ID = ?",
		product.ID,
	).Scan(
		&updated.ID, &updated.Product_name, &updated.SKU, &updated.Quantity,
		&updated.Location, &updated.Status, &updated.Created_at, &updated.Updated_at,
	)
	if err != nil {
		return model.Product{}, err
	}

	return updated, nil
}

func (r *productRepository) Delete(id int) (model.Product, error) {
	_, err := r.db.Exec("DELETE FROM Products WHERE ID = ?", id)
	if err != nil {
		return model.Product{}, err
	}
	return model.Product{ID: id}, nil
}
