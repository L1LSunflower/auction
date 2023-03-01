package users

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

func (r *Repository) Create(ctx context.Context, user *entities.User) error {
	tx, err := context_with_depends.TxFromContext(ctx)
	if err != nil {
		return err
	}

	fields := []string{
		"id",
		"phone",
		"email",
		"password",
		"first_name",
		"last_name",
		"city",
		"created_at",
		"updated_at",
	}

	query := fmt.Sprintf("insert into users (%s) values (?, ?, ?, ?, ?, ?, ?, now(), now()) returning id, created_at, updated_at", strings.Join(fields, fieldsSeparator))
	if err = tx.QueryRow(query, user.ID, user.Phone, user.Email, user.Password, user.FirstName, user.LastName, user.City).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt); err != nil {
		return err
	}

	return nil
}

func (r *Repository) User(ctx context.Context, id string) (*entities.User, error) {
	tx, err := context_with_depends.GetDb(ctx)
	if err != nil {
		return nil, err
	}

	user := &entities.User{}
	fields := []string{
		"id",
		"phone",
		"email",
		"password",
		"first_name",
		"last_name",
		"city",
		"created_at",
		"updated_at",
	}

	query := fmt.Sprintf("select %s from users where id=?", strings.Join(fields, fieldsSeparator))
	rows, err := tx.Query(query, id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		if err = rows.Scan(&user.ID, &user.Phone, &user.Email, &user.Password, &user.FirstName, &user.LastName, &user.City, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, err
		}
	}

	return user, nil
}

func (r *Repository) UserByPhone(ctx context.Context, phone string) (*entities.User, error) {
	db, err := context_with_depends.GetDb(ctx)
	if err != nil {
		return nil, err
	}

	user := &entities.User{}
	fields := []string{
		"id",
		"phone",
		"email",
		"password",
		"first_name",
		"last_name",
		"city",
		"created_at",
		"updated_at",
	}

	query := fmt.Sprintf("select %s from users where phone=?", strings.Join(fields, fieldsSeparator))
	rows, err := db.Query(query, phone)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		if err = rows.Scan(&user.ID, &user.Phone, &user.Email, &user.Password, &user.FirstName, &user.LastName, &user.City, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, err
		}
	}

	return user, nil
}

func (r *Repository) Update(ctx context.Context, user *entities.User) error {
	tx, err := context_with_depends.TxFromContext(ctx)
	if err != nil {
		return err
	}

	fields := []string{
		"first_name=?",
		"last_name=?",
		"email=?",
		"phone=?",
		"password=?",
		"updated_at=now()",
	}

	query := fmt.Sprintf("update users set %s where id=?", strings.Join(fields, fieldsSeparator))
	if _, err = tx.Exec(query, user.ID); err != nil {
		return err
	}

	if err = tx.QueryRow("select updated_at from users where id=?", user.ID).Scan(&user.UpdatedAt); err != nil {
		return err
	}

	return nil
}

func (r *Repository) Delete(ctx context.Context, id string) error {
	tx, err := context_with_depends.TxFromContext(ctx)
	if err != nil {
		return err
	}

	if _, err = tx.Exec("update users set deleted_at=now() where id=?", id); err != nil {
		return err
	}

	return nil
}

func (r *Repository) UpdatePassword(ctx context.Context, user *entities.User) error {
	tx, err := context_with_depends.TxFromContext(ctx)
	if err != nil {
		return err
	}

	if _, err = tx.Exec("update users set password=? where id=?", user.Password, user.ID); err != nil {
		return err
	}

	return nil
}
