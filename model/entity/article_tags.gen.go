// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

const TableNameArticle_tags = "article_tags"

// Article_tags mapped from table <article_tags>
type Article_tags struct {
	TagsID    uint64 `gorm:"column:tags_id;type:bigint unsigned;primaryKey" json:"tags_id"`
	ArticleID uint64 `gorm:"column:article_id;type:bigint unsigned;primaryKey;index:fk_article_tags_article,priority:1" json:"article_id"`
}

// TableName Article_tags's table name
func (*Article_tags) TableName() string {
	return TableNameArticle_tags
}