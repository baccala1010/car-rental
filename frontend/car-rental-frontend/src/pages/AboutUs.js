// src/pages/AboutUs.js
import React from 'react';

const AboutUs = () => {
    return (
        <div style={{ maxWidth: '800px', margin: 'auto', padding: '2rem' }}>
            <h1>About Us</h1>
            <p>
                Welcome to our Car Rental System! We are dedicated to providing our customers with an exceptional car rental experience. Founded in 2010, our company has grown to become a trusted name in the industry.
            </p>

            <h2>Our Services</h2>
            <p>
                We offer a wide range of vehicles—from compact cars for city driving to luxury vehicles for special occasions. Our fleet is maintained to the highest standards to ensure your safety and comfort.
            </p>

            <img
                src="https://via.placeholder.com/800x400"
                alt="Our Fleet"
                style={{ width: '100%', marginBottom: '2rem' }}
            />

            <h2>Contact Information</h2>
            <p>
                <strong>Phone:</strong> +1 (123) 456-7890<br />
                <strong>Email:</strong> info@carrentalsystem.com<br />
                <strong>Office Address:</strong> 123 Main Street, City, Country
            </p>

            <h2>Our Mission</h2>
            <p>
                Our mission is to provide quality, reliable, and affordable car rental services to ensure a seamless travel experience for all our customers.
            </p>

            <h2>Connect With Us</h2>
            <p>
                Follow us on social media for the latest updates and promotions.
            </p>
        </div>
    );
};

export default AboutUs;
