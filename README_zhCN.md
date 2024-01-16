# Results Export For WebOfScience
[English](README.md) | [繁體中文](README_zhTW.md)  
使用 Go 语言撰写的 Web of Science 搜索结果导出为纯文本文档工具（特别适用于同时有几千篇要导出时），可使用 go build 封装至不同作业环境使用，封装后则不须另外进行 Go 语言环境建置。

## 需求

1. 可以合法使用 Web of Science 的网络环境。
2. 本代码（需要 Go 语言环境）或是封装后的程序档（无须特别环境建置）。

## 使用说明

1. 于 Web of Science 搜索页面完成搜索。

2. 运行代码或封装后的程序。

3. 输入 Web of Science 搜索页面之页面 QID。

	> 举例：当网址为 "www.webofscience.com/wos/woscc/summary/104f891b-a88b-4f4a-bb8d-da4112e85882-c5b07797/relevance/1" 时，QID = 104f891b-a88b-4f4a-bb8d-da4112e85882-c5b07797。

4. 输入要导出的文献总数。

5. 等候搜索结果导出完成即自动关闭。

	> 注：有时由于网络问题会产生极小之空文档，重新运行即可。
