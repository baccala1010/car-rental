// src/pages/admin/ManageRentals.js
import React, { useEffect, useState, useCallback } from 'react';
import axios from 'axios';

const ManageRentals = () => {
    const [rentals, setRentals] = useState([]);
    const [error, setError] = useState('');
    const token = localStorage.getItem('token');

    const fetchRentals = useCallback(async () => {
        try {
            const response = await axios.get('/admin/rentals', {
                headers: { Authorization: `Bearer ${token}` }
            });
            setRentals(response.data.rentals);
        } catch (err) {
            setError('Failed to fetch rentals.');
        }
    }, [token]);

    useEffect(() => {
        fetchRentals();
    }, [fetchRentals]);

    return (
        <div style={{ padding: '2rem' }}>
            <h1>Manage Rentals</h1>
            {error && <p style={{ color: 'red' }}>{error}</p>}
            {rentals.length === 0 ? (
                <p>No rentals found.</p>
            ) : (
                <table border="1" cellPadding="8" cellSpacing="0">
                    <thead>
                    <tr>
                        <th>ID</th>
                        <th>User ID</th>
                        <th>Car ID</th>
                        <th>Start Date</th>
                        <th>End Date</th>
                    </tr>
                    </thead>
                    <tbody>
                    {rentals.map(rental => (
                        <tr key={rental.id}>
                            <td>{rental.id}</td>
                            <td>{rental.user_id}</td>
                            <td>{rental.car_id}</td>
                            <td>{new Date(rental.start_date).toLocaleString()}</td>
                            <td>{new Date(rental.end_date).toLocaleString()}</td>
                        </tr>
                    ))}
                    </tbody>
                </table>
            )}
        </div>
    );
};

export default ManageRentals;