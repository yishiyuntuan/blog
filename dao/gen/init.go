package gen

import "blog/db"

func init() {
	SetDefault(db.GetDB().Debug())
}
