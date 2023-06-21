-- +goose Up
CREATE TABLE Product (

    id INT AUTO_INCREMENT NOT NULL,
    user_id INT,
    name VARCHAR(255) NOT NULL UNIQUE,
    price FLOAT NOT NULL,
    quantity INT NOT NULL,
    created_at DATETIME NOT NULL,
    update_at DATETIME NOT NULL,
    PRIMARY KEY(id),
    FOREIGN KEY (user_id)  REFERENCES User (id) ON DELETE CASCADE

);

-- +goose Down
DROP TABLE User;