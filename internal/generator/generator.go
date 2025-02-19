package generator

import (
	"math/rand"
	"time"
)

// GeneratePassword - генерирует случайный пароль на основе заданных параметров
func GeneratePassword(length int, includeLetters, includeNumbers, includeSymbols bool) string {
	// Создаем локальный генератор соучайных чисел
	randomGenerator := rand.New(rand.NewSource(time.Now().UnixNano()))

	var letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var numbers = "0123456789"
	var symbols = "!@#$%^&*()-_=+[]{}|;:,.<>/?`~"

	// собираем строку доступных символов
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

	// Если доступные символы не определены, возвращаем пустую строку
	if len(availableChars) == 0 {
		return "доступные символы не определены"
	}

	password := make([]byte, length)
	for i := 0; i < length; i++ {
		randomIndex := randomGenerator.Intn(len(availableChars))
		password[i] = availableChars[randomIndex]
	}

	return string(password)
}
