CREATE TABLE IF NOT EXISTS movies (
id SERIAL PRIMARY KEY,
name VARCHAR(150) NOT NULL,
year INT NOT NULL,
gender VARCHAR(150) NOT NULL,
adquired DATE,
stock INT DEFAULT 10,
price DECIMAL(10,2) NOT NULL
);
