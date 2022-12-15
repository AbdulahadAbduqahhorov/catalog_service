package postgres

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/AbdulahadAbduqahhorov/gRPC/Ecommerce/catalog_service/genproto/product_service"
	"github.com/AbdulahadAbduqahhorov/gRPC/Ecommerce/catalog_service/storage/repo"
	"github.com/jmoiron/sqlx"
)

type productRepo struct {
	db *sqlx.DB
}

func NewProductRepo(db *sqlx.DB) repo.ProductRepoI {
	return productRepo{
		db: db,
	}
}

func (p productRepo) CreateProduct(id string, req *product_service.CreateProductRequest) error {
	query := `INSERT INTO product (
		id,
		title,
		description,
		quantity,
		price,
		category_id
	) 
	VALUES ($1, $2, $3,$4,$5,$6) `
	_, err := p.db.Exec(query, id, req.Title, req.Desc, req.Quantity, req.Price, req.CategoryId)
	
	if err != nil {
		return err
	}
	return nil
}
func (p productRepo) GetProductList(req *product_service.GetProductListRequest) (*product_service.GetProductListResponse, error) {
	res := &product_service.GetProductListResponse{
		Products: make([]*product_service.Product, 0),
	}
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1
	if len(req.CategoryId) > 0 {
		setValues = append(setValues, fmt.Sprintf("AND category_id=$%d", argId))
		args = append(args, req.CategoryId)
		argId++
	}
	if req.Search != "" {
		setValues = append(setValues, fmt.Sprintf("AND title ILIKE '%%' || $%d || '%%'", argId))
		args = append(args, req.Search)
		argId++
	}

	countQuery := `SELECT count(1) FROM product  WHERE true ` + strings.Join(setValues, " ")
	err := p.db.QueryRow(countQuery, args...).Scan(
		&res.Count,
	)
	if err != nil {
		return nil, err
	}
	if req.OrderBy=="Price High to Low"  {
		setValues = append(setValues, "order by price desc ")
	}else{
		setValues = append(setValues, "order by price asc ")
	}
	if req.Limit > 0 {
		setValues = append(setValues, fmt.Sprintf("limit $%d ", argId))
		args = append(args, req.Limit)
		argId++
	}
	if req.Offset >= 0 {
		setValues = append(setValues, fmt.Sprintf("offset $%d ", argId))
		args = append(args, req.Offset)
		argId++
	}
	
	s := strings.Join(setValues, " ")
	query := `SELECT
	id,
	title,
	description,
	quantity,
	price,
	category_id,
	created_at,
	updated_at
	FROM product WHERE true ` + s
	
	rows, err := p.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		obj := &product_service.Product{}
		var updatedAt sql.NullString

		err = rows.Scan(
			&obj.Id,
			&obj.Title,
			&obj.Desc,
			&obj.Quantity,
			&obj.Price,
			&obj.CategoryId,
			&obj.CreatedAt,
			&updatedAt,
		)

		if err != nil {
			return res, err
		}
		if updatedAt.Valid {
			obj.UpdatedAt = updatedAt.String
		}

		res.Products = append(res.Products, obj)
	}

	return res, nil

}
func (p productRepo) GetProductById(id string) (*product_service.GetProductByIdResponse, error) {
	res := &product_service.GetProductByIdResponse{
		Category: &product_service.GetProductByIdResponse_Category{},
	}

	var updatedAt, categoryUpdatedAt sql.NullString
	query := `
	SELECT 
		p.id,
		p.title,
		p.description,
		p.quantity,
		p.price,
		p.created_at,
		p.updated_at,
		c.id,
		c.title,
		c.created_at,
		c.updated_at
	FROM product p JOIN category c ON p.category_id=c.id WHERE p.id=$1 `

	err := p.db.QueryRow(query, id).Scan(
		&res.Id,
		&res.Title,
		&res.Desc,
		&res.Quantity,
		&res.Price,
		&res.CreatedAt,
		&updatedAt,
		&res.Category.Id,
		&res.Category.Title,
		&res.Category.CreatedAt,
		&categoryUpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	if updatedAt.Valid {
		res.UpdatedAt = updatedAt.String
	}

	if categoryUpdatedAt.Valid {
		res.Category.UpdatedAt = categoryUpdatedAt.String
	}
	return res, nil
}

func (p productRepo) UpdateProduct(req *product_service.UpdateProductRequest) (int64, error) {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1
	if len(strings.Trim(req.Title," "))>0 {
		setValues = append(setValues, fmt.Sprintf("title=$%d ", argId))
		args = append(args, req.Title)
		argId++
	}
	if len(strings.Trim(req.Desc," "))>0 {
		setValues = append(setValues, fmt.Sprintf("description=$%d ", argId))
		args = append(args, req.Desc)
		argId++
	}
	if req.Quantity>=0 {
		setValues = append(setValues, fmt.Sprintf("quantity=$%d ", argId))
		args = append(args, req.Quantity)
		argId++
	}
	if req.Price>0 {
		setValues = append(setValues, fmt.Sprintf("price=$%d ", argId))
		args = append(args, req.Price)
		argId++
	}

	s := strings.Join(setValues, ",")
	query := fmt.Sprintf(`
			UPDATE product
			SET %s ,updated_at = now()
			WHERE id = $%d`,
			s,argId)
			
	args = append(args,req.Id)

	result, err := p.db.Exec(query, args...)
	if err != nil {
		return 0,err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsAffected, nil
}
func (p productRepo) DeleteProduct(id string) (int64, error) {
	query := `DELETE FROM product WHERE id = $1`

	result, err := p.db.Exec(query, id)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil

}
