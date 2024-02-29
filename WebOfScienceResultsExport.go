package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"sync"
	"bufio"
	"runtime"
	"os/exec"
)

func SendRequestWOS(parentQid string, start int, end int, sid string) string {
	url := "https://www.webofscience.com/api/wosnx/indic/export/saveToFile"
	payload := map[string]interface{}{
		"parentQid":            parentQid,
		"sortBy":               "relevance",
		"displayTimesCited":    "true",
		"displayCitedRefs":     "true",
		"product":              "UA",
		"colName":              "WOS",
		"displayUsageInfo":     "true",
		"fileOpt":              "othersoftware",
		"action":               "saveToFieldTagged",
		"markFrom":             strconv.Itoa(start),
		"markTo":               strconv.Itoa(end),
		"view":                 "summary",
		"isRefQuery":           "false",
		"locale":               "en_US",
		"fieldList":            []string{"AUTHORS", "TITLE", "SOURCE", "CONFERENCE_INFO", "CITTIMES", "ACCESSION_NUM", "AUTHORSIDENTIFIERS", "ISSN", "PMID", "ABSTRACT", "ADDRS", "AFFILIATIONS", "DOCTYPE", "KEYWORDS", "JCR_CATEGORY", "SUBJECT_CATEGORY", "WOS_EDITIONS", "CITREF", "CITREFC", "USAGEIND", "HOT_PAPER", "HIGHLY_CITED", "FUNDING", "PUBINFO", "OPEN_ACCESS", "PAGEC", "SABBR", "IDS", "LANG"},
	}
	jsonPayload, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; rv:109.0) Gecko/20100101 Firefox/119.0")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Language", "zh-TW,en-US;q=0.7,en;q=0.3")
	req.Header.Set("Accept-Encoding", "")
	req.Header.Set("X-1P-WOS-SID", sid)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Origin", "https://www.webofscience.com")
	req.Header.Set("DNT", "1")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Referer", "https://www.webofscience.com/wos/woscc/summary/" + parentQid + "/relevance/1(overlay:export/exp)")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("TE", "trailers")

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	return string(body)
}

func SaveDataAsFile(data string, start int, end int, parentQid string) {
	if data != "" {
		dirPath := "./" + parentQid
		if _, err := os.Stat(dirPath); os.IsNotExist(err) {
            err := os.MkdirAll(dirPath, os.ModePerm)
            if err != nil {
                fmt.Println("Error creating directory: ", err)
                return
            }
        }
		filePath := fmt.Sprintf("%s/%06d_%06d.txt", dirPath, start, end)
		err := os.WriteFile(filePath, []byte(data), 0644)
		if err != nil {
			fmt.Println("Error save file: ", err)
		} else {
			fmt.Println("Complete save file: ", filePath)
		}
	} else {
		fmt.Println("Can't find data.")
	}
}

func clearScreen() {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "linux", "darwin": // Linux 和 macOS
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func GetSID() string {
	resp, err := http.Get("http://www.webofknowledge.com/")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	redirectedURL := resp.Request.URL.String()

	parsedURL, err := url.Parse(redirectedURL)
	if err != nil {
		panic(err)
	}
	queryParams := parsedURL.Query()

	referrer := queryParams.Get("referrer")
	re := regexp.MustCompile("SID=([^&]+)")
	match := re.FindStringSubmatch(referrer)

	if len(match) > 1 {
		sid := match[1]
		fmt.Println("SID from referrer:", sid)
		return sid
	} else {
		sid, exists := queryParams["SID"]
		if exists {
			fmt.Println("SID from query params:", sid[0])
			return sid[0]
		} else {
			fmt.Println("SID not found")
			return ""
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		clearScreen()
		var parentQid string
		fmt.Print("Pleace enter QID：\n")
		fmt.Print("Example: If URL = 'www.webofscience.com/wos/woscc/summary/104f891b-a88b-4f4a-bb8d-da4112e85882-c5b07797/relevance/1'\n")
		fmt.Print("Then, QID = '104f891b-a88b-4f4a-bb8d-da4112e85882-c5b07797'\n")
		fmt.Scanln(&parentQid)

		var mark int
		fmt.Print("Export quantity: ")
		fmt.Scanln(&mark)

		sid := GetSID()

		var wg sync.WaitGroup
		maxGoroutines := 8
		guard := make(chan struct{}, maxGoroutines)

		if mark < 1000 {
			wg.Add(1)
			guard <- struct{}{}
			go func() {
				defer wg.Done()
				data := SendRequestWOS(parentQid, 1, mark, sid)
				SaveDataAsFile(data, 1, mark, parentQid)
				<-guard
			}()
		} else {
			start := 1
			for start <= mark {
				end := start + 999
				if end > mark {
					end = mark
				}
				wg.Add(1)
				guard <- struct{}{}
				go func(s, e int) {
					defer wg.Done()
					data := SendRequestWOS(parentQid, s, e, sid)
					SaveDataAsFile(data, s, e, parentQid)
					<-guard
				}(start, end)
				start = end + 1
			}
		}

		wg.Wait()

		for {
			fmt.Println("Press Enter to restart or type 'exit' to quit.")
			scanner.Scan()
			input := scanner.Text()
			if input == "exit" {
				return
			} else if input == "" {
				break
			}
		}
	}
}
