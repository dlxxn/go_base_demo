package impl

import "golang.org/x/net/context"

type msgModelImpl struct{}

var MsgModelImpl = msgModelImpl{}

func (m msgModelImpl) Persist(context context.Context, msgIface interface{}) bool {
	// 具体实现
	return false
}

func (m msgModelImpl) UpdateDbContent(context context.Context, msgIface interface{}) bool {
	// 具体实现
	return false
}
