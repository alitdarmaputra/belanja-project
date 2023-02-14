CREATE TABLE products (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    qty INT DEFAULT 0 NOT NULL,
    outlets_id INT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL,
    deleted_at DATETIME,
    CONSTRAINT FK_UserProduct FOREIGN KEY (outlets_id)
    REFERENCES users(id)
)