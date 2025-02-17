// src/pages/admin/AdminDashboard.js
import React from 'react';
import { Link } from 'react-router-dom';

const AdminDashboard = () => {
    return (
        <div style={{ padding: '2rem' }}>
            <h1>Admin Dashboard</h1>
            <ul>
                <li><Link to="/admin/cars">Manage Cars</Link></li>
                <li><Link to="/admin/rentals">Manage Rentals</Link></li>
                <li><Link to="/admin/payments">Manage Payments</Link></li>
                <li><Link to="/admin/users">Manage Users</Link></li>
            </ul>
        </div>
    );
};

export default AdminDashboard;
