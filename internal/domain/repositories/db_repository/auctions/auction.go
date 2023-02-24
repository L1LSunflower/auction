package auctions

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/L1LSunflower/auction/internal/domain/entities"
	"github.com/L1LSunflower/auction/internal/tools/context_with_tx"
	"github.com/L1LSunflower/auction/internal/tools/metadata"
	"strings"
)

const (
	fieldsSeparator = ","
	whereSeparator  = " AND "
)

type Repository struct{}

func (r *Repository) Create(ctx context.Context, auction *entities.Auction) error {
	tx, err := context_with_tx.TxFromContext(ctx)
	if err != nil {
		return err
	}

	fields := []string{
		"owner_id",
		"item_id",
		"title",
		"description",
		"start_price",
		"minimal_price",
		"status",
		"created_at",
		"updated_at",
	}

	query := fmt.Sprintf("insert into auction (%s) values (?, ?, ?, ?, ?, ?, ?, now(), now()) returning id, create_at, updated_at", strings.Join(fields, fieldsSeparator))
	if err := tx.QueryRow(
		query,
		auction.OwnerID,
		auction.ItemID,
		auction.Title,
		auction.Description,
		auction.StartPrice,
		auction.MinPrice,
		auction.Status,
	).Scan(&auction.ID, &auction.CreatedAt, &auction.UpdatedAt); err != nil {
		return err
	}

	return nil
}

func (r *Repository) Auction(ctx context.Context, id int) (*entities.Auction, error) {
	tx, err := context_with_tx.TxFromContext(ctx)
	if err != nil {
		return nil, err
	}

	auction := &entities.Auction{}
	winnerID := sql.NullInt64{}
	startedAt := sql.NullTime{}
	endedAt := sql.NullTime{}
	fields := []string{
		"id",
		"owner_id",
		"winner_id",
		"item_id",
		"title",
		"description",
		"start_price",
		"minimal_price",
		"status",
		"started_at",
		"ended_at",
		"created_at",
		"updated_at",
	}

	query := fmt.Sprintf("select %s from auction where id=? and deleted_at is null", strings.Join(fields, fieldsSeparator))
	if err := tx.QueryRow(
		query,
		id,
	).Scan(
		&auction.ID,
		&auction.OwnerID,
		&winnerID,
		&auction.ItemID,
		&auction.Title,
		&auction.Description,
		&auction.StartPrice,
		&auction.MinPrice,
		&auction.Status,
		&startedAt,
		&endedAt,
		&auction.CreatedAt,
		&auction.UpdatedAt,
	); err != nil {
		return nil, err
	}
	auction.WinnerID = int(winnerID.Int64)
	auction.StartedAt = startedAt.Time
	auction.EndedAt = endedAt.Time

	return auction, nil
}

func (r *Repository) Auctions(ctx context.Context, where []string, metadata *metadata.Metadata) ([]*entities.Auction, error) {
	tx, err := context_with_tx.TxFromContext(ctx)
	if err != nil {
		return nil, err
	}

	var (
		auctions []*entities.Auction
		query    string
	)
	fields := []string{
		"id",
		"owner_id",
		"winner_id",
		"item_id",
		"title",
		"description",
		"start_price",
		"minimal_price",
		"status",
		"started_at",
		"ended_at",
		"created_at",
		"updated_at",
	}

	if len(where) > 0 {
		query = fmt.Sprintf("select %s from auction where deleted_at is null and %s limit ? offset ?", strings.Join(fields, fieldsSeparator), strings.Join(where, whereSeparator))
	} else {
		query = fmt.Sprintf("select %s from auction where deleted_at is null limit ? offset ?", strings.Join(fields, fieldsSeparator))
	}

	rows, err := tx.Query(query, metadata.Limit, metadata.Offset)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		auction := &entities.Auction{}
		winnerID := sql.NullInt64{}
		startedAt := sql.NullTime{}
		endedAt := sql.NullTime{}
		if err = rows.Scan(
			&auction.ID,
			&auction.OwnerID,
			&winnerID,
			&auction.ItemID,
			&auction.Title,
			&auction.Description,
			&auction.StartPrice,
			&auction.MinPrice,
			&auction.Status,
			&startedAt,
			&endedAt,
			&auction.CreatedAt,
			&auction.UpdatedAt,
		); err != nil {
			return nil, err
		}
		auction.WinnerID = int(winnerID.Int64)
		auction.StartedAt = startedAt.Time
		auction.EndedAt = endedAt.Time

		auctions = append(auctions, auction)
	}

	return auctions, nil
}

func (r *Repository) Update(ctx context.Context, auction *entities.Auction) error {
	tx, err := context_with_tx.TxFromContext(ctx)
	if err != nil {
		return err
	}

	fields := []string{
		"owner_id=?",
		"winner_id=?",
		"item_id=?",
		"title=?",
		"description=?",
		"start_price=?",
		"minimal_price=?",
		"status=?",
		"started_at=?",
		"ended_at=?",
		"updated_at=now()",
	}

	query := fmt.Sprintf("update auction set %s where id=?", strings.Join(fields, fieldsSeparator))
	if _, err = tx.Exec(
		query,
		auction.OwnerID,
		auction.WinnerID,
		auction.ItemID,
		auction.Title,
		auction.Description,
		auction.StartPrice,
		auction.MinPrice,
		auction.Status,
		auction.StartedAt,
		auction.EndedAt,
		auction.ID,
	); err != nil {
		return err
	}

	return nil
}

func (r *Repository) Start(ctx context.Context, auction *entities.Auction) error {
	tx, err := context_with_tx.TxFromContext(ctx)
	if err != nil {
		return err
	}

	if _, err = tx.Exec("update auction set started_at=now(), status='active' where id=?", auction.ID); err != nil {
		return err
	}

	if err = tx.QueryRow("select status, started_at from auction where id=?", auction.ID).Scan(&auction.Status, &auction.StartedAt); err != nil {
		return err
	}

	return nil
}

func (r *Repository) End(ctx context.Context, auction *entities.Auction) error {
	tx, err := context_with_tx.TxFromContext(ctx)
	if err != nil {
		return err
	}

	if _, err = tx.Exec("update auction set ended_at=now(), status='completed' where id=?", auction.ID); err != nil {
		return err
	}

	if err = tx.QueryRow("select status, ended_at from auction where id=?", auction.ID).Scan(&auction.Status, &auction.EndedAt); err != nil {
		return err
	}

	return nil
}

func (r *Repository) Delete(ctx context.Context, auction *entities.Auction) error {
	tx, err := context_with_tx.TxFromContext(ctx)
	if err != nil {
		return err
	}

	if _, err = tx.Exec("update auction set deleted_at=now() where id=?", auction.ID); err != nil {
		return err
	}

	if err = tx.QueryRow("select deleted_at from auction where id=?", auction.ID).Scan(&auction.DeletedAt); err != nil {
		return err
	}

	return nil
}

func (r *Repository) ByOwnerID(ctx context.Context, ownerID int) (*entities.Auction, error) {
	tx, err := context_with_tx.TxFromContext(ctx)
	if err != nil {
		return nil, err
	}

	auction := &entities.Auction{}
	fields := []string{
		"id",
		"owner_id",
		"winner_id",
		"item_id",
		"title",
		"description",
		"start_price",
		"minimal_price",
		"status",
		"started_at",
		"ended_at",
		"created_at",
		"updated_at",
	}

	query := fmt.Sprintf("select %s from auction where owner_id=? and deleted_at is null group by created_at desc limit 1", strings.Join(fields, fieldsSeparator))
	rows, err := tx.Query(query, ownerID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		winnerID := sql.NullInt64{}
		startedAt := sql.NullTime{}
		endedAt := sql.NullTime{}
		if err = rows.Scan(
			&auction.ID,
			&auction.OwnerID,
			&winnerID,
			&auction.ItemID,
			&auction.Title,
			&auction.Description,
			&auction.StartPrice,
			&auction.MinPrice,
			&auction.Status,
			&startedAt,
			&endedAt,
			&auction.CreatedAt,
			&auction.UpdatedAt,
		); err != nil {
			return nil, err
		}
		auction.WinnerID = int(winnerID.Int64)
		auction.StartedAt = startedAt.Time
		auction.EndedAt = endedAt.Time
	}

	return auction, nil
}
