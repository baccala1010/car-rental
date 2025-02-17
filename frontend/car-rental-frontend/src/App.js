import React, { useState, useEffect } from 'react';
import { BrowserRouter as Router, Routes, Route, Navigate } from 'react-router-dom';
import axios from 'axios';
import Navbar from './pages/Navbar';
import CarList from './pages/CarList';
import CarDetail from './pages/CarDetail';
import RentalPage from './pages/RentalPage';
import UserProfile from './pages/UserProfile';
import AboutUs from './pages/AboutUs';
import Login from './pages/Login';
import Register from './pages/Register';
import AdminDashboard from './pages/admin/AdminDashboard';
import ManageCars from './pages/admin/ManageCars';
import ManageRentals from './pages/admin/ManageRentals';
import ManagePayments from './pages/admin/ManagePayments';
import ManageUsers from './pages/admin/ManageUsers';

const App = () => {
    const [isAuthenticated, setIsAuthenticated] = useState(false);
    const [isAdmin, setIsAdmin] = useState(false);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        const checkAuth = async () => {
            const token = localStorage.getItem('token');
            if (!token) {
                setLoading(false);
                return;
            }

            try {
                const response = await axios.get('/user');
                setIsAuthenticated(true);
                setIsAdmin(response.data.role === 'ADMIN');
                localStorage.setItem('role', response.data.role);
            } catch (error) {
                handleLogout();
            } finally {
                setLoading(false);
            }
        };

        checkAuth();
    }, []);

    const handleLogout = () => {
        localStorage.removeItem('token');
        localStorage.removeItem('role');
        setIsAuthenticated(false);
        setIsAdmin(false);
    };

    if (loading) {
        return <div className="loading-screen">Loading...</div>;
    }

    return (
        <Router>
            <Navbar isAuthenticated={isAuthenticated} isAdmin={isAdmin} onLogout={handleLogout} />
            <Routes>
                <Route path="/" element={<CarList />} />
                <Route path="/cars/:id" element={<CarDetail />} />
                <Route path="/about" element={<AboutUs />} />
                <Route path="/login" element={<Login setIsAuthenticated={setIsAuthenticated} setIsAdmin={setIsAdmin} />} />
                <Route path="/register" element={<Register />} />

                <Route
                    path="/rentals"
                    element={isAuthenticated ? <RentalPage /> : <Navigate to="/login" replace />}
                />
                <Route
                    path="/profile"
                    element={isAuthenticated ? <UserProfile /> : <Navigate to="/login" replace />}
                />

                {/* Admin Routes */}
                <Route
                    path="/admin"
                    element={isAdmin ? <AdminDashboard /> : <Navigate to="/" replace />}
                />
                <Route
                    path="/admin/cars"
                    element={isAdmin ? <ManageCars /> : <Navigate to="/" replace />}
                />
                <Route
                    path="/admin/rentals"
                    element={isAdmin ? <ManageRentals /> : <Navigate to="/" replace />}
                />
                <Route
                    path="/admin/payments"
                    element={isAdmin ? <ManagePayments /> : <Navigate to="/" replace />}
                />
                <Route
                    path="/admin/users"
                    element={isAdmin ? <ManageUsers /> : <Navigate to="/" replace />}
                />

                <Route path="*" element={<Navigate to="/" replace />} />
            </Routes>
        </Router>
    );
};

export default App;