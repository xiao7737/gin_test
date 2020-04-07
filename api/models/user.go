package models

import (
	orm "gin_test/api/database"
	"strings"
)

type User struct {
	ID       int64  `json:"id"`
	Username string `form:"username" json:"username" binding:"required,NameValid"` //添加自定义验证规则
	Age      int    `form:"age" json:"age" binding:"required,gt=10,lt=120"`
	// binding中的规则不能有空格！！！此处遇坑
	//或者使用 sql.NullInt64，scanner/valuer避免0，''，false的情况
	Password string `form:"password" json:"password" binding:"required"`
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
	/*row := orm.Eloquent.Table("users").Where("username = ?", "xiaoxichuan").Select("username, age").Row()
	if err = row.Scan(&user.Username, &user.Age); err != nil {
		return
	}
	return*/
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
	if err = orm.Eloquent.First(&updateUser, id).Error; err != nil {
		return
	}
	if err = orm.Eloquent.Model(&updateUser).Update(&user).Error; err != nil {
		return
	}
	return
}

func (user *User) GetById(id int64) (users []User, err error) {
	if err = orm.Eloquent.First(&users, id).Error; err != nil {
		return
	}
	return
}

// 此处采用strings.Builder拼接like字符串
func (user *User) GetUserByName(username string) (users []User, err error) {
	var builder strings.Builder
	builder.WriteString("%" + username + "%")
	username = builder.String()
	if err = orm.Eloquent.Where("username like ?", username).Find(&users).Error; err != nil {
		return
	}
	return
}

// 所有链式方法都会创建并克隆一个新的 DB 对象 (共享一个连接池)，gorm 在多 goroutine 中是并发安全的
