package access

import (
	"restful/mysql/model"
)

func (d *_default) GetUser(userId string)(user model.User) {
	err := d.tx.First(&user, "user_id = ?", userId)

	if err != nil {
		panic(err)
	}

	return
}

func (d *_default) GetBook(UUID string)(book model.Book) {
	err := d.tx.First(&book, UUID)

	if err != nil {
		panic(err)
	}

	return
}

func (d *_default) GetBookList()(books []model.Book){
	err := d.tx.Find(&books)

	if err != nil {
		panic(err)
	}

	return
}