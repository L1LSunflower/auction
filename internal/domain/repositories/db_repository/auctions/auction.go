package auctions

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/L1LSunflower/auction/internal/domain/aggregates"
	"github.com/L1LSunflower/auction/internal/domain/entities"
	"github.com/L1LSunflower/auction/internal/tools/context_with_depends"
	"github.com/L1LSunflower/auction/internal/tools/metadata"
)

const (
	fieldsSeparator = ","
	whereSeparator  = " AND "
)

type Repository struct{}

func (r *Repository) Create(ctx context.Context, auction *entities.Auction) error {
	tx, err := context_with_depends.TxFromContext(ctx)
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
	db, err := context_with_depends.GetDb(ctx)
	if err != nil {
		return nil, err
	}

	auction := &entities.Auction{}
	winnerID := sql.NullString{}
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
	if err := db.QueryRow(
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
	auction.WinnerID = winnerID.String
	auction.StartedAt = startedAt.Time
	auction.EndedAt = endedAt.Time

	return auction, nil
}

func (r *Repository) Auctions(ctx context.Context, where []string, metadata *metadata.Metadata) (*aggregates.AuctionsItem, error) {
	db, err := context_with_depends.GetDb(ctx)
	if err != nil {
		return nil, err
	}

	var (
		auctions = &aggregates.AuctionsItem{}
		query    string
	)
	fields := []string{
		"a.id",
		"a.status",
		"i.files",
		// TODO: add params as json object in database
	}

	if len(where) > 0 {
		query = fmt.Sprintf("select %s from auction a join items i on a.item_id=i.id where a.deleted_at is null and %s limit ? offset ?", strings.Join(fields, fieldsSeparator), strings.Join(where, whereSeparator))
	} else {
		query = fmt.Sprintf("select %s from auction a join items i on a.item_id=i.id where a.deleted_at is null limit ? offset ?", strings.Join(fields, fieldsSeparator))
	}

	rows, err := db.Query(query, metadata.Limit, metadata.Offset)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		aItem := &aggregates.AuctItem{}
		if err = rows.Scan(
			&aItem.ID,
			&aItem.Status,
			&aItem.Files,
		); err != nil {
			return nil, err
		}

		auctions.AuctsItem = append(auctions.AuctsItem, aItem)
	}

	return auctions, nil
}

func (r *Repository) Update(ctx context.Context, auction *entities.Auction) error {
	tx, err := context_with_depends.TxFromContext(ctx)
	if err != nil {
		return err
	}

	fields := []string{
		"winner_id=?",
		"title=?",
		"description=?",
		"start_price=?",
		"minimal_price=?",
		"status=?",
		"updated_at=now()",
	}

	query := fmt.Sprintf("update auction set %s where id=?", strings.Join(fields, fieldsSeparator))
	if _, err = tx.Exec(
		query,
		auction.WinnerID,
		auction.Title,
		auction.Description,
		auction.StartPrice,
		auction.MinPrice,
		auction.Status,
		auction.ID,
	); err != nil {
		return err
	}

	return nil
}

func (r *Repository) Start(ctx context.Context, auction *entities.Auction) error {
	tx, err := context_with_depends.TxFromContext(ctx)
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
	tx, err := context_with_depends.TxFromContext(ctx)
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
	tx, err := context_with_depends.TxFromContext(ctx)
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
	db, err := context_with_depends.GetDb(ctx)
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
	rows, err := db.Query(query, ownerID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		winnerID := sql.NullString{}
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
		auction.WinnerID = winnerID.String
		auction.StartedAt = startedAt.Time
		auction.EndedAt = endedAt.Time
	}

	return auction, nil
}

func (r *Repository) Count(ctx context.Context) (int, error) {
	db, err := context_with_depends.GetDb(ctx)
	if err != nil {
		return 0, err
	}

	var count int
	if err = db.QueryRow("select count(*) from auction").Scan(&count); err != nil {
		return 0, err
	}

	return count, nil
}
