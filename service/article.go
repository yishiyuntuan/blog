package service

import (
	"blog/config"
	"blog/dao/mapper"
	"blog/dao/redis"
	"blog/middleware/logger"
	"blog/model/dto"
	. "blog/model/entity"
	"blog/util/tool"
	"strconv"
)

type ArticleServiceImpl struct {
	iDao mapper.ArticleDao
}

func WithArticleDao(dao mapper.ArticleDao) Option {
	return func(opts any) {
		s, ok := opts.(*ArticleServiceImpl)
		if ok {
			s.iDao = dao
		}
	}
}

func (a ArticleServiceImpl) GetsArticleList(cid, mid, tid string, size,
	num int) ([]*Article, int64) {
	result := make([]*Article, 0)
	var count int64
	if cid != "" {
		res := redis.Get[[]*Article]("cid:" + cid)
		if res == nil {
			result, count = a.iDao.GetListByCid(cid, size, num)
			// str, _ := json.Marshal(result)
			set := redis.Set[[]*Article]("cid:"+cid, result)
			cmd := redis.Set[int64]("count:"+cid, count)
			logger.Log.Debug("存redis", set, cmd)
		} else {
			result = *res
			n := redis.Get[int64]("count:" + cid)
			if n != nil {
				count = *n
			}
		}
	} else if tid != "" {
		res := redis.Get[[]*Article]("tid:" + tid)
		if res == nil {
			result, count = a.iDao.GetListByTid(tid, size, num)
			// str, _ := json.Marshal(result)
			set := redis.Set[[]*Article]("tid:"+tid, result)
			cmd := redis.Set[int64]("count:"+tid, count)
			logger.Log.Debug("存redis", set, cmd)
		} else {
			result = *res
			n := redis.Get[int64]("count:" + tid)
			if n != nil {
				count = *n
			}
		}
	} else if mid != "" && mid != "0" {
		res := redis.Get[[]*Article]("mid:" + mid)
		if res == nil {
			result, count = a.iDao.GetListByMid(mid, size, num)
			// str, _ := json.Marshal(result)
			set := redis.Set[[]*Article]("mid:"+mid, result)
			cmd := redis.Set[int64]("count:"+mid, count)
			logger.Log.Debug("存redis", set, cmd)
		} else {
			result = *res
			n := redis.Get[int64]("count:" + mid)
			if n != nil {
				count = *n
			}
		}
	} else {
		result, count = a.iDao.GetListAll(size, num)
	}
	return result, count
}

func (a ArticleServiceImpl) GetArticle(id uint64) *dto.ArticleContent {
	// 首先查询缓存
	cache := redis.Get[dto.ArticleContent]("article:" + strconv.FormatUint(id, 10))
	// 有缓存，直接返回
	if cache != nil {
		return cache
	}
	// 无缓存，查询数据库,获取文件路径
	// 1.查询文章路径
	path := a.iDao.GetArticlePathByID(id)

	// 2.进入本地git仓库
	// 3. go-git使用git pull拉取最新的文章

	// 4.有更新，更新本地仓库和数据库
	logger.Log.Debug("文章路径：" + config.BLOG_ARTICLE + path)
	content := tool.ReadFile(config.BLOG_ARTICLE + path)
	info := a.GetArticleInfo(id)
	articleContent := dto.ArticleContent{
		Article: *info,
		Content: content,
	}
	redis.Set("article:"+strconv.FormatUint(id, 10), articleContent)
	// 5.返回文章

	return &articleContent
}

// 获取文章信息
func (a ArticleServiceImpl) GetArticleInfo(id uint64) *Article {
	return a.iDao.GetArticleInfoByID(id)
}
