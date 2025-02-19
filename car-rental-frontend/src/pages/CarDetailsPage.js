import React, { useState, useEffect } from 'react';
import { getCar, listFeedbackByCar, createFeedback } from '../services/api';
import { useParams, useNavigate } from 'react-router-dom';

const CarDetailsPage = () => {
    const { id } = useParams();
    const navigate = useNavigate();
    const [car, setCar] = useState(null);
    const [feedbacks, setFeedbacks] = useState([]);
    const [rating, setRating] = useState(1);
    const [description, setDescription] = useState('');
    const [error, setError] = useState('');

    useEffect(() => {
        const fetchCar = async () => {
            try {
                const response = await getCar(id);
                if (!response) {
                    throw new Error('No car data received.');
                }
                // Since the interceptor already returns the data directly, no need to use response.data
                setCar(response);
            } catch (err) {
                console.error('Failed to fetch car details', err);
                setError('Failed to fetch car details.');
            }
        };

        const fetchFeedbacks = async () => {
            try {
                const feedbackResponse = await listFeedbackByCar(id);
                // If feedbackResponse is null, default to an empty array.
                setFeedbacks(feedbackResponse || []);
            } catch (err) {
                console.error('Failed to fetch feedbacks', err);
                setError('Failed to fetch feedbacks.');
            }
        };

        fetchCar();
        fetchFeedbacks();
    }, [id]);

    const handleFeedbackSubmit = async (e) => {
        e.preventDefault();
        const token = localStorage.getItem('token');
        if (!token) {
            navigate('/login');
            return;
        }

        try {
            await createFeedback({ car_id: id, rating, description });
            setDescription('');
            setRating(1);
            // Re-fetch feedbacks after successfully creating one.
            const feedbackResponse = await listFeedbackByCar(id);
            setFeedbacks(feedbackResponse || []);
        } catch (err) {
            console.error('Failed to create feedback', err);
            setError('Failed to create feedback.');
        }
    };

    if (error) {
        return <div style={{ color: 'red' }}>{error}</div>;
    }

    if (!car) {
        return <div>Loading...</div>;
    }

    return (
        <div>
            <h1>
                {car.brand} {car.model}
            </h1>
            <img
                src={car.photo}
                alt={`${car.brand} ${car.model}`}
                style={{ width: '400px', height: 'auto' }}
            />
            <div>Transmission: {car.transmission}</div>
            <div>Fuel Type: {car.fuel_type}</div>
            <div>Price Per Day: ${car.price_per_day}</div>
            <div>Available: {car.available ? 'Yes' : 'No'}</div>
            {car.available && <button onClick={() => navigate(`/rent/${id}`)}>Rent</button>}

            <h2>Feedback</h2>
            <form onSubmit={handleFeedbackSubmit}>
                <div>
                    <label>Rating:</label>
                    <input type="radio" name="rating" value="1" checked={rating === 1} onChange={() => setRating(1)} /> 1
                    <input type="radio" name="rating" value="2" checked={rating === 2} onChange={() => setRating(2)} /> 2
                    <input type="radio" name="rating" value="3" checked={rating === 3} onChange={() => setRating(3)} /> 3
                    <input type="radio" name="rating" value="4" checked={rating === 4} onChange={() => setRating(4)} /> 4
                    <input type="radio" name="rating" value="5" checked={rating === 5} onChange={() => setRating(5)} /> 5
                </div>
                <div>
                    <label>Description:</label>
                    <textarea
                        value={description}
                        onChange={(e) => setDescription(e.target.value)}
                        required
                    />
                </div>
                <button type="submit">Submit Feedback</button>
            </form>

            <ul>
                {feedbacks.map(feedback => (
                    <li key={feedback.id}>
                        <div>Rating: {feedback.rating}</div>
                        <div>{feedback.description}</div>
                    </li>
                ))}
            </ul>
        </div>
    );
};

export default CarDetailsPage;
