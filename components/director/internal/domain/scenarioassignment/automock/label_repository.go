// Code generated by mockery v1.0.0. DO NOT EDIT.

package automock

import context "context"
import mock "github.com/stretchr/testify/mock"
import model "github.com/kyma-incubator/compass/components/director/internal/model"

// LabelRepository is an autogenerated mock type for the LabelRepository type
type LabelRepository struct {
	mock.Mock
}

// GetRuntimeScenariosWhereLabelsMatchSelector provides a mock function with given fields: ctx, tenantID, selectorKey, selectorValue
func (_m *LabelRepository) GetRuntimeScenariosWhereLabelsMatchSelector(ctx context.Context, tenantID string, selectorKey string, selectorValue string) ([]model.Label, error) {
	ret := _m.Called(ctx, tenantID, selectorKey, selectorValue)

	var r0 []model.Label
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) []model.Label); ok {
		r0 = rf(ctx, tenantID, selectorKey, selectorValue)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Label)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string) error); ok {
		r1 = rf(ctx, tenantID, selectorKey, selectorValue)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Upsert provides a mock function with given fields: ctx, label
func (_m *LabelRepository) Upsert(ctx context.Context, label *model.Label) error {
	ret := _m.Called(ctx, label)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.Label) error); ok {
		r0 = rf(ctx, label)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}