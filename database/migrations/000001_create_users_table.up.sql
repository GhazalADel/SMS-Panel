CREATE TABLE users (
    id SERIAL,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    phone VARCHAR(255) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    national_id VARCHAR(255) NOT NULL UNIQUE,
    PRIMARY KEY (id)
);