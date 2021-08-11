// Copyright 2021 Juan Pablo Tosso
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package actions

import (
	"fmt"
	"strconv"

	engine "github.com/jptosso/coraza-waf"
)

type Phase struct{}

func (a *Phase) Init(r *engine.Rule, data string) error {
	i, err := strconv.Atoi(data)
	if data == "request" {
		i = 2
	} else if data == "response" {
		i = 4
	} else if data == "logging" {
		i = 5
	} else if err != nil || i > 5 || i < 1 {
		return fmt.Errorf("Invalid phase %s", data)
	}
	r.Phase = int(i)
	return nil
}

func (a *Phase) Evaluate(r *engine.Rule, tx *engine.Transaction) {
	// Not evaluated
}

func (a *Phase) Type() int {
	return engine.ACTION_TYPE_METADATA
}
