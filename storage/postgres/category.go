package postgres

import (
	"database/sql"
	"errors"

	"github.com/AbdulahadAbduqahhorov/gRPC/Ecommerce/ecommerce_service/genproto/category_service"
	"github.com/AbdulahadAbduqahhorov/gRPC/Ecommerce/ecommerce_service/storage/repo"
	"github.com/jmoiron/sqlx"
)

type categoryRepo struct {
	db *sqlx.DB
}

func NewCategoryRepo(db *sqlx.DB) repo.CategoryRepoI {
	return categoryRepo{
		db: db,
	}
}

// CreateCategory ...
func (c categoryRepo) CreateCategory(id string, req *category_service.CreateCategoryRequest) error {

	query := `INSERT INTO category (
		id,
		title
	) 
	VALUES ($1, $2)`

	_, err := c.db.Exec(query, id, req.Title)

	return err
}

// GetCategoryList ...
func (c categoryRepo) GetCategoryList(offset, limit int, search string) (*category_service.GetCategoryListResponse, error) {
	res := &category_service.GetCategoryListResponse{
		Categories: make([]*category_service.Category, 0),
	}

	rows, err := c.db.Queryx(`SELECT
	id,
	title,
	created_at,
	updated_at
	FROM category WHERE title ILIKE '%' || $1 || '%'
	LIMIT $2
	OFFSET $3
	`, search, limit, offset)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var updatedAt sql.NullString
		obj := &category_service.Category{}

		err := rows.Scan(
			&obj.Id,
			&obj.Title,
			&obj.CreatedAt,
			&updatedAt,
		)
		if err != nil {
			return res, err
		}

		if updatedAt.Valid {
			obj.UpdatedAt = updatedAt.String
		}

		res.Categories = append(res.Categories, obj)
	}

	return res, err

}

// GetCategoryByID ...
func (c categoryRepo) GetCategoryByID(id string) (*category_service.Category, error) {
	res := &category_service.Category{}

	var updatedAt sql.NullString
	query := `
	SELECT 
		c.id,
		c.title,
		c.created_at,
		c.updated_at,
	FROM category c  WHERE c.id=$1 `
	err := c.db.QueryRow(query, id).Scan(
		&res.Id,
		&res.Title,
		&res.CreatedAt,
		&updatedAt,
	)

	if err != nil {
		return nil, err
	}

	if updatedAt.Valid {
		res.UpdatedAt = updatedAt.String
	}

	return res, nil

}

// UpdateCategory ...
func (c categoryRepo) UpdateCategory(req *category_service.UpdateCategoryRequest) error {

	res, err := c.db.NamedExec("UPDATE category  SET title=:t, updated_at=now() WHERE  id=:id", map[string]interface{}{
		"id": req.Id,
		"t":  req.Title,
	})
	if err != nil {
		return err
	}

	n, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if n > 0 {
		return nil
	}

	return errors.New("category not found")
}

// DeleteCategory ...
func (c categoryRepo) DeleteCategory(id string) error {
	res, err := c.db.Exec("DELETE  FROM category WHERE id=$1", id)
	if err != nil {
		return err
	}

	n, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if n > 0 {
		return nil
	}

	return errors.New("category not found")
}
