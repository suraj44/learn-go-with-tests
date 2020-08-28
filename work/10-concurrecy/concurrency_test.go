package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	if url == "waat://fusballspeilen.de" {
		return false
	}
	return true
}

func slowWebsiteChecker(url string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://google.com",
		"waat://fusballspeilen.de",
	}

	want := map[string]bool{
		"http://google.com":        true,
		"waat://fusballspeilen.de": false,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}

	for i := 0; i < b.N; i++ {
		CheckWebsites(slowWebsiteChecker, urls)
	}
}
