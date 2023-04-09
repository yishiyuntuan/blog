package controller

import (
	"blog/conts"
	"blog/dao/mapper"
	"blog/middleware/logger"
	"blog/model/entity"
	"blog/model/vo"
	"blog/service"
	"blog/util/tool"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/core/router"
	"github.com/kataras/iris/v12/mvc"
)

type UserController struct {
	service service.UserService
}

func (u UserController) Configure(r router.Party) {
	u.service = service.NewService[service.UserServiceImpl](service.WithUserDao(mapper.NewUserDao()))
	// 依赖注入
	mvc.Configure(r, func(app *mvc.Application) {
		app.Handle(&u)
	})
}

// 注册用户
// @Summary 注册用户
// @Description 注册用户
// @Tags User
// @Accept  json
// @Param user formData blog.model.entity.User true "用户信息"
// @Success 200 {object} blog.model.vo.Result
// @Failure 400 {object} blog.model.vo.Result
// @Router /api/v1/user/register [post]
func (u UserController) PostRegister(user entity.User) *vo.Result {
	logger.Log.Debug(user)
	// 校验参数
	if !tool.VerifyFormat(conts.VERIFY_EXP_USERNAME, user.Username) || !tool.VerifyFormat(conts.VERIFY_EXP_PASSWORD, string(user.Password)) {
		return vo.Fail(vo.WithMessage("用户名或密码格式不正确"), vo.WithCode(400))
	}
	// 判断用户是否已经存在
	if u.service.IsExist(user.Username) {
		return vo.Fail(vo.WithMessage("用户名已被注册"), vo.WithCode(500))
	}
	// 注册用户
	if err := u.service.Register(user); err != nil {
		return vo.Fail(vo.WithMessage("注册失败"), vo.WithCode(500))
	}
	return vo.Success(vo.WithMessage("注册成功"), vo.WithCode(200))

}

func (u UserController) GetLogin(ctx iris.Context) *vo.Result {
	// TODO implement me
	// panic("implement me")
	return vo.Success(vo.WithMessage("登录成功"), vo.WithCode(200))
}
