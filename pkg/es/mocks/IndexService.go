// Code generated by mockery v1.0.0

// Copyright (c) 2017 The Jaeger Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mocks

import context "context"
import elastic "gopkg.in/olivere/elastic.v5"
import es "github.com/dashbase/jaeger/pkg/es"
import mock "github.com/stretchr/testify/mock"

// IndexService is an autogenerated mock type for the IndexService type
type IndexService struct {
	mock.Mock
}

// BodyJson provides a mock function with given fields: body
func (_m *IndexService) BodyJson(body interface{}) es.IndexService {
	ret := _m.Called(body)

	var r0 es.IndexService
	if rf, ok := ret.Get(0).(func(interface{}) es.IndexService); ok {
		r0 = rf(body)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(es.IndexService)
		}
	}

	return r0
}

// Do provides a mock function with given fields: ctx
func (_m *IndexService) Do(ctx context.Context) (*elastic.IndexResponse, error) {
	ret := _m.Called(ctx)

	var r0 *elastic.IndexResponse
	if rf, ok := ret.Get(0).(func(context.Context) *elastic.IndexResponse); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*elastic.IndexResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Id provides a mock function with given fields: id
func (_m *IndexService) Id(id string) es.IndexService {
	ret := _m.Called(id)

	var r0 es.IndexService
	if rf, ok := ret.Get(0).(func(string) es.IndexService); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(es.IndexService)
		}
	}

	return r0
}

// Index provides a mock function with given fields: index
func (_m *IndexService) Index(index string) es.IndexService {
	ret := _m.Called(index)

	var r0 es.IndexService
	if rf, ok := ret.Get(0).(func(string) es.IndexService); ok {
		r0 = rf(index)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(es.IndexService)
		}
	}

	return r0
}

// Type provides a mock function with given fields: typ
func (_m *IndexService) Type(typ string) es.IndexService {
	ret := _m.Called(typ)

	var r0 es.IndexService
	if rf, ok := ret.Get(0).(func(string) es.IndexService); ok {
		r0 = rf(typ)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(es.IndexService)
		}
	}

	return r0
}
