package cities

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/L1LSunflower/auction/internal/domain/entities"
	"github.com/L1LSunflower/auction/internal/tools/context_with_tx"
	"github.com/L1LSunflower/auction/internal/tools/errorhandler"
	"strings"
)

const (
	fieldsSeparator = ","
)

type Repository struct{}

func (r *Repository) Create(ctx context.Context, name string) (*entities.City, error) {
	tx, err := context_with_tx.TxFromContext(ctx)
	if err != nil {
		return nil, err
	}

	city := &entities.City{}
	fields := []string{
		"name",
		"created_at",
		"updated_at",
	}

	query := fmt.Sprintf("insert into cities (%s) values (?, now(), now()) returning id, created_at, updated_at", strings.Join(fields, fieldsSeparator))
	if err = tx.QueryRow(query, name).Scan(&city.ID, &city.CreatedAt, &city.UpdatedAt); err != nil {
		return nil, err
	}
	city.Name = name

	return city, nil
}

func (r *Repository) Get(ctx context.Context, id int) (*entities.City, error) {
	tx, err := context_with_tx.TxFromContext(ctx)
	if err != nil {
		return nil, err
	}

	city := &entities.City{}
	nullDeleted := sql.NullTime{}
	fields := []string{
		"id",
		"name",
		"created_at",
		"updated_at",
		"deleted_at",
	}

	query := fmt.Sprintf("select %s from cities where id=? and deleted_at not null", strings.Join(fields, fieldsSeparator))
	rows, err := tx.Query(query, id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		if err = rows.Scan(&city.ID, &city.Name, &city.CreatedAt, &city.UpdatedAt, &nullDeleted); err != nil {
			return nil, err
		}
		city.DeletedAt = nullDeleted.Time
	}

	return city, nil
}

func (r *Repository) GetAll(ctx context.Context) ([]*entities.City, error) {
	tx, err := context_with_tx.TxFromContext(ctx)
	if err != nil {
		return nil, err
	}

	cities := []*entities.City{}
	fields := []string{
		"id",
		"name",
		"created_at",
		"updated_at",
		"deleted_at",
	}

	query := fmt.Sprintf("select %s from cities where deleted_at not null", strings.Join(fields, fieldsSeparator))
	rows, err := tx.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		city := &entities.City{}
		nullDeleted := sql.NullTime{}
		if err = rows.Scan(&city.ID, &city.Name, &city.CreatedAt, &city.UpdatedAt, &nullDeleted); err != nil {
			return nil, err
		}
		city.DeletedAt = nullDeleted.Time

		cities = append(cities, city)
	}

	return cities, nil
}

func (r *Repository) Update(ctx context.Context, id int, name string) (int, error) {
	tx, err := context_with_tx.TxFromContext(ctx)
	if err != nil {
		return 0, err
	}

	tag, err := tx.Exec("update cities set name=?, created_at=now(), updated_at=now() where id=?", name, id)
	if err != nil {
		return 0, err
	}

	lid, err := tag.LastInsertId()
	if err != nil {
		return 0, nil
	}

	return int(lid), nil
}

func (r *Repository) Delete(ctx context.Context, id int) error {
	tx, err := context_with_tx.TxFromContext(ctx)
	if err != nil {
		return err
	}

	tag, err := tx.Exec("delete from cities where id=?", id)
	if err != nil {
		return err
	}

	if rowsAffected, err := tag.RowsAffected(); err != nil || rowsAffected <= 0 {
		return errorhandler.ErrDeleteCity
	}

	return nil
}
