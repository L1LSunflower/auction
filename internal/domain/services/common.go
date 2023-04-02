package services

import (
	"fmt"
	"github.com/L1LSunflower/auction/internal/domain/entities"
	"github.com/L1LSunflower/auction/internal/tools/metadata"
	"math/rand"
	"time"
)

const (
	codeLength   = 4
	tokenLength  = 32
	codeCharset  = "0123456789"
	tokenCharset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func GenerateRandomCode() string {
	return stringWithCharset(codeLength, codeCharset)
}

func GenerateToken() *entities.Tokens {
	token := stringWithCharset(tokenLength, tokenCharset)
	return &entities.Tokens{
		AccessToken:  string(token[:len(token)/2]),
		RefreshToken: string(token[len(token)/2:]),
	}
}

func stringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func GetLimitAndOffset(metadata *metadata.Metadata) error {

	metadata.LastPage = metadata.Total / metadata.PerPage
	if metadata.Total%metadata.PerPage != 0 {
		metadata.LastPage++
	}

	if metadata.CurrentPage > metadata.LastPage {
		return fmt.Errorf("failed to get page")
	}

	metadata.Limit, metadata.Offset = metadata.PerPage, (metadata.CurrentPage-1)*metadata.PerPage

	return nil
}

func CreateMapFromFiles(files []*entities.File) map[string]bool {
	filesMap := make(map[string]bool)
	for _, file := range files {
		filesMap[file.Name] = true
	}
	return filesMap
}

func CreateMapFromTags(tags []*entities.Tag) map[string]bool {
	tagsMap := make(map[string]bool)
	for _, tag := range tags {
		tagsMap[tag.Name] = true
	}
	return tagsMap
}
