package configuration

import (
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

var (
	ACCESS_USER_TOKEN_SECRET  string
	ACCESS_ADMIN_TOKEN_SECRET string

	REFRESH_USER_TOKEN_SECRET  string
	REFRESH_ADMIN_TOKEN_SECRET string

	ACCESS_USER_TOKEN_EXPIRY  time.Duration
	REFRESH_TOKEN_EXPIRY      time.Duration
	REFRESH_USER_TOKEN_EXPIRY time.Duration
)

func LoadConfig() error {
	_ = godotenv.Load(".env")
	_ = godotenv.Load(".env.development")
	_ = godotenv.Load(".env.local")

	ACCESS_USER_TOKEN_SECRET = firstNonEmpty(
		os.Getenv("ACCESS_USER_TOKEN_SECRET"),
	)
	ACCESS_ADMIN_TOKEN_SECRET = firstNonEmpty(
		os.Getenv("ACCESS_ADMIN_TOKEN_SECRET"),
		os.Getenv("ACCESS_ADMIB_TOKEN_SECRET"),
	)
	REFRESH_USER_TOKEN_SECRET = firstNonEmpty(
		os.Getenv("REFRESH_USER_TOKEN_SECRET"),
		os.Getenv("REFRESH__USER_TOKEN_SECRET"),
	)
	REFRESH_ADMIN_TOKEN_SECRET = firstNonEmpty(
		os.Getenv("REFRESH_ADMIN_TOKEN_SECRET"),
	)

	ACCESS_USER_TOKEN_EXPIRY = parseDuration(firstNonEmpty(
		os.Getenv("ACCESS_USER_TOKEN_EXPIRY"),
		"900",
	))
	REFRESH_TOKEN_EXPIRY = parseDuration(firstNonEmpty(
		os.Getenv("REFRESH_USER_TOKEN_EXPIRY"),
		os.Getenv("REFRESH__USER_TOKEN_EXPIRY"),
		os.Getenv("REFRESH_TOKEN_EXPIRY"),
		"31536000",
	))
	REFRESH_USER_TOKEN_EXPIRY = REFRESH_TOKEN_EXPIRY

	return validateConfig()
}

func firstNonEmpty(values ...string) string {
	for _, value := range values {
		if strings.TrimSpace(value) != "" {
			return strings.TrimSpace(value)
		}
	}
	return ""
}

func parseDuration(value string) time.Duration {
	value = strings.TrimSpace(value)
	if value == "" {
		return 0
	}

	if seconds, err := strconv.Atoi(value); err == nil {
		return time.Duration(seconds) * time.Second
	}

	if len(value) > 1 && value[len(value)-1] == 'd' {
		days, err := strconv.Atoi(value[:len(value)-1])
		if err != nil {
			return 0
		}
		return time.Duration(days) * 24 * time.Hour
	}

	duration, err := time.ParseDuration(value)
	if err != nil {
		return 0
	}

	return duration
}

func validateConfig() error {
	if ACCESS_USER_TOKEN_SECRET == "" {
		return os.ErrInvalid
	}
	if ACCESS_ADMIN_TOKEN_SECRET == "" {
		return os.ErrInvalid
	}
	if REFRESH_USER_TOKEN_SECRET == "" {
		return os.ErrInvalid
	}
	if REFRESH_ADMIN_TOKEN_SECRET == "" {
		return os.ErrInvalid
	}
	if ACCESS_USER_TOKEN_EXPIRY <= 0 {
		return os.ErrInvalid
	}
	if REFRESH_TOKEN_EXPIRY <= 0 {
		return os.ErrInvalid
	}
	return nil
}
