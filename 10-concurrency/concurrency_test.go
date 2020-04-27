package main

import (
	"reflect"
	"testing"
	"time"
)

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://baidu.com",
		"http://www.163.com",
		"waat://furhurterwe.geds",
	}

	actualResults := CheckWebSites(MockWebsiteChecker, websites)

	want := len(websites)
	got := len(actualResults)

	if want != got {
		t.Fatalf("Wanted %v, got %v", want, got)
	}

	expectedResult := map[string]bool{
		"http://baidu.com":        true,
		"http://163.com":          true,
		"waat://furhurterwe.geds": false,
	}

	if !reflect.DeepEqual(expectedResult, actualResults) {
		t.Fatalf("Wanted %v, got %v", expectedResult, actualResults)
	}
}

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Microsecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)

	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}

	for i := 0; i < b.N; i++ {
		CheckWebSites(slowStubWebsiteChecker, urls)
	}
}
