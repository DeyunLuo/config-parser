// Code generated by go generate; DO NOT EDIT.
package tests

import (
	"fmt"
	"strings"
	"testing"

	"github.com/haproxytech/config-parser/parsers"
)


func TestErrorFileNormal0(t *testing.T) {
	parser := &parsers.ErrorFile{}
	line := strings.TrimSpace("errorfile 400 /etc/haproxy/errorfiles/400badreq.http")
	err := ProcessLine(line, parser)
	if err != nil {
		t.Errorf(err.Error())
	}
	result, err := parser.Result(true)
	if err != nil {
		t.Errorf(err.Error())
	}
	var returnLine string
	if result[0].Comment == "" {
		returnLine = fmt.Sprintf("%s", result[0].Data)
	} else {
		returnLine = fmt.Sprintf("%s # %s", result[0].Data, result[0].Comment)
	}
	if line != returnLine {
		t.Errorf(fmt.Sprintf("error: has [%s] expects [%s]", returnLine, line))
	}
}
func TestErrorFileNormal1(t *testing.T) {
	parser := &parsers.ErrorFile{}
	line := strings.TrimSpace("errorfile 408 /dev/null # work around Chrome pre-connect bug")
	err := ProcessLine(line, parser)
	if err != nil {
		t.Errorf(err.Error())
	}
	result, err := parser.Result(true)
	if err != nil {
		t.Errorf(err.Error())
	}
	var returnLine string
	if result[0].Comment == "" {
		returnLine = fmt.Sprintf("%s", result[0].Data)
	} else {
		returnLine = fmt.Sprintf("%s # %s", result[0].Data, result[0].Comment)
	}
	if line != returnLine {
		t.Errorf(fmt.Sprintf("error: has [%s] expects [%s]", returnLine, line))
	}
}
func TestErrorFileNormal2(t *testing.T) {
	parser := &parsers.ErrorFile{}
	line := strings.TrimSpace("errorfile 403 /etc/haproxy/errorfiles/403forbid.http")
	err := ProcessLine(line, parser)
	if err != nil {
		t.Errorf(err.Error())
	}
	result, err := parser.Result(true)
	if err != nil {
		t.Errorf(err.Error())
	}
	var returnLine string
	if result[0].Comment == "" {
		returnLine = fmt.Sprintf("%s", result[0].Data)
	} else {
		returnLine = fmt.Sprintf("%s # %s", result[0].Data, result[0].Comment)
	}
	if line != returnLine {
		t.Errorf(fmt.Sprintf("error: has [%s] expects [%s]", returnLine, line))
	}
}
func TestErrorFileNormal3(t *testing.T) {
	parser := &parsers.ErrorFile{}
	line := strings.TrimSpace("errorfile 503 /etc/haproxy/errorfiles/503sorry.http")
	err := ProcessLine(line, parser)
	if err != nil {
		t.Errorf(err.Error())
	}
	result, err := parser.Result(true)
	if err != nil {
		t.Errorf(err.Error())
	}
	var returnLine string
	if result[0].Comment == "" {
		returnLine = fmt.Sprintf("%s", result[0].Data)
	} else {
		returnLine = fmt.Sprintf("%s # %s", result[0].Data, result[0].Comment)
	}
	if line != returnLine {
		t.Errorf(fmt.Sprintf("error: has [%s] expects [%s]", returnLine, line))
	}
}
func TestErrorFileFail0(t *testing.T) {
	parser := &parsers.ErrorFile{}
	line := strings.TrimSpace("errorfile")
	err := ProcessLine(line, parser)
	if err == nil {
		t.Errorf(fmt.Sprintf("error: did not throw error for line [%s]", line))
	}
	_, err = parser.Result(true)
	if err == nil {
		t.Errorf(fmt.Sprintf("error: did not throw error on result for line [%s]", line))
	}
}
func TestErrorFileFail1(t *testing.T) {
	parser := &parsers.ErrorFile{}
	line := strings.TrimSpace("---")
	err := ProcessLine(line, parser)
	if err == nil {
		t.Errorf(fmt.Sprintf("error: did not throw error for line [%s]", line))
	}
	_, err = parser.Result(true)
	if err == nil {
		t.Errorf(fmt.Sprintf("error: did not throw error on result for line [%s]", line))
	}
}
func TestErrorFileFail2(t *testing.T) {
	parser := &parsers.ErrorFile{}
	line := strings.TrimSpace("--- ---")
	err := ProcessLine(line, parser)
	if err == nil {
		t.Errorf(fmt.Sprintf("error: did not throw error for line [%s]", line))
	}
	_, err = parser.Result(true)
	if err == nil {
		t.Errorf(fmt.Sprintf("error: did not throw error on result for line [%s]", line))
	}
}