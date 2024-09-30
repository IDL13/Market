package uptrace

import (
	"context"
	"fmt"

	"github.com/uptrace/uptrace-go/uptrace"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

type Uptrace interface {
	GetSpan(nameOperation string, atr Atr) (trace.Span, context.Context)
	CheckError(err error, children trace.Span)
	Shutdown(ctx context.Context, children trace.Span)
}

type uptraceStruct struct {
	Tracer trace.Tracer
}

type Atr struct {
	Name string
}

func New(appName string) Uptrace {
	uptrace.ConfigureOpentelemetry(
		uptrace.WithDSN("https://tHQKAHt3A5tf_J3F6tWOVQ@api.uptrace.dev?grpc=4317"),
		uptrace.WithServiceVersion("1.0.0"),
	)
	return &uptraceStruct{
		Tracer: otel.Tracer(appName),
	}
}

func (upt *uptraceStruct) GetSpan(nameOperation string, atr Atr) (trace.Span, context.Context) {
	ctx := context.Background()

	_, children := upt.Tracer.Start(ctx, nameOperation)
	children.SetAttributes(
		attribute.String("Name", atr.Name),
	)

	return children, ctx
}

func (upt *uptraceStruct) CheckError(err error, children trace.Span) {
	if err != nil {
		children.RecordError(err, trace.WithStackTrace(true))
		children.SetStatus(codes.Error, err.Error())
		children.End()
	}
}

func WriteTrace(children trace.Span) {
	fmt.Printf("Trace := %s\n", uptrace.TraceURL(children))
}

func (upt *uptraceStruct) Shutdown(ctx context.Context, children trace.Span) {
	defer uptrace.Shutdown(ctx)
	defer children.End()
}

// ______ Test ______

// func f() error {
//     return errors.New("Error")
// }

// func main() {
// 	upt := uptrace.New("main")
// 	children, ctx := upt.GetSpan("function f", uptrace.Atr{Name: "f"})
// 	defer upt.Shutdown(ctx, children)

// 	err := f()
// 	upt.CheckError(err, children)
// 	uptrace.WriteTrace(children)
//}
