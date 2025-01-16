package utils

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/deigo96/e-wallet.git/config"
)

func GenerateOTP() string {
	return fmt.Sprint(time.Now().Nanosecond())[:6]
}

func IsValidOTP(otp, userOTP string) bool {
	return otp == userOTP
}

func OTPExpired() time.Time {
	return time.Now().Add(time.Minute * 5)
}

func IsExpiredOTP(expiredAt time.Time) bool {
	return time.Now().After(expiredAt)
}

func RefactorPhoneNumber(phoneNumber string) string {
	code := "62"
	if phoneNumber[0] == '0' {
		return code + phoneNumber[1:]
	}

	return phoneNumber
}

func ValidPhone(phone string) bool {
	return string(phone[0]) == "0" || phone[0:2] == "62"
}

func GenerateEmailVerification(email string) string {
	hasher := sha1.New()
	hasher.Write([]byte(email))
	hash := hex.EncodeToString(hasher.Sum(nil))
	return hash
}

func GenerateLinkEmailVerification(config *config.Configuration, email string) string {
	baseURL := config.ServiceHost + ":" + config.ServicePort + config.APIVersion + "/"
	return baseURL + "users/verify-email/" + email + "/" + GenerateEmailVerification(email)
}
