package model_mysql

import (
	"common/global"
	"errors"
	"time"
)

type Member struct {
	Id        int64     `gorm:"column:id;type:int UNSIGNED;primaryKey;not null;" json:"id"`
	Uid       int64     `gorm:"column:uid;type:int;comment:用户id;not null;" json:"uid"`                                                      // 用户id
	UserName  string    `gorm:"column:user_name;type:varchar(50);comment:用户名;not null;" json:"user_name"`                                   // 用户名
	Points    int64     `gorm:"column:points;type:int;comment:积分;not null;" json:"points"`                                                  // 积分
	LevelId   int64     `gorm:"column:level_id;type:int;comment:等级;not null;" json:"level_id"`                                              // 等级
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime(3);comment:最后更新时间;not null;default:CURRENT_TIMESTAMP(3);" json:"updated_at"` // 最后更新时间
}

func (m Member) Error() string {
	//TODO implement me
	panic("implement me")
}

func (Member) TableName() string {
	return "member"
}

func (m *Member) GetMemberInfo(uid int) (me *Member, err error) {
	err = global.DB.Debug().Where("uid = ?", uid).Limit(1).Find(&me).Error
	if err != nil {
		return nil, err
	}
	return me, nil
}

func (m *Member) CreateMemberInfo(uid int64, points int64, level int64, now time.Time) error {
	var userName string
	err := global.DB.Debug().Table("user").Select("user_name").Where("id = ?", uid).Scan(&userName).Error
	if err != nil {
		return err
	}
	me := &Member{
		Uid:       uid,
		UserName:  userName,
		Points:    points,
		LevelId:   level,
		UpdatedAt: now,
	}
	return global.DB.Debug().Create(&me).Error
}

func (m *Member) UpdateMemberInfo(uid int64, level int64, points int64, now time.Time) error {
	if level < 0 || level > 5 {
		return errors.New("invalid level range")
	}
	err := global.DB.Debug().Raw(`
        UPDATE member 
        SET level_id = ?, 
            points = ?, 
            updated_at = ? 
        WHERE uid = ?`,
		level, points, now, uid).Scan(&m).Error
	return err
}

func (m *Member) UpdateMemberPoints(uid int, points int64) error {
	err := global.DB.Debug().Model(&Member{}).Where("uid = ?", uid).Update("points", points).Error
	if err != nil {
		return err
	}
	return nil
}

// 添加消费积分记录
func (m *Member) AddConsumptionPointsRecord(uid int64, points int64, orderId string, amount float64) error {
	record := &MemberPointsRecord{
		Uid:       int32(uid),
		Points:    int32(points),
		Type:      1, // 消费类型
		OrderId:   orderId,
		Amount:    amount,
		CreatedAt: time.Now(),
	}
	return global.DB.Debug().Create(&record).Error
}

// 添加邀请积分记录
func (m *Member) AddInvitationPointsRecord(uid int64, invitedUid int64) error {
	record := &MemberPointsRecord{
		Uid:        int32(uid),
		Points:     global.POINTS_PER_INVITATION,
		Type:       2, // 邀请类型
		InvitedUid: int32(invitedUid),
		CreatedAt:  time.Now(),
	}
	return global.DB.Debug().Create(&record).Error
}

// 根据消费金额增加会员积分
func (m *Member) AddPointsByConsumption(uid int64, amount float64, orderId string) error {
	pointsToAdd := int64(amount / global.CONSUMPTION_UNIT)
	if pointsToAdd <= 0 {
		return nil
	}

	// 开启事务
	tx := global.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 更新会员积分
	member, err := m.GetMemberInfo(int(uid))
	if err != nil {
		tx.Rollback()
		return err
	}
	if member == nil {
		tx.Rollback()
		return errors.New("会员信息不存在")
	}

	newPoints := member.Points + pointsToAdd
	if err = tx.Exec(`
        UPDATE member 
        SET points = ? 
        WHERE uid = ?`,
		newPoints, uid).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 添加积分记录
	if err = m.AddConsumptionPointsRecord(uid, pointsToAdd, orderId, amount); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// 邀请注册获得积分
func (m *Member) AddPointsByInvitation(uid int64, invitedUid int64) error {
	// 开启事务
	tx := global.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 更新会员积分
	member, err := m.GetMemberInfo(int(uid))
	if err != nil {
		tx.Rollback()
		return err
	}
	if member == nil {
		tx.Rollback()
		return errors.New("会员信息不存在")
	}

	newPoints := member.Points + global.POINTS_PER_INVITATION
	if err = tx.Exec(`
        UPDATE member 
        SET points = ? 
        WHERE uid = ?`,
		newPoints, uid).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 添加积分记录
	if err = m.AddInvitationPointsRecord(uid, invitedUid); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
