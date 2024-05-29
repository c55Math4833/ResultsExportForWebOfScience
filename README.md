# Results Export For WebOfScience
[繁體中文](README_zhTW.md) | [简体中文](README_zhCN.md)

A tool written in Go for exporting Web of Science search results to plain text documents (especially useful when several thousand articles are to be shipped at once). It can be packaged for different operating environments using Go Build, and after packaging, there is no need for an additional Go language environment setup.

## Requirements

1. Access to Web of Science's online environment legally.
2. This code (requires a Go language environment) or the packaged program file (no need for a particular environment setup).

## Instructions

1. Complete the search on the Web of Science search page.
2. Run code `WebOfScienceResultsExport.go` or the packaged program.
3. Enter the page QID of the Web of Science search page.
    > For example: When the URL is "www.webofscience.com/wos/woscc/summary/104f891b-a88b-4f4a-bb8d-da4112e85882-c5b07797/relevance/1", QID = 104f891b-a88b-4f4a-bb8d-da4112e85882-c5b07797.

4. Enter the total number of documents to be exported.
5. Wait for the search results to finish exporting, and the program will automatically close.
    > Note: A tiny empty file may sometimes be produced due to network issues. Rerun the process if this occurs.

    > Note: Due to the limitations of the Web of Science web page itself, when exporting more than 100,000 documents, documents exceeding 100,000 cannot be exported.

6. [Optional] Splice data: In the folder where the search results are exported, you can run the code `WebOfScienceSplice.go` or the packaged program to merge the output data into a single text file.
