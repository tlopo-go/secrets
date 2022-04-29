package app

var cfg config

type config struct {
	databasePath   string
	masterPassword string
}

func GetDatabasePath() string {
	return cfg.databasePath
}

func GetMasterPassword() string {
	return cfg.masterPassword
}

func init() {
	cfg.databasePath = "/tmp/db.kdbx"
	cfg.masterPassword = "1234"
}
