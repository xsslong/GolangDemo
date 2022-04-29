package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"fmt"
)

// 页面表
type Page struct {
	gorm.Model
	Title       string // 标题
	Body        string // 文体内容+
	View        int    // 浏览次数
	IsPublished bool   // 是否发布
}

var GDB *gorm.DB

func main() {
	InitDB()
	//// 增
	//insertPage := &Page{
	//	Title:       "双11抢购啦",
	//	Body:        "小米MIX2只要998",
	//	View:        10086,
	//	IsPublished: true,
	//}
	//insertError := insertPage.Insert()
	//if insertError == nil {
	//	fmt.Println("插入成功")
	//}
	//
	//// 改
	//modifyPage := &Page{
	//	Title:       "双11抢购啦",
	//	Body:        "Iphone8999",
	//	View:        8888,
	//	IsPublished: true,
	//}
	//modifyError := modifyPage.Update()
	//if modifyError == nil {
	//	fmt.Println("修改成功")
	//}
	//
	// 删
	//deletePage := &Page{
	//	Title:       "双11抢购啦",
	//	Body:        "Iphone8999",
	//	View:        8888,
	//	IsPublished: true,
	//}
	//deleteError := deletePage.Delete()
	//if deleteError == nil {
	//	fmt.Println("删除成功")
	//}

	// 查
	var page Page
	GDB.First(&page, 1) // 查询id为1的product
	fmt.Println("结果是:", page)
	GDB.First(&page, "title = ?", "双11抢购啦") // 查询code为l1212的product
	fmt.Println("结果是:", page)
	//p, result := Quary(8888)
	//if result == nil {
	//	fmt.Print("结果是:", result)
	//} else {
	//	fmt.Print("无法查询:", result)
	//}
	//p.Delete()
}

// 初始化数据库
func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", "gorm_test.db")
	if err != nil {
		panic("连接数据库失败")
		return nil, err
	}
	//defer db.Close()

	// 自动迁移模式
	db.AutoMigrate(&Page{})
	GDB = db
	return db, err
}

//  增, Page表
func (page *Page) Insert() error {
	fmt.Println(page)
	return GDB.Create(page).Error
}

// 删, Page表
func (page *Page) Delete() error {
	return GDB.Delete(page).Error
}

// 改, Page表
func (page *Page) Update() error {
	return GDB.Model(page).Update(map[string]interface{}{
		"title":        page.Title,       // 标题
		"body":         page.Body,        // 文体内容
		"view":         page.View,        // 浏览次数
		"is_published": page.IsPublished, // 是否发布
	}).Error
}

// 查找, Page表
func Quary(view int) (*Page, error) {
	var page Page
	err := GDB.First(&page, "view = ?", view).Error
	return &page, err
}
