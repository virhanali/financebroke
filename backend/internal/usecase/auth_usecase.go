package usecase

import (
	"errors"
	"finance-app/backend/internal/dto"
	"finance-app/backend/internal/entity"
	"finance-app/backend/internal/repository"
	"finance-app/backend/internal/utils"
)

type AuthUsecase interface {
	Register(req *dto.RegisterRequest) (*dto.AuthResponse, error)
	Login(req *dto.LoginRequest) (*dto.AuthResponse, error)
	GetProfile(userID uint) (*entity.User, error)
}

type authUsecase struct {
	userRepo repository.UserRepository
}

func NewAuthUsecase(userRepo repository.UserRepository) AuthUsecase {
	return &authUsecase{userRepo: userRepo}
}

func (u *authUsecase) Register(req *dto.RegisterRequest) (*dto.AuthResponse, error) {
	logger := utils.GetLogger()
	logger.Info("[USECASE] Starting registration", map[string]interface{}{
		"email": req.Email,
	})

	if req.Password != req.ConfirmPassword {
		err := errors.New("passwords do not match")
		utils.LogError("AUTH_USECASE_REGISTER", err)
		return nil, err
	}

	logger.Info("[USECASE] Checking if email exists", map[string]interface{}{
		"email": req.Email,
	})
	existingUser, err := u.userRepo.FindByEmail(req.Email)
	if err == nil && existingUser.ID != 0 {
		err := errors.New("email already registered")
		utils.LogError("AUTH_USECASE_REGISTER", err)
		return nil, err
	}

	logger.Info("[USECASE] Hashing password", map[string]interface{}{
		"email": req.Email,
	})
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		utils.LogError("AUTH_USECASE_HASH", err)
		return nil, errors.New("failed to hash password")
	}

	user := entity.User{
		Email:       req.Email,
		Password:    hashedPassword,
		Name:        req.Name,
		EmailNotify: true,
	}

	logger.Info("[USECASE] Creating user in database", map[string]interface{}{
		"email": req.Email,
	})
	createdUser, err := u.userRepo.Create(user)
	if err != nil {
		utils.LogError("AUTH_USECASE_CREATE", err)
		return nil, errors.New("failed to create user")
	}

	logger.Info("[USECASE] User created successfully", map[string]interface{}{
		"user_id": createdUser.ID,
	})
	token, err := utils.GenerateToken(createdUser.ID)
	if err != nil {
		utils.LogError("AUTH_USECASE_TOKEN", err)
		return nil, errors.New("failed to generate token")
	}

	logger.Info("[USECASE] Registration completed successfully", map[string]interface{}{
		"email": req.Email,
	})
	return &dto.AuthResponse{
		Token: token,
		User:  createdUser,
	}, nil
}

func (u *authUsecase) Login(req *dto.LoginRequest) (*dto.AuthResponse, error) {
	user, err := u.userRepo.FindByEmail(req.Email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	if !utils.CheckPassword(req.Password, user.Password) {
		return nil, errors.New("invalid credentials")
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return nil, errors.New("failed to generate token")
	}

	return &dto.AuthResponse{
		Token: token,
		User:  user,
	}, nil
}

func (u *authUsecase) GetProfile(userID uint) (*entity.User, error) {
	user, err := u.userRepo.FindByID(userID)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return &user, nil
}
