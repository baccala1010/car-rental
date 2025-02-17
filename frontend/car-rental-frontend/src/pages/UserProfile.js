import React, { useState, useEffect } from 'react';
import axios from 'axios';

const UserProfile = () => {
    const [user, setUser] = useState(null);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState('');
    const [isEditing, setIsEditing] = useState(false);

    const [name, setName] = useState('');
    const [email, setEmail] = useState('');
    const [phone, setPhone] = useState('');

    const token = localStorage.getItem('token');

    useEffect(() => {
        const fetchUserProfile = async () => {
            try {
                const response = await axios.get('http://localhost:8080/api/user', {
                    headers: { Authorization: `Bearer ${token}` },
                });
                console.log(response.data); // Log the response data
                setUser (response.data);
                setName(response.data.name);
                setEmail(response.data.email);
                setPhone(response.data.phone);
            } catch (err) {
                console.error(err);
                setError('Failed to load user profile.');
            } finally {
                setLoading(false);
            }
        };

        fetchUserProfile();
    }, [token]);

    const handleEditToggle = () => {
        setIsEditing(!isEditing);
    };

    const handleUpdateProfile = async (e) => {
        e.preventDefault();
        try {
            const updatedData = { name, email, phone };
            const response = await axios.put(
                'http://localhost:8080/api/user',
                updatedData,
                {
                    headers: { Authorization: `Bearer ${token}` },
                }
            );
            setUser(response.data);
            setIsEditing(false);
        } catch (err) {
            console.error(err);
            setError('Failed to update profile.');
        }
    };

    if (loading) {
        return <p>Loading your profile...</p>;
    }
    if (error) {
        return <p style={{ color: 'red' }}>{error}</p>;
    }
    if (!user) {
        return <p>No user data available.</p>;
    }

    return (
        <div style={{ maxWidth: '600px', margin: 'auto' }}>
            <h2>User Profile</h2>
            {isEditing ? (
                <form onSubmit={handleUpdateProfile}>
                    <div style={{ marginBottom: '1rem' }}>
                        <label>Name:</label>
                        <input
                            type="text"
                            value={name}
                            onChange={(e) => setName(e.target.value)}
                            required
                            style={{ width: '100%' }}
                        />
                    </div>
                    <div style={{ marginBottom: '1rem' }}>
                        <label>Email:</label>
                        <input
                            type="email"
                            value={email}
                            onChange={(e) => setEmail(e.target.value)}
                            required
                            style={{ width: '100%' }}
                        />
                    </div>
                    <div style={{ marginBottom: '1rem' }}>
                        <label>Phone:</label>
                        <input
                            type="text"
                            value={phone}
                            onChange={(e) => setPhone(e.target.value)}
                            style={{ width: '100%' }}
                        />
                    </div>
                    <button type="submit">Update Profile</button>
                    <button type="button" onClick={handleEditToggle} style={{ marginLeft: '1rem' }}>
                        Cancel
                    </button>
                </form>
            ) : (
                <div>
                    <p><strong>ID:</strong> {user.id}</p>
                    <p><strong>Name:</strong> {user.name}</p>
                    <p><strong>Email:</strong> {user.email}</p>
                    <p><strong>Phone:</strong> {user.phone}</p>
                    <p><strong>Role:</strong> {user.role}</p>
                    <button onClick={handleEditToggle}>Edit Profile</button>
                </div>
            )}
        </div>
    );
};

export default UserProfile;