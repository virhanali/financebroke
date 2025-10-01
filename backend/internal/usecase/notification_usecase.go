package usecase

import (
	"financebroke/backend/internal/dto"
	"financebroke/backend/internal/entity"
	"financebroke/backend/internal/repository"
	"financebroke/backend/internal/services"
)

type NotificationUsecase interface {
	UpdateSettings(userID uint, req *dto.NotificationSettingsRequest) (*entity.User, error)
	TestTelegram(userID uint, req *dto.TestTelegramRequest) error
	SendBillReminder(bill entity.Bill, user entity.User) error
}

type notificationUsecase struct {
	userRepo       repository.UserRepository
	telegramSvc    *services.TelegramService
	emailSvc       *services.EmailService
}

func NewNotificationUsecase(
	userRepo repository.UserRepository,
	telegramSvc *services.TelegramService,
	emailSvc *services.EmailService,
) NotificationUsecase {
	return &notificationUsecase{
		userRepo:    userRepo,
		telegramSvc: telegramSvc,
		emailSvc:    emailSvc,
	}
}

func (u *notificationUsecase) UpdateSettings(userID uint, req *dto.NotificationSettingsRequest) (*entity.User, error) {
	chatID := ""
	if req.TelegramChatID != "" {
		chatID = req.TelegramChatID
	}

	user, err := u.userRepo.UpdateNotificationSettings(userID, chatID, req.EmailNotify, req.TelegramNotify)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *notificationUsecase) TestTelegram(userID uint, req *dto.TestTelegramRequest) error {
	user, err := u.userRepo.FindByID(userID)
	if err != nil {
		return err
	}

	if user.TelegramChatID == "" {
		return err
	}

	return u.telegramSvc.SendTestMessage(user.TelegramChatID, req.Message)
}

func (u *notificationUsecase) SendBillReminder(bill entity.Bill, user entity.User) error {
	if user.EmailNotify {
		err := u.emailSvc.SendReminder(&bill, &user)
		if err != nil {
			return err
		}
	}

	if user.TelegramNotify && user.TelegramChatID != "" {
		err := u.telegramSvc.SendReminder(&bill, &user)
		if err != nil {
			return err
		}
	}

	return nil
}