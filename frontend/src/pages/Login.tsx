import React, { useState } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import { useAuth } from '../contexts/AuthContext';
import { LoginRequest } from '../types';
import Logo from '../components/Logo';

const Login: React.FC = () => {
  const [formData, setFormData] = useState<LoginRequest>({
    email: '',
    password: '',
  });
  const [error, setError] = useState('');
  const [loading, setLoading] = useState(false);

  const { login } = useAuth();
  const navigate = useNavigate();

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setFormData({
      ...formData,
      [e.target.name]: e.target.value,
    });
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setLoading(true);
    setError('');

    try {
      await login(formData);
      navigate('/dashboard');
    } catch (err: any) {
      setError(err.response?.data?.error || 'Login failed');
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="min-h-screen bg-gradient-to-br from-yellow-200 via-pink-200 to-purple-200 flex items-center justify-center p-4">
      {/* Decorative elements */}
      <div className="absolute top-10 left-10 w-20 h-20 bg-blue-400 border-2 border-black 2"></div>
      <div className="absolute bottom-20 right-20 w-16 h-16 bg-green-400 border-2 border-black -2"></div>
      <div className="absolute top-1/3 right-1/4 w-12 h-12 bg-pink-400 border-2 border-black 5"></div>

      <div className="relative w-full max-w-md">
        {/* Main card */}
        <div className="card-brutal p-8 ">
          <div className="text-center mb-8">
            <Logo size="large" />
            <p className="text-lg font-bold uppercase tracking-wider text-gray-700 mt-4">
              Bill Reminder
            </p>
          </div>

          {error && (
            <div className="alert-brutal alert-brutal-error mb-6 ">
              {error}
            </div>
          )}

          <form onSubmit={handleSubmit} className="space-y-6">
            <div>
              <label htmlFor="email" className="block text-sm font-bold uppercase mb-2">
                Email Address
              </label>
              <input
                id="email"
                name="email"
                type="email"
                required
                value={formData.email}
                onChange={handleChange}
                className="input-brutal w-full"
                placeholder="you@email.com"
              />
            </div>

            <div>
              <label htmlFor="password" className="block text-sm font-bold uppercase mb-2">
                Password
              </label>
              <input
                id="password"
                name="password"
                type="password"
                required
                value={formData.password}
                onChange={handleChange}
                className="input-brutal w-full"
                placeholder="••••••••"
              />
            </div>

            <button
              type="submit"
              disabled={loading}
              className="btn-brutal btn-brutal-primary w-full"
            >
              {loading ? (
                <span className="flex items-center justify-center">
                  <span className="animate-pulse">Signing in...</span>
                </span>
              ) : (
                <span className="flex items-center justify-center">
                  Sign In
                  <span className="ml-2">→</span>
                </span>
              )}
            </button>
          </form>

          <div className="mt-8 text-center">
            <Link
              to="/register"
              className="inline-block font-bold text-black border-b-2 border-black  "
            >
              No account? Create one!
            </Link>
          </div>
        </div>

        {/* Floating badge */}
        <div className="absolute -top-4 -right-4 bg-gradient-to-r from-pink-400 to-purple-400 text-white px-4 py-2 border-2 border-black font-bold text-sm uppercase 2 shadow-[4px_4px_0px_0px_rgba(0,0,0,1)]">
          v1.0
        </div>
      </div>
    </div>
  );
};

export default Login;