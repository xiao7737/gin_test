package models

import (
	orm "gin_test/api/database"
)

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Age      int    `json:"age"` //或者使用 sql.NullInt64，scanner/valuer避免0，''，false的情况
	Password string `json:"password"`
}

// 默认表名为struct的复数，本model对应的表名为users
// 可以指定表名，或者关掉这种规则
/*func (User) TableName() string {
	return "user"
}*/
/*func init()  {
	orm.Eloquent.SingularTable(true)
}*/

// insert data
func (user User) Insert() (id int64, err error) {
	result := orm.Eloquent.Create(&user)
	id = user.ID
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

// get user list
func (user *User) Users() (users []User, err error) {
	if err = orm.Eloquent.Find(&users).Error; err != nil {
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

func (user *User) Update(id int64) (updateUser User, err error) {
	if err = orm.Eloquent.Select([]string{"id", "username"}).First(&updateUser, id).Error; err != nil {
		return
	}
	if err = orm.Eloquent.Model(&updateUser).Update(&user).Error; err != nil {
		return
	}
	return
}
