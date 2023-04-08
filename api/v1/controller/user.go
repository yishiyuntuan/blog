package controller

import (
	"blog/conts"
	"blog/dao/gen"
	mapper_impl "blog/dao/mapper"
	"blog/middleware/logger"
	"blog/model/entity"
	"blog/model/vo"
	"blog/service"
	"blog/util/tool"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/core/router"
	"github.com/kataras/iris/v12/mvc"
)

func (u UserController) Configure(r router.Party) {
	userService := service.NewUserService(service.WithUserDao(mapper_impl.NewUserDao(gen.User)))
	r.RegisterDependency(userService)
	// 依赖注入
	mvc.Configure(r, func(app *mvc.Application) {
		app.Handle(&u)
	})
}

// 注册用户
// @Summary 注册用户
// @Param   username    query    string  true        "用户名"
// @Param   password    query    string  true        "密码"
// @Router /user/register [post]
func (u UserController) PostRegister(ctx iris.Context, user entity.User) *vo.Result {
	logger.Log.Debug(user)
	// 校验参数
	if !tool.VerifyFormat(conts.VERIFY_EXP_USERNAME, user.Username) || !tool.VerifyFormat(conts.VERIFY_EXP_PASSWORD, string(user.Password)) {
		return vo.Fail(vo.WithMessage("用户名或密码格式不正确"), vo.WithCode(400))
	}
	// 判断用户是否已经存在
	if u.Service.IsExist(user.Username) {
		return vo.Fail(vo.WithMessage("用户名已被注册"), vo.WithCode(500))
	}
	// 注册用户
	if err := u.Service.Register(user); err != nil {
		return vo.Fail(vo.WithMessage("注册失败"), vo.WithCode(500))
	}
	return vo.Success(vo.WithMessage("注册成功"), vo.WithCode(200))

}

func (u UserController) GetLogin(ctx iris.Context) *vo.Result {
	// TODO implement me
	// panic("implement me")
	return vo.Success(vo.WithMessage("登录成功"), vo.WithCode(200))
}
