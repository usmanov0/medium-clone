package errors

const (
	ErrUserNotFound              = Err("user not found")
	ErrUserUpdateFailed          = Err("failed to update user")
	ErrInvalidPassword           = Err("invalid password")
	ErrEmptyUserName             = Err("user name is empty")
	ErrEmptyMail                 = Err("email address is empty")
	ErrInvalidEmailFormat        = Err("invalid email format")
	ErrBadCredentials            = Err("fad credentials")
	ErrIdScanFailed              = Err("failed to scan id")
	ErrScanRows                  = Err("failed to scan rows")
	ErrShouldBeDifferentName     = Err("username should be different from before")
	ErrShouldBeDifferentPassword = Err("password should be different from before")
	ErrShouldBeDifferentBio      = Err("bio should be different from before")
	ErrFailedDeleteAccount       = Err("failed to delete user account")
	ErrArticleTitleEmpty         = Err("article title can't be empty")
	ErrArticleBodyEmpty          = Err("article body can't be empty")
	ErrArticleExceedTitle        = Err("article title exceeds maximum length")
	ErrArticleExceedBody         = Err("article body exceeds maximum length")
	ErrCreationToken             = Err("Error creating token")
	ErrFailedExecuteQuery        = Err("failed execute query")
)

type Err string

func (e Err) Error() string {
	return string(e)
}
