CREATE TABLE orders (
    order_id VARCHAR(20) NOT NULL,
    product_code VARCHAR(20),
    product_name VARCHAR(255),
    quantity INT,
    order_date DATETIME NOT NULL,
    price DECIMAL(10,2),
    customer_id INT NULL,
    PRIMARY KEY (order_id, product_code)
);