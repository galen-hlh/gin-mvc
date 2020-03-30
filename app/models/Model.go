package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Model struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt int64
	UpdatedAt int64
}

func (u *Model) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("UpdatedAt", int(time.Now().Unix()))
	return nil
}

//BeforeCreate is a hook to set the created_at column to UNIX timestamp int.
func (u *Model) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("UpdatedAt", int(time.Now().Unix()))
	scope.SetColumn("CreatedAt", int(time.Now().Unix()))
	return nil
}
