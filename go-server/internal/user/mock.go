package user

type MockUser struct{}

func (mu MockUser) Get(email string) (*User, error) {
	return nil, nil
}

func (mu MockUser) Store(u *User) error {
	return nil
}
