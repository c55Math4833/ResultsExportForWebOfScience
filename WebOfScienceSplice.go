package main

import (
	"bufio"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// processFile 根據指定條件處理每個文件。
func processFile(filePath string, isFirst bool, isLast bool) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

    // 刪除除最後一個文件外之 "EF" 行。
	if !isLast && lines[len(lines)-1] == "EF" {
		lines = lines[:len(lines)-1]
	}

    // 刪除除第一個文件外之 "FN"、"VR" 行。
	if !isFirst && strings.Join(lines[:2], "\n") == "\ufeffFN Clarivate Analytics Web of Science\nVR 1.0" {
		lines = lines[2:]
	}

	return lines, scanner.Err()
}

// mergeFiles 將文件合併為單個文件。
func mergeFiles(fileList []string, outputFile string) error {
	outFile, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer outFile.Close()

	writer := bufio.NewWriter(outFile)
	for i, filePath := range fileList {
		lines, err := processFile(filePath, i == 0, i == len(fileList)-1)
		if err != nil {
			return err
		}

		for _, line := range lines {
			_, err := writer.WriteString(line + "\n")
			if err != nil {
				return err
			}
		}
	}
	return writer.Flush()
}

func main() {
	// 尋找文件。
	pattern := "[0-9][0-9][0-9][0-9][0-9][0-9]_[0-9][0-9][0-9][0-9][0-9][0-9].txt"
	files, err := filepath.Glob(pattern)
	if err != nil {
		panic(err)
	}

    if len(files) == 0 {
		panic(err) // 無法匹配正確文件。
	}

	// 排序文件。
	sort.Strings(files)

	// 定義輸出文件名稱。
	outputFile := "T_" + files[0][:7] + files[len(files)-1][7:17]

	err = mergeFiles(files, outputFile)
	if err != nil {
		panic(err) 
	}
}
