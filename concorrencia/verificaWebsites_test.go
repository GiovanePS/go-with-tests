package concorrencia

import (
	"reflect"
	"testing"
	"time"
)

func mockVerificadorWebsites(url string) bool {
	return !(url == "waat://furhurterwe.geds")
}

func TestVerificadorWebsites(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://blog.gypsydave5.com",
		"waat://furhurterwe.geds",
	}

	expected := map[string]bool{
		"http://google.com":          true,
		"http://blog.gypsydave5.com": true,
		"waat://furhurterwe.geds":    false,
	}

	output := VerificaWebsites(mockVerificadorWebsites, websites)

	if !reflect.DeepEqual(output, expected) {
		t.Fatalf("output %v, expected %v", output, expected)
	}
}

func slowStubVerificadorWebsite(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkVerificaWebsites(b *testing.B) {
	urls := make([]string, 100)

	for i := 0; i < len(urls); i++ {
		urls[i] = "url"
	}

	for i := 0; i < b.N; i++ {
		VerificaWebsites(slowStubVerificadorWebsite, urls)
	}
}
