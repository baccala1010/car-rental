// src/pages/LoginPage.js
import React, { useState } from 'react';
import { login } from '../services/api';
import { useNavigate, Link } from 'react-router-dom';
import { parseJwt } from '../utils'; // Correct import from src/utils.js

const LoginPage = () => {
    const [credentials, setCredentials] = useState({ email: '', password: '' });
    const [errorMessage, setErrorMessage] = useState('');
    const navigate = useNavigate();

    const handleChange = e => {
        const { name, value } = e.target;
        setCredentials(prev => ({ ...prev, [name]: value }));
    };

    const handleSubmit = async e => {
        e.preventDefault();
        try {
            const loginResponse = await login(credentials);
            console.log('Login response:', loginResponse);

            if (!loginResponse.token) {
                setErrorMessage('No token received from the server. Please check your credentials.');
                return;
            }
            // Save the token.
            localStorage.setItem('token', loginResponse.token);

            // Decode the token to get user data.
            const userData = parseJwt(loginResponse.token);
            if (!userData) {
                setErrorMessage('Failed to decode token.');
                return;
            }
            // Save the decoded data (it may only contain user_id and role)
            localStorage.setItem('user', JSON.stringify(userData));

            navigate('/cars');
        } catch (error) {
            console.error('Login failed:', error);
            setErrorMessage('Login failed. Please check your credentials and try again.');
        }
    };

    return (
        <div>
            <h1>Login</h1>
            {errorMessage && <p style={{ color: 'red' }}>{errorMessage}</p>}
            <form onSubmit={handleSubmit}>
                <div>
                    <label>Email:</label>
                    <input
                        type="email"
                        name="email"
                        value={credentials.email}
                        onChange={handleChange}
                        required
                    />
                </div>
                <div>
                    <label>Password:</label>
                    <input
                        type="password"
                        name="password"
                        value={credentials.password}
                        onChange={handleChange}
                        required
                    />
                </div>
                <button type="submit">Login</button>
            </form>
            <p>Don't have an account? <Link to="/register">Register here</Link></p>
        </div>
    );
};

export default LoginPage;
