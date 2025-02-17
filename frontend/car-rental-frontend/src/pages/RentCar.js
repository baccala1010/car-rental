// src/pages/RentCar.js
import React, { useState } from 'react';
import axios from 'axios';
import { useParams, useNavigate } from 'react-router-dom';

const RentCar = () => {
    const { id } = useParams(); // car id
    const navigate = useNavigate();
    const [startDate, setStartDate] = useState('');
    const [endDate, setEndDate]     = useState('');
    const [error, setError]         = useState('');
    const [success, setSuccess]     = useState('');

    const handleRent = async (e) => {
        e.preventDefault();
        const token = localStorage.getItem('token');
        try {
            await axios.post(
                'http://localhost:8080/api/rentals',
                {
                    car_id: parseInt(id, 10),
                    start_date: startDate,
                    end_date: endDate,
                },
                {
                    headers: { Authorization: `Bearer ${token}` },
                }
            );
            setSuccess('Car rented successfully!');
            // Redirect after a short delay (or immediately)
            setTimeout(() => navigate('/'), 2000);
        } catch (err) {
            console.error(err);
            setError('Failed to rent car. Please try again.');
        }
    };

    return (
        <div style={{ maxWidth: '500px', margin: 'auto' }}>
            <h2>Rent Car</h2>
            {error && <p style={{ color: 'red' }}>{error}</p>}
            {success && <p style={{ color: 'green' }}>{success}</p>}
            <form onSubmit={handleRent}>
                <div style={{ marginBottom: '1rem' }}>
                    <label>Start Date (ISO format):</label>
                    <input
                        type="datetime-local"
                        value={startDate}
                        onChange={(e) => setStartDate(e.target.value)}
                        required
                        style={{ width: '100%' }}
                    />
                </div>
                <div style={{ marginBottom: '1rem' }}>
                    <label>End Date (ISO format):</label>
                    <input
                        type="datetime-local"
                        value={endDate}
                        onChange={(e) => setEndDate(e.target.value)}
                        required
                        style={{ width: '100%' }}
                    />
                </div>
                <button type="submit">Confirm Rental</button>
            </form>
        </div>
    );
};

export default RentCar;
