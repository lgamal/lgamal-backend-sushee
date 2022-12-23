package usecase

import (
	"errors"
	"final-project-backend/entity"
	"final-project-backend/errorlist"
	"final-project-backend/repository"

	"gorm.io/gorm"
)

type MenuUsecase interface {
	GetMenu(entity.MenuQuery) (*[]entity.Menu, int, error)

}

type menuUsecaseImpl struct {
	menuRepository   repository.MenuRepository
}

type MenuUsecaseConfig struct {
	MenuRepository   repository.MenuRepository
}

func NewMenuUsecase(c MenuUsecaseConfig) MenuUsecase {
	return &menuUsecaseImpl{
		menuRepository:   c.MenuRepository,
	}
}


func (u *menuUsecaseImpl) GetMenu(q entity.MenuQuery) (*[]entity.Menu, int, error) {
	menus, rows, err := u.menuRepository.GetMenu(q)
	if errors.Is(err, gorm.ErrRecordNotFound) || len(*menus) == 0 {
		return nil, 0, errorlist.BadRequestError("no menu found", "NO_MENU_FOUND")
	}
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, 0, errorlist.InternalServerError()
	}
	
	maxPage := (rows + q.Limit - 1)/ q.Limit
	return menus, maxPage, nil
}