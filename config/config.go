package config

func loadConfig(path string) (string, error) {
	return "", nil
}

type Database struct {
	Host string `json:"host"`
	Port string `json:"port"`
}
