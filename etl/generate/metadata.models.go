package generate

import (
	"encoding/xml"
)

type xml_quran struct {
	XMLName   xml.Name    `xml:"quran"`
	Type      string      `xml:"type,attr"`
	Version   string      `xml:"version,attr"`
	Copyright string      `xml:"copyright,attr"`
	Suras     xml_Suras   `xml:"suras"`
	Juzs      xml_Juzs    `xml:"juzs"`
	Hizbs     xml_Hizbs   `xml:"hizbs"`
	Manzils   xml_Manzils `xml:"manzils"`
	Rukus     xml_Rukus   `xml:"rukus"`
	Pages     xml_Pages   `xml:"pages"`
	Sajdas    xml_Sajdas  `xml:"sajdas"`
}

type xml_Suras struct {
	XMLName xml.Name   `xml:"suras"`
	Alias   string     `xml:"alias,attr"`
	Sura    []xml_Sura `xml:"sura"`
}

type xml_Sura struct {
	XMLName         xml.Name `xml:"sura"`
	Id              int      `xml:"index,attr"`
	Ayas            int      `xml:"ayas,attr"`
	StartLine       int      `xml:"start,attr"`
	NameArabic      string   `xml:"name,attr"`
	NameEnglish     string   `xml:"ename,attr"`
	RevelationCity  string   `xml:"type,attr"`
	RevelationOrder int      `xml:"order,attr"`
	Rukus           int      `xml:"rukus,attr"`
}

type xml_Juzs struct {
	XMLName xml.Name  `xml:"juzs"`
	Alias   string    `xml:"alias,attr"`
	Juz     []xml_Juz `xml:"juz"`
}

type xml_Juz struct {
	XMLName xml.Name `xml:"juz"`
	Id      int      `xml:"index,attr"`
	Sura    int      `xml:"sura,attr"`
	Aya     int      `xml:"aya,attr"`
}

type xml_Hizbs struct {
	XMLName xml.Name   `xml:"hizbs"`
	Alias   string     `xml:"alias,attr"`
	Hizb    []xml_Hizb `xml:"quarter"`
}

type xml_Hizb struct {
	XMLName xml.Name `xml:"quarter"`
	Id      int      `xml:"index,attr"`
	Sura    int      `xml:"sura,attr"`
	Aya     int      `xml:"aya,attr"`
}

type xml_Manzils struct {
	XMLName xml.Name     `xml:"manzils"`
	Alias   string       `xml:"alias,attr"`
	Manzil  []xml_Manzil `xml:"manzil"`
}

type xml_Manzil struct {
	XMLName xml.Name `xml:"manzil"`
	Id      int      `xml:"index,attr"`
	Sura    int      `xml:"sura,attr"`
	Aya     int      `xml:"aya,attr"`
}

type xml_Rukus struct {
	XMLName xml.Name   `xml:"rukus"`
	Alias   string     `xml:"alias,attr"`
	Ruku    []xml_Ruku `xml:"ruku"`
}

type xml_Ruku struct {
	XMLName xml.Name `xml:"ruku"`
	Id      int      `xml:"index,attr"`
	Sura    int      `xml:"sura,attr"`
	Aya     int      `xml:"aya,attr"`
}

type xml_Pages struct {
	XMLName xml.Name   `xml:"pages"`
	Alias   string     `xml:"alias,attr"`
	Page    []xml_Page `xml:"page"`
}

type xml_Page struct {
	XMLName xml.Name `xml:"page"`
	Id      int      `xml:"index,attr"`
	Sura    int      `xml:"sura,attr"`
	Aya     int      `xml:"aya,attr"`
}

type xml_Sajdas struct {
	XMLName xml.Name    `xml:"sajdas"`
	Alias   string      `xml:"alias,attr"`
	Sajda   []xml_Sajda `xml:"sajda"`
}

type xml_Sajda struct {
	XMLName xml.Name `xml:"sajda"`
	Id      int      `xml:"index,attr"`
	Sura    int      `xml:"sura,attr"`
	Aya     int      `xml:"aya,attr"`
}
