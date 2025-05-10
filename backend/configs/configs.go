package configs

type Config struct {
	PostgreSQL PostgreSQLConfig
	App        Fiber
	JWT        JWT
}

type PostgreSQLConfig struct {
	Host     string
	Port     string
	Protocol string
	Username string
	Password string
	Database string
	SSLMode  string
}

type Fiber struct {
	Host string
	Port string
}

type JWT struct {
	Secret        string
	Expire        int
	RefreshSecret string
	RefreshExpire int
}
