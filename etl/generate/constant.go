package generate

import (
	"os"
)

const output_path = "output/"
const dataload = output_path + "_dataload.sql"
const build_database = output_path + "_build_database.sh"

func createFile(filePath string, content string) {
	f, _ := os.Create(filePath)
	f.WriteString(content)
	f.Close()
}

func appendFile(filePath string, content string) {
	f, _ := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0644)
	f.WriteString(content)
	f.Close()
}
