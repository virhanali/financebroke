# Personal Finance & Bill Reminder App

Phase 1 Implementation:
- ✅ User Authentication (Register/Login)
- ✅ Bill Management (CRUD)
- ✅ Simple Reminder Services (Telegram + Email)
- ✅ Dashboard with Bill List

## Tech Stack
- **Backend**: Go with Gin framework, PostgreSQL, GORM
- **Frontend**: React with TypeScript, Tailwind CSS
- **Authentication**: JWT tokens
- **Notifications**: Telegram Bot API, SMTP Email

## Setup Instructions

### Backend
1. Navigate to backend directory
```bash
cd backend
```

2. Install dependencies
```bash
go mod download
```

3. Setup environment variables
```bash
cp .env.example .env
# Edit .env with your configuration
```

4. Run the server
```bash
go run cmd/server/main.go
```

### Frontend
1. Navigate to frontend directory
```bash
cd frontend
```

2. Install dependencies
```bash
npm install
```

3. Start development server
```bash
npm start
```

## API Endpoints

### Authentication
- POST `/api/v1/register` - User registration
- POST `/api/v1/login` - User login
- GET `/api/v1/profile` - Get user profile (protected)

### Bills
- GET `/api/v1/bills` - Get all bills
- POST `/api/v1/bills` - Create new bill
- GET `/api/v1/bills/:id` - Get specific bill
- PUT `/api/v1/bills/:id` - Update bill
- DELETE `/api/v1/bills/:id` - Delete bill
- GET `/api/v1/bills/upcoming` - Get upcoming bills

### Dashboard
- GET `/api/v1/dashboard` - Get dashboard data

### Notifications
- PUT `/api/v1/notifications/settings` - Update notification settings
- POST `/api/v1/notifications/test-telegram` - Test Telegram notification

## Features

### User Management
- Register new account
- Login with JWT authentication
- Profile management

### Bill Management
- Add, view, update, delete bills
- Mark bills as paid/unpaid/overdue
- Set reminder days before due date

### Dashboard
- View bill statistics
- See upcoming bills
- Quick actions for bill management

### Notifications
- Email reminders
- Telegram bot reminders
- Configurable notification settings

## Next Steps for Phase 2
- Implement finance tracker (income/expense)
- Add data visualization with charts
- Export to CSV functionality
- Admin dashboard
- Enhanced notification scheduler