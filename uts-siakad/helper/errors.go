package helper

// AppError khusus untuk melempar status HTTP dan pesan
type AppError struct {
	Code    int
	Message string
}

func (e *AppError) Error() string {
	return e.Message
}

// Kumpulan fungsi pembuat Error
func BadRequest(message string) *AppError {
	return &AppError{Code: 400, Message: message}
}

func Unauthorized(message string) *AppError {
	return &AppError{Code: 401, Message: message}
}

func Forbidden(message string) *AppError {
	return &AppError{Code: 403, Message: message}
}

func NotFound(message string) *AppError {
	return &AppError{Code: 404, Message: message}
}

func Conflict(message string) *AppError {
	return &AppError{Code: 409, Message: message}
}

func ValidationError(message string) *AppError {
	return &AppError{Code: 422, Message: message}
}
