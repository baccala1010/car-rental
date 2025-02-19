import React from 'react';
import { Link, useNavigate } from 'react-router-dom';
import { logout } from '../services/api';

const AdminNavbar = () => {
    const navigate = useNavigate();

    const handleLogout = async () => {
        try {
            await logout();
            localStorage.removeItem('token');
            localStorage.removeItem('user');
            navigate('/');
        } catch (error) {
            console.error('Failed to logout', error);
        }
    };

    return (
        <nav>
            <div>
                <Link to="/admin">Admin Dashboard</Link>
                <Link to="/about">About Us</Link>
                <Link to="/cars">Cars</Link>
            </div>
            <div>
                <button onClick={handleLogout}>Logout</button>
            </div>
        </nav>
    );
};

export default AdminNavbar;
