package repository

import (
	"database/sql"
	"financebroke/backend/internal/entity"
	"financebroke/backend/internal/utils"
)

type UserRepository interface {
	Create(user entity.User) (entity.User, error)
	FindByEmail(email string) (entity.User, error)
	FindByID(id uint) (entity.User, error)
	UpdateNotificationSettings(id uint, chatID string, emailNotify, telegramNotify bool) (entity.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user entity.User) (entity.User, error) {
	logger := utils.GetLogger()
	logger.Info("[REPO] Creating user", map[string]interface{}{
		"email": user.Email,
	})

	query := `
		INSERT INTO users (email, password, name)
		VALUES ($1, $2, $3)
		RETURNING id, email, name, telegram_chat_id, email_notify, telegram_notify, created_at, updated_at
	`

	var telegramChatID sql.NullString
	err := r.db.QueryRow(query, user.Email, user.Password, user.Name).Scan(
		&user.ID, &user.Email, &user.Name, &telegramChatID,
		&user.EmailNotify, &user.TelegramNotify, &user.CreatedAt, &user.UpdatedAt,
	)

	if telegramChatID.Valid {
		user.TelegramChatID = telegramChatID.String
	} else {
		user.TelegramChatID = ""
	}

	if err != nil {
		utils.LogError("REPO_USER_CREATE", err)
		return entity.User{}, err
	}

	logger.Info("[REPO] User created successfully", map[string]interface{}{
		"user_id": user.ID,
	})
	return user, nil
}

func (r *userRepository) FindByEmail(email string) (entity.User, error) {
	logger := utils.GetLogger()
	logger.Info("[REPO] Finding user by email", map[string]interface{}{
		"email": email,
	})

	query := `
		SELECT id, email, password, name, telegram_chat_id, email_notify, telegram_notify, created_at, updated_at
		FROM users
		WHERE email = $1
	`

	var user entity.User
	var telegramChatID sql.NullString
	err := r.db.QueryRow(query, email).Scan(
		&user.ID, &user.Email, &user.Password, &user.Name, &telegramChatID,
		&user.EmailNotify, &user.TelegramNotify, &user.CreatedAt, &user.UpdatedAt,
	)

	if telegramChatID.Valid {
		user.TelegramChatID = telegramChatID.String
	} else {
		user.TelegramChatID = ""
	}

	if err != nil {
		if err == sql.ErrNoRows {
			logger.Info("[REPO] User not found", map[string]interface{}{
				"email": email,
			})
		} else {
			utils.LogError("REPO_USER_FIND_EMAIL", err)
		}
		return entity.User{}, err
	}

	logger.Info("[REPO] User found", map[string]interface{}{
		"user_id": user.ID,
	})
	return user, nil
}

func (r *userRepository) FindByID(id uint) (entity.User, error) {
	logger := utils.GetLogger()
	query := `
		SELECT id, email, name, telegram_chat_id, email_notify, telegram_notify, created_at, updated_at
		FROM users
		WHERE id = $1
	`

	var user entity.User
	var telegramChatID sql.NullString
	err := r.db.QueryRow(query, id).Scan(
		&user.ID, &user.Email, &user.Name, &telegramChatID,
		&user.EmailNotify, &user.TelegramNotify, &user.CreatedAt, &user.UpdatedAt,
	)

	if telegramChatID.Valid {
		user.TelegramChatID = telegramChatID.String
	} else {
		user.TelegramChatID = ""
	}

	if err != nil {
		if err == sql.ErrNoRows {
			logger.Info("[REPO] User not found", map[string]interface{}{
				"user_id": id,
			})
		} else {
			utils.LogError("REPO_USER_FIND_ID", err)
		}
		return entity.User{}, err
	}

	logger.Info("[REPO] User found", map[string]interface{}{
		"user_id": user.ID,
	})
	return user, nil
}

func (r *userRepository) UpdateNotificationSettings(id uint, chatID string, emailNotify, telegramNotify bool) (entity.User, error) {
	query := `
		UPDATE users
		SET telegram_chat_id = $1, email_notify = $2, telegram_notify = $3, updated_at = CURRENT_TIMESTAMP
		WHERE id = $4
		RETURNING id, email, name, telegram_chat_id, email_notify, telegram_notify, created_at, updated_at
	`

	var user entity.User
	err := r.db.QueryRow(query, chatID, emailNotify, telegramNotify, id).Scan(
		&user.ID, &user.Email, &user.Name, &user.TelegramChatID,
		&user.EmailNotify, &user.TelegramNotify, &user.CreatedAt, &user.UpdatedAt,
	)

	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}