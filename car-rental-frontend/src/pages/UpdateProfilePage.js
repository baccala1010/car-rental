import React, { useState, useEffect } from 'react';
import { getUser, updateUser } from '../services/api';
import { useNavigate } from 'react-router-dom';

const UpdateProfilePage = () => {
    const [user, setUser] = useState({ name: '', email: '', phone: '' });
    const navigate = useNavigate();

    useEffect(() => {
        const fetchUser = async () => {
            try {
                const response = await getUser();
                setUser(response.data);
            } catch (error) {
                console.error('Failed to fetch user data', error);
            }
        };

        fetchUser();
    }, []);

    const handleChange = (e) => {
        const { name, value } = e.target;
        setUser({ ...user, [name]: value });
    };

    const handleSubmit = async (e) => {
        e.preventDefault();
        try {
            await updateUser(user);
            navigate('/profile');
        } catch (error) {
            console.error('Failed to update user data', error);
        }
    };

    return (
        <div>
            <h1>Update Profile</h1>
            <form onSubmit={handleSubmit}>
                <div>
                    <label>Name:</label>
                    <input type="text" name="name" value={user.name} onChange={handleChange} required />
                </div>
                <div>
                    <label>Email:</label>
                    <input type="email" name="email" value={user.email} onChange={handleChange} required />
                </div>
                <div>
                    <label>Phone:</label>
                    <input type="text" name="phone" value={user.phone} onChange={handleChange} />
                </div>
                <button type="submit">Update</button>
            </form>
        </div>
    );
};

export default UpdateProfilePage;