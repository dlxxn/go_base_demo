package orm

import "github.com/jinzhu/gorm"

var db *gorm.DB

func Conn() {
	/*db, err = gorm.Open("mysql", "<user>:<password>/<database>?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}*/
}
