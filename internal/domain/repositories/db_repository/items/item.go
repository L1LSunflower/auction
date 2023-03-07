package items

import (
	"context"
	"fmt"
	"github.com/L1LSunflower/auction/internal/domain/entities"
	"github.com/L1LSunflower/auction/internal/tools/context_with_depends"
	"strings"
	"time"
)

const (
	fieldsSeparator = ","
)

type Repository struct{}

func (r *Repository) Create(ctx context.Context, item *entities.Item) error {
	tx, err := context_with_depends.TxFromContext(ctx)
	if err != nil {
		return err
	}

	fields := []string{
		"user_id",
		"name",
		"description",
		"created_at",
		"updated_at",
	}

	now := time.Now()
	item.CreatedAt = now
	item.UpdatedAt = now
	query := fmt.Sprintf("insert into items (%s) values (?, ?, ?, ?, ?)", strings.Join(fields, fieldsSeparator))
	tag, err := tx.Exec(query, item.UserID, item.Name, item.Description, item.CreatedAt, item.UpdatedAt)
	if err != nil {
		return err
	}

	id, err := tag.LastInsertId()
	if err != nil {
		return err
	}
	item.ID = int(id)

	return nil
}

func (r *Repository) Item(ctx context.Context, id int) (*entities.Item, error) {
	db, err := context_with_depends.GetDb(ctx)
	if err != nil {
		return nil, err
	}

	item := &entities.Item{}
	fields := []string{
		"id",
		"user_id",
		"name",
		"description",
		"created_at",
		"updated_at",
	}

	query := fmt.Sprintf("select %s from items where id=? and deleted_at is null", strings.Join(fields, fieldsSeparator))
	if err = db.QueryRow(query, id).Scan(&item.ID, &item.UserID, &item.Name, &item.Description, &item.CreatedAt, &item.UpdatedAt); err != nil {
		return nil, err
	}

	return item, nil
}

func (r *Repository) Update(ctx context.Context, item *entities.Item) error {
	tx, err := context_with_depends.TxFromContext(ctx)
	if err != nil {
		return err
	}

	fields := []string{
		"name=?",
		"description=?",
		"updated_at=now()",
	}

	query := fmt.Sprintf("update items set %s where id=?", strings.Join(fields, fieldsSeparator))
	if _, err = tx.Exec(
		query,
		item.Name,
		item.Description,
		item.ID,
	); err != nil {
		return err
	}

	return nil
}

func (r *Repository) Delete(ctx context.Context, id int) error {
	tx, err := context_with_depends.TxFromContext(ctx)
	if err != nil {
		return err
	}

	if _, err = tx.Exec("update items set deleted_at=now() where id=?", id); err != nil {
		return err
	}

	return nil
}
