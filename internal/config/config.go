package config

import "github.com/spf13/viper"

// AppConfig contains application settings.
type AppConfig struct {
	App    AppSettings
	Server ServerSettings
}

// AppSettings contains general application settings.
type AppSettings struct {
	Name        string
	Version     string
	Environment string
}

// ServerSettings contains server settings.
type ServerSettings struct {
	Host string
	Port int
}

// Load reads application configuration from config.yaml.
func Load() (*AppConfig, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	viper.SetDefault("app.name", "Lab 2 Go Tooling")
	viper.SetDefault("app.version", "1.0.0")
	viper.SetDefault("app.environment", "development")
	viper.SetDefault("server.host", "localhost")
	viper.SetDefault("server.port", 8080)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	cfg := &AppConfig{
		App: AppSettings{
			Name:        viper.GetString("app.name"),
			Version:     viper.GetString("app.version"),
			Environment: viper.GetString("app.environment"),
		},
		Server: ServerSettings{
			Host: viper.GetString("server.host"),
			Port: viper.GetInt("server.port"),
		},
	}

	return cfg, nil
}
