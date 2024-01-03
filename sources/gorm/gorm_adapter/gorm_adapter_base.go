package gormadapter

import "gorm.io/gorm"

type gormAdapter struct {
	*gorm.DB
}

// Initiate New adapter gorm
func NewGormAdapter(db *gorm.DB) *gormAdapter {
	return &gormAdapter{db}
}