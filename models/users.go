package models

import (
	"fmt"
	"time"

	"github.com/gethinyan/auth/internal/util"
	"gorm.io/gorm"
)

// TableName 指定用户表表名
func (User) TableName() string {
	return "users"
}

// User 定义用户表对应的结构
type User struct {
	// 用户 ID
	ID uint `json:"id" gorm:"primaryKey;not null;autoIncrement"`
	// 手机号
	Phone string `json:"phone" gorm:"uniqueIndex:uk_phone;type:char(11);not null;default:''"`
	// 邮箱
	Email string `json:"email" gorm:"uniqueIndex:uk_email;size:50;not null;default:''"`
	// 用户名
	Username string `json:"username" gorm:"uniqueIndex:uk_username;size:50;not null;default:''"`
	// 密码
	Password string `json:"password" gorm:"type:char(60);not null;default:''"`
	// 昵称
	Nickname string `json:"nickname" gorm:"size:50;not null;default:''"`
	// 地址
	Address string `json:"address" gorm:"size:200;not null;default:''"`
	// 注册 IP
	RegIP string `json:"reg_ip" gorm:"uniqueIndex:uk_ip;size:50;not null;default:''"`
	// 注册地址（省市区）
	RegAddr string `json:"reg_addr" gorm:"size:50;not null;default:''"`
	// 性别（1 男、2 女）
	Gender int8 `json:"gender" gorm:"not null;default:1"`
	// 生日（格式 2020-01-01）
	Birth time.Time `json:"birth" gorm:"uniqueIndex:uk_birth;null;type:date;default:null"`
	// 头像地址
	AvatarURL string `json:"avatar_url" gorm:"size:200;not null;default:''"`
	// 创建时间
	CreatedAt time.Time `json:"created_at" gorm:"null;default:null"`
	// 更新时间
	UpdatedAt time.Time `json:"updated_at" gorm:"null;default:null"`
	// 删除时间
	DeletedAt time.Time `json:"deleted_at" gorm:"null;default:null"`
}

// UserRequestBody 用户请求参数
// swagger:parameters UserRequestBody
type UserRequestBody struct {
	// 手机号
	// Required: true
	Phone string `json:"phone" binding:"required"`
	// 邮箱地址
	// Required: true
	Email string `json:"email" binding:"required,email"`
	// 用户名
	// Required: true
	Username string `json:"username" binding:"required"`
	// 密码
	// Required: true
	Password string `json:"password" binding:"required"`
	// 昵称
	Nickname string `json:"nickname"`
	// 地址
	Address string `json:"address"`
	// 性别（0：女；1：男）
	// Required: true
	Gender int8 `json:"gender" binding:"required"`
	//生日
	Birth time.Time `json:"birth"`
	// 头像地址
	AvatarURL string `json:"avatar_url"`
	// 验证码（注册时必填）
	Code int `json:"code"`
}

// UserResponseBody 用户响应参数
// swagger:parameters UserResponseBody
type UserResponseBody struct {
	// 用户 ID
	ID uint `json:"id"`
	// 手机号
	Phone string `json:"phone"`
	// 邮箱地址
	Email string `json:"email"`
	// 用户名
	Username string `json:"username"`
	// 地址
	Address string `json:"address"`
	// 性别（0：女；1：男）
	Gender int8 `json:"gender"`
	// 生日
	Birth time.Time `json:"birth"`
	// 头像地址
	AvatarURL string `json:"avatar_url"`
}

// GetUserByID 通过 ID 获取用户信息
func GetUserByID(id uint) (user User, err error) {
	if err = dbConn.Where(User{ID: id}).Find(&user).Error; err != nil {
		return
	}
	return
}

// GetUserByEmail 通过邮箱获取用户信息
func GetUserByEmail(email string) (user User, err error) {
	if err = dbConn.Where(User{Email: email}).Find(&user).Error; err != nil {
		return
	}
	return
}

// BeforeSave 保存用户信息前执行逻辑
func (user *User) BeforeSave(dbConn *gorm.DB) (err error) {
	if user.Password, err = util.GeneratePassword(user.Password); err != nil {
		return
	}
	return
}

// CreateUser 创建用户
func CreateUser(user *User) error {
	user.CreatedAt = time.Now()
	if err := dbConn.Create(user).Error; err != nil {
		return err
	}
	fmt.Println(user)
	return nil
}

// UpdateUser 更新用户信息
func UpdateUser(user *User) error {
	user.UpdatedAt = time.Now()
	if err := dbConn.Model(user).Updates(user).Error; err != nil {
		return err
	}
	fmt.Println(user)
	return nil
}

// DeleteUserByID 通过 ID 删除用户（逻辑删除）
func DeleteUserByID(id uint) error {
	user := &User{ID: id}
	if err := dbConn.Delete(user).Error; err != nil {
		return err
	}
	fmt.Println(user)
	return nil
}

// ConvertToResponse 转换为响应参数
func (user *User) ConvertToResponse() UserResponseBody {
	response := UserResponseBody{
		ID:        user.ID,
		Phone:     user.Phone,
		Email:     user.Email,
		Username:  user.Username,
		Address:   user.Address,
		Gender:    user.Gender,
		Birth:     user.Birth,
		AvatarURL: user.AvatarURL,
	}
	return response
}
