package simple

import (
	"fmt"

	"github.com/haproxytech/config-parser/errors"
)

type SimpleTime struct {
	Enabled    bool
	Value      string
	Name       string
	SearchName string
}

func (s *SimpleTime) Init() {
	s.Enabled = false
	s.SearchName = s.Name
}

func (s *SimpleTime) GetParserName() string {
	return s.SearchName
}

func (s *SimpleTime) Parse(line string, parts, previousParts []string) (changeState string, err error) {
	if parts[0] == s.SearchName {
		if len(parts) < 2 {
			return "", &errors.ParseError{Parser: "SimpleTime", Line: line, Message: "Parse error"}
		}
		s.Enabled = true
		s.Value = parts[1]
		return "", nil
	}
	return "", &errors.ParseError{Parser: s.SearchName, Line: line}
}

func (s *SimpleTime) Valid() bool {
	if s.Enabled {
		return true
	}
	return false
}

func (s *SimpleTime) Result(AddComments bool) []string {
	if s.Enabled {
		return []string{fmt.Sprintf("  %s %s", s.SearchName, s.Value)}
	}
	return []string{}
}