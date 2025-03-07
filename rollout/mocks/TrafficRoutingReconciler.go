// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	trafficrouting "github.com/argoproj/argo-rollouts/rollout/trafficrouting"
	mock "github.com/stretchr/testify/mock"
)

// TrafficRoutingReconciler is an autogenerated mock type for the TrafficRoutingReconciler type
type TrafficRoutingReconciler struct {
	mock.Mock
}

// SetWeight provides a mock function with given fields: desiredWeight, additionalDestinations
func (_m *TrafficRoutingReconciler) SetWeight(desiredWeight int32, additionalDestinations ...trafficrouting.WeightDestination) error {
	_va := make([]interface{}, len(additionalDestinations))
	for _i := range additionalDestinations {
		_va[_i] = additionalDestinations[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, desiredWeight)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(int32, ...trafficrouting.WeightDestination) error); ok {
		r0 = rf(desiredWeight, additionalDestinations...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Type provides a mock function with given fields:
func (_m *TrafficRoutingReconciler) Type() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// UpdateHash provides a mock function with given fields: canaryHash, stableHash
func (_m *TrafficRoutingReconciler) UpdateHash(canaryHash string, stableHash string) error {
	ret := _m.Called(canaryHash, stableHash)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(canaryHash, stableHash)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// VerifyWeight provides a mock function with given fields: desiredWeight
func (_m *TrafficRoutingReconciler) VerifyWeight(desiredWeight int32) (bool, error) {
	ret := _m.Called(desiredWeight)

	var r0 bool
	if rf, ok := ret.Get(0).(func(int32) bool); ok {
		r0 = rf(desiredWeight)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int32) error); ok {
		r1 = rf(desiredWeight)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
