// src/pages/CarDetail.js
import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { useParams, useNavigate } from 'react-router-dom';

const CarDetail = () => {
    const { id } = useParams(); // Car ID from the route parameter
    const navigate = useNavigate();

    // States for car details and error message
    const [car, setCar] = useState(null);
    const [error, setError] = useState('');

    // Fetch car details (public endpoint, no authentication required)
    useEffect(() => {
        axios
            .get(`http://localhost:8080/api/cars/${id}`)
            .then((response) => setCar(response.data))
            .catch((err) => {
                console.error(err);
                setError('Failed to load car details.');
            });
    }, [id]);

    // Handler for "Rent This Car"
    const handleRent = () => {
        const token = localStorage.getItem('token');
        if (!token) {
            // If no token, redirect to login
            navigate('/login');
        } else {
            // If token exists, navigate to rental page for this car
            navigate(`/car/${car.id}/rent`);
        }
    };

    if (error) {
        return <p style={{ color: 'red' }}>{error}</p>;
    }
    if (!car) {
        return <p>Loading car details...</p>;
    }

    return (
        <div style={{ maxWidth: '800px', margin: 'auto' }}>
            <h2>{car.brand} {car.model}</h2>
            <img
                src={car.photo}
                alt={`${car.brand} ${car.model}`}
                style={{ width: '100%', height: '300px', objectFit: 'cover' }}
            />
            <p><strong>Transmission:</strong> {car.transmission}</p>
            <p><strong>Fuel:</strong> {car.fuel_type}</p>
            <p><strong>Price per Day:</strong> ${car.price_per_day}</p>
            <p><strong>Available:</strong> {car.available ? 'Yes' : 'No'}</p>
            <button onClick={handleRent} disabled={!car.available}>
                Rent This Car
            </button>
            {/* You may also include other public details such as feedback, etc. */}
        </div>
    );
};

export default CarDetail;
