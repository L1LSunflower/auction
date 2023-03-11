package responses

import (
	"github.com/L1LSunflower/auction/internal/domain/aggregates"
	"github.com/L1LSunflower/auction/internal/domain/entities"
	"github.com/L1LSunflower/auction/internal/responses/structs"
	"github.com/gofiber/fiber/v2"
	"time"
)

func SuccessSignUp(ctx *fiber.Ctx, user *entities.User) error {
	return ctx.JSON(structs.SignUp{
		ID:    user.ID,
		Phone: user.Phone,
		Date:  time.Now().Format(time.RFC3339),
	})
}

func SuccessSignIn(ctx *fiber.Ctx, userToken *aggregates.UserToken) error {
	return ctx.JSON(structs.UserToken{
		ID:           userToken.User.ID,
		AccessToken:  userToken.Token.AccessToken,
		RefreshToken: userToken.Token.RefreshToken,
	})
}

func SuccessConfirm(ctx *fiber.Ctx, userToken *aggregates.UserToken) error {
	return ctx.JSON(structs.UserToken{
		ID:           userToken.User.ID,
		AccessToken:  userToken.Token.AccessToken,
		RefreshToken: userToken.Token.RefreshToken,
	})
}

func SuccessGetUser(ctx *fiber.Ctx, user *entities.User) error {
	return ctx.JSON(structs.User{
		ID:        user.ID,
		Phone:     user.Phone,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		City:      user.City,
	})
}

func SuccessChangePassword(ctx *fiber.Ctx, userToken *aggregates.UserToken) error {
	return ctx.JSON(structs.UserToken{
		ID:           userToken.User.ID,
		AccessToken:  userToken.Token.AccessToken,
		RefreshToken: userToken.Token.RefreshToken,
	})
}

func RefreshTokens(ctx *fiber.Ctx, tokens *entities.Tokens) error {
	return ctx.JSON(structs.Tokens{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	})
}

func SuccessSendOtp(ctx *fiber.Ctx, status, message string) error {
	return ctx.JSON(structs.OtpSent{
		Status:  status,
		Message: message,
	})
}

func UpdateUser(ctx *fiber.Ctx, user *entities.User) error {
	return ctx.JSON(structs.UpdateUser{
		Status:    successStatus,
		ID:        user.ID,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		City:      user.City,
	})
}

func DeleteUser(ctx *fiber.Ctx, user *entities.User) error {
	return ctx.JSON(structs.DeleteUser{
		Status: successStatus,
		ID:     user.ID,
	})
}

func UserProfile(ctx *fiber.Ctx, userProfile *aggregates.ProfileAggregation) error {
	userProfileResponse := &structs.Profile{Status: successStatus, Balance: userProfile.Balance.Balance}

	for _, auction := range userProfile.Auctions {
		file := GetFirstVideoOrImage(auction.Files)
		userProfileResponse.Auctions = append(userProfileResponse.Auctions, structs.AuctionWithFile{
			ID:               auction.Auction.ID,
			Status:           auction.Auction.Status,
			ShortDescription: auction.Auction.ShortDescription,
			Category:         auction.Auction.Category,
			Files:            file,
		})
	}

	return ctx.JSON(userProfileResponse)
}

func UserProfileHistory(ctx *fiber.Ctx, userProfile *aggregates.ProfileHistoryAggregation) error {
	userProfileResponse := &structs.ProfileHistory{Status: successStatus}

	for _, auction := range userProfile.Auctions {
		file := GetFirstVideoOrImage(auction.Files)
		userProfileResponse.Auctions = append(userProfileResponse.Auctions, structs.AuctionWithFile{
			ID:               auction.Auction.ID,
			Status:           auction.Auction.Status,
			ShortDescription: auction.Auction.ShortDescription,
			Category:         auction.Auction.Category,
			Files:            file,
		})
	}

	return ctx.JSON(userProfileResponse)
}
