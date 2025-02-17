// src/pages/admin/ManageCars.js
import React, { useState, useEffect, useCallback } from 'react';
import axios from 'axios';

const ManageCars = () => {
    const [cars, setCars] = useState([]);
    const [error, setError] = useState('');
    const [newCar, setNewCar] = useState({
        brand: '',
        model: '',
        transmission: '',
        fuel_type: '',
        price_per_day: 0,
        available: true,
        photo: ''
    });
    const [editingCar, setEditingCar] = useState(null);
    const token = localStorage.getItem('token');

    // Wrap fetchCars in useCallback so its reference is stable.
    const fetchCars = useCallback(async () => {
        try {
            const response = await axios.get('/cars', {
                headers: { Authorization: `Bearer ${token}` },
            });
            setCars(response.data);
        } catch (err) {
            setError('Failed to fetch cars.');
        }
    }, [token]);

    // Now include fetchCars in the dependency array.
    useEffect(() => {
        fetchCars();
    }, [fetchCars]);

    const handleCreateCar = async () => {
        try {
            await axios.post('/admin/cars', newCar, {
                headers: { Authorization: `Bearer ${token}` }
            });
            fetchCars();
            setNewCar({ brand: '', model: '', transmission: '', fuel_type: '', price_per_day: 0, available: true, photo: '' });
        } catch (err) {
            setError('Failed to create car.');
        }
    };

    const handleDeleteCar = async (id) => {
        try {
            await axios.delete(`/admin/cars/${id}`, {
                headers: { Authorization: `Bearer ${token}` }
            });
            fetchCars();
        } catch (err) {
            setError('Failed to delete car.');
        }
    };

    const handleUpdateCar = async () => {
        try {
            await axios.put(`/admin/cars/${editingCar.id}`, editingCar, {
                headers: { Authorization: `Bearer ${token}` }
            });
            setEditingCar(null);
            fetchCars();
        } catch (err) {
            setError('Failed to update car.');
        }
    };

    return (
        <div style={{ padding: '2rem' }}>
            <h1>Manage Cars</h1>
            {error && <p style={{ color: 'red' }}>{error}</p>}

            <table border="1" cellPadding="8" cellSpacing="0">
                <thead>
                <tr>
                    <th>ID</th>
                    <th>Brand</th>
                    <th>Model</th>
                    <th>Transmission</th>
                    <th>Fuel Type</th>
                    <th>Price/Day</th>
                    <th>Available</th>
                    <th>Photo</th>
                    <th>Actions</th>
                </tr>
                </thead>
                <tbody>
                {cars.map(car => (
                    <tr key={car.id}>
                        <td>{car.id}</td>
                        <td>{car.brand}</td>
                        <td>{car.model}</td>
                        <td>{car.transmission}</td>
                        <td>{car.fuel_type}</td>
                        <td>{car.price_per_day}</td>
                        <td>{car.available ? 'Yes' : 'No'}</td>
                        <td><img src={car.photo} alt={`${car.brand} ${car.model}`} width="100"/></td>
                        <td>
                            <button onClick={() => setEditingCar(car)}>Edit</button>
                            <button onClick={() => handleDeleteCar(car.id)}>Delete</button>
                        </td>
                    </tr>
                ))}
                </tbody>
            </table>

            <h2>Create New Car</h2>
            <div>
                <input type="text" placeholder="Brand" value={newCar.brand} onChange={(e) => setNewCar({ ...newCar, brand: e.target.value })}/>
                <input type="text" placeholder="Model" value={newCar.model} onChange={(e) => setNewCar({ ...newCar, model: e.target.value })}/>
                <input type="text" placeholder="Transmission" value={newCar.transmission} onChange={(e) => setNewCar({ ...newCar, transmission: e.target.value })}/>
                <input type="text" placeholder="Fuel Type" value={newCar.fuel_type} onChange={(e) => setNewCar({ ...newCar, fuel_type: e.target.value })}/>
                <input type="number" placeholder="Price per Day" value={newCar.price_per_day} onChange={(e) => setNewCar({ ...newCar, price_per_day: parseFloat(e.target.value) })}/>
                <input type="text" placeholder="Photo URL" value={newCar.photo} onChange={(e) => setNewCar({ ...newCar, photo: e.target.value })}/>
                <button onClick={handleCreateCar}>Create Car</button>
            </div>

            {editingCar && (
                <div>
                    <h2>Edit Car</h2>
                    <input type="text" placeholder="Brand" value={editingCar.brand} onChange={(e) => setEditingCar({ ...editingCar, brand: e.target.value })}/>
                    <input type="text" placeholder="Model" value={editingCar.model} onChange={(e) => setEditingCar({ ...editingCar, model: e.target.value })}/>
                    <input type="text" placeholder="Transmission" value={editingCar.transmission} onChange={(e) => setEditingCar({ ...editingCar, transmission: e.target.value })}/>
                    <input type="text" placeholder="Fuel Type" value={editingCar.fuel_type} onChange={(e) => setEditingCar({ ...editingCar, fuel_type: e.target.value })}/>
                    <input type="number" placeholder="Price per Day" value={editingCar.price_per_day} onChange={(e) => setEditingCar({ ...editingCar, price_per_day: parseFloat(e.target.value) })}/>
                    <input type="text" placeholder="Photo URL" value={editingCar.photo} onChange={(e) => setEditingCar({ ...editingCar, photo: e.target.value })}/>
                    <button onClick={handleUpdateCar}>Update Car</button>
                    <button onClick={() => setEditingCar(null)}>Cancel</button>
                </div>
            )}
        </div>
    );
};

export default ManageCars;
