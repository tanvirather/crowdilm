package generate

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"os"
)

func MetaData() {
	data, _ := os.ReadFile("tanzil/quran-data.xml")
	var quran xml_quran
	xml.Unmarshal(data, &quran)

	createFile(dataload, "")
	generateSura(&quran)
	generateAya(&quran)
}

func generateSura(quran *xml_quran) {
	var b bytes.Buffer

	b.WriteString(`/************************************************** sura **************************************************/
insert ignore into sura(id, ayas, name_arabic, name_english, revelation_city, revelation_order) values
`)

	for i := 0; i < len(quran.Suras.Sura); i++ {
		var sura = quran.Suras.Sura[i]
		b.WriteString(fmt.Sprintf("(%d, %d, '%s', '%s', '%s', %d),\n",
			sura.Id, sura.Ayas, sura.NameArabic, sura.NameEnglish, sura.RevelationCity, sura.RevelationOrder))
	}
	// add the last quran with a semi-colon instead of comma and new line
	var sura = quran.Suras.Sura[len(quran.Suras.Sura)-1]
	b.WriteString(fmt.Sprintf("(%d, %d, '%s', '%s', '%s', %d);\n",
		sura.Id, sura.Ayas, sura.NameArabic, sura.NameEnglish, sura.RevelationCity, sura.RevelationOrder))

	// save to file
	appendFile(dataload, b.String())
}

func generateAya(quran *xml_quran) {
	var b bytes.Buffer

	// aya_id, surah_id
	matrix := [6236][8]int{}

	b.WriteString(`/************************************************** line **************************************************/
insert ignore into line(id, surah, aya, manzil, juz, hizb, ruku, page) values
`)

	for i := 0; i < len(matrix); i++ {
		matrix[i][0] = i
	}

	// sura
	var start_line = 0
	var ayas = 0
	for suraIndex := 0; suraIndex < len(quran.Suras.Sura); suraIndex++ {
		var sura = quran.Suras.Sura[suraIndex]
		start_line = sura.StartLine
		ayas = sura.Ayas
		var sura_line = 0
		for matrixIndex := start_line; matrixIndex < ayas+start_line; matrixIndex++ {
			matrix[matrixIndex][1] = sura.Id
			sura_line++
			matrix[matrixIndex][2] = sura_line
		}
		start_line += ayas
	}

	// manzil
	start_line = 0
	ayas = 0
	var counter = 0
	for i := 0; i < len(matrix); i++ {
		if counter < len(quran.Manzils.Manzil) &&
			matrix[i][1] == quran.Manzils.Manzil[counter].Sura &&
			matrix[i][2] == quran.Manzils.Manzil[counter].Aya {
			counter++
		}
		matrix[i][3] = counter
	}

	// juz
	start_line = 0
	ayas = 0
	counter = 0
	for i := 0; i < len(matrix); i++ {
		if counter < len(quran.Juzs.Juz) &&
			matrix[i][1] == quran.Juzs.Juz[counter].Sura &&
			matrix[i][2] == quran.Juzs.Juz[counter].Aya {
			counter++
		}
		matrix[i][4] = counter
	}

	// hizb
	start_line = 0
	ayas = 0
	counter = 0
	for i := 0; i < len(matrix); i++ {
		if counter < len(quran.Hizbs.Hizb) &&
			matrix[i][1] == quran.Hizbs.Hizb[counter].Sura &&
			matrix[i][2] == quran.Hizbs.Hizb[counter].Aya {
			counter++
		}
		matrix[i][5] = counter
	}

	// ruku
	start_line = 0
	ayas = 0
	counter = 0
	for i := 0; i < len(matrix); i++ {
		if counter < len(quran.Rukus.Ruku) &&
			matrix[i][1] == quran.Rukus.Ruku[counter].Sura &&
			matrix[i][2] == quran.Rukus.Ruku[counter].Aya {
			counter++
		}
		matrix[i][6] = counter
	}

	// page
	start_line = 0
	ayas = 0
	counter = 0
	for i := 0; i < len(matrix); i++ {
		if counter < len(quran.Pages.Page) &&
			matrix[i][1] == quran.Pages.Page[counter].Sura &&
			matrix[i][2] == quran.Pages.Page[counter].Aya {
			counter++
		}
		matrix[i][7] = counter
	}
	for i := 0; i < len(matrix)-1; i++ {
		var item = matrix[i]
		b.WriteString(fmt.Sprintf("(%d, %d, %d, %d, %d, %d, %d, %d),\n", item[0]+1, item[1], item[2], item[3], item[4], item[5], item[6], item[7]))
	}
	var item = matrix[len(matrix)-1]
	b.WriteString(fmt.Sprintf("(%d, %d, %d, %d, %d, %d, %d, %d);\n", item[0]+1, item[1], item[2], item[3], item[4], item[5], item[6], item[7]))

	// save to file
	appendFile(dataload, b.String())
}
