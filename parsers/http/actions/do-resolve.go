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

package actions

import (
	"fmt"
	"strings"

	"github.com/deyunluo/config-parser/v4/common"
)

type DoResolve struct {
	Var       string
	Resolvers string
	Protocol  string
	Expr      common.Expression
	Cond      string
	CondTest  string
	Comment   string
}

func (f *DoResolve) Parse(parts []string, comment string) error {
	if len(parts) < 3 {
		return fmt.Errorf("not enough params")
	}

	data := strings.TrimPrefix(parts[1], "do-resolve(")
	data = strings.TrimRight(data, ")")
	d := strings.SplitN(data, ",", 3)

	if len(d) < 2 {
		return fmt.Errorf("not enough params")
	}

	f.Var = d[0]
	f.Resolvers = d[1]
	if len(d) == 3 {
		f.Protocol = d[2]
	}

	command, condition := common.SplitRequest(parts[2:]) // 2 not 3 !

	if len(command) > 0 {
		expr := common.Expression{}
		err := expr.Parse(command)
		if err != nil {
			return fmt.Errorf("not enough params")
		}
		f.Expr = expr
	}

	if len(condition) > 1 {
		f.Cond = condition[0]
		f.CondTest = strings.Join(condition[1:], " ")
	}

	return nil
}

func (f *DoResolve) String() string {
	var stmt string
	if len(f.Protocol) > 0 {
		stmt = fmt.Sprintf("%s,%s,%s", f.Var, f.Resolvers, f.Protocol)
	} else {
		stmt = fmt.Sprintf("%s,%s", f.Var, f.Resolvers)
	}

	if f.Cond == "" {
		return fmt.Sprintf("do-resolve(%s) %s", stmt, f.Expr.String())
	}
	return fmt.Sprintf("do-resolve(%s) %s %s %s", stmt, f.Expr.String(), f.Cond, f.CondTest)
}

func (f *DoResolve) GetComment() string {
	return f.Comment
}
