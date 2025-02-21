package cmd

import (
	"log/slog"
	"os"

	"github.com/c0smIcs/RanPass_CLI/internal/generator"
	"github.com/spf13/cobra"
)

var (
	passwordLength int    // Длина пароля
	includeLetters bool   // Включать ли буквы
	includeNumbers bool   // Включать ли цифры
	includeSymbols bool   // Включать ли специальные символы
	outputFile     string // Флаг для сохранения пароля в файл
)

var GenerateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Генерация случайного пароля",
	Long: `
Генерация случайного пароля с настраиваемыми параметрами.

Примеры использования:
	# Генерация пароля длинной 16 символов с буквами и цифрами:
	./RunPass generate --length 16 --letters --numbers

	# Генерация пароля длинной 20 символов с буквами, цифрами и специальными символами:
	./RunPass generate --length 20 --letters --numbers --symbols

	# Сохранение пароля в файл:
	./RunPass generate --length 12 --letters --numbers --output password.txt
`,
	Run: processPasswordCreation,
}

func processPasswordCreation(cmd *cobra.Command, args []string) {
	// Проверка минимальной длины пароля
	if passwordLength < 4 {
		slog.Error("Ошибка: длина пароля должна быть не менее 4 символов")
		return
	}
	
	// проверка, что хотя бы один тип символа выбран
	if !includeLetters && !includeNumbers && !includeSymbols {
		slog.Error("Ошибка: Выберите хотя бы один тип симолов (--letters, --numbers, --symbols)")
		return
	}

	// Генерация пароля
	password := generator.GeneratePassword(passwordLength, includeLetters, includeNumbers, includeSymbols)
	slog.Info("Сгенерированный пароль", "password", password)
	if outputFile != "" {
		err := os.WriteFile(outputFile, []byte(password), 0644)
		if err != nil {
			slog.Error("Ошибка при сохранении пароля в файл", "file", outputFile, "error", err)
			return
		}
		slog.Info("Пароль успешно сохранен в файл", "file", outputFile)
	}
}

func init() {
	GenerateCmd.Flags().IntVarP(&passwordLength, "length", "l", 12, "Длина пароля")
	GenerateCmd.Flags().BoolVarP(&includeLetters, "letters", "a", false, "Включать буквы")
	GenerateCmd.Flags().BoolVarP(&includeNumbers, "numbers", "n", false, "Включать цифры")
	GenerateCmd.Flags().BoolVarP(&includeSymbols, "symbols", "s", false, "Включать специальные символы")
	GenerateCmd.Flags().StringVarP(&outputFile, "output", "o", "", "Сохранить пароль в файл (опционально)")
}
