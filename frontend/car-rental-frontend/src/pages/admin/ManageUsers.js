// src/pages/admin/ManageUsers.js
import React, { useEffect, useState, useCallback } from 'react';
import axios from 'axios';

const ManageUsers = () => {
    const [users, setUsers] = useState([]);
    const [searchEmail, setSearchEmail] = useState('');
    const [error, setError] = useState('');
    const token = localStorage.getItem('token');

    const fetchUsers = useCallback(async () => {
        try {
            const response = await axios.get('/admin/users', {
                headers: { Authorization: `Bearer ${token}` }
            });
            setUsers(response.data);
        } catch (err) {
            setError('Failed to fetch users.');
        }
    }, [token]);

    useEffect(() => {
        fetchUsers();
    }, [fetchUsers]);

    const handleSearch = async () => {
        try {
            const response = await axios.get(`/admin/users?email=${searchEmail}`, {
                headers: { Authorization: `Bearer ${token}` }
            });
            setUsers(response.data);
        } catch (err) {
            setError('Search failed.');
        }
    };

    return (
        <div style={{ padding: '2rem' }}>
            <h1>Manage Users</h1>
            <div>
                <input
                    type="text"
                    placeholder="Search by email"
                    value={searchEmail}
                    onChange={(e) => setSearchEmail(e.target.value)}
                />
                <button onClick={handleSearch}>Search</button>
            </div>
            {error && <p style={{ color: 'red' }}>{error}</p>}
            {users.length === 0 ? (
                <p>No users found.</p>
            ) : (
                <table border="1" cellPadding="8" cellSpacing="0">
                    <thead>
                    <tr>
                        <th>ID</th>
                        <th>Name</th>
                        <th>Email</th>
                        <th>Phone</th>
                        <th>Role</th>
                    </tr>
                    </thead>
                    <tbody>
                    {users.map(user => (
                        <tr key={user.id}>
                            <td>{user.id}</td>
                            <td>{user.name}</td>
                            <td>{user.email}</td>
                            <td>{user.phone}</td>
                            <td>{user.role}</td>
                        </tr>
                    ))}
                    </tbody>
                </table>
            )}
        </div>
    );
};

export default ManageUsers;