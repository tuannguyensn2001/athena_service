package config

type Config struct {
	DbUrl string
	Port  string
}

func Get() (Config, error) {
	result := Config{
		Port: "15000",
	}

	result.DbUrl = "host=localhost port=5432 user=tuannguyen password='' dbname=athena_go sslmode=disable"

	return result, nil
}
