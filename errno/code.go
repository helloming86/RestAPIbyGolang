package errno

var (
	// common err
	OK = &Errno{Code:0, Message:"OK"}
	InternalServerError = &Errno{Code:10001, Message:"Internal server error"}
	ErrBind = &Errno{Code:10002, Message:"Error occurred while binding the request body to the struct"}

	// user err
	ErrUserNotFound = &Errno{Code:20102, Message:"The user was not found"}
)
