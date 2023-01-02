package usecase

import (
	"errors"
	"final-project-backend/entity"
	"final-project-backend/errorlist"
	"final-project-backend/repository"

	"gorm.io/gorm"
)

type CouponUsecase interface {
	AddCoupon(username string, reqBody *entity.CouponAddReqBody) (*entity.Coupon, error)
	GetCoupon() (*[]entity.Coupon, error)
	UpdateCoupon(username string, couponId int, reqBody *entity.CouponEditReqBody) (*entity.Coupon, error)
	DeleteCoupon(couponId int) (*entity.Coupon, error)
	GetUserCouponByUsername(username string) (*[]entity.UserCoupon, error)
	GetUserCouponByCouponCode(userId int, couponCode string) (*entity.UserCoupon, error)
}

type couponUsecaseImpl struct {
	couponRepository   repository.CouponRepository
	userRepository repository.UserRepository
}

type CouponUsecaseConfig struct {
	CouponRepository   repository.CouponRepository
	UserRepository repository.UserRepository
}

func NewCouponUsecase(c CouponUsecaseConfig) CouponUsecase {
	return &couponUsecaseImpl{
		couponRepository: c.CouponRepository,
		userRepository: c.UserRepository,
	}
}

func (u *couponUsecaseImpl) AddCoupon(username string, reqBody *entity.CouponAddReqBody) (*entity.Coupon, error) {
	admin, err := u.userRepository.GetUserByEmailOrUsername(username)
	if err != nil {
		return nil, errorlist.InternalServerError()
	}
	
	newCoupon := entity.Coupon {
		AdminId: int(admin.ID),
		Description: reqBody.Description,
		DiscountAmount: reqBody.DiscountAmount,
		QuotaInitial: reqBody.QuotaInitial,
		QuotaLeft: reqBody.QuotaLeft,
	}

	coupons, err :=  u.couponRepository.AddCoupon(&newCoupon)
	if err != nil {
		return nil, errorlist.InternalServerError()
	}
	
	return coupons, nil
}

func (u *couponUsecaseImpl) GetCoupon() (*[]entity.Coupon, error) {
	coupons, err :=  u.couponRepository.GetCoupon()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errorlist.BadRequestError("no coupon available", "NO_COUPON_EXIST")
	}
	if err != nil {
		return nil, errorlist.InternalServerError()
	}
	
	return coupons, nil
}

func (u *couponUsecaseImpl) UpdateCoupon(username string, couponId int, reqBody *entity.CouponEditReqBody) (*entity.Coupon, error) {
	admin, err := u.userRepository.GetUserByEmailOrUsername(username)
	if err != nil {
		return nil, errorlist.InternalServerError()
	}
	
	newCoupon := entity.Coupon{
		AdminId: int(admin.ID),
		Description: reqBody.Description,
		DiscountAmount: reqBody.DiscountAmount,
		QuotaInitial: reqBody.QuotaInitial,
		QuotaLeft: reqBody.QuotaLeft,
	}
	
	err =  u.couponRepository.UpdateCouponById(couponId, &newCoupon)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errorlist.BadRequestError("no coupon available", "NO_COUPON_EXIST")
	}
	if err != nil {
		return nil, errorlist.InternalServerError()
	}
	
	coupon, err := u.couponRepository.GetCouponById(couponId)
	if err != nil {
		return nil, errorlist.InternalServerError()
	}

	return coupon, nil
}

func (u *couponUsecaseImpl) DeleteCoupon(couponId int) (*entity.Coupon, error) {
	coupons, err :=  u.couponRepository.DeleteCouponById(couponId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errorlist.BadRequestError("no coupon available", "NO_COUPON_EXIST")
	}
	if err != nil {
		return nil, errorlist.InternalServerError()
	}
	
	return coupons, nil
}

func (u *couponUsecaseImpl) GetUserCouponByUsername(username string) (*[]entity.UserCoupon, error) {
	userCoupons, r, err :=  u.couponRepository.GetUserCouponByUsername(username)
	if errors.Is(err, gorm.ErrRecordNotFound) || r == 0 {
		return nil, errorlist.BadRequestError("no coupon available", "NO_COUPON_EXIST")
	}
	if err != nil {
		return nil, errorlist.InternalServerError()
	}
	
	return userCoupons, nil
}


func (u *couponUsecaseImpl) GetUserCouponByCouponCode(userId int, couponCode string) (*entity.UserCoupon, error) {
	userCoupons, r, err :=  u.couponRepository.GetUserCouponByCouponCode(userId, couponCode)
	if errors.Is(err, gorm.ErrRecordNotFound) || r == 0 {
		return nil, errorlist.BadRequestError("no coupon available", "NO_COUPON_EXIST")
	}
	if err != nil {
		return nil, errorlist.InternalServerError()
	}
	return userCoupons, nil
}
