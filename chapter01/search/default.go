package search

// 기본 검색기를 구현할 defaultMatcher 타입
type defaultMatcher struct{}

// init 함수에서는 기본 검색기를 프로그램에 등록한다.
func init() {
	var matcher defaultMatcher

	Register("default", matcher)
}

func (m defaultMatcher) Search(feed *Feed, searchTerm string) ([]*Result, error) {
	return nil, nil
}
