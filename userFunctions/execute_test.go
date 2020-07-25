package userFunctions

import (
	"regexp"
	"testing"
)

func TestExecuteCopy(t *testing.T) {
	_ = "/Users/300069858/work/personal/options-viewer/backend/config/customData.go"
	data2 := "customData.go"
	regex, err := regexp.Compile(`custom`)
	if err != nil {
		t.Error(err)
	}
	if regex.MatchString(data2) {
		t.Log(regex.FindString(data2))
	} else {
		t.Fail()
	}
}
