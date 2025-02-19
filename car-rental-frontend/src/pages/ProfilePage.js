// src/pages/ProfilePage.js
import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { parseJwt } from '../utils'; // Correct import

const ProfilePage = () => {
    const [user, setUser] = useState(null);
    const [errorMessage, setErrorMessage] = useState('');
    const navigate = useNavigate();

    useEffect(() => {
        const token = localStorage.getItem('token');
        if (token) {
            // Decode the token for user info.
            const userData = parseJwt(token);
            if (userData) {
                setUser(userData);
            } else {
                setErrorMessage('Failed to load user data.');
            }
        } else {
            setErrorMessage('No token found. Please log in.');
            navigate('/login');
        }
    }, [navigate]);

    if (errorMessage) {
        return <div style={{ color: 'red' }}>{errorMessage}</div>;
    }
    if (!user) {
        return <div>Loading...</div>;
    }

    return (
        <div>
            <h1>Profile</h1>
            {/* The token payload may only contain user_id and role.
          To display full details, the backend must include them in GET /api/user or in the token. */}
            <div>User ID: {user.user_id}</div>
            <div>Role: {user.role}</div>
            <button onClick={() => navigate('/update-profile')}>Update</button>
        </div>
    );
};

export default ProfilePage;
