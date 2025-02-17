// src/pages/admin/ManagePayments.js
import React, { useState, useEffect, useCallback } from 'react';
import axios from 'axios';

const ManagePayments = () => {
    const [payments, setPayments] = useState([]);
    const [error, setError] = useState('');
    const token = localStorage.getItem('token');

    // Wrap fetchPayments in useCallback so its reference remains stable.
    const fetchPayments = useCallback(async () => {
        try {
            const response = await axios.get('/admin/payments', {
                headers: { Authorization: `Bearer ${token}` },
            });
            setPayments(response.data);
        } catch (err) {
            setError('Failed to fetch payments.');
        }
    }, [token]);

    // Include fetchPayments in the dependency array
    useEffect(() => {
        fetchPayments();
    }, [fetchPayments]);

    return (
        <div style={{ padding: '2rem' }}>
            <h1>Manage Payments</h1>
            {error && <p style={{ color: 'red' }}>{error}</p>}
            {payments.length === 0 ? (
                <p>No payments found.</p>
            ) : (
                <table border="1" cellPadding="8" cellSpacing="0">
                    <thead>
                    <tr>
                        <th>ID</th>
                        <th>Rental ID</th>
                        <th>Status</th>
                        {/* Add actions columns if needed */}
                    </tr>
                    </thead>
                    <tbody>
                    {payments.map(payment => (
                        <tr key={payment.id}>
                            <td>{payment.id}</td>
                            <td>{payment.rental_id}</td>
                            <td>{payment.status}</td>
                        </tr>
                    ))}
                    </tbody>
                </table>
            )}
        </div>
    );
};

export default ManagePayments;
