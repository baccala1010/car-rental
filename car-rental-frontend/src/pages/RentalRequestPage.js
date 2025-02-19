import React, { useState } from 'react';
import { rentCar } from '../services/api';
import { useParams, useNavigate } from 'react-router-dom';

const RentalRequestPage = () => {
    const { id } = useParams();
    const navigate = useNavigate();
    const [startDate, setStartDate] = useState('');
    const [endDate, setEndDate] = useState('');

    const handleSubmit = async (e) => {
        e.preventDefault();
        // Check token explicitly
        const token = localStorage.getItem('token');
        if (!token) {
            navigate('/login');
            return;
        }
        try {
            const rentalPayload = {
                car_id: parseInt(id, 10),
                start_date: new Date(startDate).toISOString(),
                end_date: new Date(endDate).toISOString(),
            };
            await rentCar(rentalPayload);
            alert('Car rented successfully!');
            navigate('/cars');
        } catch (error) {
            console.error('Failed to rent car', error);
            alert('Failed to rent car');
        }
    };

    return (
        <div>
            <h1>Rental Request</h1>
            <form onSubmit={handleSubmit}>
                <div>
                    <label>Start Date:</label>
                    <input
                        type="date"
                        value={startDate}
                        onChange={(e) => setStartDate(e.target.value)}
                        required
                    />
                </div>
                <div>
                    <label>End Date:</label>
                    <input
                        type="date"
                        value={endDate}
                        onChange={(e) => setEndDate(e.target.value)}
                        required
                    />
                </div>
                <button type="submit">Submit</button>
            </form>
        </div>
    );
};

export default RentalRequestPage;
