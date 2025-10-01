import React, { useState } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import { useAuth } from '../contexts/AuthContext';
import { RegisterRequest } from '../types';
import Logo from '../components/Logo';

const Register: React.FC = () => {
  const [formData, setFormData] = useState<RegisterRequest>({
    name: '',
    email: '',
    password: '',
    confirm_password: '',
  });
  const [error, setError] = useState('');
  const [loading, setLoading] = useState(false);

  const { register } = useAuth();
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

    if (formData.password !== formData.confirm_password) {
      setError('Passwords do not match');
      setLoading(false);
      return;
    }

    try {
      await register(formData);
      navigate('/dashboard');
    } catch (err: any) {
      setError(err.response?.data?.error || 'Registration failed');
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="min-h-screen bg-gradient-to-br from-green-200 via-blue-200 to-purple-200 flex items-center justify-center p-4">
      {/* Decorative elements */}
      <div className="absolute top-20 left-20 w-24 h-24 bg-yellow-400 border-2 border-black -2"></div>
      <div className="absolute bottom-10 right-10 w-20 h-20 bg-pink-400 border-2 border-black 2"></div>
      <div className="absolute top-1/4 left-1/4 w-16 h-16 bg-blue-400 border-2 border-black 5"></div>

      <div className="relative w-full max-w-md">
        {/* Main card */}
        <div className="card-brutal p-8 ">
          <div className="text-center mb-8">
            <Logo size="large" />
            <p className="text-lg font-bold uppercase tracking-wider text-gray-700 mt-4">
              Register
            </p>
          </div>

          {error && (
            <div className="alert-brutal alert-brutal-error mb-6 ">
              {error}
            </div>
          )}

          <form onSubmit={handleSubmit} className="space-y-4">
            <div>
              <label htmlFor="name" className="block text-sm font-bold uppercase mb-2">
                Full Name
              </label>
              <input
                id="name"
                name="name"
                type="text"
                required
                value={formData.name}
                onChange={handleChange}
                className="input-brutal w-full"
                placeholder="John Doe"
              />
            </div>

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

            <div>
              <label htmlFor="confirm_password" className="block text-sm font-bold uppercase mb-2">
                Confirm Password
              </label>
              <input
                id="confirm_password"
                name="confirm_password"
                type="password"
                required
                value={formData.confirm_password}
                onChange={handleChange}
                className="input-brutal w-full"
                placeholder="••••••••"
              />
            </div>

            <button
              type="submit"
              disabled={loading}
              className="btn-brutal btn-brutal-primary w-full mt-6"
            >
              {loading ? (
                <span className="flex items-center justify-center">
                  <span className="animate-pulse">Creating Account...</span>
                </span>
              ) : (
                <span className="flex items-center justify-center">
                  Create Account
                  <span className="ml-2">✨</span>
                </span>
              )}
            </button>
          </form>

          <div className="mt-8 text-center">
            <Link
              to="/login"
              className="inline-block font-bold text-black border-b-2 border-black  "
            >
              Already have an account? Sign In!
            </Link>
          </div>
        </div>

        {/* Floating badge */}
        <div className="absolute -top-4 -left-4 bg-gradient-to-r from-green-400 to-blue-400 text-white px-4 py-2 border-2 border-black font-bold text-sm uppercase 2 shadow-[4px_4px_0px_0px_rgba(0,0,0,1)]">
          New
        </div>
      </div>
    </div>
  );
};

export default Register;