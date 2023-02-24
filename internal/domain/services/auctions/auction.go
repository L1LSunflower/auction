package auctions

//func Create(parentCtx context.Context, db *sql.DB, request auctionReq.Create) (*entities.Auction, error) {
//	ctx, err := context_with_tx.ContextWithTx(parentCtx, db)
//	if err != nil {
//		return nil, err
//	}
//	defer context_with_tx.TxRollback(ctx)
//
//	auction, err := db_repository.AuctionInterface.ByOwnerID(ctx, request.OwnerID)
//	if err != nil {
//		return nil, err
//	}
//
//	if !auction.CretedAt.IsZero() {
//		return nil, errors.New("auction limit on auction owner id")
//	}
//
//	item := &entities.Item{}
//	if err = db_repository.ItemInterface.Create(item); err != nil {
//		return nil, errors.New("failed to create item for auction")
//	}
//
//	auction = &entities.Auction{
//		OwnerID:     request.OwnerID,
//		ItemID:      item.ID,
//		Title:       request.Title,
//		Description: request.Description,
//		Status:      "inactive",
//	}
//	if err = db_repository.AuctionInterface.Create(ctx, auction); err != nil {
//		return nil, err
//	}
//
//	context_with_tx.TxCommit(ctx)
//	return auction, nil
//}
//
//func Auction(parentCtx context.Context, db *sql.DB, request auctionReq.Auction) (*entities.Auction, error) {
//	ctx, err := context_with_tx.ContextWithTx(parentCtx, db)
//	if err != nil {
//		return nil, err
//	}
//	defer context_with_tx.TxRollback(ctx)
//
//	auction, err := db_repository.AuctionInterface.Auction(ctx, request.ID)
//	if err != nil {
//		return nil, err
//	}
//
//	if auction.CreatedAt.IsZero() {
//		return nil, errors.New("that auction does not exist")
//	}
//
//	item, err := db_repository.ItemInterface.Item(ctx, auction.ItemID)
//	if err != nil {
//		return nil, err
//	}
//
//	if item.CreatedAt.IsZero() {
//		return nil, errors.New("failed to get auction item")
//	}
//
//	context_with_tx.TxCommit(ctx)
//	return auction, nil
//}
//
//func Auctions(parentCtx context.Context, db *sql.DB, request auctionReq.Auction) ([]*entities.Auction, error) {
//	ctx, err := context_with_tx.ContextWithTx(parentCtx, db)
//	if err != nil {
//		return nil, err
//	}
//	defer context_with_tx.TxRollback(ctx)
//
//	auctions, err := db_repository.AuctionInterface.Auctions(ctx, request.Where, request.Metadata)
//	if err != nil {
//		return nil, err
//	}
//
//	context_with_tx.TxCommit(ctx)
//	return auctions, nil
//}
//
//func Update(parentCtx context.Context, db *sql.DB, request auctionReq.Update) (*entities.Auction, error) {
//	ctx, err := context_with_tx.ContextWithTx(parentCtx, db)
//	if err != nil {
//		return nil, err
//	}
//	defer context_with_tx.TxRollback(ctx)
//
//	auction, err := db_repository.AuctionInterface.Auction(ctx, request.ID)
//	if err != nil {
//		return nil, err
//	}
//
//	if auction.CreatedAt.IsZero() {
//		return nil, errors.New("that auction does not exist")
//	}
//
//	context_with_tx.TxCommit(ctx)
//	return auction, nil
//}
//
//func Start(parentCtx context.Context, db *sql.DB, request auctionReq.Start) (*entities.Auction, error) {
//	ctx, err := context_with_tx.ContextWithTx(parentCtx, db)
//	if err != nil {
//		return nil, err
//	}
//	defer context_with_tx.TxRollback(ctx)
//
//	auction, err := db_repository.AuctionInterface.Auction(request.ID)
//	if err != nil {
//		return nil, err
//	}
//
//	if auction.CreatedAt.IsZero() {
//		return nil, errors.New("that auction does not exist")
//	}
//
//	if err = db_repository.AuctionInterface.Start(ctx, auction); err != nil {
//		return nil, errors.New("failed to start auction")
//	}
//
//	context_with_tx.TxCommit(ctx)
//	return auction, nil
//}
//
//func End(parentCtx context.Context, db *sql.DB, request auctionReq.End) (*entities.Auction, error) {
//	ctx, err := context_with_tx.ContextWithTx(parentCtx, db)
//	if err != nil {
//		return nil, err
//	}
//	defer context_with_tx.TxRollback(ctx)
//
//	auction, err := db_repository.AuctionInterface.Auction(request.ID)
//	if err != nil {
//		return nil, err
//	}
//
//	if auction.CreatedAt.IsZero() {
//		return nil, errors.New("that auction does not exist")
//	}
//
//	if err = db_repository.AuctionInterface.End(ctx, auction); err != nil {
//		return nil, errors.New("failed to end auction")
//	}
//
//	context_with_tx.TxCommit(ctx)
//	return auction, nil
//}
//
//func Delete(parentCtx context.Context, db *sql.DB, request auctionReq.Delete) (*entities.Auction, error) {
//	ctx, err := context_with_tx.ContextWithTx(parentCtx, db)
//	if err != nil {
//		return nil, err
//	}
//	defer context_with_tx.TxRollback(ctx)
//
//	auction, err := db_repository.AuctionInterface.Auction(request.ID)
//	if err != nil {
//		return nil, err
//	}
//
//	if auction.CreatedAt.IsZero() {
//		return nil, errors.New("that auction does not exist")
//	}
//
//	if err = db_repository.AuctionInterface.End(ctx, auction); err != nil {
//		return nil, errors.New("failed to end auction")
//	}
//
//	context_with_tx.TxCommit(ctx)
//	return auction, nil
//}
