# Results Export For WebOfScience
[英文](README.md) | [简体中文](README_zhCN.md)

使用 Go 語言撰寫的 Web of Science 搜尋結果匯出為純文字文件工具（特別適用於同時有幾千篇要匯出時），可使用 go build 封裝至不同作業環境使用，封裝後則不須另外進行 Go 語言環境建置。

## 需求

1. 可以合法使用 Web of Science 的網路環境。
2. 本代碼（需要 Go 語言環境）或是封裝後的程式檔（無須特別環境建置）。

## 使用說明

1. 於 Web of Science 搜尋頁面完成搜尋。

2. 執行代碼或封裝後的程式。

3. 輸入 Web of Science 搜尋頁面之頁面 QID。

	> 舉例：當網址為 "www.webofscience.com/wos/woscc/summary/104f891b-a88b-4f4a-bb8d-da4112e85882-c5b07797/relevance/1" 時，QID = 104f891b-a88b-4f4a-bb8d-da4112e85882-c5b07797。

4. 輸入要匯出的文獻總數。

5. 等候搜尋結果匯出完成即自動關閉。

	> 註：有時由於網路問題會產生極小之空文件，重新執行即可。

