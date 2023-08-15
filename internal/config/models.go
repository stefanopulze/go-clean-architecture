package config

type Config struct {
	Log      LogOpts      `yaml:"log"`
	Database DatabaseOpts `yaml:"database"`
}

type LogOpts struct {
	Level string `yaml:"level"`
	Type  string `yaml:"type"`
}

type DatabaseOpts struct {
	Uri  string `yaml:"uri" env:"DATABASE_URI"`
	Name string `yaml:"name" env:"DATABASE_NAME"`
}
