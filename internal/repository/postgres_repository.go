package repository

import (
	"github.com/jmoiron/sqlx"
	"gitlab.com/advanced-programing/car-rental-system/internal/domain"
)

type PostgresRepository struct {
	db *sqlx.DB
}

func NewPostgresRepository(db *sqlx.DB) Repository {
	return &PostgresRepository{db: db}
}

// -- User Methods --

func (r *PostgresRepository) CreateUser(user domain.User) (int64, error) {
	var id int64
	query := `INSERT INTO users (name, email, phone, role, password, created_at)
	          VALUES ($1, $2, $3, $4, $5, NOW()) RETURNING id`
	err := r.db.QueryRow(query, user.Name, user.Email, user.Phone, user.Role, user.Password).Scan(&id)
	return id, err
}

func (r *PostgresRepository) GetUserByEmail(email string) (domain.User, error) {
	var user domain.User
	query := `SELECT * FROM users WHERE email=$1`
	err := r.db.Get(&user, query, email)
	return user, err
}

func (r *PostgresRepository) GetAllUsers() ([]domain.User, error) {
	var users []domain.User
	query := `SELECT * FROM users`
	err := r.db.Select(&users, query)
	return users, err
}

// -- Car Methods --

func (r *PostgresRepository) CreateCar(car domain.Car) (int64, error) {
	query := `INSERT INTO cars (brand, model, transmission, fuel_type, price_per_day, available, photo) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`
	var id int64
	err := r.db.QueryRow(query, car.Brand, car.Model, car.Transmission, car.FuelType, car.PricePerDay, car.Available, car.Photo).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *PostgresRepository) UpdateCar(car domain.Car) error {
	query := `UPDATE cars SET brand=$1, model=$2, transmission=$3, fuel_type=$4, price_per_day=$5, available=$6, photo=$7 WHERE id=$8`
	_, err := r.db.Exec(query, car.Brand, car.Model, car.Transmission, car.FuelType, car.PricePerDay, car.Available, car.Photo, car.ID)
	return err
}

func (r *PostgresRepository) DeleteCar(id int64) error {
	query := `DELETE FROM cars WHERE id=$1`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *PostgresRepository) GetCarByID(id int64) (domain.Car, error) {
	var car domain.Car
	query := `SELECT * FROM cars WHERE id=$1`
	err := r.db.Get(&car, query, id)
	return car, err
}

func (r *PostgresRepository) ListCars() ([]domain.Car, error) {
	var cars []domain.Car
	query := `SELECT * FROM cars`
	err := r.db.Select(&cars, query)
	return cars, err
}

func (r *PostgresRepository) ListCarsByCriteria(criteria map[string]interface{}) ([]domain.Car, error) {
	// Simplified example: filter by brand if provided.
	var cars []domain.Car
	query := `SELECT * FROM cars WHERE 1=1`
	args := []interface{}{}
	argIndex := 1

	if brand, ok := criteria["brand"]; ok {
		query += " AND brand=$1"
		args = append(args, brand)
		argIndex++
	}
	if model, ok := criteria["model"]; ok {
		query += " AND model=$2"
		args = append(args, model)
		argIndex++
	}
	if transmission, ok := criteria["transmission"]; ok {
		query += " AND transmission=$3"
		args = append(args, transmission)
		argIndex++
	}
	if fuelType, ok := criteria["fuel_type"]; ok {
		query += " AND fuel_type=$4"
		args = append(args, fuelType)
		argIndex++
	}
	if mileage, ok := criteria["mileage"]; ok {
		query += " AND mileage<=$5"
		args = append(args, mileage)
		argIndex++
	}
	if pricePerDay, ok := criteria["price_per_day"]; ok {
		query += " AND price_per_day<=$6"
		args = append(args, pricePerDay)
		argIndex++
	}

	err := r.db.Select(&cars, query, args...)
	return cars, err
}

// -- Rental Methods --

func (r *PostgresRepository) CreateRental(rental domain.Rental) (int64, error) {
	var id int64
	query := `INSERT INTO rentals (user_id, car_id, payment_id, start_date, end_date)
	          VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err := r.db.QueryRow(query, rental.UserID, rental.CarID, rental.PaymentID, rental.StartDate, rental.EndDate).Scan(&id)
	return id, err
}

func (r *PostgresRepository) UpdateRental(rental domain.Rental) error {
	query := `UPDATE rentals SET user_id=$1, car_id=$2, payment_id=$3, start_date=$4, end_date=$5 WHERE id=$6`
	_, err := r.db.Exec(query, rental.UserID, rental.CarID, rental.PaymentID, rental.StartDate, rental.EndDate, rental.ID)
	return err
}

func (r *PostgresRepository) GetRentalByID(id int64) (domain.Rental, error) {
	var rental domain.Rental
	query := `SELECT * FROM rentals WHERE id=$1`
	err := r.db.Get(&rental, query, id)
	return rental, err
}

func (r *PostgresRepository) GetAllRentals() ([]domain.Rental, error) {
	var rentals []domain.Rental
	err := r.db.Select(&rentals, "SELECT * FROM rentals")
	return rentals, err
}

func (r *PostgresRepository) GetRentalsByUser(userID int64) ([]domain.Rental, error) {
	var rentals []domain.Rental
	err := r.db.Select(&rentals, "SELECT * FROM rentals WHERE user_id = $1", userID)
	return rentals, err
}

// -- Payment Methods --

func (r *PostgresRepository) CreatePayment(payment domain.Payment) (int64, error) {
	var id int64
	query := `INSERT INTO payments (rental_id, status) VALUES ($1, $2) RETURNING id`
	err := r.db.QueryRow(query, payment.RentalID, payment.Status).Scan(&id)
	return id, err
}

func (r *PostgresRepository) UpdatePayment(payment domain.Payment) error {
	query := `UPDATE payments SET rental_id=$1, status=$2 WHERE id=$3`
	_, err := r.db.Exec(query, payment.RentalID, payment.Status, payment.ID)
	return err
}

func (r *PostgresRepository) GetPaymentByID(id int64) (domain.Payment, error) {
	var payment domain.Payment
	query := `SELECT * FROM payments WHERE id=$1`
	err := r.db.Get(&payment, query, id)
	return payment, err
}

func (r *PostgresRepository) DeletePayment(id int64) error {
	query := `DELETE FROM payments WHERE id=$1`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *PostgresRepository) ListAllPayments() ([]domain.Payment, error) {
	var payments []domain.Payment
	query := `SELECT * FROM payments`
	err := r.db.Select(&payments, query)
	return payments, err
}

// -- Feedback Methods --

func (r *PostgresRepository) CreateFeedback(feedback domain.Feedback) (int64, error) {
	var id int64
	query := `INSERT INTO feedback (user_id, car_id, rating, description) VALUES ($1, $2, $3, $4) RETURNING id`
	err := r.db.QueryRow(query, feedback.UserID, feedback.CarID, feedback.Rating, feedback.Description).Scan(&id)
	return id, err
}

func (r *PostgresRepository) UpdateFeedback(feedback domain.Feedback) error {
	query := `UPDATE feedback SET user_id=$1, car_id=$2, rating=$3, description=$4 WHERE id=$5`
	_, err := r.db.Exec(query, feedback.UserID, feedback.CarID, feedback.Rating, feedback.Description, feedback.ID)
	return err
}

func (r *PostgresRepository) DeleteFeedback(id int64) error {
	query := `DELETE FROM feedback WHERE id=$1`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *PostgresRepository) GetFeedbackByID(id int64) (domain.Feedback, error) {
	var feedback domain.Feedback
	query := `SELECT * FROM feedback WHERE id=$1`
	err := r.db.Get(&feedback, query, id)
	return feedback, err
}

func (r *PostgresRepository) ListFeedbackByCar(carID int64) ([]domain.Feedback, error) {
	var feedbacks []domain.Feedback
	query := `SELECT * FROM feedback WHERE car_id=$1`
	err := r.db.Select(&feedbacks, query, carID)
	return feedbacks, err
}

func (r *PostgresRepository) ListFeedbackByUser(userID int64) ([]domain.Feedback, error) {
	var feedbacks []domain.Feedback
	query := `SELECT * FROM feedback WHERE user_id=$1`
	err := r.db.Select(&feedbacks, query, userID)
	return feedbacks, err
}

func (r *PostgresRepository) ListAllFeedback() ([]domain.Feedback, error) {
	var feedbacks []domain.Feedback
	query := `SELECT * FROM feedback`
	err := r.db.Select(&feedbacks, query)
	return feedbacks, err
}

func (r *PostgresRepository) UpdateUser(user domain.User) error {
	query := `UPDATE users SET name=$1, email=$2, phone=$3, password=$4 WHERE id=$5`
	_, err := r.db.Exec(query, user.Name, user.Email, user.Phone, user.Password, user.ID)
	return err
}

func (r *PostgresRepository) GetUserByID(id int64) (*domain.User, error) {
	var user domain.User
	err := r.db.Get(&user, "SELECT * FROM users WHERE id=$1", id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
