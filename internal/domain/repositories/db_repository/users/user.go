package users

import (
	"context"
	"fmt"
	"github.com/L1LSunflower/auction/internal/domain/entities"
	"github.com/L1LSunflower/auction/internal/tools/context_with_tx"
	"strings"
)

const (
	fieldsSeparator = ","
)

type Repository struct{}

func (r *Repository) Create(ctx context.Context, user *entities.User) error {
	tx, err := context_with_tx.TxFromContext(ctx)
	if err != nil {
		return err
	}

	fields := []string{
		"id",
		"password",
		"email",
		"first_name",
		"last_name",
		"phone",
		"city",
		"is_active",
		"created_at",
		"updated_at",
	}

	query := fmt.Sprintf("insert into users (%s) values (?, ?, ?, ?, ?, ?, ?, 1, now(), now()) returning id, created_at, updated_at", strings.Join(fields, fieldsSeparator))
	if err = tx.QueryRow(query, user.ID, user.Password, user.Email, user.FirstName, user.LastName, user.Phone, user.City).Scan(&user.ID, &user.CreateAt, &user.UpdatedAt); err != nil {
		return err
	}

	return nil
}

func (r *Repository) User(ctx context.Context, id string) (*entities.User, error) {
	tx, err := context_with_tx.TxFromContext(ctx)
	if err != nil {
		return nil, err
	}

	user := &entities.User{}
	fields := []string{
		"id",
		"email",
		"first_name",
		"last_name",
		"phone",
		"city",
		"is_active",
		"created_at",
		"updated_at",
	}

	query := fmt.Sprintf("select %s from users where id=? and is_active=1 and deleted_at not null", strings.Join(fields, fieldsSeparator))
	rows, err := tx.Query(query, id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		if err = rows.Scan(&user.ID, &user.Email, &user.FirstName, &user.LastName, &user.Phone, &user.City, &user.IsActive, &user.CreateAt, &user.UpdatedAt); err != nil {
			return nil, err
		}
	}

	return user, nil
}

func (r *Repository) UserByPhone(ctx context.Context, phone string) (*entities.User, error) {
	tx, err := context_with_tx.TxFromContext(ctx)
	if err != nil {
		return nil, err
	}

	user := &entities.User{}
	fields := []string{
		"id",
		"email",
		"first_name",
		"last_name",
		"phone",
		"city",
		"is_active",
		"created_at",
		"updated_at",
	}

	query := fmt.Sprintf("select %s from users where phone=? and is_active=1 and deleted_at is null", strings.Join(fields, fieldsSeparator))
	rows, err := tx.Query(query, phone)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		if err = rows.Scan(&user.ID, &user.Email, &user.FirstName, &user.LastName, &user.Phone, &user.City, &user.IsActive, &user.CreateAt, &user.UpdatedAt); err != nil {
			return nil, err
		}
	}

	return user, nil
}
