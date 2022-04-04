package internal

import (
	"context"
	"database/sql"
	"errors"
)

type Repository interface {
	GetAll(ctx context.Context) ([]Product, error)
	GetByID(ctx context.Context, id int) (Product, error)
	Save(ctx context.Context, p Product) (int, error)
	Update(ctx context.Context, p Product) error
	Delete(ctx context.Context, id int) error
	Exist(ctx context.Context, id int) (bool, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db}
}

func (r *repository) GetAll(ctx context.Context) ([]Product, error) {
	rows, err := r.db.Query(`SELECT id, name, description, price, stock, status FROM products`)
	if err != nil {
		return nil, err
	}

	var products []Product

	for rows.Next() {
		var p Product
		err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Stock, &p.Status)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}

func (r *repository) GetByID(ctx context.Context, id int) (Product, error) {
	row := r.db.QueryRow(`SELECT id, name, description, price, stock, status FROM products WHERE id = ?`, id)

	var p Product
	err := row.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Stock, &p.Status)
	if err != nil {
		return Product{}, err
	}

	return p, nil
}

func (r *repository) Save(ctx context.Context, p Product) (int, error) {
	stmt, err := r.db.Prepare(`INSERT INTO products (name, description, price, stock, status) VALUES (?, ?, ?, ?, ?) `)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(p.Name, p.Description, p.Price, p.Stock, p.Status)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *repository) Update(ctx context.Context, p Product) error {
	stmt, err := r.db.Prepare(`UPDATE products SET name = ?, description = ?, price = ?, stock = ?, status = ? WHERE id = ?`)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(p.Name, p.Description, p.Price, p.Stock, p.Status, p.ID)
	if err != nil {
		return err
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affect < 1 {
		return errors.New("Product not found")
	}

	return nil
}

func (r *repository) Delete(ctx context.Context, id int) error {
	stmt, err := r.db.Prepare(`DELETE FROM products WHERE id = ?`)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(id)
	if err != nil {
		return err
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affect < 1 {
		return errors.New("Product not found")
	}

	return nil
}

func (r *repository) Exist(ctx context.Context, id int) (bool, error) {
	row := r.db.QueryRow(`SELECT id FROM products WHERE id = ?`, id)
	err := row.Scan(&id)
	if err != nil {
		return false, err
	}
	return true, nil
}
