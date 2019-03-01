//
// Copyright (c) 2019 Intel Corporation
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
//

package runtime

import (
	"strconv"

	"github.com/edgexfoundry/go-mod-core-contracts/models"

	"github.com/edgexfoundry/app-functions-sdk-go/pkg/excontext"
)

// GolangRuntime represents the golang runtime environment
type GolangRuntime struct {
	Transforms []func(excontext.Context, ...interface{}) (bool, interface{})
}

// ProcessEvent handles processing the event
func (gr GolangRuntime) ProcessEvent(edgexcontext excontext.Context, event models.Event) error {
	edgexcontext.LoggingClient.Info("Processing Event: " + strconv.Itoa(len(gr.Transforms)) + " Transforms")
	var result interface{}
	var continuePipeline = true
	for _, trxFunc := range gr.Transforms {
		if result != nil {
			continuePipeline, result = trxFunc(edgexcontext, result)
		} else {
			continuePipeline, result = trxFunc(edgexcontext, event)
		}
		if continuePipeline != true {
			if result != nil {
				if result, ok := result.(error); ok {
					edgexcontext.LoggingClient.Error((result).(error).Error())
				}
			}
			break
		}
	}
	return nil
}
