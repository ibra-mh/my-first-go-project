package utils

// import "fmt"

var appProperties = applicationProperties{}

type applicationProperties struct {
	Database database `yaml:"database"`
	Vault    vault    `yaml:"vault"`
}

type database struct {
	Name string
	Pwd  string
	Host string
	Port int
}

type vault struct {
	Path string
}

func (appsProperties *applicationProperties) LoadProperties() {
	// appProperties = fmt.Sprint()
}