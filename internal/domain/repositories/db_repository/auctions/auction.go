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
)

type Repository struct{}

func (r *Repository) Create(ctx context.Context, auction *entities.Auction) error {
	tx, err := context_with_depends.TxFromContext(ctx)
	if err != nil {
		return err
	}

	// TODO: set visit_status
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
		"visit_status",
	}

	query := fmt.Sprintf("insert into auctions (%s) values (?, ?, ?, ?, ?, ?, ?, ?, now(), now(), ?)", strings.Join(fields, fieldsSeparator))
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
		entities.VisitNotSet,
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
	price := sql.NullFloat64{}
	startedAt := sql.NullTime{}
	endedAt := sql.NullTime{}
	visitStatus := sql.NullString{}
	visitStartedDate := sql.NullTime{}
	visitEndDate := sql.NullTime{}
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
		"price",
		"created_at",
		"updated_at",
		"visit_status",
		"visit_start_date",
		"visit_end_date",
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
		&price,
		&auction.CreatedAt,
		&auction.UpdatedAt,
		&visitStatus,
		&visitStartedDate,
		&visitEndDate,
	); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	auction.ShortDescription = shortDescription.String
	auction.WinnerID = winnerID.String
	auction.Price = price.Float64
	auction.StartedAt = startedAt.Time
	auction.EndedAt = endedAt.Time
	auction.VisitStatus = visitStatus.String
	auction.VisitStartDate = visitStartedDate.Time
	auction.VisitEndDate = visitEndDate.Time

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
		query += fmt.Sprintf(" join item_tags it on a.item_id = it.item_id join tags t on it.tag_id = t.id where t.name in (%s) and deleted_at is null and status!='%s'", tags, entities.CompletedStatus)
	} else {
		query += fmt.Sprintf(" where deleted_at is null and status!='%s'", entities.CompletedStatus)
	}

	if len(where) > 0 {
		query += fmt.Sprintf(" and %s and status!='%s'", where, entities.CompletedStatus)
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
		"short_description=?",
		"started_at=?",
		"visit_start_date=?",
		"visit_end_date=?",
		"updated_at=now()",
	}

	query := fmt.Sprintf("update auctions set %s where id=?", strings.Join(fields, fieldsSeparator))
	if _, err = tx.Exec(
		query,
		auction.StartPrice,
		auction.MinPrice,
		auction.ShortDescription,
		auction.StartedAt,
		auction.VisitStartDate,
		auction.VisitEndDate,
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

	if _, err = tx.Exec("update auctions set started_at=now(), status='active', ended_at=?, updated_at=now() where id=?", endedDate, id); err != nil {
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

func (r *Repository) UpdatePrice(ctx context.Context, id int, price float64) error {
	tx, err := context_with_depends.TxFromContext(ctx)
	if err != nil {
		return err
	}

	if _, err = tx.Exec("update auctions set price=? where id=?", price, id); err != nil {
		return err
	}

	return nil
}

func (r *Repository) ActivateAuctions(ctx context.Context) error {
	db, err := context_with_depends.GetDb(ctx)
	if err != nil {
		return err
	}

	if _, err = db.Exec("update auctions set status=?, ended_at=now() + interval 12 hour where status=? and started_at <= now() and deleted_at is null", entities.ActiveStatus, entities.InactiveStatus); err != nil {
		return err
	}

	return nil
}

func (r *Repository) EndAuctions(ctx context.Context) error {
	db, err := context_with_depends.GetDb(ctx)
	if err != nil {
		return err
	}

	if _, err = db.Exec("update auctions set status=? where status=? and ended_at <= now() and deleted_at is null", entities.CompletedStatus, entities.ActiveStatus); err != nil {
		return err
	}

	return nil
}

func (r *Repository) SetPrice(ctx context.Context, auctionID int, userID string, price float64) error {
	tx, err := context_with_depends.TxFromContext(ctx)
	if err != nil {
		return err
	}

	if err = tx.QueryRow("update auctions set price=?, winner_id=? where id=?", price, userID, auctionID).Err(); err != nil {
		return err
	}

	if err = tx.QueryRow("update auction_members set price=? where auction_id=? and participant_id=?", price, auctionID, userID).Err(); err != nil {
		return err
	}

	return nil
}

func (r *Repository) Completed(ctx context.Context, userID string) ([]*entities.Auction, error) {
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

	query := fmt.Sprintf(`select %s from auctions a where winner_id=? and status=?`, strings.Join(fields, fieldsSeparator))
	rows, err := db.Query(query, userID, entities.CompletedStatus)
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

func (r *Repository) SetVisit(ctx context.Context, visit *entities.AuctionVisit) error {
	tx, err := context_with_depends.TxFromContext(ctx)
	if err != nil {
		return err
	}

	fields := []string{
		"visit_status=?",
		"visit_start_date=?",
		"visit_end_date=?",
	}

	query := fmt.Sprintf("update auctions set %s where id=?", strings.Join(fields, fieldsSeparator))
	if _, err = tx.Exec(query, entities.VisitSet, visit.StartDate, visit.EndDate, visit.AuctionID); err != nil {
		return err
	}

	return nil
}

func (r *Repository) Visitor(ctx context.Context, auctionID int, userID string) (*entities.AuctionVisitor, error) {
	db, err := context_with_depends.GetDb(ctx)
	if err != nil {
		return nil, err
	}

	fields := []string{
		"auction_id",
		"user_id",
		"first_name",
		"last_name",
		"phone",
	}

	auctionVisitor := &entities.AuctionVisitor{}

	query := fmt.Sprintf("select %s from auction_visitors where auction_id=? and user_id=?", strings.Join(fields, fieldsSeparator))
	firstName := sql.NullString{}
	LastName := sql.NullString{}
	if err = db.QueryRow(query, auctionID, userID).Scan(&auctionVisitor.AuctionID, &auctionVisitor.UserID, &firstName, &LastName, &auctionVisitor.Phone); err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	auctionVisitor.FirstName = firstName.String
	auctionVisitor.LastName = LastName.String

	return auctionVisitor, nil
}

func (r *Repository) Visit(ctx context.Context, visitor *entities.AuctionVisitor) error {
	tx, err := context_with_depends.TxFromContext(ctx)
	if err != nil {
		return err
	}

	fields := []string{
		"auction_id",
		"user_id",
		"first_name",
		"last_name",
		"phone",
	}

	query := fmt.Sprintf("insert into auction_visitors (%s) values (?, ?, ?, ?, ?)", strings.Join(fields, fieldsSeparator))
	if err = tx.QueryRow(query, visitor.AuctionID, visitor.UserID, visitor.FirstName, visitor.LastName, visitor.Phone).Err(); err != nil {
		return err
	}

	return nil
}

func (r *Repository) Visitors(ctx context.Context, auctionID int) ([]*entities.AuctionVisitor, error) {
	db, err := context_with_depends.GetDb(ctx)
	if err != nil {
		return nil, err
	}

	fields := []string{
		"auction_id",
		"user_id",
		"first_name",
		"last_name",
		"phone",
	}

	var visitors []*entities.AuctionVisitor

	query := fmt.Sprintf("select %s from auction_visitors where auction_id=?", strings.Join(fields, fieldsSeparator))
	rows, err := db.Query(query, auctionID)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	for rows.Next() {
		visitor := &entities.AuctionVisitor{}
		firstName := sql.NullString{}
		lastName := sql.NullString{}

		if err = rows.Scan(&visitor.AuctionID, &visitor.UserID, &firstName, &lastName, &visitor.Phone); err != nil {
			return nil, err
		}

		visitor.FirstName = firstName.String
		visitor.LastName = lastName.String

		visitors = append(visitors, visitor)
	}

	return visitors, nil
}

func (r *Repository) Unvisit(ctx context.Context, auctionID int, userID string) error {
	tx, err := context_with_depends.TxFromContext(ctx)
	if err != nil {
		return err
	}

	if err = tx.QueryRow("delete from auction_visitors where auction_id=? and user_id=?", auctionID, userID).Err(); err != nil {
		return err
	}

	return nil
}

func (r *Repository) VisitorsCount(ctx context.Context, auctionID int) (int, error) {
	db, err := context_with_depends.GetDb(ctx)
	if err != nil {
		return 0, err
	}

	var visitorsCount int
	if err = db.QueryRow("select count(*) from auction_visitors where auction_id=?", auctionID).Scan(&visitorsCount); err != nil && err != sql.ErrNoRows {
		return 0, err
	}

	return visitorsCount, nil
}

func (r *Repository) StartVisit(ctx context.Context) error {
	db, err := context_with_depends.GetDb(ctx)
	if err != nil {
		return err
	}

	if _, err = db.Exec("update auctions set visit_status=? where status=? and visit_start_date <= now() and deleted_at is null", entities.VisitSet, entities.VisitOpened); err != nil {
		return err
	}

	return nil
}

func (r *Repository) EndVisit(ctx context.Context) error {
	db, err := context_with_depends.GetDb(ctx)
	if err != nil {
		return err
	}

	if _, err = db.Exec("update auctions set visit_status=? where status=? and visit_end_date <= now() and deleted_at is null", entities.VisitOpened, entities.VisitClosed); err != nil {
		return err
	}

	return nil
}

func (r *Repository) Owner(ctx context.Context, userID string) (bool, error) {
	db, err := context_with_depends.GetDb(ctx)
	if err != nil {
		return false, err
	}

	var owner int
	if err = db.QueryRow("select exists(select * from auctions where owner_id=?)").Scan(owner); err != nil && err != sql.ErrNoRows {
		return false, err
	}

	if owner == 0 {
		return false, nil
	} else {
		return true, nil
	}
}
