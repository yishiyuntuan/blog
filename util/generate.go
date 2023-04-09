package main

import (
	"blog/db"
	"fmt"
	"strings"

	"github.com/goccy/go-json"
	"gorm.io/gen"
	"gorm.io/gorm"
)

type CommonMethod struct {
}

// 序列化
func (c CommonMethod) MarshalBinary() (data []byte, err error) {
	fmt.Println("MarshalBinary")
	return json.Marshal(c)
}

// 反序列化
func (c CommonMethod) UnmarshalBinary(data []byte) error {
	fmt.Println("UnmarshalBinary")
	return json.Unmarshal(data, c)

}
func GenerateTableStruct(db *gorm.DB) {
	// 根据配置实例化 gen
	g := gen.NewGenerator(gen.Config{
		// 文件生成路径
		ModelPkgPath: "model/entity",
		OutPath:      "dao/gen",
		// OutFile:      "auto.go",
		// WithDefaultQuery 生成默认查询结构体(作为全局变量使用), 即`Q`结构体和其字段(各表模型)
		// WithoutContext 生成没有context调用限制的代码供查询
		// WithQueryInterface 生成interface形式的查询代码(可导出), 如`Where()`方法返回的就是一个可导出的接口类型
		Mode: gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,

		// 表字段可为 null 值时, 对应结体字段使用指针类型
		FieldNullable: false, // generate pointer when field is nullable

		// 表字段默认值与模型结构体字段零值不一致的字段, 在插入数据时需要赋值该字段值为零值的, 结构体字段须是指针类型才能成功, 即`FieldCoverable:true`配置下生成的结构体字段.
		// 因为在插入时遇到字段为零值的会被GORM赋予默认值. 如字段`age`表默认值为10, 即使你显式设置为0最后也会被GORM设为10提交.
		// 如果该字段没有上面提到的插入时赋零值的特殊需要, 则字段为非指针类型使用起来会比较方便.
		FieldCoverable: true, // generate pointer when field has default value,
		// to fix problem zero value cannot be assign: https://gorm.io/docs/create.html#Default-Values

		// 模型结构体字段的数字类型的符号表示是否与表字段的一致, `false`指示都用有符号类型
		FieldSignable: true, // detect integer field's unsigned type,
		// adjust generated data type
		// 生成 gorm 标签的字段索引属性
		FieldWithIndexTag: true, // generate with gorm index tag
		// 生成 gorm 标签的字段类型属性
		FieldWithTypeTag: true, // generate with gorm column type tag
		WithUnitTest:     false,
	})

	// 使用数据库的实例
	g.UseDB(db)
	// 模型结构体的命名规则
	g.WithModelNameStrategy(func(tableName string) (modelName string) {
		fmt.Println("xxxxxxxxx", tableName)
		if strings.HasSuffix(tableName, ".gen") {
			tableName = tableName[:len(tableName)-4]
		}

		if strings.HasPrefix(tableName, "tb_") {
			return firstUpper(tableName[3:])
		}
		if strings.HasPrefix(tableName, "table") {
			return firstUpper(tableName[5:])
		}
		return firstUpper(tableName)
	})
	// 模型文件的命名规则
	g.WithFileNameStrategy(func(tableName string) (fileName string) {
		fmt.Println(">>>>", tableName)
		if strings.HasPrefix(tableName, "tb_") {
			return firstLower(tableName[3:])
		}
		if strings.HasPrefix(tableName, "table") {
			return firstLower(tableName[5:])
		}
		return tableName
	})

	// 数据类型映射
	dataMap := map[string]func(detailType gorm.ColumnType) (dataType string){
		"int": func(detailType gorm.ColumnType) (dataType string) {
			//if strings.Contains(detailType, "unsigned") {
			//   return "uint64"
			//}
			return "int64"
		},
		"bigint": func(detailType gorm.ColumnType) (dataType string) {
			//if strings.Contains(detailType, "unsigned") {
			//   return "uint64"
			//}
			return "int64"
		},
	}

	// 使用上面的类型映射
	// 要先于`ApplyBasic`执行
	g.WithDataTypeMap(dataMap)

	//g.ApplyBasic(g.GenerateAllTable()...)
	g.ApplyBasic(
		// 密码字段属性为[]byte
		g.GenerateModelAs("user", "User", gen.FieldType("password", "[]byte")),
		g.GenerateModelAs("article", "Article"),
		g.GenerateModelAs("article_path", "AriclePath"),
		g.GenerateModelAs("article_tags", "ArticleTags"),
		g.GenerateModelAs("category", "Category"),
		g.GenerateModelAs("menuchild", "MenuChild"),
		g.GenerateModelAs("tags", "Tags"),
	)

	// 执行
	g.Execute()
}

// 字符串第一位改成小写
func firstLower(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToLower(s[:1]) + s[1:]
}

// 字符串第一位改成大写
func firstUpper(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(s[:1]) + s[1:]
}
func main() {
	// gormdb, _ := gorm.Open(mysql.Open("root:123456@(127.0.0.1:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local"))

	GenerateTableStruct(db.GetDB())
}
