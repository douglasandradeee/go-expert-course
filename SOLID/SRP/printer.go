package singleResponsibilityPrinciple

import (
	"io/fs"
	"os"
)

func PrintNews(filename string, journal Journal) {
	_ = os.WriteFile(filename, []byte(journal.String()), fs.FileMode(0644))
}
