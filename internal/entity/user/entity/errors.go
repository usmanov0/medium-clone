package entity

const (
	ErrUserNotFound       = Err("User not found")
	ErrInvalidPassword    = Err("Invalid password")
	ErrEmptyUserName      = Err("User name is empty")
	ErrEmptyMail          = Err("Email address is empty")
	ErrInvalidEmailFormat = Err("Invalid email format")
	ErrBadCredentials     = Err("Bad credentials")
)

type Err string

func (e Err) Error() string {
	return string(e)
}
