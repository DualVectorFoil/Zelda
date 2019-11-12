package controller

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"strconv"
)

type UserCtrl struct {

}

func NewUserCtrl() *UserCtrl {
	return &UserCtrl{

	}
}

func (ctrl *UserCtrl) Login(ctx *gin.Context) {

}

func (ctrl *UserCtrl) Register(ctx *gin.Context) {

}

func (ctrl *UserCtrl) VerifyCode(ctx *gin.Context) {
	phoneNum := ctx.PostForm("phone_num")
	verifyCode := strconv.Itoa(rand.Intn(899999) + 100009)

	ctx.JSON(http.StatusOK, gin.H{
		"status": 200,
		"phoneNum": phoneNum + "~	~xixi",
		"code": verifyCode,
	})
}

func (ctrl *UserCtrl) ModifyPassword(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": 200,
		"word": "haha111",
	})
}
