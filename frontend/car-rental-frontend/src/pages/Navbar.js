import React from 'react';
import { Link, useNavigate } from 'react-router-dom';

const Navbar = ({ isAuthenticated, isAdmin, onLogout }) => {
    const navigate = useNavigate();

    const handleProtectedNavigation = (path) => {
        if (!isAuthenticated) {
            navigate('/login');
        } else {
            navigate(path);
        }
    };

    return (
        <nav style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', padding: '1rem', borderBottom: '1px solid #ddd' }}>
            <div>
                <Link to="/" style={{ marginRight: '1rem' }}>All Cars</Link>
                <Link to="/about" style={{ marginRight: '1rem' }}>About Us</Link>
                {isAdmin && (
                    <>
                        <Link to="/admin" style={{ marginRight: '1rem' }}>Admin Dashboard</Link>
                        <Link to="/admin/cars" style={{ marginRight: '1rem' }}>Manage Cars</Link>
                        <Link to="/admin/rentals" style={{ marginRight: '1rem' }}>Manage Rentals</Link>
                        <Link to="/admin/payments" style={{ marginRight: '1rem' }}>Manage Payments</Link>
                        <Link to="/admin/users" style={{ marginRight: '1rem' }}>Manage Users</Link>
                    </>
                )}
            </div>
            <div>
                {isAuthenticated ? (
                    <>
                        <button onClick={() => handleProtectedNavigation('/profile')} style={{ marginRight: '1rem' }}>Profile</button>
                        <button onClick={() => handleProtectedNavigation('/rentals')} style={{ marginRight: '1rem' }}>My Rentals</button>
                        <button onClick={onLogout}>Logout</button>
                    </>
                ) : (
                    <>
                        <Link to="/login" style={{ marginRight: '1rem' }}>Login</Link>
                        <Link to="/register">Register</Link>
                    </>
                )}
            </div>
        </nav>
    );
};

export default Navbar;