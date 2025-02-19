import React, { useEffect, useState } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import { logout } from '../services/api';
import AdminNavbar from './AdminNavbar';

const Navbar = () => {
    const navigate = useNavigate();
    const token = localStorage.getItem('token');
    const [user, setUser] = useState(null);

    useEffect(() => {
        if (token) {
            const storedUser = localStorage.getItem('user');
            if (storedUser) {
                setUser(JSON.parse(storedUser));
            }
        }
    }, [token]);

    const handleLogout = async () => {
        try {
            await logout();
            localStorage.removeItem('token');
            localStorage.removeItem('user');
            setUser(null);
            navigate('/');
        } catch (error) {
            console.error('Failed to logout', error);
        }
    };

    // Render the admin navbar if the user is an admin.
    if (user && user.role === 'ADMIN') {
        return <AdminNavbar />;
    }

    return (
        <nav>
            <div>
                <Link to="/about">About Us</Link>
                <Link to="/cars">Cars</Link>
            </div>
            <div>
                {!token ? (
                    <>
                        <Link to="/login">Login</Link>
                        <Link to="/register">Register</Link>
                    </>
                ) : (
                    <>
                        <Link to="/rentals">My Rentals</Link>
                        <Link to="/profile">Profile</Link>
                        <button onClick={handleLogout}>Logout</button>
                    </>
                )}
            </div>
        </nav>
    );
};

export default Navbar;
