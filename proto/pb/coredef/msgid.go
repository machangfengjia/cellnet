// Generated by github.com/davyxu/cellnet/protoc-gen-msg
// DO NOT EDIT!
// Source: coredef.proto

package coredef

import (
	"github.com/davyxu/cellnet"
	"reflect"
	_ "github.com/davyxu/cellnet/codec/pb"
)

func init() {

	// coredef.proto
	cellnet.RegisterMessageMeta("pb", "coredef.SessionAccepted", reflect.TypeOf((*SessionAccepted)(nil)).Elem(), 3495179174)
	cellnet.RegisterMessageMeta("pb", "coredef.SessionConnected", reflect.TypeOf((*SessionConnected)(nil)).Elem(), 3551021301)
	cellnet.RegisterMessageMeta("pb", "coredef.SessionAcceptFailed", reflect.TypeOf((*SessionAcceptFailed)(nil)).Elem(), 3277953230)
	cellnet.RegisterMessageMeta("pb", "coredef.SessionConnectFailed", reflect.TypeOf((*SessionConnectFailed)(nil)).Elem(), 3980285497)
	cellnet.RegisterMessageMeta("pb", "coredef.SessionClosed", reflect.TypeOf((*SessionClosed)(nil)).Elem(), 3480086952)
	cellnet.RegisterMessageMeta("pb", "coredef.RemoteCallACK", reflect.TypeOf((*RemoteCallACK)(nil)).Elem(), 2811469770)

}
