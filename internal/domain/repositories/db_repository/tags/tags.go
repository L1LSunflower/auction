package tags

import (
	"context"

	"github.com/L1LSunflower/auction/internal/domain/entities"
	"github.com/L1LSunflower/auction/internal/tools/context_with_depends"
)

type Repository struct{}

func (r *Repository) Create(ctx context.Context, tagName string) (*entities.Tag, error) {
	tx, err := context_with_depends.TxFromContext(ctx)
	if err != nil {
		return nil, err
	}

	tag := &entities.Tag{Name: tagName}
	if err = tx.QueryRow("insert into tags (name) values (?) returning id", tag.Name).Scan(&tag.ID); err != nil {
		return nil, err
	}

	return tag, nil
}

func (r *Repository) ByName(ctx context.Context, tagName string) (*entities.Tag, error) {
	db, err := context_with_depends.GetDb(ctx)
	if err != nil {
		return nil, err
	}

	tag := &entities.Tag{Name: tagName}
	if err = db.QueryRow("select id from tags where name=?", tag.Name).Scan(&tag.ID); err != nil {
		return nil, err
	}

	return tag, nil
}

func (r *Repository) CreateLink(ctx context.Context, itemID, fileID int) (*entities.ItemTags, error) {
	tx, err := context_with_depends.TxFromContext(ctx)
	if err != nil {
		return nil, err
	}

	itemTag := &entities.ItemTags{
		ItemID: itemID,
		TagID:  fileID,
	}
	if _, err = tx.Exec("insert into item_tags (item_id, tag_id) values  (?, ?)", itemTag.ItemID, itemTag.TagID); err != nil {
		return nil, err
	}

	return itemTag, nil
}

func (r *Repository) Tags(ctx context.Context, itemID int) ([]*entities.Tag, error) {
	db, err := context_with_depends.GetDb(ctx)
	if err != nil {
		return nil, err
	}

	var tags []*entities.Tag
	rows, err := db.Query("select t.id, t.name from tags t join item_tags it on t.id = it.tag_id where it.item_id=?", itemID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		tag := new(entities.Tag)
		if err = rows.Scan(&tag.ID, &tag.Name); err != nil {
			return nil, err
		}
		tags = append(tags, tag)
	}

	return tags, nil
}

func (r *Repository) DeleteItemTags(ctx context.Context, itemID, tagID int) error {
	tx, err := context_with_depends.TxFromContext(ctx)
	if err != nil {
		return err
	}

	if _, err = tx.Exec("delete from item_tags where item_id=? and tag_id=?", itemID, tagID); err != nil {
		return err
	}

	return nil
}

func (r *Repository) DeleteItemLinks(ctx context.Context, itemID int) error {
	tx, err := context_with_depends.TxFromContext(ctx)
	if err != nil {
		return err
	}

	if _, err = tx.Exec("delete from item_tags where item_id=?", itemID); err != nil {
		return err
	}

	return nil
}
