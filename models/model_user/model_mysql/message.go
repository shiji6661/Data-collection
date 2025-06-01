package model_mysql

import (
	"common/global"
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	SendId     int32  `gorm:"column:send_id;type:int;comment:发送人id;not null;" json:"send_id"`         // 发送人id
	ReceiverId int32  `gorm:"column:receiver_id;type:int;comment:接收人id;not null;" json:"receiver_id"` // 接收人id
	Context    string `gorm:"column:context;type:varchar(255);comment:内容;not null;" json:"context"`   // 内容
}

func (m *Message) TableName() string {
	return "message"
}

func (m *Message) UserSendToReceiver() error {
	return global.DB.Create(&m).Error
}

func FindMessage(receiveId int) (result []*Message, err error) {
	err = global.DB.Where("receiver_id = ?", receiveId).Find(&result).Error
	return
}
