package usecase

import (
	"errors"
	"final-project-backend/entity"
	"final-project-backend/errorlist"
	"final-project-backend/repository"

	"github.com/jackc/pgtype"
	"gorm.io/gorm"
)

type CartUsecase interface {
	GetCart(username string) (*[]entity.Cart, error)
	GetCartByCartId(username string, cartId int) (*entity.Cart, error)
	AddCart(username string, c *entity.CartReqBody) (*entity.Cart, error)
	DeleteCart(username string) (error)
	DeleteCartByCartId(username string, cartId int) (error)
	UpdateCartByCartId(username string, cartId int, updatePremises *entity.CartEditDetailsReqBody) (*entity.Cart, error)
}

type cartUsecaseImpl struct {
	cartRepository   repository.CartRepository
	userRepository repository.UserRepository
}

type CartUsecaseConfig struct {
	CartRepository   repository.CartRepository
	UserRepository repository.UserRepository
}

func NewCartUsecase(c CartUsecaseConfig) CartUsecase {
	return &cartUsecaseImpl{
		cartRepository:   c.CartRepository,
		userRepository: c.UserRepository,
	}
}

func guardNullJSON(j pgtype.JSON) pgtype.JSON {
	if j.Status == 0 {
		return pgtype.JSON{Bytes:[]byte{34,34}, Status:2}
	}
	return j
}

func validateCartOwnership(cart *entity.Cart, user *entity.User) error {
	if cart.UserId != int(user.ID) {
		return errorlist.UnauthorizedError()
	} 
	return nil
}

func (u *cartUsecaseImpl) GetCart(username string) (*[]entity.Cart, error) {
	carts, err := u.cartRepository.GetCartByUsername(username)
	if errors.Is(err, gorm.ErrRecordNotFound) || len(*carts) == 0 {
		return nil, errorlist.BadRequestError("no cart found", "NO_CART_FOUND")
	}
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errorlist.InternalServerError()
	}

	return carts, nil
}

func (u *cartUsecaseImpl) GetCartByCartId(username string, cartId int) (*entity.Cart, error) {
	cart, err := u.cartRepository.GetCartByCartId(cartId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errorlist.BadRequestError("cart not found", "NO_CART_FOUND")
	}
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errorlist.InternalServerError()
	}

	user, err := u.userRepository.GetUserByEmailOrUsername(username)
	if err != nil {
		return nil, errorlist.InternalServerError()
	}

	err = validateCartOwnership(cart, user)
	if err != nil {
		return nil, err
	}

	return cart, nil
}

func (u *cartUsecaseImpl) AddCart(username string, c *entity.CartReqBody) (*entity.Cart, error) {
	user, err := u.userRepository.GetUserByEmailOrUsername(username)
	if err != nil {
		return nil, errorlist.InternalServerError()
	}
	
	newCartItem := entity.Cart{
		UserId: int(user.ID),
		MenuId: c.MenuId,
		PromotionId: c.PromotionId,
		Quantity: c.Quantity,
		MenuOption: guardNullJSON(c.MenuOption),
	}

	cart, err := u.cartRepository.AddItemToCart(&newCartItem)
	if err != nil {
		return nil, errorlist.InternalServerError()
	}
	
	return cart, nil
}


func (u *cartUsecaseImpl) DeleteCart(username string) (error) {
	err := u.cartRepository.DeleteCart(username)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errorlist.BadRequestError("no cart found", "NO_CART_FOUND")
	}
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return errorlist.InternalServerError()
	}
	
	return nil
}

func (u *cartUsecaseImpl) DeleteCartByCartId(username string, cartId int) (error) {
	cart, err := u.cartRepository.GetCartByCartId(cartId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errorlist.BadRequestError("cart not found", "NO_CART_FOUND")
	}
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return errorlist.InternalServerError()
	}

	user, err := u.userRepository.GetUserByEmailOrUsername(username)
	if err != nil {
		return errorlist.InternalServerError()
	}

	err = validateCartOwnership(cart, user)
	if err != nil {
		return err
	}

	err = u.cartRepository.DeleteCartByCartId(cartId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errorlist.BadRequestError("cart item not found", "NO_CART_FOUND")
	}
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return errorlist.InternalServerError()
	}
	
	return nil
}


func (u *cartUsecaseImpl) UpdateCartByCartId(username string, cartId int, reqBody *entity.CartEditDetailsReqBody) (*entity.Cart, error) {
	cart, err := u.cartRepository.GetCartByCartId(cartId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errorlist.BadRequestError("cart not found", "NO_CART_FOUND")
	}
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errorlist.InternalServerError()
	}

	user, err := u.userRepository.GetUserByEmailOrUsername(username)
	if err != nil {
		return nil, errorlist.InternalServerError()
	}

	err = validateCartOwnership(cart, user)
	if err != nil {
		return nil, err
	}
	
	
	newCart := entity.Cart{
		Quantity: reqBody.Quantity,
		MenuOption: guardNullJSON(reqBody.MenuOption),
	}

	err = u.cartRepository.UpdateCartByCartId(cartId, &newCart)
	if err != nil {
		return nil, errorlist.InternalServerError()
	}

	cart, err = u.cartRepository.GetCartByCartId(cartId)
	if err != nil {
		return nil, errorlist.InternalServerError()
	}

	return cart, err
}
