package log

import (
	"context"
	"fmt"
	"runtime"
	"strings"

	"github.com/segmentio/ksuid"
)

// Info is general info log
func Info(ctx context.Context, message string, args ...interface{}) {
	getLogImpl().Info(ctx, composeAdditionalInfo(ctx, 3)+message, args...)
}

// Error is General error log
func Error(ctx context.Context, message string, args ...interface{}) {
	getLogImpl().Error(ctx, composeAdditionalInfo(ctx, 3)+message, args...)
}

// InfoRequest is specific request log with additional REQUEST words in the log prefix
func InfoRequest(ctx context.Context, data interface{}) {
	getLogImpl().Info(ctx, composeAdditionalInfo(ctx, 3)+fmt.Sprintf("REQUEST  %s", data))
}

// InfoResponse is specific response with additional RESPONSE word in the log prefix
func InfoResponse(ctx context.Context, data interface{}) {
	getLogImpl().Info(ctx, composeAdditionalInfo(ctx, 3)+fmt.Sprintf("RESPONSE %s", data))
}

// ErrorResponse is specific error for any response with additional RESPONSE word in the log prefix
func ErrorResponse(ctx context.Context, err error) {
	getLogImpl().Error(ctx, composeAdditionalInfo(ctx, 3)+fmt.Sprintf("RESPONSE %s", err.Error()))
}

// UseRotateFile called if we need to have log file capability
func UseRotateFile(path, name string, maxAgeInDays int) {
	setFile(path, name, maxAgeInDays)
}

// ContextWithOperationID is always called in the beginning request of controller
// This method will generate the operation ID that will distributed to the next method call via context
func ContextWithOperationID(ctx context.Context) context.Context {
	opIDInterface := ctx.Value(operationIDField)
	if opIDInterface == nil {
		return context.WithValue(ctx, operationIDField, initOperationID())
	}
	return ctx
}

// getOperationID extratc the operation id from context
func getOperationID(ctx context.Context) string {

	operationIDInterface := ctx.Value(operationIDField)
	if operationIDInterface == nil {
		return "000000000000000000000000000"
	}

	operationID, ok := operationIDInterface.(string)
	if !ok {
		return "000000000000000000000000000"
	}

	return operationID
}

// initOperationID generate operation id
func initOperationID() string {
	return ksuid.New().String()
}

// getFunctionCall get the function information like filename and line number
// skip is the parameter that need to adjust if we add new method layer
func getFunctionCall(skip int) string {
	pc, _, line, ok := runtime.Caller(skip)
	if !ok {
		return ""
	}
	funcName := runtime.FuncForPC(pc).Name()
	x := strings.LastIndex(funcName, "/")
	return fmt.Sprintf("%s:%d", funcName[x+1:], line)
}

// composeAdditionalInfo append operation id and function information
func composeAdditionalInfo(ctx context.Context, skip int) string {
	return fmt.Sprintf("%s %s ", getOperationID(ctx), getFunctionCall(skip))
}
