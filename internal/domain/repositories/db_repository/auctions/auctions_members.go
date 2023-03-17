package auctions

import (
	"context"
	"github.com/L1LSunflower/auction/internal/domain/entities"
	"github.com/L1LSunflower/auction/internal/tools/context_with_depends"
)

func (r *Repository) CreateMember(ctx context.Context, auctionID int, userID string) (*entities.AuctionMember, error) {
	db, err := context_with_depends.GetDb(ctx)
	if err != nil {
		return nil, err
	}

	auctionMember := &entities.AuctionMember{AuctionID: auctionID, ParticipantID: userID}
	rows, err := db.Query("insert into auction_members (auction_id, participant_id) values (?, ?)", auctionMember.AuctionID, auctionMember.ParticipantID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		if err = rows.Scan(&auctionMember.AuctionID, &auctionMember.ParticipantID); err != nil {
			return nil, err
		}
	}

	return auctionMember, nil
}

func (r *Repository) Member(ctx context.Context, auctionID int, userID string) (*entities.AuctionMember, error) {
	db, err := context_with_depends.GetDb(ctx)
	if err != nil {
		return nil, err
	}

	auctionMember := &entities.AuctionMember{}
	row := db.QueryRow("select auction_id, participant_id from auction_members where auction_id=? and participant_id=?", auctionID, userID)
	if err = row.Scan(&auctionMember.AuctionID, &auctionMember.ParticipantID); err != nil {
		return nil, err
	}

	return auctionMember, nil
}
