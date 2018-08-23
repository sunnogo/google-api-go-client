// Copyright 2016 Google Inc. All Rights Reserved.
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

// +build appengine

package grpc

import (
	"net"
	"time"

	"github.com/sunnogo/net/context"
	"github.com/sunnogo/appengine"
	"github.com/sunnogo/appengine/socket"
	"github.com/sunnogo/grpc-go"
)

func init() {
	// NOTE: dev_appserver doesn't currently support SSL.
	// When it does, this code can be removed.
	if appengine.IsDevAppServer() {
		return
	}

	appengineDialerHook = func(ctx context.Context) grpc.DialOption {
		return grpc.WithDialer(func(addr string, timeout time.Duration) (net.Conn, error) {
			return socket.DialTimeout(ctx, "tcp", addr, timeout)
		})
	}
}
