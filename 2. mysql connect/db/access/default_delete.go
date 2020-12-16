package access

import "restful/mysql/model"

func (d *_default) DeleteUser(userId string) (err error) {
	err = d.tx.Where("user_id = ?", userId).Delete(&model.User{}).Error
	return
}

func (d *_default) DeleteBook(UUID string) (err error) {
	err = d.tx.Where("uuid = ?", UUID).Delete(&model.Book{}).Error
	return
}