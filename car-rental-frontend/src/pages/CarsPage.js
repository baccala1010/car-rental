import React, { useState, useEffect } from 'react';
import { listCars } from '../services/api';
import { useNavigate } from 'react-router-dom';

const CarsPage = () => {
    const [cars, setCars] = useState([]); // Initialize as an empty array
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState('');
    const navigate = useNavigate();

    useEffect(() => {
        const fetchCars = async () => {
            try {
                // Fetch the list of cars from the API
                const response = await listCars();
                console.log('Cars response:', response);

                // If the response is wrapped in an object (e.g. { cars: [...] }),
                // then adjust accordingly. For now, we assume it is an array.
                const carsData = Array.isArray(response) ? response : response.cars;
                setCars(carsData || []);
            } catch (err) {
                console.error('Error fetching cars:', err);
                setError('Error fetching cars.');
            } finally {
                setLoading(false);
            }
        };

        fetchCars();
    }, []);

    if (loading) {
        return <div>Loading...</div>;
    }
    if (error) {
        return <div style={{ color: 'red' }}>{error}</div>;
    }

    return (
        <div>
            <h1>Cars</h1>
            {cars.length > 0 ? (
                <ul style={{ listStyle: 'none', padding: 0 }}>
                    {cars.map(car => (
                        <li
                            key={car.id}
                            style={{
                                border: '1px solid #ccc',
                                marginBottom: '1rem',
                                padding: '1rem',
                                display: 'flex',
                                alignItems: 'center'
                            }}
                        >
                            <img
                                src={car.photo}
                                alt={`${car.brand} ${car.model}`}
                                style={{ width: '150px', height: 'auto', marginRight: '1rem' }}
                            />
                            <div style={{ flex: '1' }}>
                                <h2>
                                    {car.brand} {car.model}
                                </h2>
                                <p>Fuel Type: {car.fuel_type || 'N/A'}</p>
                                <p>Transmission: {car.transmission}</p>
                                <p>Price per day: ${car.price_per_day}</p>
                                <p>Available: {car.available ? 'Yes' : 'No'}</p>
                            </div>
                            <div>
                                {/* Details button to navigate to the car's detail page */}
                                <button onClick={() => navigate(`/cars/${car.id}`)}>
                                    Details
                                </button>
                            </div>
                        </li>
                    ))}
                </ul>
            ) : (
                <p>No cars available.</p>
            )}
        </div>
    );
};

export default CarsPage;
