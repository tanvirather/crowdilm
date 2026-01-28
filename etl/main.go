package main

import (
	"crowdilm.com/generate"
)

func main() {
	generate.Database()
	generate.MetaData()
	generate.Quran()
	generate.QuranLine()
}
