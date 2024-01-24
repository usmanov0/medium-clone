package errors

const (
	ErrUserNotFound        = Err("User not found")
	ErrUserUpdateFailed    = Err("Failed to update user")
	ErrInvalidPassword     = Err("Invalid password")
	ErrEmptyUserName       = Err("User name is empty")
	ErrEmptyMail           = Err("Email address is empty")
	ErrInvalidEmailFormat  = Err("Invalid email format")
	ErrBadCredentials      = Err("Bad credentials")
	ErrIdScanFailed        = Err("Failed to scan id")
	ErrScanRows            = Err("Failed to scan rows")
	ErrFailedDeleteAccount = Err("Failed to delete user account")
	ErrArticleTitleEmpty   = Err("article title can't be empty")
	ErrArticleBodyEmpty    = Err("article body can't be empty")
	ErrArticleExceedTitle  = Err("article title exceeds maximum length")
	ErrArticleExceedBody   = Err("article body exceeds maximum length")
	ErrCreationToken       = Err("Error creating token")
)

type Err string

func (e Err) Error() string {
	return string(e)
}
