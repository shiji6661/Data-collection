package dao_mysql

func UpdateUserPoint(userId uint64, point int) error {
	err := U.UpdateUserPoints(int64(userId), int64(point))
	if err != nil {
		return err
	}
	return nil
}
