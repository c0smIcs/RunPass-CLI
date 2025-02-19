package cmd

import (
	"os"

	"github.com/c0smIcs/RanPass_CLI/internal/generator"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// Глобальные переменные для флагов
var (
	passwordLength int    // Длина пароля
	includeLetters bool   // Включать ли буквы 
	includeNumbers bool   // Включать ли цифры
	includeSymbols bool   // Включать ли специальные символы
	outputFile     string // Флаг для сохранения пароля в файл
)

// RootCmd - представляет команду generate
var RootCmd = &cobra.Command{
	Use: "generate",
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
		color.Red("Ошибка: длина пароля должна быть не менее 4 символов")
		return
	}
		
	// проверка, что хотя бы один тип символа выбран
	if !includeLetters && !includeNumbers && !includeSymbols {
		color.Red("Ошибка: Выберите хотя бы один тип симолов (--letters, --numbers, --symbols)")
		return
	}

	// Генерация пароля
	password := generator.GeneratePassword(passwordLength, includeLetters, includeNumbers, includeSymbols)
	color.Green("Сгенерированный пароль: %s\n", password)
			// Сохранение пароля в файл, если указан флаг --output
	if outputFile != "" {
		err := os.WriteFile(outputFile, []byte(password), 0644)
		if err != nil {
			color.Red("Ошибка при сохранении пароля в файл: %v\n", err)
			return
		}
		color.Green("Пароль успешно сохранен в файл: %s\n", outputFile)
	}
}

func init() {
	// Добавляем флаги к команде
	RootCmd.Flags().IntVarP(&passwordLength, "length", "l", 12, "Длина пароля")
	RootCmd.Flags().BoolVarP(&includeLetters, "letters", "a", false, "Включать буквы")
	RootCmd.Flags().BoolVarP(&includeNumbers, "numbers", "n", false, "Включать цифры")
	RootCmd.Flags().BoolVarP(&includeSymbols, "symbols", "s", false, "Включать специальные символы")
	RootCmd.Flags().StringVarP(&outputFile, "output", "o", "", "Сохранить пароль в файл (опционально)")
}
