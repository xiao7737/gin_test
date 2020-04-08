package models

import (
	orm "gin_test/api/database"
	"github.com/jinzhu/gorm"
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
// Db后面加Debug()打印sql
func (user *User) Users() (users []User, count int, err error) {
	if err = orm.Eloquent.Debug().Order("id desc").Find(&users).Count(&count).Error; err != nil {
		return
	}
	return
	/*row := orm.Eloquent.Table("users").Where("username = ?", "xiao_admin").Select("username, age").Row()
	if err = row.Scan(&user.Username, &user.Age); err != nil {
		return
	}
	return*/
}

// delete
func (user *User) Destroy(id int64) (err error) {
	if err = orm.Eloquent.Select("id").First(&user, id).Error; err != nil {
		return
	}
	if err = orm.Eloquent.Delete(&user).Error; err != nil {
		return
	}
	return
}

//忽略更新字段 Omit
//指定更新字段 Select
func (user *User) Update(id int64) (users User, err error) {
	if err = orm.Eloquent.First(&users, id).Error; err != nil {
		return
	}
	if err = orm.Eloquent.Model(&users).Omit("username").Update(&user).Error; err != nil {
		return
	}
	return
}

func (user *User) GetById(id int64) (users []User, err error) {
	err = orm.Eloquent.First(&users, id).Error
	if err != nil && err != gorm.ErrRecordNotFound { // 查不到记录也是一种错
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
