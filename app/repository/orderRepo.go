package repository

import (
	"HalykProject/app/models"
	"HalykProject/util/exception"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus" // Импортируем logrus
)

type OrderRepo struct {
	db *pgxpool.Pool
}

func NewOrderRepo(db *pgxpool.Pool) *OrderRepo {
	return &OrderRepo{db: db}
}

func (r *OrderRepo) CreateOrder(ctx context.Context, user models.Order) *exception.AppError {
	logrus.Infof("Создание заказа для пользователя %v", user.UserId)
	_, err := r.db.Exec(ctx,
		`INSERT INTO orders 
        (user_id, boxes_id, status, booking_date_time, rental_period, total_price) 
        VALUES ($1, $2, $3, $4, $5, $6)`,
		user.UserId, user.BoxesId, user.Status,
		user.BookingDateTime, user.RentalPeriod, user.TotalPrice,
	)
	if err != nil {
		return exception.Cast(400, fmt.Sprintf("CreateOrder: не удалось создать заказ для пользователя %v: %w", user.UserId, err))
	}
	return nil
}

func (r *OrderRepo) GetAllOrders(ctx context.Context) ([]*models.Order, *exception.AppError) {
	logrus.Info("Получение всех заказов") // Логируем начало получения всех заказов
	rows, err := r.db.Query(ctx, "SELECT * FROM orders")
	if err != nil {
		return nil, exception.Cast(500, fmt.Sprintf("GetAllOrders: не удалось получить ордеры: %v", err))
	}
	defer rows.Close()

	var orders []*models.Order
	for rows.Next() {
		order, err := rowParser(rows)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	return orders, nil
}

func (r *OrderRepo) GetOrderByID(ctx context.Context, id string) (*models.Order, *exception.AppError) {
	logrus.Infof("Получение всех заказов по Id: %s", id)

	row := r.db.QueryRow(ctx, "SELECT * FROM orders WHERE id=$1", id)

	var order models.Order
	err := row.Scan(&order.Id,
		&order.UserId,
		&order.BoxesId,
		&order.Status,
		&order.BookingDateTime,
		&order.RentalPeriod,
		&order.TotalPrice,
	)
	if err != nil {
		return nil, exception.Cast(500, fmt.Sprintf("rowParser: не удалось запарсить строку: %w", err))
	}
	return &order, nil
}

func (r *OrderRepo) GetOrdersByUserId(ctx context.Context, id string) ([]*models.Order, *exception.AppError) {
	logrus.Infof("Получение всех заказов по userId: %s", id)

	rows, err := r.db.Query(ctx, "SELECT * FROM orders WHERE user_id=$1", id)
	if err != nil {
		return nil, exception.Cast(500, fmt.Sprintf("GetOrdersByUserId: не удалось получить ордеры для пользователя %v: %w", id, err))
	}
	defer rows.Close()

	var orders []*models.Order
	for rows.Next() {
		order, err := rowParser(rows)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	return orders, nil
}

func (r *OrderRepo) GetOrderByBoxId(ctx context.Context, id string) ([]*models.Order, *exception.AppError) {
	logrus.Infof("Получение всех заказов по Id: boxId: %s", id)

	rows, err := r.db.Query(ctx, "SELECT * FROM orders WHERE boxes_id @> $1::jsonb", fmt.Sprintf(`["%s"]`, id))

	if err != nil {
		return nil, exception.Cast(500, fmt.Sprintf("GetByBoxId: Ордер с boxId %v не найден. ERROR: %v", id, err))
	}
	defer rows.Close()

	var orders []*models.Order
	for rows.Next() {
		order, err := rowParser(rows)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	return orders, nil
}

func (r *OrderRepo) GetOrderByStatus(ctx context.Context, status string) ([]*models.Order, *exception.AppError) {
	logrus.Infof("Все ордеры с статусом %s", status)

	// todo: Нужно исправить запрос
	rows, err := r.db.Query(ctx, "SELECT * FROM orders WHERE status = $1", status)
	if err != nil {
		return nil, exception.Cast(500, fmt.Sprintf("GetOrderByStatus: Ошибка при пойске ордера с статусом %s  %v", status, err))
	}
	defer rows.Close()

	var orders []*models.Order
	for rows.Next() {
		order, err := rowParser(rows)
		if err != nil {
			return nil, err
		}
		print(order)
		orders = append(orders, order)
	}
	return orders, nil
}

// todo: нужно исправить
func (r *OrderRepo) Complete(ctx context.Context, id string) *exception.AppError {
	logrus.Infof("Complete: завершение ордера с id: %s", id)

	_, err := r.db.Exec(ctx, "UPDATE orders SET status = $1 WHERE id = $2", "EXPIRED", id)
	if err != nil {
		return exception.Cast(500, fmt.Sprintf("Completing: Не удалось завершить ордер. Id: %s \n %e", id, err))
	}
	return nil
}

// todo: нужно исправить
func (r *OrderRepo) Update(ctx context.Context, order models.Order) *exception.AppError {
	_, err := r.db.Exec(ctx, "UPDATE orders SET user_id = $1, boxes_id = $2, status = $3, booking_date_time = $4, rental_period = $5, total_price = $6  WHERE id = $7",
		order.UserId, order.BoxesId, order.Status, order.BookingDateTime, order.RentalPeriod, order.TotalPrice, order.Id)

	if err != nil {
		return exception.Cast(500, fmt.Sprintf("Update: Не удалось обнавить ордер. Id: %s \n %e", order.UserId, err))
	}

	return nil
}

func (r *OrderRepo) Extend(ctx context.Context, id string, rentalPeriod int) *exception.AppError {
	logrus.Infof("Extend: Продление ордера на %d недель. id: %s ", rentalPeriod, id)

	order, err := r.GetOrderByID(ctx, id)
	if err != nil {
		return exception.Cast(500, fmt.Sprintf("Extend: Не удалось продлить ордер. Id: %s \n %s", id, err))
	}
	totalRentalPeriod := order.RentalPeriod + rentalPeriod

	_, erro := r.db.Exec(ctx, "UPDATE orders SET rental_period = $1 WHERE id = $2", totalRentalPeriod, id)
	if erro != nil {
		return exception.Cast(500, fmt.Sprintf("Extend: Не удалось продлить ордер. Id: %s \n %s", id, erro))
	}
	return nil
}

func rowParser(row pgx.Rows) (*models.Order, *exception.AppError) {
	var order models.Order
	err := row.Scan(&order.Id,
		&order.UserId,
		&order.BoxesId,
		&order.Status,
		&order.BookingDateTime,
		&order.RentalPeriod,
		&order.TotalPrice,
	)
	if err != nil {
		return nil, exception.Cast(500, fmt.Sprintf("rowParser: не удалось запарсить строку: %w", err))
	}
	return &order, nil
}
