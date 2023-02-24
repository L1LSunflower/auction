package auction

//func Create(ctx *fiber.Ctx) error {
//	request, ok := ctx.Locals(requests.RequestKey).(*auctionReq.Create)
//	if !ok {
//		return responses.NewFailedResponse(ctx, errorhandler.ErrParseRequest)
//	}
//
//	auction, err := auctionService.Create(context.Background(), request)
//	if err != nil {
//		return responses.NewFailedResponse(ctx, err)
//	}
//
//	return auctionResp.Create(auction)
//}
//
//func Auction(ctx *fiber.Ctx) error {
//	request, ok := ctx.Locals(requests.RequestKey).(*auctionReq.Auction)
//	if !ok {
//		return responses.NewFailedResponse(ctx, errorhandler.ErrParseRequest)
//	}
//
//	metadata, err := metadata.GetParams(ctx)
//	if err != nil {
//		return err
//	}
//
//	auction, err := auctionService.Auction(context.Background(), request)
//	if err != nil {
//		return responses.NewFailedResponse(ctx, err)
//	}
//
//	return auctionResp.Auction(auction)
//}
//
//func Auctions(ctx *fiber.Ctx) error {
//	request, ok := ctx.Locals(requests.RequestKey).(*auctionReq.Auctions)
//	if !ok {
//		return responses.NewFailedResponse(ctx, errorhandler.ErrParseRequest)
//	}
//
//	auctions, err := auctionService.Auctions(context.Background(), request)
//	if err != nil {
//		return responses.NewFailedResponse(ctx, err)
//	}
//
//	return auctionResp.Auctions(auctions)
//}
//
//func Update(ctx *fiber.Ctx) error {
//	request, ok := ctx.Locals(requests.RequestKey).(*auctionReq.Update)
//	if !ok {
//		return responses.NewFailedResponse(ctx, errorhandler.ErrParseRequest)
//	}
//
//	auction, err := auctionService.Update(context.Background(), request)
//	if err != nil {
//		return responses.NewFailedResponse(ctx, err)
//	}
//
//	return auctionResp.Update(auction)
//}
//
//func Start(ctx *fiber.Ctx) error {
//	request, ok := ctx.Locals(requests.RequestKey).(*auctionReq.Start)
//	if !ok {
//		return responses.NewFailedResponse(ctx, errorhandler.ErrParseRequest)
//	}
//
//	auction, err := auctionService.Start(context.Background(), request)
//	if err != nil {
//		return responses.NewFailedResponse(ctx, err)
//	}
//
//	return auctionResp.Start(action)
//}
//
//func End(ctx *fiber.Ctx) error {
//	request, ok := ctx.Locals(requests.RequestKey).(*auctionReq.End)
//	if !ok {
//		return responses.NewFailedResponse(ctx, errorhandler.ErrParseRequest)
//	}
//
//	auction, err := auctionService.End(context.Background(), request)
//	if err != nil {
//		return responses.NewFailedResponse(ctx, err)
//	}
//
//	return auctionResp.End(auction)
//}
//
//func Delete(ctx *fiber.Ctx) error {
//	request, ok := ctx.Locals(requests.RequestKey).(*auctionReq.Delete)
//	if !ok {
//		return responses.NewFailedResponse(ctx, errorhandler.ErrParseRequest)
//	}
//
//	auction, err := auctionService.Delete(context.Background(), request)
//	if err != nil {
//		return responses.NewFailedResponse(ctx, err)
//	}
//
//	return auctionResp.Delete(auction)
//}
