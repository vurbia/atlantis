package matchers

import (
	"reflect"

	models "github.com/hootsuite/atlantis/server/events/models"
	"github.com/petergtz/pegomock"
)

func AnySliceOfModelsProjectlock() []models.ProjectLock {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*([]models.ProjectLock))(nil)).Elem()))
	var nullValue []models.ProjectLock
	return nullValue
}

func EqSliceOfModelsProjectlock(value []models.ProjectLock) []models.ProjectLock {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue []models.ProjectLock
	return nullValue
}
