package access

import (
	"restful/mysql/db/access/errors"
	"restful/mysql/model"
)

func (d *_default) CreateUser(user *model.User) (*model.User, error) {
	result := d.tx.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}

	result.Error = errors.UserAssertionError
	return nil, result.Error
}

func (d *_default) CreateBook(book *model.Book) (*model.Book, error) {
	result := d.tx.Create(book)
	if result.Error != nil {
		return nil, result.Error
	}

	result.Error = errors.BookAssertionError
	return nil, result.Error
}