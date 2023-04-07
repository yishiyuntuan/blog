package mapper_impl

import (
	"blog/dao/gen"
	"blog/dao/mapper"
	"blog/logger"
	"blog/model/entity"

	"golang.org/x/sync/singleflight"
)

type userDao struct {
	dao *gen.UserExec
	// cache cache.ArticleCache
	sfg *singleflight.Group
}

// NewUserDao creating the dao interface
func NewUserDao(dao *gen.UserExec) mapper.UserDao {
	return &userDao{
		dao: dao,
		sfg: new(singleflight.Group)}
}

func (u userDao) IsExist(username string) bool {
	first, err := u.dao.Where(u.dao.Username.Eq(username)).First()
	if err != nil || first == nil {
		return false
	}
	logger.Log.Debug(first) //nolint:wsl
	return true             //nolint:wsl
}

func (u userDao) Register(user *entity.User) error {
	err := u.dao.Create(user)
	if err != nil {
		logger.Log.Debug(err)
		return err
	}
	return nil
}
