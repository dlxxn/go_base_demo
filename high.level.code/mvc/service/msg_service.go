package service

import (
	"go_base_demo/high.level.code/mvc/impl"
	"go_base_demo/high.level.code/mvc/model"
)

// 定义一个msgService struct包含了model里面的UserModel和MsgModel两个model
type msgService struct {
	msgModel model.MsgModel
}

// 定义一个MsgService的变量，并初始化，这样通过MsgService，就能引用并访问model的所有方法
var (
	MsgService = msgService{
		msgModel: impl.MsgModelImpl,
	}
)
