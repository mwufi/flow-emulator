/*
 * Flow Emulator
 *
 * Copyright 2019-2020 Dapper Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package convert

import (
	"github.com/onflow/flow-go/fvm"

	sdkConvert "github.com/onflow/flow-emulator/convert/sdk"
	"github.com/onflow/flow-emulator/types"
)

func VMTransactionResultToEmulator(tp *fvm.TransactionProcedure, txIndex int) types.TransactionResult {
	txID := sdkConvert.FlowIdentifierToSDK(tp.ID)

	sdkEvents := sdkConvert.RuntimeEventsToSDK(tp.Events, txID, txIndex)

	return types.TransactionResult{
		TransactionID: txID,
		Error:         VMErrorToEmulator(tp.Err),
		Logs:          tp.Logs,
		Events:        sdkEvents,
	}
}

func VMErrorToEmulator(vmError fvm.Error) error {
	if vmError == nil {
		return nil
	}

	return &types.FlowError{FlowError: vmError}
}
