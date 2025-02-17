import React, { useState, useEffect } from 'react';
import axios from 'axios';

const RentalPage = () => {
    const [rentals, setRentals] = useState([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState('');

    const token = localStorage.getItem('token');

    useEffect(() => {
        const fetchRentals = async () => {
            try {
                const response = await axios.get('http://localhost:8080/api/rentals', {
                    headers: { Authorization: `Bearer ${token}` },
                });
                setRentals(response.data || []);
                setLoading(false);
            } catch (err) {
                console.error(err);
                setError('Failed to load rentals.');
                setLoading(false);
            }
        };

        fetchRentals();
    }, [token]);

    const handleReturnRental = async (rentalId) => {
        try {
            await axios.post(
                `http://localhost:8080/api/rentals/${rentalId}/return`,
                {},
                {
                    headers: { Authorization: `Bearer ${token}` },
                }
            );
            const response = await axios.get('http://localhost:8080/api/rentals', {
                headers: { Authorization: `Bearer ${token}` },
            });
            setRentals(response.data || []);
        } catch (err) {
            console.error(err);
            setError('Failed to return the rental.');
        }
    };

    if (loading) {
        return <p>Loading your rentals...</p>;
    }
    if (error) {
        return <p style={{ color: 'red' }}>{error}</p>;
    }
    if (!rentals.length) {
        return <p>You have no current rentals.</p>;
    }

    return (
        <div style={{ maxWidth: '800px', margin: 'auto' }}>
            <h2>Your Rentals</h2>
            <table style={{ width: '100%', borderCollapse: 'collapse' }}>
                <thead>
                <tr>
                    <th style={{ border: '1px solid #ccc', padding: '0.5rem' }}>Rental ID</th>
                    <th style={{ border: '1px solid #ccc', padding: '0.5rem' }}>Car ID</th>
                    <th style={{ border: '1px solid #ccc', padding: '0.5rem' }}>Rental Period</th>
                    <th style={{ border: '1px solid #ccc', padding: '0.5rem' }}>Actions</th>
                </tr>
                </thead>
                <tbody>
                {rentals.map((rental) => (
                    <tr key={rental.id}>
                        <td style={{ border: '1px solid #ccc', padding: '0.5rem' }}>{rental.id}</td>
                        <td style={{ border: '1px solid #ccc', padding: '0.5rem' }}>{rental.car_id}</td>
                        <td style={{ border: '1px solid #ccc', padding: '0.5rem' }}>
                            {new Date(rental.start_date).toLocaleString()} &mdash; {new Date(rental.end_date).toLocaleString()}
                        </td>
                        <td style={{ border: '1px solid #ccc', padding: '0.5rem' }}>
                            <button onClick={() => handleReturnRental(rental.id)}>
                                Return Car
                            </button>
                        </td>
                    </tr>
                ))}
                </tbody>
            </table>
        </div>
    );
};

export default RentalPage;