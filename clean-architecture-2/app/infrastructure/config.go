package infrastructure

type Config struct {
	DB struct {
		Host     string
		Username string
		Password string
		DBName   string
	}
}

func NewConfig() *Config { // This should be written in .env // Do not hard coding
	c := new(Config)

	c.DB.Host = "localhost"
	c.DB.Username = "username"
	c.DB.Password = "password"
	c.DB.DBName = "db_name"

	return c
}
