package api

type Config struct {
	Host string
	Username string
	Password string
	Database string
}

func GetConfig() Config {
	return Config{
		Host: "remotemysql.com",
		Password: "6KhsmP3paT",
		Database: "D4EuHqdPic",
		Username: "D4EuHqdPic",
	}
}

