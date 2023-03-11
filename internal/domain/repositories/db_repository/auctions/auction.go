package auctions

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/L1LSunflower/auction/internal/domain/entities"
	"github.com/L1LSunflower/auction/internal/tools/context_with_depends"
	"github.com/L1LSunflower/auction/internal/tools/metadata"
)

const (
	fieldsSeparator = ","
	whereSeparator  = " AND "
	dateFormat      = "2006-01-02 15:04:05"
)

type Repository struct{}

func (r *Repository) Create(ctx context.Context, auction *entities.Auction) error {
	tx, err := context_with_depends.TxFromContext(ctx)
	if err != nil {
		return err
	}

	now := time.Now()
	auction.CreatedAt = now
	auction.UpdatedAt = now
	fields := []string{
		"category",
		"owner_id",
		"item_id",
		"short_description",
		"start_price",
		"minimal_price",
		"status",
		"started_at",
		"created_at",
		"updated_at",
	}

	query := fmt.Sprintf("insert into auctions (%s) values (?, ?, ?, ?, ?, ?, ?, ?, now(), now())", strings.Join(fields, fieldsSeparator))
	tag, err := tx.Exec(
		query,
		auction.Category,
		auction.OwnerID,
		auction.ItemID,
		auction.ShortDescription,
		auction.StartPrice,
		auction.MinPrice,
		auction.Status,
		auction.StartedAt,
	)
	if err != nil {
		return err
	}

	id, err := tag.LastInsertId()
	if err != nil {
		return err
	}
	auction.ID = int(id)

	return nil
}

func (r *Repository) Auction(ctx context.Context, id int) (*entities.Auction, error) {
	db, err := context_with_depends.GetDb(ctx)
	if err != nil {
		return nil, err
	}

	auction := &entities.Auction{}

	winnerID := sql.NullString{}
	shortDescription := sql.NullString{}
	startedAt := sql.NullTime{}
	endedAt := sql.NullTime{}
	fields := []string{
		"id",
		"category",
		"owner_id",
		"winner_id",
		"item_id",
		"short_description",
		"start_price",
		"minimal_price",
		"status",
		"started_at",
		"ended_at",
		"created_at",
		"updated_at",
	}

	query := fmt.Sprintf("select %s from auctions where id=? and deleted_at is null", strings.Join(fields, fieldsSeparator))
	if err = db.QueryRow(
		query,
		id,
	).Scan(
		&auction.ID,
		&auction.Category,
		&auction.OwnerID,
		&winnerID,
		&auction.ItemID,
		&shortDescription,
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
	auction.ShortDescription = shortDescription.String
	auction.WinnerID = winnerID.String
	auction.StartedAt = startedAt.Time
	auction.EndedAt = endedAt.Time

	return auction, nil
}

func (r *Repository) Auctions(ctx context.Context, where, tags, orderBy string, metadata *metadata.Metadata) ([]*entities.Auction, error) {
	db, err := context_with_depends.GetDb(ctx)
	if err != nil {
		return nil, err
	}

	var (
		auctions []*entities.Auction
		query    string
	)

	fields := []string{
		"a.id",
		"a.status",
		"a.short_description",
		"a.item_id",
		"a.category",
	}

	query = fmt.Sprintf(`select %s from auctions a`, strings.Join(fields, fieldsSeparator))
	if len(tags) > 0 {
		query += fmt.Sprintf(" join item_tags it on a.item_id = it.item_id join tags t on it.tag_id = t.id where t.name in (%s) and deleted_at is null", tags)
	} else {
		query += " where deleted_at is null"
	}

	if len(where) > 0 {
		query += fmt.Sprintf(" and %s", where)
	}

	if len(orderBy) > 0 {
		query += " order by " + orderBy
	}
	query += " limit ? offset ?"

	rows, err := db.Query(query, metadata.Limit, metadata.Offset)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		auction := &entities.Auction{}
		shortDescription := sql.NullString{}
		if err = rows.Scan(
			&auction.ID,
			&auction.Status,
			&shortDescription,
			&auction.ItemID,
			&auction.Category,
		); err != nil {
			return nil, err
		}
		auction.ShortDescription = shortDescription.String
		auctions = append(auctions, auction)
	}

	return auctions, nil
}

func (r *Repository) Update(ctx context.Context, auction *entities.Auction) error {
	tx, err := context_with_depends.TxFromContext(ctx)
	if err != nil {
		return err
	}

	fields := []string{
		"start_price=?",
		"minimal_price=?",
		"status=?",
		"updated_at=now()",
	}

	query := fmt.Sprintf("update auctions set %s where id=?", strings.Join(fields, fieldsSeparator))
	if _, err = tx.Exec(
		query,
		auction.StartPrice,
		auction.MinPrice,
		auction.Status,
		auction.ID,
	); err != nil {
		return err
	}

	return nil
}

func (r *Repository) Start(ctx context.Context, id int, endedDate time.Time) error {
	tx, err := context_with_depends.TxFromContext(ctx)
	if err != nil {
		return err
	}

	if _, err = tx.Exec("update auctions set started_at=now(), status='active', ended_at=?, updated_at=now() where id=?", id, endedDate.Format(dateFormat)); err != nil {
		return err
	}

	return nil
}

func (r *Repository) End(ctx context.Context, id int) error {
	tx, err := context_with_depends.TxFromContext(ctx)
	if err != nil {
		return err
	}

	if _, err = tx.Exec("update auctions set ended_at=now(), status='completed', updated_at=now() where id=?", id); err != nil {
		return err
	}

	return nil
}

func (r *Repository) Delete(ctx context.Context, auction *entities.Auction) error {
	tx, err := context_with_depends.TxFromContext(ctx)
	if err != nil {
		return err
	}

	if _, err = tx.Exec("update auctions set deleted_at=now() where id=?", auction.ID); err != nil {
		return err
	}

	if err = tx.QueryRow("select deleted_at from auctions where id=?", auction.ID).Scan(&auction.DeletedAt); err != nil {
		return err
	}

	return nil
}

func (r *Repository) ByOwnerID(ctx context.Context, ownerID string) ([]*entities.Auction, error) {
	db, err := context_with_depends.GetDb(ctx)
	if err != nil {
		return nil, err
	}

	var auctions []*entities.Auction
	fields := []string{
		"a.id",
		"a.status",
		"a.short_description",
		"a.item_id",
		"a.category",
	}

	query := fmt.Sprintf(`select %s from auctions a where owner_id=?`, strings.Join(fields, fieldsSeparator))
	rows, err := db.Query(query, ownerID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		auction := &entities.Auction{}
		shortDescription := sql.NullString{}
		if err = rows.Scan(
			&auction.ID,
			&auction.Status,
			&shortDescription,
			&auction.ItemID,
			&auction.Category,
		); err != nil {
			return nil, err
		}
		auction.ShortDescription = shortDescription.String
		auctions = append(auctions, auction)
	}

	return auctions, nil
}

func (r *Repository) Count(ctx context.Context) (int, error) {
	db, err := context_with_depends.GetDb(ctx)
	if err != nil {
		return 0, err
	}

	var count int
	if err = db.QueryRow("select count(*) from auctions").Scan(&count); err != nil {
		return 0, err
	}

	return count, nil
}

func (r *Repository) ActiveAuction(ctx context.Context, ownerID string) (*entities.Auction, error) {
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
		"start_price",
		"minimal_price",
		"status",
		"started_at",
		"ended_at",
		"created_at",
		"updated_at",
	}

	query := fmt.Sprintf("select %s from auctions where owner_id=? and status=? and deleted_at is null group by created_at limit 1", strings.Join(fields, fieldsSeparator))
	rows, err := db.Query(query, ownerID, entities.ActiveStatus)
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

func (r *Repository) CountInactiveAuctions(ctx context.Context, ownerID string) (int, error) {
	db, err := context_with_depends.GetDb(ctx)
	if err != nil {
		return 0, err
	}

	var count int
	if err = db.QueryRow("select count(*) from auctions where owner_id=? and status=?", ownerID, entities.InactiveStatus).Scan(&count); err != nil && err.Error() != "sql: no rows in result set" {
		return 0, err
	}

	return count, nil
}

func (r *Repository) DeleteByOwnerID(ctx context.Context, ownerID string) error {
	tx, err := context_with_depends.TxFromContext(ctx)
	if err != nil {
		return err
	}

	if _, err = tx.Exec("update auctions set deleted_at=now() where owner_id=?", ownerID); err != nil {
		return err
	}

	return nil
}

func (r *Repository) ActiveAuctions(ctx context.Context, ownerID string) ([]*entities.Auction, error) {
	db, err := context_with_depends.GetDb(ctx)
	if err != nil {
		return nil, err
	}

	var auctions []*entities.Auction
	fields := []string{
		"a.id",
		"a.status",
		"a.short_description",
		"a.item_id",
		"a.category",
	}

	query := fmt.Sprintf(`select %s from auctions a join auction_members am on a.id=am.auction_id where am.participant_id=? and a.status=?`, strings.Join(fields, fieldsSeparator))
	rows, err := db.Query(query, ownerID, entities.ActiveStatus)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		auction := &entities.Auction{}
		shortDescription := sql.NullString{}
		if err = rows.Scan(
			&auction.ID,
			&auction.Status,
			&shortDescription,
			&auction.ItemID,
			&auction.Category,
		); err != nil {
			return nil, err
		}
		auction.ShortDescription = shortDescription.String
		auctions = append(auctions, auction)
	}

	return auctions, nil
}
