package main

import (
	"log"
	"os"

	"main.go/search"
)

func main() {
	// 지정된 검색어로 검색을 수행
	search.Run("Sherlock Holmes")
}

func init() {
	// 표준 출력으로 로그를 출력하도록 변경한다.
	log.SetOutput(os.Stdout)
}
