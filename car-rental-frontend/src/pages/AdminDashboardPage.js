import React, { useState, useEffect } from 'react';
import {
    createCar, updateCar, deleteCar,
    listPayments, updatePayment, deletePayment,
    listAllRentals, deleteFeedbackAdmin, listAllFeedback, listCars
} from '../services/api';

const AdminDashboardPage = () => {
    const [cars, setCars] = useState([]);
    const [payments, setPayments] = useState([]);
    const [rentals, setRentals] = useState([]);
    const [feedbacks, setFeedbacks] = useState([]);
    const [newCar, setNewCar] = useState({ brand: '', model: '', transmission: '', fuel_type: '', price_per_day: 0, photo: '' });

    useEffect(() => {
        // Fetch initial data for cars, payments, rentals, and feedbacks
        fetchCars();
        fetchPayments();
        fetchRentals();
        fetchFeedbacks();
    }, []);

    const fetchCars = async () => {
        try {
            const response = await listCars();
            setCars(response.data);
        } catch (error) {
            console.error('Failed to fetch cars', error);
        }
    };

    const fetchPayments = async () => {
        try {
            const response = await listPayments();
            setPayments(response.data);
        } catch (error) {
            console.error('Failed to fetch payments', error);
        }
    };

    const fetchRentals = async () => {
        try {
            const response = await listAllRentals();
            setRentals(response.data);
        } catch (error) {
            console.error('Failed to fetch rentals', error);
        }
    };

    const fetchFeedbacks = async () => {
        try {
            const response = await listAllFeedback();
            setFeedbacks(response.data);
        } catch (error) {
            console.error('Failed to fetch feedbacks', error);
        }
    };

    const handleCreateCar = async () => {
        try {
            await createCar(newCar);
            fetchCars();
        } catch (error) {
            console.error('Failed to create car', error);
        }
    };

    const handleUpdateCar = async (id, updatedCar) => {
        try {
            await updateCar(id, updatedCar);
            fetchCars();
        } catch (error) {
            console.error('Failed to update car', error);
        }
    };

    const handleDeleteCar = async (id) => {
        try {
            await deleteCar(id);
            fetchCars();
        } catch (error) {
            console.error('Failed to delete car', error);
        }
    };

    const handleUpdatePayment = async (id, updatedPayment) => {
        try {
            await updatePayment(id, updatedPayment);
            fetchPayments();
        } catch (error) {
            console.error('Failed to update payment', error);
        }
    };

    const handleDeletePayment = async (id) => {
        try {
            await deletePayment(id);
            fetchPayments();
        } catch (error) {
            console.error('Failed to delete payment', error);
        }
    };

    const handleDeleteFeedback = async (id) => {
        try {
            await deleteFeedbackAdmin(id);
            fetchFeedbacks();
        } catch (error) {
            console.error('Failed to delete feedback', error);
        }
    };

    return (
        <div>
            <h1>Admin Dashboard</h1>
            <section>
                <h2>Cars</h2>
                <form onSubmit={handleCreateCar}>
                    <input type="text" placeholder="Brand" value={newCar.brand} onChange={(e) => setNewCar({ ...newCar, brand: e.target.value })} required />
                    <input type="text" placeholder="Model" value={newCar.model} onChange={(e) => setNewCar({ ...newCar, model: e.target.value })} required />
                    <input type="text" placeholder="Transmission" value={newCar.transmission} onChange={(e) => setNewCar({ ...newCar, transmission: e.target.value })} required />
                    <input type="text" placeholder="Fuel Type" value={newCar.fuel_type} onChange={(e) => setNewCar({ ...newCar, fuel_type: e.target.value })} required />
                    <input type="number" placeholder="Price Per Day" value={newCar.price_per_day} onChange={(e) => setNewCar({ ...newCar, price_per_day: parseFloat(e.target.value) })} required />
                    <input type="text" placeholder="Photo URL" value={newCar.photo} onChange={(e) => setNewCar({ ...newCar, photo: e.target.value })} required />
                    <button type="submit">Create Car</button>
                </form>
                <ul>
                    {cars.map(car => (
                        <li key={car.id}>
                            <div>{car.brand} {car.model}</div>
                            <button onClick={() => handleUpdateCar(car.id, car)}>Update</button>
                            <button onClick={() => handleDeleteCar(car.id)}>Delete</button>
                        </li>
                    ))}
                </ul>
            </section>
            <section>
                <h2>Payments</h2>
                <ul>
                    {payments.map(payment => (
                        <li key={payment.id}>
                            <div>Rental ID: {payment.rental_id}</div>
                            <div>Status: {payment.status}</div>
                            <button onClick={() => handleUpdatePayment(payment.id, payment)}>Update</button>
                            <button onClick={() => handleDeletePayment(payment.id)}>Delete</button>
                        </li>
                    ))}
                </ul>
            </section>
            <section>
                <h2>Rentals</h2>
                <ul>
                    {rentals.map(rental => (
                        <li key={rental.id}>
                            <div>Car ID: {rental.car_id}</div>
                            <div>User ID: {rental.user_id}</div>
                            <div>Start Date: {new Date(rental.start_date).toLocaleDateString()}</div>
                            <div>End Date: {new Date(rental.end_date).toLocaleDateString()}</div>
                        </li>
                    ))}
                </ul>
            </section>
            <section>
                <h2>Feedback</h2>
                <ul>
                    {feedbacks.map(feedback => (
                        <li key={feedback.id}>
                            <div>User ID: {feedback.user_id}</div>
                            <div>Car ID: {feedback.car_id}</div>
                            <div>Rating: {feedback.rating}</div>
                            <div>Description: {feedback.description}</div>
                            <button onClick={() => handleDeleteFeedback(feedback.id)}>Delete</button>
                        </li>
                    ))}
                </ul>
            </section>
        </div>
    );
};

export default AdminDashboardPage;