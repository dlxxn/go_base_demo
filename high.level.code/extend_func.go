package high_level_code

import (
	"fmt"
)

/**
Golang 里面没有 C++ 、Java 那种继承的实现方式，但是，我们可以通过 Golang 的匿名组合来实现继承，这里要注意，
这个是实际编程中经常用到的一种姿势。具体实现就是一个 struct 里面包含一个匿名的 struct，也就是通过匿名组合，
这最基础的基类就是一个 struct 结构，然后定义相关成员变量，然后再定义一个子类，也是一个 struct，里面包含前面的 struct，
即可实现继承。
*/

// 【基类】
//定义一个最基础的struct类MsgModel，里面包含一个成员变量msgId
type MsgModel struct {
	msgId   int
	msgType int
}

// MsgModel的一个成员方法，用来设置msgId
func (msg *MsgModel) SetId(msgId int) {
	msg.msgId = msgId
}

func (msg *MsgModel) SetType(msgType int) {
	msg.msgType = msgType
}

//【子类】
// 再定义一个struct为GroupMsgModel，包含了MsgModel，即组合，但是并没有给定MsgModel任何名字，因此是匿名组合
type GroupMsgModel struct {
	MsgModel

	// 如果子类也包含一个基类的一样的成员变量，那么通过子类设置和获取得到的变量都是基类的
	msgId int

	msgDes string
}

func (group *GroupMsgModel) GetId() int {
	return group.msgId
}

func (msg *GroupMsgModel) GetDes(msgDes string) {
	msg.msgDes = msgDes
}

func (msg *GroupMsgModel) SetDes(msgDes string) {
	msg.msgDes = msgDes
}

/*
func (group *GroupMsgModel) SetId(msgId int) {
 group.msgId = msgId
}
*/

func Extend() {

	group := &GroupMsgModel{}

	group.SetId(123)
	group.SetType(1)
	group.SetDes("test group des")

	fmt.Println("group.msgId =", group.msgId, "\tgroup.MsgModel.msgId =", group.MsgModel.msgId)
	fmt.Println("group.msgType =", group.msgType, "\tgroup.MsgModel.msgType =", group.MsgModel.msgType)
	fmt.Println("group.MsgModel.msgDes =", group.msgDes)

}
