package repository

import (
	"final-project-backend/entity"

	"gorm.io/gorm"
)

type CartRepository interface {
	AddItemToCart(c *entity.Cart) (*entity.Cart, error)
	GetCartByUsername(username string) (*[]entity.Cart, error)
	GetCartByCartId(cartId int) (*entity.Cart, error)
	GetCartByCartIds(cartIds []int) (*[]entity.Cart, error)
	DeleteCart(username string) (error)
	DeleteCartByCartId(cartId int) (error)
	UpdateCartByCartId(cartId int, updatePremises *entity.Cart) (error)
	UpdateCartByCartIds(cartIds []int, updatePremises *entity.Cart) (error)
	GetCartTotalPriceByCartIds(cartIds []int) (float64, error) 
}

type CartRepositoryImpl struct {
	db *gorm.DB
}

type CartRepositoryConfig struct {
	DB *gorm.DB
}

func NewCartRepository(c CartRepositoryConfig) CartRepository {
	return &CartRepositoryImpl{
		db: c.DB,
	}
}

func (r *CartRepositoryImpl) AddItemToCart(c *entity.Cart) (*entity.Cart, error) {
	err := r.db.
		Create(&c).Error
	return c, err
}

func (r *CartRepositoryImpl) GetCartByUsername(username string) (*[]entity.Cart, error) {
	var carts []entity.Cart

	userSQ := r.db.
		Select("id").
		Where("username = (?)", username).
		Table("users")
	menuSubQuery := r.db.
		Where("is_ordered != (?)", true).
		Table("carts")
	query := r.db.
		Model(&entity.Cart{}).
		Preload("Menu").
		Table("(?) as th", menuSubQuery).
		Where("user_id = (?)", userSQ).
		Order("created_at desc").
		Find(&carts)
		
	err := query.Error
	return &carts, err
}

func (r *CartRepositoryImpl) GetCartByCartId(cartId int) (*entity.Cart, error) {
	var cart entity.Cart

	query := r.db.
		Where("id = (?)", cartId).
		First(&cart)
		
	err := query.Error
	return &cart, err
}

func (r *CartRepositoryImpl) GetCartByCartIds(cartIds []int) (*[]entity.Cart, error) {
	var cart []entity.Cart

	query := r.db.
		Where("id in (?)", cartIds).
		Find(&cart)
		
	err := query.Error
	return &cart, err
}

func (r *CartRepositoryImpl) DeleteCart(username string) (error) {
	var carts []entity.Cart

	userSQ := r.db.
		Select("id").
		Where("username = (?)", username).
		Table("users")
	query := r.db.
		Where("user_id = (?) AND is_ordered != (?)", userSQ, true).
		Delete(&carts)

	err := query.Error
	return err
}  

func (r *CartRepositoryImpl) DeleteCartByCartId(cartId int) (error) {
	var carts []entity.Cart

	query := r.db.
		Where("id = (?) AND is_ordered != (?)", cartId, true).
		Delete(&carts)

	err := query.Error
	return err
}

func (r *CartRepositoryImpl) UpdateCartByCartId(cartId int, newCart *entity.Cart) (error) {
	err := r.db.
		Where("id = ?", cartId).
		Updates(newCart).
		Debug().Error
	return err
}

func (r *CartRepositoryImpl) UpdateCartByCartIds(cartIds []int, newCart *entity.Cart) (error) {
	err := r.db.
		Where("id in (?)", cartIds).
		Updates(newCart).
		Debug().Error
	return err
}

func (r *CartRepositoryImpl) GetCartTotalPriceByCartIds(cartIds []int) (float64, error) {
	var result float64
	sq := r.db.
		Table("carts").
		Select("COALESCE(carts.promotion_price, menus.price)*carts.quantity as cart_total_price").
		Where("carts.id in (?)", cartIds).
		Joins("JOIN menus on carts.menu_id = menus.id")

	err := r.db.
		Table("(?) as sq_cart_prices", sq).
		Select("sum(cart_total_price) as total_price").
		Scan(&result).
		Error
	
	return result, err
}
