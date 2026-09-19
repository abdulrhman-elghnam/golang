package enum
type SystemRole string

const (
	USER  SystemRole = "USER"
	ADMIN SystemRole = "ADMIN"
)

type TokenType string

const (
	ACCESS  TokenType = "ACCESS"
	REFRESH TokenType = "REFRESH"
)