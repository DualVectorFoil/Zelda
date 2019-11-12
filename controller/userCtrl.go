package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/DualVectorFoil/Zelda/app/conf"
	"github.com/DualVectorFoil/Zelda/client/httpclient"
	"github.com/DualVectorFoil/Zelda/dao"
	"github.com/DualVectorFoil/Zelda/model"
	jsonUtils "github.com/DualVectorFoil/Zelda/utils/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
)

type UserCtrl struct {
	VerifyCodeDao dao.IVerifyCodeDao
}

func NewUserCtrl(verifyCodeDao dao.IVerifyCodeDao) *UserCtrl {
	return &UserCtrl{
		VerifyCodeDao: verifyCodeDao,
	}
}

func (ctrl *UserCtrl) Login(ctx *gin.Context) {

}

func (ctrl *UserCtrl) Register(ctx *gin.Context) {

}

func (ctrl *UserCtrl) VerifyCode(ctx *gin.Context) {
	phoneNum := ctx.PostForm("phone_num")
	verifyCode := strconv.Itoa(rand.Intn(899999) + 100009)

	err := sendVerifyCode(phoneNum, verifyCode)
	if err != nil {
		ctx.String(http.StatusInternalServerError, jsonUtils.JsonResp(http.StatusInternalServerError, "verify code send failed, server error occured."))
		return
	}

	// TODO 将phoneNum、verifyCode通过dao存入redis中，过期5分钟，本地cookie内存中存储，不用存到redis
	err = ctrl.VerifyCodeDao.SetVerifyCodeInfo(ctx.Request.Context(), phoneNum, verifyCode)
	if err != nil {
		ctx.String(http.StatusInternalServerError, jsonUtils.JsonResp(http.StatusInternalServerError, "set verify code in redis failed."))
		return
	}

	ctx.String(http.StatusOK, jsonUtils.JsonResp(http.StatusOK, "verify code has sended."))
}

func (ctrl *UserCtrl) ModifyPassword(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": 200,
		"word": "haha111",
	})
}

func sendVerifyCode(phoneNum string, verifyCode string) error {
	response, err := httpclient.GetHttpClientInstance().PostForm(conf.MMS_URL, url.Values{
		"appId": {conf.APP_ID},
		"appSecret": {conf.APP_SECRETE},
		"message": {fmt.Sprintf(conf.MMS_CONTENT, verifyCode)},
		"number": {phoneNum},
	})

	if err != nil {
		fmt.Println("sendVerifyCode failed.")
		return err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		fmt.Println("sendVerifyCode parse response body failed.")
		return err
	}

	mmsResponse := &model.MMSModel{}
	err = json.Unmarshal(body, mmsResponse)
	if err != nil {
		fmt.Println("sendVerifyCode unmarshall response body failed.")
	}

	if mmsResponse.Code != 0 {
		fmt.Println("response err: " + mmsResponse.Data)
		return errors.New(mmsResponse.Data)
	}

	return nil
}
