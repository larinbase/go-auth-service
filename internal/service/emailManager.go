package service

import (
	"auth-service/internal/utils"
	"fmt"
	"log"
	"sync"
	"time"
)

type EmailCode struct {
	Code int
	Date time.Time
}

type EmailManager struct {
	emailService   *EmailService
	restoreCodes   sync.Map
	restoreCodeTtl time.Duration // minutes
}

func NewEmailManager(codeTtlMinutes int, emailService *EmailService) *EmailManager {
	return &EmailManager{
		restoreCodeTtl: time.Duration(codeTtlMinutes) * time.Minute,
		emailService:   emailService,
	}
}

func (rm *EmailManager) SendEmailCode(email string) error {
	code := utils.GenerateEmailCode()
	err := rm.emailService.sendMesage(fmt.Sprintf("Your restore code: %d", code), email)
	if err != nil {
		return err
	}
	rm.restoreCodes.Store(email, EmailCode{
		Code: code,
		Date: time.Now(),
	})
	return nil
}

func (rm *EmailManager) ValidateCode(email string, code int) bool {
	val, ok := rm.restoreCodes.Load(email)
	if !ok {
		log.Println("Wrong restore code: %s", email)
		return false
	}

	storedCode := val.(EmailCode)
	if storedCode.Code != code {
		log.Println("Wrong restore code: %s", email)
		return false
	}

	if storedCode.Date.Before(time.Now().Add(-rm.restoreCodeTtl)) {
		log.Println("Restore code is outdated: %s", email)
		return false
	}

	return true
}
