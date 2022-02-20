package model

import "golang.org/x/net/context"

type MsgModel interface {
	Persist(context context.Context, msg interface{}) bool
	UpdateDbContent(context context.Context, msgIface interface{}) bool
}
