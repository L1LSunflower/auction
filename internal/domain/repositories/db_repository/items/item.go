package items

import (
	"context"
	"fmt"
	"github.com/L1LSunflower/auction/internal/domain/entities"
	"github.com/L1LSunflower/auction/internal/tools/context_with_depends"
	"strings"
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
		"tag1",
		"tag2",
		"tag3",
		"tag4",
		"tag5",
		"tag6",
		"tag7",
		"tag8",
		"tag9",
		"tag10",
		"files",
		"description",
		"created_at",
		"updated_at",
	}

	query := fmt.Sprintf("insert into items (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, now(), now()) returnin id, created_at, updated_at", strings.Join(fields, fieldsSeparator))
	if err = tx.QueryRow(query).Scan(&item.ID, &item.CreatedAt, &item.UpdatedAt); err != nil {
		return err
	}

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
		"files",
		"description",
		"created_at",
		"updated_at",
	}

	query := fmt.Sprintf("select %s from items where id=? and deleted_at is null", strings.Join(fields, fieldsSeparator))
	if err = db.QueryRow(query, id).Scan(&item.ID, &item.UserID, &item.Name, &item.Images, &item.Description, &item.CreatedAt, &item.UpdatedAt); err != nil {
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
		"tag1=?",
		"tag2=?",
		"tag3=?",
		"tag4=?",
		"tag5=?",
		"tag6=?",
		"tag7=?",
		"tag8=?",
		"tag9=?",
		"tag10=?",
		"files=?",
		"description=?",
		"updated_at=now()",
	}

	query := fmt.Sprintf("update items set %s where id=?", strings.Join(fields, fieldsSeparator))
	if _, err = tx.Exec(
		query,
		item.Name,
		item.Tag1,
		item.Tag2,
		item.Tag3,
		item.Tag4,
		item.Tag5,
		item.Tag6,
		item.Tag7,
		item.Tag8,
		item.Tag9,
		item.Tag10,
		item.Images,
		item.Description,
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
