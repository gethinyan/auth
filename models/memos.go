package models

import (
	"fmt"
	"time"
)

// MemoModule 备忘录/便笺表模块
const MemoModule = "memo"

// 状态码
const (
	StatusNormal  = 0
	StatusArchive = 1
	StatusDeleted = 2
	StatusTrashed = 3
)

// TableName 指定备忘录/便笺表表名
func (Memo) TableName() string {
	return "memos"
}

// Memo 定义备忘录/便笺表对应的结构
type Memo struct {
	ID        uint      `json:"id" gorm:"primary_key;not null;auto_increment"`
	UserID    uint      `json:"user_id" gorm:"not null;default:0"`
	Name      string    `json:"name" gorm:"size:200;not null;default:''"`
	Content   string    `json:"content" gorm:"not null;default:''"`
	Status    int8      `json:"status" gorm:"not null;default:0"`
	CreatedAt time.Time `json:"created_at" gorm:"not null;default:current_timestamp"`
	UpdatedAt time.Time `json:"updated_at" gorm:"not null;default:current_timestamp"`
}

// MemoRequestBody 备忘录/便笺请求参数
type MemoRequestBody struct {
	// 备忘录/便笺 ID（有就传）
	ID uint `json:"id"`
	// 备忘录/便笺名
	Name string `json:"name"`
	// 备忘录/便笺内容
	Content string `json:"content"`
}

// MemoResponseBody 备忘录/便笺响应参数
type MemoResponseBody struct {
	// 备忘录/便笺 ID
	ID string `json:"id"`
	// 备忘录/便笺名
	Name string `json:"name"`
	// 备忘录/便笺内容
	Content string `json:"content"`
	// 备忘录/便笺状态（0正常1回收站2已删除）
	Status int8 `json:"status" gorm:"not null;default:0"`
	// 备忘录/便笺创建时间
	CreatedAt string `json:"created_at"`
	// 备忘录/便笺更新时间
	UpdatedAt string `json:"updated_at"`
}

// GetMemoList 获取备忘录/便笺列表
func GetMemoList() (memos []Memo) {
	dbConn.Find(&memos)
	return
}

// SaveMemo 保存备忘录/便笺信息，包括创建/更新
func SaveMemo(memo *Memo) error {
	if err := dbConn.Save(memo).Error; err != nil {
		return err
	}
	fmt.Println(memo)
	return nil
}

// InsertBatchMemo 批量新增备忘录/便笺
func InsertBatchMemo(memos []MemoRequestBody) bool {
	sql := "INSERT INTO memos (name, content) VALUES "
	// 循环 memos 数组，组合 sql 语句
	for key, memo := range memos {
		if len(memos)-1 == key {
			// 最后一条数据以分号结尾
			sql += fmt.Sprintf("('%s', '%s');", memo.Name, memo.Content)
		} else {
			sql += fmt.Sprintf("('%s', '%s'),", memo.Name, memo.Content)
		}
	}
	dbConn.Exec(sql)
	return true
}

// UpdateBatchMemo 批量更新备忘录/便笺
func UpdateBatchMemo(memos []MemoRequestBody) bool {
	fmt.Println(memos)
	// 循环 memos 数组一个一个更新
	for _, memo := range memos {
		dbConn.Table("memos").Where("id = ? AND status IN (?)", memo.ID, []int8{StatusNormal}).Updates(map[string]interface{}{"name": memo.Name, "content": memo.Content, "updated_at": time.Now()})
	}
	return true
}

// RestoreBatchMemo 批量恢复备忘录/便笺
func RestoreBatchMemo(memos []MemoRequestBody) bool {
	idArr := []uint{}
	for _, memo := range memos {
		if memo.ID > 0 {
			idArr = append(idArr, memo.ID)
		}
	}
	if len(idArr) > 0 {
		dbConn.Table("memos").Where("id IN (?) AND status IN (?)", idArr, []int8{StatusArchive, StatusDeleted}).Updates(map[string]interface{}{"status": StatusNormal, "updated_at": time.Now()})
	}
	return true
}

// ArchiveBatchMemo 批量归档备忘录/便笺
func ArchiveBatchMemo(memos []MemoRequestBody) bool {
	idArr := []uint{}
	for _, memo := range memos {
		if memo.ID > 0 {
			idArr = append(idArr, memo.ID)
		}
	}
	if len(idArr) > 0 {
		dbConn.Table("memos").Where("id IN (?) AND status IN (?)", idArr, []int8{StatusNormal}).Updates(map[string]interface{}{"status": StatusArchive, "updated_at": time.Now()})
	}
	return true
}

// DeleteBatchMemo 批量删除备忘录/便笺
func DeleteBatchMemo(memos []MemoRequestBody) bool {
	idArr := []uint{}
	for _, memo := range memos {
		if memo.ID > 0 {
			idArr = append(idArr, memo.ID)
		}
	}
	if len(idArr) > 0 {
		dbConn.Table("memos").Where("id IN (?)", idArr).Updates(map[string]interface{}{"status": StatusDeleted, "updated_at": time.Now()})
	}
	return true
}
