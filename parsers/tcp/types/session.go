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

package types

import (
	"fmt"
	"strings"

	"github.com/deyunluo/config-parser/v4/common"
	"github.com/deyunluo/config-parser/v4/errors"
	"github.com/deyunluo/config-parser/v4/parsers/tcp/actions"
	"github.com/deyunluo/config-parser/v4/types"
)

type Session struct {
	Action   types.TCPAction
	Cond     string
	CondTest string
	Comment  string
}

func (f *Session) ParseAction(action types.TCPAction, parts []string) error {
	err := action.Parse(parts)
	if err != nil {
		return &errors.ParseError{Parser: "TCPRequestSession", Line: ""}
	}

	f.Action = action

	return nil
}

func (f *Session) Parse(parts []string, comment string) error {
	if comment != "" {
		f.Comment = comment
	}

	if len(parts) >= 3 {

		command, condition := common.SplitRequest(parts[2:])

		var err error

		if len(command) > 0 {

			switch command[0] {
			case "accept":
				err = f.ParseAction(&actions.Accept{}, command)
			case "reject":
				err = f.ParseAction(&actions.Reject{}, command)
			case "track-sc0":
				err = f.ParseAction(&actions.TrackSc0{}, command)
			case "track-sc1":
				err = f.ParseAction(&actions.TrackSc1{}, command)
			case "track-sc2":
				err = f.ParseAction(&actions.TrackSc2{}, command)
			case "silent-drop":
				err = f.ParseAction(&actions.SilentDrop{}, command)
			default:
				switch {
				case strings.HasPrefix(command[0], "sc-inc-gpc0"):
					err = f.ParseAction(&actions.ScIncGpc0{}, command)
				case strings.HasPrefix(command[0], "sc-inc-gpc1"):
					err = f.ParseAction(&actions.ScIncGpc1{}, command)
				case strings.HasPrefix(command[0], "sc-set-gpt0"):
					err = f.ParseAction(&actions.ScSetGpt0{}, command)
				case strings.HasPrefix(command[0], "set-var"):
					err = f.ParseAction(&actions.SetVar{}, command)
				case strings.HasPrefix(command[0], "unset-var"):
					err = f.ParseAction(&actions.UnsetVar{}, command)
				default:
					return err
				}
			}

			if err != nil {
				return err
			}

		} else {
			return fmt.Errorf("not enough params")
		}

		if len(condition) > 1 {
			f.Cond = condition[0]
			f.CondTest = strings.Join(condition[1:], " ")
		}

		return nil
	}
	return fmt.Errorf("not enough params")
}

func (f *Session) String() string {
	var result strings.Builder

	result.WriteString("session")
	result.WriteString(" ")
	result.WriteString(f.Action.String())

	if f.Cond != "" {
		result.WriteString(" ")
		result.WriteString(f.Cond)
		result.WriteString(" ")
		result.WriteString(f.CondTest)
	}

	if f.Comment != "" {
		result.WriteString(" # ")
		result.WriteString(f.Comment)
	}

	return result.String()
}

func (f *Session) GetComment() string {
	return f.Comment
}
