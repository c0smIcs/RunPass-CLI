package generator

import (
	"log/slog"
	"math/rand"
	"regexp"
	"time"
)

// Константы с regex-шаблонами для каждого набора символов
const (
	lettersRegex = `[a-zA-Z]` // все английские символы
	numbersRegex = `\d` // все цифры
	symbolsRegex = `[!@#$%^&*()\-_=+[\]{}|;:,.<>/?` + "`~]" // Экранированные спецсимволы
)

// Готовые наборы символов (генерируются автоматически)
var (
	letters = generateChars('a', 'z') + generateChars('A', 'Z')
	numbers = generateChars('0', '9')
	symbols = "!@#$%^&*()-_=+[]{}|;:,.<>/?`~"
)

// Проверочны регулярные выражения
// var (
// 	letterRegex = regexp.MustCompile(LettersRegex)
// 	numberRegex = regexp.MustCompile(NumbersRegex)
// 	symbolRegex = regexp.MustCompile(SymbolsRegex)
// )

// generateChars - создает последовательность символов от start до end
func generateChars(start, end rune) string {
	var chars string
	for c := start; c <= end; c++ {
		chars += string(c)
	}

	return chars
}

// ValidateChar проверяет, содержит ли строка хотя бы один символ из набора
func ValidateChar(s string, pattern string) bool {
	re := regexp.MustCompile(pattern)
	return re.MatchString(s)
}

// GeneratePassword - генерирует случайный пароль на основе заданных параметров
func GeneratePassword(length int, includeLetters, includeNumbers, includeSymbols bool) string {
	randomGenerator := rand.New(rand.NewSource(time.Now().UnixNano()))

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

	// Проверяем, что выбранные наборы содержат нужные символы
	if includeLetters && !ValidateChar(availableChars, lettersRegex) {
		slog.Error("Ошибка: набор букв не содержит буквенных символов")
	}

	if includeNumbers && !ValidateChar(availableChars, numbersRegex) {
		slog.Error("Ошибка: набор цифр не содержит спецсимволов")
	}

	if includeSymbols && !ValidateChar(availableChars, symbolsRegex) {
		slog.Error("Ошибка: набор символов не содержит спецсимволы")
	}

	if len(availableChars) == 0 {
		slog.Warn("Нет доступных символов", "letters", includeLetters, "numbers", includeNumbers, "symbols", includeSymbols)
	}

	password := make([]byte, length)
	for i := 0; i < length; i++ {
		randomIndex := randomGenerator.Intn(len(availableChars))
		password[i] = availableChars[randomIndex]
	}

	slog.Debug("Пароль успешно сгенерирован внутри функции", "password", string(password))
	return string(password)
}
