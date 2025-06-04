package model_mysql

import (
	"common/global"
	"gorm.io/gorm"
	"sort"
	"user_srv/proto_user/user"
)

type User struct {
	gorm.Model
	UserName           string  `gorm:"type:varchar(20);comment:'用户名'"`
	UserPassword       string  `gorm:"type:varchar(100);comment:'密码'"`
	UserPhone          string  `gorm:"type:char(11);comment:'手机号码'"`
	UserEmail          string  `gorm:"type:varchar(50);comment:'用户邮箱'"`
	UserState          int64   `gorm:"type:bigint(1);comment:'用户状态 0异常 1正常';default:1"`
	InviteCode         string  `gorm:"type:char(8);comment:'用户的邀请码'"`
	UserInviteId       int64   `gorm:"type:int;comment:'邀请人id'"`
	UserVip            int64   `gorm:"type:tinyint(1);comment:'是否会员 0否 1是';default:0"`
	UserPid            int64   `gorm:"type:int;comment:''用户的上级id'分销'"`
	UserPoint          int64   `gorm:"type:int;comment:'用户积分'"`
	SubUserPrice       float64 `gorm:"type:decimal(10,2);comment:''拉人奖励 提成'分销'"`
	SubCode            string  `gorm:"type:char(8);comment:'分销邀请码'"`
	UserAddress        string  `gorm:"type:varchar(200);comment:'用户地址'"`
	ConsecutiveSignIns int32   `gorm:"column:consecutive_sign_ins;type:int;comment:连续签到次数;not null;default:0;"` // 连续签到次数
}

func (u *User) TableName() string {
	return "user"
}
func (u *User) Register() error {
	return global.DB.Create(&u).Error
}

func (u *User) Login(account string) error {
	return global.DB.Where("account = ?", account).Find(&u).Error
}

// todo: 根据用户id查找用户
func (u *User) FindUserById(userId int64) error {
	return global.DB.Debug().Where("id=?", userId).Find(&u).Error
}

func (u *User) FindUserByIdBargain(userId int64) (use *User, err error) {
	err = global.DB.Where("id = ?", userId).Find(&use).Error
	return
}

// todo 查找用户是否存在
func (u *User) FindUser(username string) (users *User, err error) {
	err = global.DB.Debug().Model(&User{}).Where("user_name = ?", username).Limit(1).Find(&u).Error
	if err != nil {
		return nil, err
	}
	return u, nil

}

// todo:查找该邀请码是否真实
func (u *User) FindUserByInviteCode(inviteCode string) (user *User, err error) {
	err = global.DB.Where("invite_code=?", inviteCode).Find(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

// TODO 事务
func (u *User) BeginTransaction() *gorm.DB {
	return global.DB
}

// TODO 创建新用户
func (u *User) CreateUser() error {
	return global.DB.Create(&u).Error
}

// todo 更新 MySQL 中的积分
func (u *User) UpdateUserPoints(userId, point int64) (err error) {
	err = global.DB.Model(&u).Where("id = ?", userId).Update("user_point", point).Error
	if err != nil {
		return err
	}
	return nil
}

// todo: 查询用户邮箱是否存在
func (u *User) FindUserByEmail(email string) (user *User, err error) {
	err = global.DB.Where("user_email=?", email).Limit(1).Find(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

// todo: 用户手机号登录
func (u *User) FindUserByPhone(userPhone string) (user *User, err error) {
	err = global.DB.Where("user_phone=?", userPhone).Limit(1).Find(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

// todo:查看个人信息
func (u *User) GetUserInfo(id int64) (user *user.UserInfoResponse, err error) {
	err = global.DB.Debug().Raw("SELECT * FROM `user` JOIN user_detail ON `user`.id = user_detail.id WHERE `user`.id = ?", id).Scan(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

// todo:查询用户表和地址表
func (u *User) FindUserModify(id int64) (user *User, err error) {
	err = global.DB.Debug().Raw("SELECT * FROM `user` JOIN address ON `user`.id = address.user_id WHERE `user`.id = ?", id).Scan(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *User) GetUserUseRights(uid int64) (list []*user.UseRights, err error) {
	err = global.DB.Debug().Raw("SELECT u.user_name,r.rights_name,r.img,r.`explain`,ur.created_at FROM  `user` u JOIN user_rights ur ON u.Id = ur.Uid JOIN rights r ON ur.rights_id = r.Id WHERE u.Id = ?", uid).Scan(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (u *User) UpdateUserIntegral(id int64, point int64) error {
	return global.DB.Debug().Model(&User{}).Where("id = ?", id).Update("user_point", gorm.Expr("user_point + ?", point)).Error

}

func (u *User) RestoreUserSignInTimes(context int64, uid int64) error {
	err := global.DB.Model(&User{}).Where("id = ?", uid).Update("consecutive_sign_ins", context).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *User) CommissionList(page, size int) (list []*User, err error) {
	offset := (page - 1) * size
	err = global.DB.Order("sub_user_price desc").Offset(offset).Limit(size).Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

// 邀请用户排行
func (u *User) InviteUserList() (list []*user.InviteUserList, err error) {
	var us []User
	err = global.DB.Find(&us).Error
	if err != nil {
		return nil, err
	}
	var userList []*user.InviteUserList
	for _, users := range us {
		var inviteCount int64
		if err = global.DB.Model(&User{}).Where("user_invite_id = ?", users.ID).Count(&inviteCount).Error; err != nil {
			return nil, err
		}
		userList = append(userList, &user.InviteUserList{
			UserId:    int64(users.ID),
			UserName:  users.UserName,
			InviteNum: inviteCount,
		})
	}
	// 对 userList 按照 InviteNum 进行排序
	sort.Slice(userList, func(i, j int) bool {
		return userList[i].InviteNum > userList[j].InviteNum
	})
	return userList, nil
}
