package dao_redis

import "models/model_user/model_redis"

func IsCheckedIn(userId int64, day int) (bool, error) {
	in, err := model_redis.IsCheckIn(userId, day)
	if err != nil {
		return false, err
	}

	return in == 1, nil
}
