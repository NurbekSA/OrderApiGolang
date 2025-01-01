package starter

import (
	"HalykProject/api"
	"HalykProject/app/repository"
	"HalykProject/app/service"
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

type ApiServer struct {
	Сonfig       *Config
	Db           *pgxpool.Pool
	OrderService *service.OrderService
	OrderRepo    *repository.OrderRepo
}

func NewApiServer(config *Config) *ApiServer {
	return &ApiServer{
		Сonfig: config,
		Db:     nil,
	}
}

func (apiServer *ApiServer) Run() {

	level, err := logrus.ParseLevel(apiServer.Сonfig.App.LogLevel)
	if err != nil {
		logrus.Fatalf("Run. Error running server: %v", err)
	}

	// Настройка стандартного логгера
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableTimestamp: true, // Убираем время из заголовка лога
		ForceColors:      true, // Включаем цветной вывод
	})
	logrus.SetLevel(level)      // Уровень логирования
	logrus.SetOutput(os.Stdout) // Вывод в стандартный поток

	logrus.Info("Run. Запуск API Server")

	// Initialize database
	ctx := context.Background()
	err = NewDatabase(ctx, apiServer)
	if err != nil {
		logrus.Fatalf("Run. Error connecting to database: %v", err)
	}
	logrus.Info("Run. База данных успешно инициализирована")

	// Initialize repository, service, and API
	apiServer.OrderRepo = repository.NewOrderRepo(apiServer.Db)
	apiServer.OrderService = service.NewOrderService(apiServer.OrderRepo)
	orderHandler := api.NewUserHandler(apiServer.OrderService)

	app := fiber.New()

	api := app.Group("/api/order")

	api.Use(LogrusMiddleware(logrus.New()))
	api.Post("/create", orderHandler.CreateOrder)
	api.Get("/all", orderHandler.GetAllOrder)
	api.Get("/:id", orderHandler.GetOrderByID)
	api.Get("/user/:userId", orderHandler.GetOrderByUserId)
	api.Get("/box/:boxId", orderHandler.GetOrderByBoxId)
	api.Get("/status/:status", orderHandler.GetOrderByStatus)
	api.Patch("/complete/:id", orderHandler.Complete)
	api.Post("/update", orderHandler.UpdateOrder)
	api.Patch("/extension/:id", orderHandler.Extend)

	// Запуск сервера
	logrus.Fatal(app.Listen(":8080"))
}

// CustomFormatter задает цвет всей строки в зависимости от уровня
type CustomFormatter struct{}

func (f *CustomFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var color string

	// Задаем цвет строки в зависимости от уровня
	switch entry.Level {
	case logrus.InfoLevel:
		color = "\033[32m" // Зеленый для INFO
	case logrus.WarnLevel:
		color = "\033[33m" // Желтый для WARN
	case logrus.ErrorLevel:
		color = "\033[31m" // Красный для ERROR
	default:
		color = "\033[36m" // Голубой для остальных уровней
	}

	// Форматируем строку с применением цвета ко всей строке
	log := fmt.Sprintf("%s[%s] %s: %v%s\n",
		color, entry.Level.String(), entry.Message, entry.Data, "\033[0m",
	)

	return []byte(log), nil
}

// LogrusMiddleware создает middleware для логирования запросов и ответов
func LogrusMiddleware(logger *logrus.Logger) fiber.Handler {
	// Устанавливаем кастомный форматтер
	logger.SetFormatter(&CustomFormatter{})

	return func(c *fiber.Ctx) error {
		start := time.Now()

		// Выполняем обработку запроса
		err := c.Next()

		// Статус ответа
		status := c.Response().StatusCode()

		// Подготовка логируемых данных
		entry := logger.WithFields(logrus.Fields{
			"method":     c.Method(),
			"path":       c.Path(),
			"status":     status,
			"latency":    time.Since(start),
			"client_ip":  c.IP(),
			"user_agent": string(c.Request().Header.UserAgent()),
		})

		// Логирование в зависимости от статуса
		switch {
		case status >= 500:
			entry.Error("Internal server error")
		case status >= 400:
			entry.Warn("Client error")
		default:
			entry.Info("Request handled successfully")
		}

		return err
	}
}
