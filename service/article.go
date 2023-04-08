package service

import (
	"blog/dao/redis"
	"blog/middleware/logger"
	. "blog/model/entity"
)

func (a articleServiceImpl) GetsArticleList(cid, mid, tid string, size,
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
	}
	return result, count
}

func (a articleServiceImpl) GetArticle(id uint64) *Article {
	return a.iDao.GetArticleByID(id)
}
