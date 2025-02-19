import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import HomePage from './pages/HomePage';
import LoginPage from './pages/LoginPage';
import RegisterPage from './pages/RegisterPage';
import CarsPage from './pages/CarsPage';
import CarDetailsPage from './pages/CarDetailsPage';
import RentalRequestPage from './pages/RentalRequestPage';
import ProfilePage from './pages/ProfilePage';
import UpdateProfilePage from './pages/UpdateProfilePage';
import RentalsPage from './pages/RentalsPage';
import Navbar from './components/Navbar';
import AdminDashboardPage from './pages/AdminDashboardPage';
import AboutUsPage from './pages/AboutUsPage';

const App = () => {
    return (
        <Router>
            <Navbar />
            <Routes>
                <Route path="/" element={<HomePage />} />
                <Route path="/login" element={<LoginPage />} />
                <Route path="/register" element={<RegisterPage />} />
                <Route path="/cars" element={<CarsPage />} />
                <Route path="/cars/:id" element={<CarDetailsPage />} />
                <Route path="/rent/:id" element={<RentalRequestPage />} />
                <Route path="/profile" element={<ProfilePage />} />
                <Route path="/update-profile" element={<UpdateProfilePage />} />
                <Route path="/rentals" element={<RentalsPage />} />
                <Route path="/admin" element={<AdminDashboardPage />} />
                <Route path="/about" element={<AboutUsPage />} />
            </Routes>
        </Router>
    );
};

export default App;
