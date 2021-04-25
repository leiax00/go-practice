package task

func QueryAll(userId int) (*User, error) {
	user, err := QueryData(userId)
	if err != nil {
		return nil, err
	}
	return user, err
}