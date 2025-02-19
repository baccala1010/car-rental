import axios from 'axios';

const api = axios.create({
    baseURL: 'http://localhost:8080/api',
});

// Request interceptor: Attach the JWT token with "Bearer " prefix.
api.interceptors.request.use(
    config => {
        const token = localStorage.getItem('token');
        if (token) {
            // The backend expects: Authorization: Bearer <token>
            config.headers.Authorization = `Bearer ${token}`;
        }
        return config;
    },
    error => Promise.reject(error)
);

// Response interceptor: Return only the response data.
api.interceptors.response.use(
    response => response.data,
    error => {
        if (error.response && error.response.status === 401) {
            console.error("Unauthorized – check your token or credentials.");
        }
        return Promise.reject(error);
    }
);

// Public endpoints
export const apiStatus = () => api.get('/');
export const listCars = () => api.get('/cars');
export const getCar = id => api.get(`/cars/${id}`);
export const searchCars = criteria => api.get('/cars/search', { params: criteria });
export const listFeedbackByCar = carId => api.get(`/feedback/car/${carId}`);
export const listFeedbackByUser = userId => api.get(`/feedback/user/${userId}`);
export const listAllFeedback = () => api.get('/feedback');
export const login = credentials => api.post('/login', credentials);
export const register = user => api.post('/register', user);

// Protected endpoints (JWT required)
export const rentCar = rental => api.post('/rentals', rental);
export const returnCar = id => api.post(`/rentals/${id}/return`);
export const createFeedback = feedback => api.post('/feedback', feedback);
export const updateUser = user => api.put('/user', user);
export const getUser = () => api.get('/user'); // (Currently this endpoint returns 401)
export const deleteFeedbackClient = id => api.delete(`/feedback/${id}`);
export const getPayment = id => api.get(`/payments/${id}`);
export const logout = () => api.post('/logout');
export const listRentalsByUser = () => api.get('/rentals');

// Admin endpoints (JWT + admin role)
export const createCar = car => api.post('/admin/cars', car);
export const updateCar = (id, car) => api.put(`/admin/cars/${id}`, car);
export const deleteCar = id => api.delete(`/admin/cars/${id}`);
export const deleteFeedbackAdmin = id => api.delete(`/admin/feedback/${id}`);
export const updatePayment = (id, payment) => api.put(`/admin/payments/${id}`, payment);
export const deletePayment = id => api.delete(`/admin/payments/${id}`);
export const listPayments = () => api.get('/admin/payments');
export const listAllRentals = () => api.get('/admin/rentals');

export default api;
