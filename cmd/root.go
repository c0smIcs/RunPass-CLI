/* 
	Данный файл определяет команду для Cobra "ranpass".
	В init() добавляет GenerateCmd как подкоманду.
*/

package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "ranpass",
	Short: "CLI-приложение для генерации случайных паролей",
	Long:  `Это простое CLI-приложение, которое генерирует случайные пароли с настраиваемыми параметрами`,
}

func init() {
	RootCmd.AddCommand(GenerateCmd)
}
