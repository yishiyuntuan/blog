package service

import (
	"blog/middleware/logger"
	"blog/model/entity"
	"time"

	"github.com/bwmarrin/snowflake"
	"golang.org/x/crypto/bcrypt"
)

func (u userServiceImpl) IsExist(username string) bool {
	return u.iDao.IsExist(username)
}

func (u userServiceImpl) Register(user entity.User) error { // password 加密
	hash, err := bcrypt.GenerateFromPassword(user.Password, bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = hash
	// id生成
	node, err := snowflake.NewNode(1)
	if err != nil {
		logger.Log.Error(err)
		return err
	}
	user.ID = uint64(node.Generate())
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	// todo: 其他字段

	// 保存用户信息
	return u.iDao.Register(&user)
}
