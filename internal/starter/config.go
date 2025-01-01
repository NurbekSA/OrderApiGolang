package starter

type Config struct {
	App      AppConfig      `json:"app"`
	Database DatabaseConfig `json:"database"`
	Kafka    KafkaConfig    `json:"kafka"`
	Redis    RedisConfig    `json:"redis"` // Новая структура для Redis
}

type AppConfig struct {
	AppName  string `json:"app_name"`  // Имя приложения
	BindAddr string `json:"bind_addr"` // Адрес, на который будет привязано приложение
	BindPort string `json:"bind_port"` // Порт, на котором будет запущено приложение
	LogLevel string `json:"log_level"` // Уровень логирования
}

type DatabaseConfig struct {
	Host         string `json:"host"`     // Хост для подключения к базе данных
	Port         string `json:"port"`     // Порт для подключения к базе данных
	Username     string `json:"username"` // Имя пользователя для подключения
	Password     string `json:"password"` // Пароль для подключения
	DatabaseName string `json:"database"` // Имя базы данных
}

type KafkaConfig struct {
	Brokers []string `json:"brokers"` // Список брокеров Kafka
}

type RedisConfig struct {
	Host     string `json:"host"`     // Хост для подключения к Redis
	Port     string `json:"port"`     // Порт для подключения к Redis
	Password string `json:"password"` // Пароль для подключения (если установлен)
	DB       int    `json:"db"`       // Номер базы данных Redis
}

// NewConfig Возвращает конфигурацию по умолчанию
func NewConfig() *Config {
	return &Config{
		App: AppConfig{
			AppName:  "amanzat_user_api", // Имя приложения
			BindAddr: "127.0.0.1",        // Адрес для подключения
			BindPort: "8080",             // Порт для подключения
			LogLevel: "debug",            // Уровень логирования
		},
		Database: DatabaseConfig{
			Host:         "localhost",    // Хост для подключения к базе данных
			Port:         "5432",         // Порт для подключения
			Username:     "postgres",     // Имя пользователя для подключения
			Password:     "postgres",     // Пароль для подключения
			DatabaseName: "amanzat_user", // Имя базы данных
		},
		Kafka: KafkaConfig{
			Brokers: []string{"localhost:9092"}, // Список брокеров Kafka
		},
		Redis: RedisConfig{
			Host:     "localhost", // Хост для Redis
			Port:     "6379",      // Порт для Redis
			Password: "",          // Пароль для Redis
			DB:       0,           // Номер базы данных Redis
		},
	}
}
