// src/pages/CarList.js
import React, { useEffect, useState } from 'react';
import axios from 'axios';
import { Link } from 'react-router-dom';

const CarList = () => {
    const [cars, setCars] = useState([]);
    const [error, setError] = useState('');

    useEffect(() => {
        // Get token from localStorage
        const token = localStorage.getItem('token');
        axios
            .get('http://localhost:8080/api/cars', {
                headers: { Authorization: `Bearer ${token}` },
            })
            .then((response) => setCars(response.data))
            .catch((err) => {
                console.error(err);
                setError('Failed to load cars');
            });
    }, []);

    return (
        <div>
            <h2>Available Cars</h2>
            {error && <p style={{ color: 'red' }}>{error}</p>}
            <div style={{ display: 'flex', flexWrap: 'wrap', gap: '1rem' }}>
                {cars && cars.length > 0 ? (
                    cars.map((car) => (
                        <div
                            key={car.id}
                            style={{
                                border: '1px solid #ccc',
                                padding: '1rem',
                                width: '250px',
                                borderRadius: '5px',
                            }}
                        >
                            <img
                                src={car.photo}
                                alt={`${car.brand} ${car.model}`}
                                style={{ width: '100%', height: '150px', objectFit: 'cover' }}
                            />
                            <h3>
                                {car.brand} {car.model}
                            </h3>
                            <p>
                                {car.transmission} | {car.fuel_type}
                            </p>
                            <p>
                                Price/Day: ${car.price_per_day}
                            </p>
                            {/* Link to car detail page */}
                            <Link to={`/car/${car.id}`}>
                                <button>View Details</button>
                            </Link>
                        </div>
                    ))
                ) : (
                    <p>No cars available</p>
                )}
            </div>
        </div>
    );
};

export default CarList;