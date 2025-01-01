CREATE TABLE orders (
                        id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),  -- Уникальный идентификатор в формате UUID
                        user_id VARCHAR(255) NOT NULL,  -- ID пользователя, также UUID
                        boxes_id JSONB NOT NULL,  -- Массив идентификаторов коробок в формате JSONB
                        status VARCHAR(50) NOT NULL,  -- Статус заказа (например, "pending", "completed")
                        booking_date_time BIGINT NOT NULL,  -- Время бронирования (в формате Unix timestamp)
                        rental_period INT NOT NULL,  -- Срок аренды (в днях)
                        total_price DECIMAL(10, 2) NOT NULL  -- Общая стоимость (например, 99999.99)
);
