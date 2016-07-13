package igetui

import "github.com/zhaijian/getui/template"

type IGtMessage struct {
	IsOffline         bool
	OfflineExpireTime int32
	Data              template.ITemplate
}

type IGtSingleMessage struct {
	IGtMessage
}

func NewIGtSingleMessage(isoffline bool, offlineexpiretime int32, templatee template.ITemplate) *IGtSingleMessage {
	return &IGtSingleMessage{
		IGtMessage: IGtMessage{
			IsOffline:         isoffline,
			OfflineExpireTime: offlineexpiretime,
			Data:              templatee,
		},
	}
}

type IGtListMessage struct {
	IGtMessage
}

func NewIGtListMessage(isoffline bool, offlineexpiretime int32, templatee template.ITemplate) *IGtListMessage {
	return &IGtListMessage{
		IGtMessage: IGtMessage{
			IsOffline:         isoffline,
			OfflineExpireTime: offlineexpiretime,
			Data:              templatee,
		},
	}
}

type IGtAppMessage struct {
	IGtMessage
	AppIdList     []string
	PhoneTypeList []string
	ProvinceList  []string
}

func NewIGtAppMessage(isoffline bool, offlineexpiretime int32, templatee template.ITemplate) *IGtAppMessage {
	return &IGtAppMessage{
		IGtMessage: IGtMessage{
			IsOffline:         isoffline,
			OfflineExpireTime: offlineexpiretime,
			Data:              templatee,
		},
	}
}
