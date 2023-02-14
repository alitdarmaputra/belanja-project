CREATE TABLE orders (
    id INT PRIMARY KEY AUTO_INCREMENT,
    users_id INT,
    outlets_id INT,
    shipper_id INT,
    status VARCHAR(20),
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL,
    deleted_at DATETIME,
    CONSTRAINT FK_UserOrder FOREIGN KEY (users_id)
    REFERENCES users(id),
    CONSTRAINT FK_OutletOrder FOREIGN KEY (outlets_id)
    REFERENCES users(id)
) ENGINE=InnoDB;
