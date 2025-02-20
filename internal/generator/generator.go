package generator

import (
	"log/slog"
	"math/rand"
	"time"
)

// GeneratePassword - генерирует случайный пароль на основе заданных параметров
func GeneratePassword(length int, includeLetters, includeNumbers, includeSymbols bool) string {
	randomGenerator := rand.New(rand.NewSource(time.Now().UnixNano()))

	var letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var numbers = "0123456789"
	var symbols = "!@#$%^&*()-_=+[]{}|;:,.<>/?`~"

	var availableChars string
	if includeLetters {
		availableChars += letters
	}
	if includeNumbers {
		availableChars += numbers
	}
	if includeSymbols {
		availableChars += symbols
	}

	if len(availableChars) == 0 {
		slog.Warn("доступные символы не определены", "letters", includeLetters, "numbers", includeNumbers, "symbols", includeSymbols)
	}

	password := make([]byte, length)
	for i := 0; i < length; i++ {
		randomIndex := randomGenerator.Intn(len(availableChars))
		password[i] = availableChars[randomIndex]
	}

	slog.Debug("Пароль успешно сгенерирован внутри функции", "password", string(password))
	return string(password)
}
