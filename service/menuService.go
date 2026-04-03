package service

import (
	"github.com/Nathanaelaxl/qresto-api/model"
	"github.com/Nathanaelaxl/qresto-api/repository"
)

type MenuService interface {
	GetAllMenus() ([]model.Menu, error)
}

type menuService struct {
	menuRepo repository.MenuRepository
}

func NewMenuService(repo repository.MenuRepository) MenuService {
	return &menuService{menuRepo: repo}
}

func (s *menuService) GetAllMenus() ([]model.Menu, error) {
	return s.menuRepo.FindAllMenus()
}
