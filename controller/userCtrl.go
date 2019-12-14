package controller

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/DualVectorFoil/Zelda/app/conf"
	"github.com/DualVectorFoil/Zelda/client/httpclient"
	"github.com/DualVectorFoil/Zelda/dao"
	"github.com/DualVectorFoil/Zelda/model"
	jsonUtils "github.com/DualVectorFoil/Zelda/utils/json"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type UserCtrl struct {
	RegisterDao   dao.IRegisterUserDao
	VerifyCodeDao dao.IVerifyCodeDao
}

func NewUserCtrl(registerDao dao.IRegisterUserDao, verifyCodeDao dao.IVerifyCodeDao) *UserCtrl {
	return &UserCtrl{
		RegisterDao:   registerDao,
		VerifyCodeDao: verifyCodeDao,
	}
}

func (ctrl *UserCtrl) Login(ctx *gin.Context) {

}

func (ctrl *UserCtrl) Register(ctx *gin.Context) {
	registerInfoBytes, err := base64.StdEncoding.DecodeString(ctx.PostForm("register_info"))
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"request_register_info": ctx.PostForm("register_info"),
			"err":                   err,
		}).Error("base64.StdEncoding.DecodeString register info failed.")
		ctx.String(http.StatusBadRequest, jsonUtils.JsonResp(http.StatusBadRequest, &model.RegisterRespModel{
			IsSuccess: false,
			Error:     "base64.StdEncoding.DecodeString register info failed.",
		}))
		return
	}

	registerInfo := strings.Split(string(registerInfoBytes), ":")
	if len(registerInfo) != 4 {
		logrus.WithFields(logrus.Fields{
			"request_register_info": ctx.PostForm("register_info"),
			"err":                   err,
		}).Error("Uncorrected register info.")
		ctx.String(http.StatusBadRequest, jsonUtils.JsonResp(http.StatusBadRequest, &model.RegisterRespModel{
			IsSuccess: false,
			Error:     "Uncorrected register info.",
		}))
		return
	}

	phoneNum, userName, pwd, verifyCode := registerInfo[0], registerInfo[1], registerInfo[2], registerInfo[3]
	if phoneNum == "" || userName == "" || pwd == "" || verifyCode == "" {
		logrus.WithFields(logrus.Fields{
			"request_register_info": ctx.PostForm("register_info"),
			"err":                   err,
		}).Error("Uncorrected register info.")
		ctx.String(http.StatusBadRequest, jsonUtils.JsonResp(http.StatusBadRequest, &model.RegisterRespModel{
			IsSuccess: false,
			Error:     "Uncorrected register info.",
		}))
		return
	}

	err = ctrl.RegisterDao.RegisterUser(ctx, phoneNum, userName, pwd, verifyCode)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"phoneNum":   phoneNum,
			"userName":   userName,
			"pwd_length": len(pwd),
			"verifyCode": verifyCode,
		}).Error("Register failed, err: " + err.Error())
		ctx.JSON(http.StatusBadRequest, jsonUtils.JsonResp(http.StatusBadRequest, &model.RegisterRespModel{
			IsSuccess: false,
			Error:     err.Error(),
		}))
		return
	}

	ctx.JSON(http.StatusOK, jsonUtils.JsonResp(http.StatusOK, &model.RegisterRespModel{IsSuccess: true}))
}

func (ctrl *UserCtrl) VerifyCode(ctx *gin.Context) {
	phoneNumBytes, err := base64.StdEncoding.DecodeString(ctx.PostForm("phone_num"))
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"request_phone_num": ctx.PostForm("phone_num"),
			"err":               err,
		}).Error("base64.StdEncoding.DecodeString phone number failed.")
		ctx.String(http.StatusBadRequest, jsonUtils.JsonResp(http.StatusBadRequest, "Uncorrected phone number."))
		return
	}
	phoneNum := string(phoneNumBytes)
	verifyCode := strconv.Itoa(rand.Intn(899999) + 100009)

	// TODO open when need sendVerifyCode
	fmt.Println("verifyCode: " + verifyCode)
	//err = ctrl.sendVerifyCode(phoneNum, verifyCode)
	if err != nil {
		ctx.String(http.StatusInternalServerError, jsonUtils.JsonResp(http.StatusInternalServerError, "verify code send failed, server error occured."))
		return
	}

	err = ctrl.VerifyCodeDao.SetVerifyCodeInfo(ctx.Request.Context(), phoneNum, verifyCode)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"phoneNum":   phoneNum,
			"verifyCode": verifyCode,
		}).Error("SetVerifyCodeInfo failed, err: " + err.Error())
		ctx.String(http.StatusInternalServerError, jsonUtils.JsonResp(http.StatusInternalServerError, "set verify code in redis failed."))
		return
	}

	ctx.String(http.StatusOK, jsonUtils.JsonResp(http.StatusOK, "verify code has sended."))
}

func (ctrl *UserCtrl) ModifyPassword(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": 200,
		"word":   "haha111",
	})
}

func (ctrl *UserCtrl) sendVerifyCode(phoneNum string, verifyCode string) error {
	response, err := httpclient.GetHttpClientInstance().PostForm(conf.MMS_URL, url.Values{
		"appId":     {conf.APP_ID},
		"appSecret": {conf.APP_SECRETE},
		"message":   {fmt.Sprintf(conf.MMS_CONTENT, verifyCode)},
		"number":    {phoneNum},
	})

	if err != nil {
		return err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return err
	}

	mmsResponse := &model.MMSModel{}
	err = json.Unmarshal(body, mmsResponse)
	if err != nil {
		return err
	}

	if mmsResponse.Code != 0 {
		return errors.New(mmsResponse.Data)
	}

	return nil
}
