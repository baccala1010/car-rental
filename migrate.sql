-- Create Users Table
CREATE TABLE IF NOT EXISTS users (
                                     id SERIAL PRIMARY KEY,
                                     name VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    phone VARCHAR(20),
    role VARCHAR(20) NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
    );

-- Create Cars Table
CREATE TABLE IF NOT EXISTS cars (
                                    id SERIAL PRIMARY KEY,
                                    brand VARCHAR(50) NOT NULL,
    model VARCHAR(50) NOT NULL,
    transmission VARCHAR(20) NOT NULL,
    fuel_type VARCHAR(20) NOT NULL,
    price_per_day NUMERIC(10,2) NOT NULL,
    available BOOLEAN DEFAULT TRUE,
    photo TEXT
    );

-- Create Rentals Table
CREATE TABLE IF NOT EXISTS rentals (
                                       id SERIAL PRIMARY KEY,
                                       user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    car_id INTEGER NOT NULL REFERENCES cars(id) ON DELETE CASCADE,
    payment_id INTEGER,
    start_date TIMESTAMP NOT NULL,
    end_date TIMESTAMP NOT NULL
    );

-- Create Payments Table
CREATE TABLE IF NOT EXISTS payments (
                                        id SERIAL PRIMARY KEY,
                                        rental_id INTEGER NOT NULL REFERENCES rentals(id) ON DELETE CASCADE,
    status VARCHAR(20) NOT NULL
    );

-- Create Feedback Table
CREATE TABLE IF NOT EXISTS feedback (
                                        id SERIAL PRIMARY KEY,
                                        user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    car_id INTEGER NOT NULL REFERENCES cars(id) ON DELETE CASCADE,
    rating INTEGER NOT NULL,
    description TEXT
    );
