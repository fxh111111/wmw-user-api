package services

import (
	"strings"
	"time"
	"wmw-user-api/dao"
	"wmw-user-api/model"

	"github.com/fxh111111/utility/jwt"
	"github.com/fxh111111/utility/response"
	"github.com/fxh111111/utility/wmwerrors"
	"github.com/gin-gonic/gin"
)

const MaxAge = time.Hour

type user struct{}

var User *user

func (u *user) Login(c *gin.Context) {
	req := new(model.UserLoginReq)
	err := c.ShouldBind(req)
	if err != nil {
		response.ErrorExit(c, wmwerrors.BadReq(err))
		return
	}
	var (
		me     *model.User
		findMe = dao.User.FindByMobile
	)
	if strings.Contains(req.Username, "@") {
		findMe = dao.User.FindByEmail
	}
	me, err = findMe(c, req.Username)
	if err != nil {
		response.ErrorExit(c, wmwerrors.Internal(err))
		return
	}
	if !strings.EqualFold(me.Password, req.Password) {
		response.ErrorExit(c, wmwerrors.BadReq("incorrect account"))
		return
	}
	var token string
	token, err = jwt.SignAToken(me.ID.Hex())
	if err != nil {
		response.ErrorExit(c, wmwerrors.Internal(err))
		return
	}
	c.SetCookie("wmw-token", token, int(MaxAge.Seconds()), "", "", true, true)
	response.DataExit(c, nil)
}

func (u *user) Register(c *gin.Context) {
	me := new(model.UserRegisterReq)
	err := c.ShouldBind(me)
	if err != nil {
		response.ErrorExit(c, wmwerrors.BadReq(err))
		return
	}
	var id string
	id, err = dao.User.Add(c, me)
	if err != nil {
		response.ErrorExit(c, wmwerrors.Internal(err))
		return
	}
	var token string
	token, err = jwt.SignAToken(id)
	if err != nil {
		response.ErrorExit(c, wmwerrors.Internal(err))
		return
	}
	c.SetCookie("wmw-token", token, int(MaxAge.Seconds()), "", "", true, true)
	response.DataExit(c, nil)
}

func (u *user) Info(c *gin.Context) {
	me, err := dao.User.FindByID(c, c.GetString("uid"))
	if err != nil {
		response.ErrorExit(c, wmwerrors.Internal(err))
		return
	}
	response.DataExit(c, me)
}
