package generate

import (
	"bufio"
	"bytes"
	"fmt"
	"html"
	"os"
)

func QuranLine() {
	var b bytes.Buffer

	for i := 0; i < len(quranLists); i++ {
		var item = quranLists[i]
		b.WriteString("insert ignore into quran_line(quran_id, line_id, text) values\n")

		file, _ := os.Open("tanzil/" + item.id + ".txt")
		defer file.Close()
		scanner := bufio.NewScanner(file)
		var counter = 0
		for scanner.Scan() {
			counter++
			line := scanner.Text()
			// add ',' for each line. for the last line add ';'
			if counter < 6236 {
				b.WriteString(fmt.Sprintf("('%s', %d, '%s'),\n", item.id, counter, html.EscapeString(line)))
			} else {
				b.WriteString(fmt.Sprintf("('%s', %d, '%s');\n", item.id, counter, html.EscapeString(line)))
				break
			}
		}
		// write to file
		createFile(output_path+item.id+".sql", b.String())
		b.Truncate(0)

		// create sql run shell script
		appendFile(build_database, "docker exec --interactive $DOCKER_CONTAINER mysql --host $HOST --database $DATABASE --user $USER --password=$PASSWORD --default-character-set=utf8 < "+item.id+".sql\n")
	}
}
