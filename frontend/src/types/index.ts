export interface User {
  id: number;
  email: string;
  name: string;
  telegram_chat_id: string;
  email_notify: boolean;
  telegram_notify: boolean;
  created_at: string;
  updated_at: string;
}

export interface Bill {
  id: number;
  user_id: number;
  name: string;
  amount: number;
  due_date: string;
  description: string;
  status: 'unpaid' | 'paid' | 'overdue';
  remind_before: number;
  created_at: string;
  updated_at: string;
}

export interface LoginRequest {
  email: string;
  password: string;
}

export interface RegisterRequest {
  name: string;
  email: string;
  password: string;
  confirm_password: string;
}

export interface AuthResponse {
  token: string;
  user: User;
}

export interface BillCreateRequest {
  name: string;
  amount: number;
  due_date: string;
  description?: string;
  remind_before?: number;
}

export interface BillUpdateRequest {
  name?: string;
  amount?: number;
  due_date?: string;
  description?: string;
  status?: 'unpaid' | 'paid' | 'overdue';
  remind_before?: number;
}

export interface DashboardResponse {
  total_bills: number;
  paid_bills: number;
  unpaid_bills: number;
  overdue_bills: number;
  upcoming_bills: Bill[];
  recent_bills: Bill[];
  summary: {
    total_amount: number;
    paid_amount: number;
    unpaid_amount: number;
    overdue_amount: number;
  };
}