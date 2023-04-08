package service

import "blog/dao/mapper"

type articleServiceImpl struct {
	iDao mapper.ArticleDao
}

func NewArticleService(opt ...Option) ArticleService {
	a := &articleServiceImpl{}
	for _, f := range opt {
		f(a)
	}
	return a
}

func WithArticleDao(dao mapper.ArticleDao) Option {
	return func(opts any) {
		s, ok := opts.(*articleServiceImpl)
		if ok {
			s.iDao = dao
		}
	}
}

type userServiceImpl struct {
	iDao mapper.UserDao
}

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
