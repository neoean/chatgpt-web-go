package env

import "os"

const (
	KEY = "env"

	KeyDev  = "dev"
	KeyTest = "test"
	KeyProd = "prod"
)

func GetEnv() string {
	e := os.Getenv(KEY)
	if e == "" {
		return KeyDev
	}
	return e
}

func IsDevelop() bool {
	return GetEnv() == KeyDev || GetEnv() == ""
}

func IsTest() bool {
	return GetEnv() == KeyTest
}

func IsProduction() bool {
	return GetEnv() == KeyProd
}
