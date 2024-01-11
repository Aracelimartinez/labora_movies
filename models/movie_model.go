package models

type Movie struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Year     int     `json:"year"`
	Gender   string  `json:"gender"`
	Adquired string  `json:"adquired"`
	Stock    int     `json:"stock"`
	Price    float64 `json:"price"`
}

// CREATE TABLE IF NOT EXISTS movies (
// 	id SERIAL PRIMARY KEY,
// 	name VARCHAR(150) NO NULL,
// 	year INT NOT NULL,
// 	gender VARCHAR(150) NOT NULL,
// 	adquired DATE,
// 	stock INT DEFAULT 10,
// 	price DECIMAL(2) NOT NULL
// 	);
