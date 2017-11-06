package matchers

import (
	"reflect"

	logging "github.com/hootsuite/atlantis/server/logging"
	"github.com/petergtz/pegomock"
)

func AnyPtrToLoggingSimplelogger() *logging.SimpleLogger {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(*logging.SimpleLogger))(nil)).Elem()))
	var nullValue *logging.SimpleLogger
	return nullValue
}

func EqPtrToLoggingSimplelogger(value *logging.SimpleLogger) *logging.SimpleLogger {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue *logging.SimpleLogger
	return nullValue
}
