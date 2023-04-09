package auctions

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/L1LSunflower/auction/internal/domain/entities"
	"github.com/L1LSunflower/auction/internal/tools/context_with_depends"
	"strings"
)

func (r *Repository) CreateMember(ctx context.Context, member *entities.AuctionMember) error {
	tx, err := context_with_depends.TxFromContext(ctx)
	if err != nil {
		return err
	}

	fields := []string{
		"auction_id",
		"participant_id",
		"price",
		"first_name",
		"last_name",
	}

	query := fmt.Sprintf("insert into auction_members (%s) values (?, ?, ?, ?, ?)", strings.Join(fields, fieldsSeparator))
	if err = tx.QueryRow(query, member.AuctionID, member.ParticipantID, member.Price, member.FirstName, member.LastName).Err(); err != nil {
		return err
	}

	return nil
}

func (r *Repository) Member(ctx context.Context, auctionID int, userID string) (*entities.AuctionMember, error) {
	db, err := context_with_depends.GetDb(ctx)
	if err != nil {
		return nil, err
	}

	auctionMember := &entities.AuctionMember{}
	row := db.QueryRow("select auction_id, participant_id from auction_members where auction_id=? and participant_id=?", auctionID, userID)
	if err = row.Scan(&auctionMember.AuctionID, &auctionMember.ParticipantID); err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return auctionMember, nil
}

func (r *Repository) Members(ctx context.Context, auctionID int) ([]*entities.AuctionMember, error) {
	db, err := context_with_depends.GetDb(ctx)
	if err != nil {
		return nil, err
	}

	fields := []string{
		"auction_id",
		"participant_id",
		"price",
		"first_name",
		"last_name",
	}

	var auctionMembers []*entities.AuctionMember

	query := fmt.Sprintf("select %s from auction_members where auction_id=? order by price desc", strings.Join(fields, fieldsSeparator))
	rows, err := db.Query(query, auctionID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		auctionMember := &entities.AuctionMember{}
		firstName := sql.NullString{}
		lastName := sql.NullString{}

		if err = rows.Scan(&auctionMember.AuctionID, &auctionMember.ParticipantID, &auctionMember.Price, &firstName, &lastName); err != nil {
			return nil, err
		}

		auctionMember.FirstName = firstName.String
		auctionMember.LastName = lastName.String

		auctionMembers = append(auctionMembers, auctionMember)
	}

	return auctionMembers, nil
}
