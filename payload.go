package caracolines

type Config struct {
	Parameters map[string]interface{} `yaml:"parameters"`
}

type Services struct {
	Items map[string]Service `yaml:"services"`
}

type Service struct {
	Package     string   `yaml:"package"`
	Constructor string   `yaml:"constructor"`
	Arguments   []string `yaml:"arguments"`
}
