package file_analyzer

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"read_files/structs"
	"read_files/util"
	"read_files/util/constants"
	"strings"

	"github.com/ledongthuc/pdf"
)

func containsAllKeywords(text string, keywords []string) bool {
	textUpper := strings.ToUpper(text)

	for _, keyword := range keywords {
		upperKeyword := strings.ToUpper(keyword)
		if !strings.Contains(textUpper, upperKeyword) {
			return false
		}
	}

	return true
}

func SearchKeywordsInPdfFiles(file multipart.File, filename string, keywords []string, results chan<- structs.FileReader) error {
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		util.CustomLogger(constants.Error, fmt.Sprintf("io.ReadAll: %v", err))
		return fmt.Errorf("io.ReadAll: %v", err)
	}

	readerAt := bytes.NewReader(data)
	pdfReader, err := pdf.NewReader(readerAt, int64(len(data)))
	if err != nil {
		util.CustomLogger(constants.Error, fmt.Sprintf("NewReader: %v", err))
		return fmt.Errorf("NewReader: %v", err)
	}

	numPages := pdfReader.NumPage()

	for pageIndex := 1; pageIndex <= numPages; pageIndex++ {
		page := pdfReader.Page(pageIndex)
		if page.V.IsNull() {
			continue
		}

		text, err := page.GetPlainText(nil)
		if err != nil {
			util.CustomLogger(constants.Error, fmt.Sprintf("GetPlainText page %d: %v", pageIndex, err))
			continue
		}

		if containsAllKeywords(text, keywords) {
			results <- structs.FileReader{Filename: filename, Reader: bytes.NewBuffer(data)}
			break
		}
	}

	return nil
}
