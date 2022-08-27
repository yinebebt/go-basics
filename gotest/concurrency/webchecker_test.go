package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func mockwebsiteChecker(url string) bool {
	if url == "waat://hacker.nfm" {
		return false
	}
	return true
}

func TestWebstatus(t *testing.T) {
	websites := []string{
		"http://www.facebook.com",
		"https://www.ebc.com",
		"waat://hacker.nfm",
	}
	want := map[string]bool{
		"http://www.facebook.com": true,
		"https://www.ebc.com":     true,
		"waat://hacker.nfm":       false,
	}

	got := checkWebsites(mockwebsiteChecker, websites)

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v want %v", got, want)
	}
}

//let we check the speed, and following implement concurrency

func fakewebchecker(_ string) bool { // for benchmark testing purpose, just for comaprission
	time.Sleep(20 * time.Millisecond) //just to have a very small duration of time
	return true
}

func BenchmarkWebchecheck(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}
	b.ResetTimer() // just to have a time forward this, after the test execution we needed is started
	for i := 0; i < b.N; i++ {
		checkWebsites(fakewebchecker, urls)
	}

}
