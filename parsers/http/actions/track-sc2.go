/*
Copyright 2019 HAProxy Technologies

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

//nolint:dupl
package actions

import (
	"fmt"
	"strings"

	"github.com/deyunluo/config-parser/v4/common"
)

type TrackSc2 struct {
	Key      string
	Table    string
	Comment  string
	Cond     string
	CondTest string
}

func (f *TrackSc2) Parse(parts []string, comment string) error {
	if len(parts) < 3 {
		return fmt.Errorf("not enough params")
	}

	command, condition := common.SplitRequest(parts[2:])

	f.Key = command[0]

	if len(command) == 3 && command[1] == "table" {
		f.Table = command[2]
	}
	if len(condition) > 1 {
		f.Cond = condition[0]
		f.CondTest = strings.Join(condition[1:], " ")
	}

	return nil
}

func (f *TrackSc2) String() string {
	var result strings.Builder
	result.WriteString("track-sc2 ")
	result.WriteString(f.Key)
	if f.Table != "" {
		result.WriteString(" table ")
		result.WriteString(f.Table)
	}
	if f.Cond != "" {
		result.WriteString(" ")
		result.WriteString(f.Cond)
		result.WriteString(" ")
		result.WriteString(f.CondTest)
	}
	return result.String()
}

func (f *TrackSc2) GetComment() string {
	return f.Comment
}
