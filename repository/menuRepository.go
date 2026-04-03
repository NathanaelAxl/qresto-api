package repository

import (
	"github.com/Nathanaelaxl/qresto-api/model"
	"gorm.io/gorm"
)

type MenuRepository interface {
	FindAllMenus() ([]model.Menu, error)
}

type menuRepository struct {
	dbConn *gorm.DB
}

func NewMenuRepository(db *gorm.DB) MenuRepository {
	return &menuRepository{dbConn: db}
}

func (r *menuRepository) FindAllMenus() ([]model.Menu, error) {
	var menuData []model.Menu
	err := r.dbConn.Find(&menuData).Error
	return menuData, err
}
