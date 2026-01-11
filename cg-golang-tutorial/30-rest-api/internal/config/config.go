package config

// to copy config yaml variables when thinsg start
type Config struct { // needs to be used publically, so caps
	// ythe annotations shows file type and the variable name to be copied from
	// sturct tags
	Env        string               `yaml:"env"`                         // for multiple `yaml:"env env:"ENV env-required:"true" to validate if the variable is there or not`
	Db_path    string               `yaml:"db_path" env-required:"true"` // for defualt values env-default
	HttpServer `yaml:"http_server"` // this nested so  calling another struct, embedding
}

// to hvae nested struct for Config struct
type HttpServer struct {
	Addr string `yaml:"address"`
}
