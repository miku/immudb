/*
Copyright 2022 Codenotary Inc. All rights reserved.

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

package client

import (
	"context"

	"github.com/codenotary/immudb/pkg/api/schema"
)

func (c *immuClient) ExportTx(ctx context.Context, req *schema.ExportTxRequest) (schema.ImmuService_ExportTxClient, error) {
	if req == nil {
		return nil, ErrIllegalArguments
	}

	if !c.IsConnected() {
		return nil, ErrNotConnected
	}

	return c.ServiceClient.ExportTx(ctx, req)
}

func (c *immuClient) ReplicateTx(ctx context.Context) (schema.ImmuService_ReplicateTxClient, error) {
	if !c.IsConnected() {
		return nil, ErrNotConnected
	}

	return c.ServiceClient.ReplicateTx(ctx)
}
