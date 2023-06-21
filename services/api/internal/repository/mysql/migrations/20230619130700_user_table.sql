-- +goose Up
CREATE TABLE User (

    id INT AUTO_INCREMENT NOT NULL,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    phone VARCHAR(255) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    created_at DATETIME NOT NULL,
    update_at DATETIME NOT NULL,
    PRIMARY KEY(id)

);

-- +goose Down
DROP TABLE User;

