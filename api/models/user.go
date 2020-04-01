package models

import (
	orm "gin_test/api/database"
)

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Age      int64  `json:"age"`
	Password string `json:"password"`
}

var Users []User

// insert data
func (user User) Insert(id int64, err error) {
	if err = orm.Eloquent.Create(&user).Error; err != nil {
		id = user.ID
		return
	}
	return
}

// get user list
func (user *User) Users(users []User, err error) {
	if res, err := orm.Eloquent.Find(&user).Error; err != nil {
		return
	}
	return
}

// delete
func (user *User) Destroy(id int64) (Result User, err error) {
	if err = orm.Eloquent.Select([]string{"id"}).First(&user, id).Error; err != nil {
		return
	}
	if err = orm.Eloquent.Delete(&user).Error; err != nil {
		return
	}
	Result = *user
	return
}

func Update() {

}
