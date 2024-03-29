// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package noop provides noop tracing implementations for tracer and span.
package noop

import (
	"context"

	"go.opentelemetry.io/otel/api/trace"
)

var (
	// Tracer is a noop tracer that starts noop spans.
	Tracer trace.Tracer

	// Span is a noop Span.
	Span trace.Span
)

func init() {
	Tracer = trace.NoopTracerProvider().Tracer("")
	_, Span = Tracer.Start(context.Background(), "")
}
