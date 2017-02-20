package rpc

import (
	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/proto/pb/coredef"
)

type UnboxHandler struct {
	cellnet.BaseEventHandler

	feedbackHandler cellnet.EventHandler
}

func (self *UnboxHandler) Call(ev *cellnet.SessionEvent) error {

	wrapper := ev.Msg.(*coredef.RemoteCallACK)

	ev.MsgID = wrapper.MsgID
	ev.Data = wrapper.Data

	ev.SendHandler = self.feedbackHandler

	return self.CallNext(ev)
}

func NewUnboxHandler(feedbackHandler cellnet.EventHandler) cellnet.EventHandler {
	return &UnboxHandler{
		feedbackHandler: feedbackHandler,
	}

}
