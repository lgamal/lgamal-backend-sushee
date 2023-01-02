package usecase

import (
	"errors"
	"final-project-backend/entity"
	"final-project-backend/errorlist"
	"final-project-backend/repository"
	"time"

	"gorm.io/gorm"
)

type OrderUsecase interface {
	GetPaymentOption() (*[]entity.PaymentOption, error)
	AddOrder(username string, reqBody *entity.OrderReqBody) (*entity.Order, error)
	GetOrderHistory(username string) (*[]entity.Order, error)
}

type orderUsecaseImpl struct {
	orderRepository   repository.OrderRepository
	userRepository repository.UserRepository
	cartRepository repository.CartRepository
	couponRepository repository.CouponRepository
}

type OrderUsecaseConfig struct {
	OrderRepository   repository.OrderRepository
	UserRepository repository.UserRepository
	CartRepository repository.CartRepository
	CouponRepository repository.CouponRepository

}

func NewOrderUsecase(c OrderUsecaseConfig) OrderUsecase {
	return &orderUsecaseImpl{
		orderRepository:   c.OrderRepository,
		userRepository: c.UserRepository,
		cartRepository: c.CartRepository,
		couponRepository: c.CouponRepository,
	}
}

func validateCartOwnershipAndAvailability(cart *entity.Cart, user *entity.User) error {
	if cart.UserId != int(user.ID) {
		return errorlist.UnauthorizedError()
	} 
	if cart.IsOrdered {
		return errorlist.BadRequestError("the cart has been ordered before", "INVALID_CART")
	} 

	return nil
}


func (u *orderUsecaseImpl) GetPaymentOption() (*[]entity.PaymentOption, error) {
	paymentOptions, err :=  u.orderRepository.GetPaymentOption()
	if err != nil {
		return nil, errorlist.InternalServerError()
	}
	
	return paymentOptions, nil
}

func (u *orderUsecaseImpl) AddOrder(username string, reqBody *entity.OrderReqBody) (*entity.Order, error) {
	if len(reqBody.CartIdList) == 0 {
		return nil, errorlist.BadRequestError("Please order something", "INVALID_ORDER")
	}
	user, err := u.userRepository.GetUserByEmailOrUsername(username)
	if err != nil {
		return nil, errorlist.InternalServerError()
	}

	carts, err :=  u.cartRepository.GetCartByCartIds(reqBody.CartIdList)
	if err != nil {
		return nil, errorlist.InternalServerError()
	}
	if len(*carts) != len(reqBody.CartIdList) {
		return nil, errorlist.BadRequestError("Some cart IDs are not found", "INVALID_CART_IDS")
	}
	for _, cart := range *carts {
		err = validateCartOwnershipAndAvailability(&cart, user)
		if err != nil {
			return nil, err
		}
	}

	newOrder := entity.Order{
		UserId: int(user.ID),
		OrderDate: time.Now(),
		PaymentOptionId: reqBody.PaymentOptionId,
	}

	if reqBody.CouponCode != "" {
		coupon, r, err := u.couponRepository.GetUserCouponByCouponCode(int(user.ID), reqBody.CouponCode)
		if errors.Is(err, gorm.ErrRecordNotFound) || r == 0{
			return nil, errorlist.BadRequestError("coupon code invalid", "INVALID_COUPON_CODE")
		}
		if err != nil {
			return nil, errorlist.InternalServerError()
		}
		couponId := int(coupon.ID) 
		newOrder.CouponId = &couponId
		newOrder.DiscountAmount = coupon.DiscountAmount
	}
	
	totalPrice, err :=  u.cartRepository.GetCartTotalPriceByCartIds(reqBody.CartIdList)
	if err != nil {
		return nil, err
	}
	newOrder.GrossAmount = totalPrice
	newOrder.NetAmount = newOrder.GrossAmount-newOrder.DiscountAmount

	order, err := u.orderRepository.AddOrder(&newOrder)

	if err != nil {
		return nil, errorlist.InternalServerError()
	}


	var newOrderedMenus []entity.OrderedMenu
	for _, cart := range *carts {
		o := entity.OrderedMenu{
			OrderId: int(order.ID),
			Quantity: cart.Quantity,
			MenuOption: cart.MenuOption,
		}
		if cart.MenuId != nil {
			o.MenuId = cart.MenuId
		}
		if cart.PromotionId != nil {
			o.PromotionId = cart.PromotionId

		}
		newOrderedMenus = append(newOrderedMenus, o)
	}

	_, err = u.orderRepository.AddOrderedMenu(&newOrderedMenus)
	if err != nil {
		return nil, errorlist.InternalServerError()
	}

	
	u.cartRepository.UpdateCartByCartIds(reqBody.CartIdList, &entity.Cart{
		IsOrdered: true,
	})

	for _, id := range reqBody.CartIdList {
		u.cartRepository.DeleteCartByCartId(id)
	}


	

	// todo: delivery

	return order, nil
}


func (u *orderUsecaseImpl) GetOrderHistory(username string) (*[]entity.Order, error) {
	user, err := u.userRepository.GetUserByEmailOrUsername(username)
	if err != nil {
		return nil, errorlist.InternalServerError()
	}

	order, err := u.orderRepository.GetOrderHistory(int(user.ID))
	if err != nil {
		return nil, errorlist.InternalServerError()
	}

	return order, nil
}

// todo: set delivery status
// plan: Prepared --> Sending --> Received
// Details:
// P -> default val
// S -> admin confirmed
// S -> admin confirmed