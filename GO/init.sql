-- Таблица для подсчета овец
CREATE TABLE IF NOT EXISTS sheep_counts (
    id SERIAL PRIMARY KEY,
    sheep_array BOOLEAN[] NOT NULL,
    sheep_count INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Таблица для преобразований массивов
CREATE TABLE IF NOT EXISTS array_operations (
    id SERIAL PRIMARY KEY,
    original_array INTEGER[] NOT NULL,
    transformed_array INTEGER[] NOT NULL,
    operation_type VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);