import axios from 'axios';
import { LoginRequest, RegisterRequest, AuthResponse, Bill, BillCreateRequest, BillUpdateRequest, DashboardResponse } from '../types';

const API_BASE_URL = process.env.REACT_APP_API_URL || 'http://localhost:8080/api/v1';

const api = axios.create({
  baseURL: API_BASE_URL,
  headers: {
    'Content-Type': 'application/json',
  },
});

// Add token to requests
api.interceptors.request.use((config) => {
  const token = localStorage.getItem('token');
  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
});

// Handle token expiration
api.interceptors.response.use(
  (response) => response,
  (error) => {
    if (error.response?.status === 401) {
      localStorage.removeItem('token');
      localStorage.removeItem('user');
      window.location.href = '/login';
    }
    return Promise.reject(error);
  }
);

export const authApi = {
  login: (data: LoginRequest): Promise<AuthResponse> => api.post('/login', data).then(res => res.data),
  register: (data: RegisterRequest): Promise<AuthResponse> => api.post('/register', data).then(res => res.data),
  getProfile: (): Promise<any> => api.get('/profile').then(res => res.data),
};

export const billApi = {
  getBills: (): Promise<Bill[]> => api.get('/bills').then(res => res.data),
  getBill: (id: number): Promise<Bill> => api.get(`/bills/${id}`).then(res => res.data),
  createBill: (data: BillCreateRequest): Promise<Bill> => api.post('/bills', data).then(res => res.data),
  updateBill: (id: number, data: BillUpdateRequest): Promise<Bill> => api.put(`/bills/${id}`, data).then(res => res.data),
  deleteBill: (id: number): Promise<void> => api.delete(`/bills/${id}`).then(res => res.data),
  getUpcomingBills: (): Promise<Bill[]> => api.get('/bills/upcoming').then(res => res.data),
};

export const dashboardApi = {
  getDashboard: (): Promise<DashboardResponse> => api.get('/dashboard').then(res => res.data),
};

export default api;