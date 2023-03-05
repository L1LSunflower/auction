package files

import (
	"context"

	"github.com/L1LSunflower/auction/internal/domain/entities"
	"github.com/L1LSunflower/auction/internal/tools/context_with_depends"
)

type Repository struct{}

func (r *Repository) Create(ctx context.Context, filename string) (*entities.File, error) {
	tx, err := context_with_depends.TxFromContext(ctx)
	if err != nil {
		return nil, err
	}

	file := &entities.File{Name: filename}
	if err = tx.QueryRow("insert into item_images (filename) values (?) returning item_id", file.Name).Scan(&file.ID); err != nil {
		return nil, err
	}

	return file, nil
}

func (r *Repository) CreateLink(ctx context.Context, itemID int, filename string) (*entities.ItemFiles, error) {
	tx, err := context_with_depends.TxFromContext(ctx)
	if err != nil {
		return nil, err
	}

	itemFile := &entities.ItemFiles{
		ItemID:   itemID,
		Filename: filename,
	}
	if _, err = tx.Exec("insert into item_images (item_id, filename) values (?, ?)", itemFile.ItemID, itemFile.Filename); err != nil {
		return nil, err
	}

	return itemFile, nil
}

func (r *Repository) Files(ctx context.Context, itemID int) ([]*entities.File, error) {
	db, err := context_with_depends.GetDb(ctx)
	if err != nil {
		return nil, err
	}

	var files []*entities.File
	rows, err := db.Query("select item_id, filename from item_images where item_id=?", itemID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		file := new(entities.File)
		if err = rows.Scan(&file.ID, &file.Name); err != nil {
			return nil, err
		}
		files = append(files, file)
	}

	return files, nil
}

func (r *Repository) Delete(ctx context.Context, itemID int, filename string) error {
	tx, err := context_with_depends.TxFromContext(ctx)
	if err != nil {
		return err
	}

	if _, err = tx.Exec("delete from item_images where item_id=? and filename=?", itemID, filename); err != nil {
		return err
	}

	return nil
}

func (r *Repository) DeleteAll(ctx context.Context, itemId int) error {
	tx, err := context_with_depends.TxFromContext(ctx)
	if err != nil {
		return err
	}

	if _, err = tx.Exec("delete from item_tags where item_id=?", itemId); err != nil {
		return err
	}

	return nil
}
