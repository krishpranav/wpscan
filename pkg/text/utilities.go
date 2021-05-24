package text

import (
	"bufio"
	"encoding/csv"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/krishpranav/wpscan/pkg/printer"
	"golang.org/x/net/html/charset"
)

// Decode :: This function will convert any string to UTF-8.
func Decode(text string) string {
	r, err := charset.NewReader(strings.NewReader(text), "latin1")

	if err != nil {
		log.Fatal(err)
	}

	result, err := ioutil.ReadAll(r)

	if err != nil {
		log.Fatal(err)
	}

	return string(result)
}

// GetOneImportantFile ::
func GetOneImportantFile(raw string) string {
	rex := regexp.MustCompile("<a href=\"([[[R|r]eadme|EADME]|[C|c]hangelog|HANGELOG]|[R|r]elease_log].*?)\">.*?</a>")

	submatchall := rex.FindStringSubmatch(raw)

	if len(submatchall) > 0 {
		return submatchall[1]
	}

	return ""
}

// GetFileExtensions :: This function searches for files by their extension, within an index of.
func GetFileExtensions(url string, raw string) [][][]byte {
	rex := regexp.MustCompile("<a href=\"(.*?.[sql|db|zip|tar|tar.gz])\">.*?</a>")
	submatchall := rex.FindAllSubmatch([]byte(raw), -1)

	return submatchall
}

// GetVersionStableTag :: This function searches for the version of the plugin or theme.
func GetVersionStableTag(raw string) []string {
	rex := regexp.MustCompile("[S|s]table [T|t]ag.*?([0-9.-]+)")

	return rex.FindStringSubmatch(raw)
}

// GetVersionChangelog :: This function searches for the version of the plugin or theme.
func GetVersionChangelog(raw string) []string {
	rex := regexp.MustCompile("=+\\s+(?:v(?:ersion)?\\s*)?([0-9.-]+)[ \ta-z0-9().\\-/]*=+")

	return rex.FindStringSubmatch(raw)
}

// GetVersionReleaseLog :: This function searches for the version of the plugin or theme.
func GetVersionReleaseLog(raw string) []string {
	rex := regexp.MustCompile("[v|V]ersion.*?([0-9.-]+)")

	return rex.FindStringSubmatch(raw)
}

// ReadAllFile :: This function will be responsible for reading the files.
func ReadAllFile(filename string) (chars []string, count int) {
	file, err := os.Open(filename)

	if err != nil {
		printer.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		chars = append(chars, scanner.Text())
	}

	return chars, len(chars)
}

func ReadCSVFile(filename string) [][]string {
	file, err := os.Open(filename)

	if err != nil {
		printer.Fatal(err)
	}

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()

	if err != nil {
		printer.Fatal(err)
	}

	return records
}
