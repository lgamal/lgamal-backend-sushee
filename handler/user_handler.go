package handler

import (
	"final-project-backend/config"
	"final-project-backend/errorlist"
	"final-project-backend/utils"
	"fmt"
	"strconv"

	"final-project-backend/entity"
	"final-project-backend/handler/router_helper"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Register(c *gin.Context) {
	var reqBody entity.UserRegisterReqBody

	if err := c.BindJSON(&reqBody); err != nil {
		router_helper.GenerateErrorMessage(c, errorlist.BadRequestError("something wrong with the request content", "INPUT_INCOMPLETE"))
		return
	}

	u, err := h.userUsecase.Register(&reqBody)

	if err != nil {
		router_helper.GenerateErrorMessage(c, err)
		return
	}

	router_helper.GenerateResponseMessage(c, u)
}

func (h *Handler) Login(c *gin.Context) {
	var reqBody entity.UserLoginReqBody
	if err := c.BindJSON(&reqBody); err != nil {
		router_helper.GenerateErrorMessage(c, errorlist.BadRequestError("should provide identifier and password", "LOGIN_INPUT_INCOMPLETE"))
		return
	}

	AccessToken, refreshToken, err := h.userUsecase.Login(reqBody.Identifier, reqBody.Password)
	if err != nil {
		router_helper.GenerateErrorMessage(c, err)
		return
	}

	expirationLimit, _ := strconv.Atoi(config.Config.AuthConfig.TimeLimitRefreshToken)
	c.SetCookie(
		"refreshToken",
		refreshToken,
		expirationLimit,
		"/",
		"localhost",
		false,
		true,
	)

	router_helper.GenerateResponseMessage(c, AccessToken)
}


func (h *Handler) Refresh(c *gin.Context) {
	refreshToken, err := c.Cookie("refreshToken")
	if err != nil {
		router_helper.GenerateErrorMessage(c, errorlist.UnauthorizedError())
		return
	}

	accessToken, err := h.userUsecase.Refresh(refreshToken)
	if err != nil {
		router_helper.GenerateErrorMessage(c, err)
		return
	}
	router_helper.GenerateResponseMessage(c, accessToken)
}

func (h *Handler) UserDetail(c *gin.Context) {
	username := c.GetString("username")

	userInfo, err := h.userUsecase.GetDetailUserByUsername(username)
	if err != nil {
		router_helper.GenerateErrorMessage(c, err)
		return
	}

	router_helper.GenerateResponseMessage(c, userInfo)
}

func (h *Handler) UpdateUserProfile(c *gin.Context) {
	username := c.GetString("username")
	var formFile entity.UserProfileUploadReqBody
	fmt.Println("-=-=-=-=-=-=-=-=-=-=-=-=-=-")
	fmt.Println("BEFORE BIND:", formFile)
	
	if err := c.ShouldBind(&formFile); err != nil {
		router_helper.GenerateErrorMessage(c, errorlist.BadRequestError("invalid input", "UPDATE_INPUT_INCOMPLETE"))
	}

	fmt.Println("-=-=-=-=-=-=-=-=-=-=-=-=-=-")
	fmt.Println("AFTER BIND:", formFile)
	fmt.Println("-=-=-=-=-=-=-=-=-=-=-=-=-=-")

	var uploadUrl string
	var err error
	if formFile.Img != nil {
		uploadUrl, err = utils.NewMediaUpload().FileUpload(username, formFile)
		if err != nil {
			router_helper.GenerateErrorMessage(c, err)
			return
		}
	}

	updateProfileReq := entity.UserEditDetailsReqBody{
		FullName: formFile.FullName,
		Phone: formFile.Phone,
		Email: formFile.Email,
		Password: formFile.Password,
		ProfilePicture: uploadUrl,
	}

	userInfo, err := h.userUsecase.UpdateUserDetailsByUsername(username, updateProfileReq)
	if err != nil {
		router_helper.GenerateErrorMessage(c, err)
		return
	}

	router_helper.GenerateResponseMessage(c, userInfo)
}