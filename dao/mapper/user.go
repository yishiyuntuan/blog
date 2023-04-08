package mapper

import (
	"blog/middleware/logger"
	"blog/model/entity"
)

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
