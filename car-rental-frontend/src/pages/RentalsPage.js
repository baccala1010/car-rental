import React, { useState, useEffect } from 'react';
import { listRentalsByUser, returnCar } from '../services/api';
import { useNavigate } from 'react-router-dom';

const RentalsPage = () => {
    const [rentals, setRentals] = useState([]);
    const [errorMessage, setErrorMessage] = useState('');
    const navigate = useNavigate();

    // Define the handleReturn function that uses returnCar.
    const handleReturn = async rentalId => {
        try {
            await returnCar(rentalId);
            setRentals(prev => prev.filter(rental => rental.id !== rentalId));
        } catch (error) {
            console.error('Failed to return car', error);
        }
    };

    useEffect(() => {
        const fetchRentals = async () => {
            try {
                const response = await listRentalsByUser();
                console.log('Rentals response:', response);
                // If the response is an array or wrapped inside an object.
                const rentalsData = Array.isArray(response) ? response : response.rentals;
                setRentals(rentalsData || []);
            } catch (error) {
                console.error('Failed to fetch rentals', error);
                setErrorMessage('You are not authorized to view rentals. Please contact support or try logging in again.');
                // Optionally, redirect to login:
                // navigate('/login');
            }
        };

        fetchRentals();
    }, [navigate]);

    if (errorMessage) {
        return <div style={{ color: 'red' }}>{errorMessage}</div>;
    }

    return (
        <div>
            <h1>My Rentals</h1>
            {rentals && rentals.length > 0 ? (
                <ul>
                    {rentals.map(rental => (
                        <li key={rental.id}>
                            <div>Car ID: {rental.car_id}</div>
                            <div>Payment ID: {rental.payment_id}</div>
                            <div>Start Date: {new Date(rental.start_date).toLocaleDateString()}</div>
                            <div>End Date: {new Date(rental.end_date).toLocaleDateString()}</div>
                            <button onClick={() => handleReturn(rental.id)}>Return</button>
                        </li>
                    ))}
                </ul>
            ) : (
                <p>No rentals found.</p>
            )}
        </div>
    );
};

export default RentalsPage;
