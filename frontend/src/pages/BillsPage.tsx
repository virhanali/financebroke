import React, { useState, useEffect } from 'react';
import { Link, useLocation } from 'react-router-dom';
import { Bill, BillCreateRequest } from '../types';
import { billApi } from '../services/api';
import Logo from '../components/Logo';

const BillsPage: React.FC = () => {
  const location = useLocation();
  const [bills, setBills] = useState<Bill[] | null>(null);
  const [loading, setLoading] = useState(true);
  const [showAddForm, setShowAddForm] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [formData, setFormData] = useState<BillCreateRequest>({
    name: '',
    amount: 0,
    due_date: '',
    description: '',
    remind_before: 3,
  });

  useEffect(() => {
    fetchBills();

    // Check if we should show add form (from navigation)
    if (location.state?.showAddForm) {
      setShowAddForm(true);
    }
  }, [location.state]);

  const fetchBills = async () => {
    try {
      setError(null);
      const data = await billApi.getBills();
      setBills(data);
    } catch (err: any) {
      console.error('Failed to fetch bills:', err);
      setError(err.response?.data?.error || 'Failed to fetch bills');
      setBills([]); // Set to empty array on error
    } finally {
      setLoading(false);
    }
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      setError(null);
      await billApi.createBill(formData);
      setFormData({
        name: '',
        amount: 0,
        due_date: '',
        description: '',
        remind_before: 3,
      });
      setShowAddForm(false);
      fetchBills();
    } catch (err: any) {
      console.error('Failed to create bill:', err);
      setError(err.response?.data?.error || 'Failed to create bill');
    }
  };

  const handleDelete = async (id: number) => {
    if (window.confirm('Are you sure you want to delete this bill?')) {
      try {
        setError(null);
        await billApi.deleteBill(id);
        fetchBills();
      } catch (err: any) {
        console.error('Failed to delete bill:', err);
        setError(err.response?.data?.error || 'Failed to delete bill');
      }
    }
  };

  const markAsPaid = async (id: number) => {
    try {
      setError(null);
      await billApi.updateBill(id, { status: 'paid' });
      fetchBills();
    } catch (err: any) {
      console.error('Failed to update bill:', err);
      setError(err.response?.data?.error || 'Failed to update bill');
    }
  };

  const formatDate = (dateString: string) => {
    return new Date(dateString).toLocaleDateString('id-ID', {
      year: 'numeric',
      month: 'long',
      day: 'numeric',
    });
  };

  const getStatusColor = (status: string) => {
    switch (status) {
      case 'paid':
        return 'badge-brutal-green';
      case 'overdue':
        return 'badge-brutal-red';
      default:
        return 'badge-brutal-yellow';
    }
  };

  if (loading) {
    return (
      <div className="min-h-screen bg-gradient-to-br from-blue-100 via-purple-100 to-pink-100 flex items-center justify-center">
        <div className="card-brutal p-8 ">
          <div className="text-2xl font-bold uppercase">Loading Bills...</div>
        </div>
      </div>
    );
  }

  return (
    <div className="min-h-screen bg-gradient-to-br from-blue-100 via-purple-100 to-pink-100">
      {/* Navigation */}
      <nav className="nav-brutal sticky top-0 z-50">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="flex justify-between h-16 items-center">
            <div className="flex items-center space-x-4">
              <Link to="/dashboard" className="btn-brutal">
                ← Dashboard
              </Link>
            </div>
            <div className="flex items-center space-x-4">
              <div className="flex items-center space-x-2">
                <Logo size="medium" />
                <span className="text-2xl font-black">BILLS.</span>
              </div>
              <span className="badge-brutal badge-brutal-blue">Manage</span>
            </div>
          </div>
        </div>
      </nav>

      {/* Main content */}
      <main className="max-w-7xl mx-auto py-8 px-4 sm:px-6 lg:px-8">
        {/* Decorative elements */}
        <div className="absolute top-32 right-20 w-20 h-20 bg-yellow-400 border-2 border-black -z-10"></div>
        <div className="absolute bottom-40 left-10 w-16 h-16 bg-green-400 border-2 border-black -z-10"></div>

        {/* Error Display */}
        {error && (
          <div className="card-brutal p-6 mb-8  bg-red-100 border-red-500">
            <div className="flex items-center">
              <span className="text-2xl mr-3">⚠️</span>
              <div>
                <h3 className="font-bold text-lg uppercase">Error!</h3>
                <p className="font-medium">{error}</p>
              </div>
            </div>
          </div>
        )}

        {/* Add Bill Button */}
        <div className="flex justify-between items-center mb-8">
          <div></div>
          <button
            onClick={() => setShowAddForm(!showAddForm)}
            className="btn-brutal btn-brutal-primary"
          >
            <span className="flex items-center">
              {showAddForm ? 'Cancel' : 'Add New Bill'}
              <span className="ml-2">{showAddForm ? '✕' : '+'}</span>
            </span>
          </button>
        </div>

        {/* Add Bill Form */}
        {showAddForm && (
          <div className="card-brutal p-8 mb-8 ">
            <h2 className="text-2xl font-bold uppercase mb-6">Add New Bill</h2>
            <form onSubmit={handleSubmit} className="space-y-6">
              <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
                <div>
                  <label className="block text-sm font-bold uppercase mb-2">Bill Name</label>
                  <input
                    type="text"
                    required
                    value={formData.name}
                    onChange={(e) => setFormData({ ...formData, name: e.target.value })}
                    className="input-brutal w-full"
                    placeholder="Electricity Bill"
                  />
                </div>
                <div>
                  <label className="block text-sm font-bold uppercase mb-2">Amount (Rp)</label>
                  <input
                    type="number"
                    required
                    min="0"
                    step="0.01"
                    value={formData.amount === 0 ? '' : formData.amount}
                    onChange={(e) => setFormData({ ...formData, amount: parseFloat(e.target.value) })}
                    className="input-brutal w-full"
                    placeholder="150000"
                  />
                </div>
                <div>
                  <label className="block text-sm font-bold uppercase mb-2">Due Date</label>
                  <input
                    type="date"
                    required
                    value={formData.due_date}
                    onChange={(e) => setFormData({ ...formData, due_date: e.target.value })}
                    className="input-brutal w-full"
                  />
                </div>
                <div>
                  <label className="block text-sm font-bold uppercase mb-2">Remind Before (days)</label>
                  <input
                    type="number"
                    min="0"
                    value={formData.remind_before || 3}
                    onChange={(e) => setFormData({ ...formData, remind_before: parseInt(e.target.value) })}
                    className="input-brutal w-full"
                  />
                </div>
              </div>
              <div>
                <label className="block text-sm font-bold uppercase mb-2">Description (Optional)</label>
                <textarea
                  value={formData.description}
                  onChange={(e) => setFormData({ ...formData, description: e.target.value })}
                  className="input-brutal w-full"
                  rows={3}
                  placeholder="Additional notes..."
                />
              </div>
              <div className="flex justify-end space-x-4">
                <button
                  type="button"
                  onClick={() => setShowAddForm(false)}
                  className="btn-brutal"
                >
                  Cancel
                </button>
                <button
                  type="submit"
                  className="btn-brutal btn-brutal-success"
                >
                  Save Bill
                </button>
              </div>
            </form>
          </div>
        )}

        {/* Bills List */}
        <div className="card-brutal ">
          <div className="p-6">
            <div className="flex justify-between items-center mb-6">
              <h2 className="text-2xl font-bold uppercase">My Bills</h2>
              <span className="badge-brutal badge-brutal-blue">
                {bills?.length || 0} Total
              </span>
            </div>

            {bills && bills.length > 0 ? (
              <div className="space-y-4">
                {bills.map((bill, index) => (
                  <div
                    key={bill.id}
                    className="bg-white border-2 border-black p-6"
                  >
                    <div className="flex items-start justify-between">
                      <div className="flex-1">
                        <div className="flex items-center space-x-3 mb-3">
                          <h3 className="font-bold text-xl">{bill.name}</h3>
                          <span className={`badge-brutal ${getStatusColor(bill.status)} `}>
                            {bill.status}
                          </span>
                        </div>

                        <div className="grid grid-cols-1 md:grid-cols-3 gap-4 mb-3">
                          <div className="bg-gray-100 border-2 border-black p-3">
                            <div className="text-xs font-bold uppercase text-gray-600">Amount</div>
                            <div className="text-lg font-black">Rp {bill.amount.toLocaleString('id-ID')}</div>
                          </div>
                          <div className="bg-gray-100 border-2 border-black p-3">
                            <div className="text-xs font-bold uppercase text-gray-600">Due Date</div>
                            <div className="text-lg font-black">{formatDate(bill.due_date)}</div>
                          </div>
                          <div className="bg-gray-100 border-2 border-black p-3">
                            <div className="text-xs font-bold uppercase text-gray-600">Reminder</div>
                            <div className="text-lg font-black">{bill.remind_before} days before</div>
                          </div>
                        </div>

                        {bill.description && (
                          <div className="bg-yellow-100 border-2 border-black p-3 mt-3">
                            <div className="text-xs font-bold uppercase text-gray-600 mb-1">Notes</div>
                            <div className="font-medium">{bill.description}</div>
                          </div>
                        )}
                      </div>

                      <div className="ml-6 flex flex-col space-y-2">
                        {bill.status !== 'paid' && (
                          <button
                            onClick={() => markAsPaid(bill.id)}
                            className="btn-brutal btn-brutal-success"
                          >
                            Mark Paid
                          </button>
                        )}
                        <button
                          onClick={() => handleDelete(bill.id)}
                          className="btn-brutal btn-brutal-danger"
                        >
                          Delete
                        </button>
                      </div>
                    </div>
                  </div>
                ))}
              </div>
            ) : (
              <div className="text-center py-12">
                <div className="text-6xl mb-4">📭</div>
                <p className="font-bold text-2xl uppercase mb-2">No Bills Found</p>
                <p className="text-gray-600 mb-6">Add your first bill to get started!</p>
                <button
                  onClick={() => setShowAddForm(true)}
                  className="btn-brutal btn-brutal-primary"
                >
                  Add Your First Bill
                </button>
              </div>
            )}
          </div>
        </div>
      </main>
    </div>
  );
};

export default BillsPage;