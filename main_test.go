package main

import (
	"regexp"
	"testing"
)

func Test_Version(t *testing.T) {
	matched, err := regexp.MatchString(`^\d\.\d\.\d$`, version)
	if err != nil {
		t.Fatal(err)
	}
	if !matched {
		t.Logf("%s is not sementically versioned.", version)
		t.Fail()
	}
}
