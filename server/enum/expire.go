package enum

import (
	"time"
)

const (
	// authorization_code 有效时间十分钟
	AUTHORIZATION_CODE_EXPIRE time.Duration = 10
	// Access Token的有效期为30天
	ACCESS_TOKEN_EXPIRE time.Duration = 30
	// Refresh Token的有效期为365天
	REFRESH_TOKEN_EXPIRE time.Duration = 365
)
