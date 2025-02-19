import React from 'react';
import { Link } from 'react-router-dom';

const HomePage = () => {
    return (
        <div>
            <h1>Welcome to the Car Rental System</h1>
            <p>Find the best cars for rent at affordable prices.</p>
            <img src="https://myskillsconnect.com/uploads/posts/2023-07/1689244487_myskillsconnect-com-p-foto-avtosalon-shevrole-1.jpg" alt="premium cars"/>
            <Link to="/cars">View Cars</Link>
        </div>
    );
};

export default HomePage;