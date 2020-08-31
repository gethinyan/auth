package models

import (
	"fmt"
)

// TableName 指定版本表表名
func (Version) TableName() string {
	return "versions"
}

// Version 定义版本表对应的结构
type Version struct {
	ID      uint   `json:"id" gorm:"primaryKey;not null;autoIncrement"`
	Module  string `json:"module" gorm:"size:50;not null;default:''"`
	Version int    `json:"version" gorm:"not null;default:0"`
}

// GetVersionList 获取版本列表
func GetVersionList() (versions []Version) {
	dbConn.Find(&versions)
	return
}

// SaveVersion 保存版本信息，包括创建/更新
func SaveVersion(version *Version) error {
	if err := dbConn.Save(version).Error; err != nil {
		return err
	}
	fmt.Println(version)
	return nil
}
