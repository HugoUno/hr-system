package errors

import "errors"

var (
	ErrNotFound      = errors.New("記錄未找到")
	ErrInvalidInput  = errors.New("無效的輸入")
	ErrUnauthorized  = errors.New("未授權的訪問")
	ErrInternalError = errors.New("內部服務器錯誤")
)
