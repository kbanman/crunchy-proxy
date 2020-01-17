/*
Copyright 2017 Crunchy Data Solutions, Inc.
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

package grpcutil

import (
	"io"
	"strings"

	"golang.org/x/net/context"
	"golang.org/x/net/http2"

	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func IsClosedConnection(err error) bool {
	err = errors.Cause(err)

	if err == context.Canceled ||
		status.Code(err) == codes.Canceled ||
		status.Code(err) == codes.Unavailable ||
		strings.Contains(err.Error(), "is closing") ||
		strings.Contains(err.Error(), "tls: use of closed connection") ||
		strings.Contains(err.Error(), "use of closed network connection") ||
		strings.Contains(err.Error(), io.ErrClosedPipe.Error()) ||
		strings.Contains(err.Error(), io.EOF.Error()) {
		return true
	}

	if streamErr, ok := err.(http2.StreamError); ok && streamErr.Code == http2.ErrCodeCancel {
		return true
	}

	return false
}
