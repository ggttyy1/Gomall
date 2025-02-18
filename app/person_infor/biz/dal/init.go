package dal

import (
	"github.com/cloudwego/biz-demo/gomall/app/person_infor/biz/dal/mysql"
	"github.com/cloudwego/biz-demo/gomall/app/person_infor/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
