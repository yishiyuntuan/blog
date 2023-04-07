package service

import (
	"blog/dao/mapper"
	"blog/model/entity"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type userServiceImpl struct {
	iDao mapper.UserDao
}

// type Option func(*userServiceImpl)

// NewUserService 创建用户服务
func NewUserService(opt ...Option) UserService {
	u := &userServiceImpl{}
	for _, f := range opt {
		f(u)
	}
	return u
}

func WithUserDao(iDao mapper.UserDao) Option {
	return func(u any) {
		impl, ok := u.(*userServiceImpl)
		if ok {
			impl.iDao = iDao
		}
	}
}

func (u userServiceImpl) IsExist(username string) bool {
	return u.iDao.IsExist(username)
}

func (u userServiceImpl) Register(username string, password []byte) error {

	// password 加密
	password = append(password, []byte("yishiyuntuan")...)
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	// 构造用户结构体
	user := entity.User{
		ID:        2222,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Username:  username,
		Password:  hash,
		Relation:  "关系",
		Role:      -1,
		Avatar:    "nil",
		NickName:  "nil",
	}
	// 保存用户信息
	return u.iDao.Register(&user)
}
