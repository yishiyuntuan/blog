package service

import (
	"blog/dao/mapper"
	"blog/dao/redis"
	"blog/logger"
	. "blog/model/entity"
)

type ArticleService interface {
	GetsArticleList(cid, mid, tid string, size, num int) ([]*Article, int64)
}

type articleServiceImpl struct {
	iDao mapper.ArticleDao
}

type Option func(opts *articleServiceImpl)

func NewArticleService(opt ...Option) ArticleService {
	a := &articleServiceImpl{}
	for _, f := range opt {
		f(a)
	}
	return a
}

func WithArticleDao(dao mapper.ArticleDao) Option {
	return func(opts *articleServiceImpl) {
		opts.iDao = dao
	}
}

func (a articleServiceImpl) GetsArticleList(cid, mid, tid string, size,
	num int) ([]*Article, int64) {

	// a.iDao.GetList()
	// 存在缓存中的json数据
	// var dataJson []byte
	// 缓存中的格式
	// var data Coding
	// 需要返回的Articles
	// res := make([]Article, 0)

	if cid != "" {
		res := redis.Get[[]*Article]("cid:" + cid)
		if res == nil {
			result, count := a.iDao.GetList(cid, mid, tid, size, num)

			// str, _ := json.Marshal(result)
			set := redis.Set[[]*Article]("cid:"+cid, result)
			cmd := redis.Set[int64]("count:"+cid, count)
			logger.Log.Debug("存redis", set, cmd)
			return result, count
		}
		get := redis.Get[int64]("count:" + cid)
		return *res, *get

	}
	return nil, 0

	// 根据key和条件，存进缓存
	// articlesHset := func(key string, where map[string]interface{}) {
	//	Db.
	//		Model(Article{}).
	//		Where(where).
	//		Order("created_at desc").
	//		Pluck("id", &data.Ids).
	//		Count(&data.Total)
	//	dataJson, _ = json.Marshal(data)
	//	db.Rdb.HSet(ctx, "articles", key, dataJson)
	// }
	// 根据分类或者菜单获取文章列表
	// if cid != 0 {
	//	cidKey := fmt.Sprintf("cid:%d", cid)
	//	dataJson, err = db.Rdb.HGet(ctx, "articles", cidKey).Bytes()
	//	if err != nil {
	//		articlesHset(cidKey, map[string]interface{}{"cid": cid})
	//	}
	// } else if tid != 0 {
	//	tidKey := fmt.Sprintf("tid:%d", tid)
	//	dataJson, err = db.Rdb.HGet(ctx, "articles", tidKey).Bytes()
	//	if err != nil {
	//		var tdata Tags
	//		Db.Preload("Article.Tags").Take(&tdata, tid)
	//		for _, v := range tdata.Article {
	//			data.Ids = append(data.Ids, strconv.Itoa(int(v.ID)))
	//		}
	//		dataJson, _ = json.Marshal(data)
	//		db.Rdb.HSet(ctx, "articles", tidKey, dataJson)
	//	}
	// } else {
	//	midKey := fmt.Sprintf("mid:%d", mid)
	//	dataJson, err = db.Rdb.HGet(ctx, "articles", midKey).Bytes()
	//	if err != nil {
	//		if mid == -2 {
	//			articlesHset(midKey, nil)
	//		} else {
	//			articlesHset(midKey, map[string]interface{}{"cid": GetMidCid(mid)})
	//		}
	//	}
	// }

	// _ = json.Unmarshal(dataJson, &data)
	// //防止越界
	// ids := tool.PageIds(pageNum, pageSize, data.Ids)
	// for _, v := range ids {
	//	var d Articles
	//	var dd Article
	//	articleJson, err := db.Rdb.HGet(ctx, "article", v).Bytes()
	//	articleUv, _ := db.Rdb.PFCount(ctx, fmt.Sprintf("article/uv/aid:%s;", v)).Result()
	//	if err != nil {
	//		Db.Preload("Tags").Take(&dd, v)
	//		articleJson, _ = json.Marshal(dd)
	//		db.Rdb.HSet(ctx, "article", v, articleJson)
	//	} else {
	//		_ = json.Unmarshal(articleJson, &dd)
	//	}
	//	//优化，列表不返回文章内容
	//	d = Articles{dd.ID, dd.CreatedAt, dd.UpdatedAt, dd.Title, dd.Img, dd.Desc, dd.Cid, dd.Tags, articleUv}
	//	res = append(res, d)
	// }

	// return res, data.Total
}
