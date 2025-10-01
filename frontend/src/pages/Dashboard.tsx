import React, { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import { useAuth } from '../contexts/AuthContext';
import { Bill, DashboardResponse } from '../types';
import { dashboardApi, billApi } from '../services/api';
import Logo from '../components/Logo';

const Dashboard: React.FC = () => {
  const { user, logout } = useAuth();
  const [dashboard, setDashboard] = useState<DashboardResponse | null>(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    fetchDashboard();
  }, []);

  const fetchDashboard = async () => {
    try {
      const data = await dashboardApi.getDashboard();
      setDashboard(data);
    } catch (err) {
      console.error('Failed to fetch dashboard:', err);
    } finally {
      setLoading(false);
    }
  };

  const markAsPaid = async (billId: number) => {
    try {
      await billApi.updateBill(billId, { status: 'paid' });
      fetchDashboard();
    } catch (err) {
      console.error('Failed to update bill:', err);
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

  if (loading || !dashboard) {
    return (
      <div className="min-h-screen bg-gradient-to-br from-purple-100 via-pink-100 to-yellow-100 flex items-center justify-center">
        <div className="text-center">
          <div className="card-brutal p-8 ">
            <div className="text-2xl font-bold uppercase">Loading...</div>
          </div>
        </div>
      </div>
    );
  }

  return (
    <div className="min-h-screen bg-gradient-to-br from-purple-100 via-pink-100 to-yellow-100">
      {/* Navigation */}
      <nav className="nav-brutal sticky top-0 z-50">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="flex justify-between h-16 items-center">
            <div className="flex items-center space-x-4">
              <Logo size="medium" />
              <span className="badge-brutal badge-brutal-blue">Dashboard</span>
            </div>
            <div className="flex items-center space-x-4">
              <span className="font-bold uppercase text-sm">Welcome, {user?.name}</span>
              <Link
                to="/bills"
                className="btn-brutal"
              >
                My Bills
              </Link>
              <button
                onClick={logout}
                className="btn-brutal btn-brutal-danger"
              >
                Logout
              </button>
            </div>
          </div>
        </div>
      </nav>

      {/* Main content */}
      <main className="max-w-7xl mx-auto py-8 px-4 sm:px-6 lg:px-8">
        {/* Decorative elements */}
        <div className="absolute top-32 right-10 w-16 h-16 bg-blue-400 border-2 border-black -z-10"></div>
        <div className="absolute bottom-20 left-10 w-12 h-12 bg-pink-400 border-2 border-black -z-10"></div>

        {/* Stats Grid */}
        {dashboard && (
          <div className="grid grid-cols-1 gap-6 sm:grid-cols-2 lg:grid-cols-4 mb-12">
            <div className="stat-card-brutal ">
              <div className="text-center">
                <div className="text-4xl font-black mb-2">{dashboard.total_bills}</div>
                <div className="text-sm font-bold uppercase">Total Bills</div>
              </div>
            </div>
            <div className="stat-card-brutal ">
              <div className="text-center">
                <div className="text-4xl font-black mb-2 text-green-600">{dashboard.paid_bills}</div>
                <div className="text-sm font-bold uppercase">Paid Bills</div>
              </div>
            </div>
            <div className="stat-card-brutal ">
              <div className="text-center">
                <div className="text-4xl font-black mb-2 text-yellow-600">{dashboard.unpaid_bills}</div>
                <div className="text-sm font-bold uppercase">Unpaid Bills</div>
              </div>
            </div>
            <div className="stat-card-brutal ">
              <div className="text-center">
                <div className="text-4xl font-black mb-2 text-red-600">{dashboard.overdue_bills}</div>
                <div className="text-sm font-bold uppercase">Overdue Bills</div>
              </div>
            </div>
          </div>
        )}

        {/* Summary Section */}
        {dashboard && (
          <div className="card-brutal p-6 mb-8 ">
            <h2 className="text-xl font-bold uppercase mb-4">Financial Summary</h2>
            <div className="grid grid-cols-1 md:grid-cols-4 gap-4">
              <div className="text-center p-4 bg-gray-100 border-2 border-black">
                <div className="text-lg font-bold">Total Amount</div>
                <div className="text-2xl font-black">
                  Rp {dashboard.summary.total_amount.toLocaleString('id-ID')}
                </div>
              </div>
              <div className="text-center p-4 bg-green-100 border-2 border-black">
                <div className="text-lg font-bold">Paid Amount</div>
                <div className="text-2xl font-black text-green-600">
                  Rp {dashboard.summary.paid_amount.toLocaleString('id-ID')}
                </div>
              </div>
              <div className="text-center p-4 bg-yellow-100 border-2 border-black">
                <div className="text-lg font-bold">Unpaid Amount</div>
                <div className="text-2xl font-black text-yellow-600">
                  Rp {dashboard.summary.unpaid_amount.toLocaleString('id-ID')}
                </div>
              </div>
              <div className="text-center p-4 bg-red-100 border-2 border-black">
                <div className="text-lg font-bold">Overdue Amount</div>
                <div className="text-2xl font-black text-red-600">
                  Rp {dashboard.summary.overdue_amount.toLocaleString('id-ID')}
                </div>
              </div>
            </div>
          </div>
        )}

        {/* Upcoming Bills */}
        <div className="card-brutal mb-8 ">
          <div className="p-6">
            <div className="flex justify-between items-center mb-6">
              <h2 className="text-2xl font-bold uppercase">Upcoming Bills</h2>
              <span className="badge-brutal badge-brutal-yellow ">Attention</span>
            </div>

            {dashboard?.upcoming_bills && dashboard.upcoming_bills.length > 0 ? (
              <div className="space-y-4">
                {dashboard.upcoming_bills.map((bill: Bill) => (
                  <div key={bill.id} className="bg-white border-2 border-black p-4 ">
                    <div className="flex items-center justify-between">
                      <div className="flex-1">
                        <div className="flex items-center space-x-3 mb-2">
                          <h3 className="font-bold text-lg">{bill.name}</h3>
                          <span className={`badge-brutal ${getStatusColor(bill.status)}`}>
                            {bill.status}
                          </span>
                        </div>
                        <div className="space-y-1 text-sm">
                          <p className="font-bold">Rp {bill.amount.toLocaleString('id-ID')}</p>
                          <p>Due: {formatDate(bill.due_date)}</p>
                          {bill.description && (
                            <p className="text-gray-600">{bill.description}</p>
                          )}
                        </div>
                      </div>
                      <div className="ml-4">
                        {bill.status !== 'paid' && (
                          <button
                            onClick={() => markAsPaid(bill.id)}
                            className="btn-brutal btn-brutal-success"
                          >
                            Mark as Paid
                          </button>
                        )}
                      </div>
                    </div>
                  </div>
                ))}
              </div>
            ) : (
              <div className="text-center py-12">
                <p className="font-bold text-lg">No upcoming bills!</p>
                <p className="text-gray-600 mt-2">You're all caught up 🎉</p>
              </div>
            )}
          </div>
        </div>

        {/* Quick Actions */}
        <div className="flex flex-wrap gap-4">
          <Link
            to="/bills"
            state={{ showAddForm: true }}
            className="btn-brutal btn-brutal-primary"
          >
            <span className="flex items-center">
              Add New Bill
              <span className="ml-2">+</span>
            </span>
          </Link>
          <button className="btn-brutal" onClick={fetchDashboard}>
            <span className="flex items-center">
              Refresh
              <span className="ml-2">🔄</span>
            </span>
          </button>
        </div>
      </main>
    </div>
  );
};

export default Dashboard;