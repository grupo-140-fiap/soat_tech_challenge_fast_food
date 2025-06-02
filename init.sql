USE soat_fast_food;

CREATE TABLE customers (
    id          BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    first_name  VARCHAR(100),
    last_name   VARCHAR(100),
    cpf         VARCHAR(14) UNIQUE,
    email       VARCHAR(255) UNIQUE,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE orders (
    id          BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    customer_id BIGINT UNSIGNED NULL REFERENCES customers(id),
    cpf         VARCHAR(14) NULL,
    status      ENUM('received', 'preparation', 'ready', 'completed'),
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE products (
    id          BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    name        VARCHAR(100) NOT NULL,
    description TEXT,
    price       FLOAT NOT NULL,
    category    ENUM('burger', 'side', 'drink', 'dessert'),
    image_url   VARCHAR(255) NULL,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE order_items (
    id           BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    order_id     BIGINT UNSIGNED NOT NULL REFERENCES orders(id) ON DELETE CASCADE,
    product_id   BIGINT UNSIGNED NOT NULL REFERENCES products(id),
    quantity     INTEGER NOT NULL DEFAULT 1,
    price        FLOAT NOT NULL,
    created_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);