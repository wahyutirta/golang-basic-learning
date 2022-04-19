package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDefaultEngine_MaxSpeed(t *testing.T) {
	testCases := []struct {
		name string
		want int
	}{
		{name: "default engine", want: 50},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			e := DefaultEngine{}
			assert.Equal(t, e.MaxSpeed(), tc.want)
		})

	}
}

func TestTurboEngine_MaxSpeed(t *testing.T) {
	testCases := []struct {
		name string
		want int
	}{
		{name: "turbo engine", want: 100},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			e := TurboEngine{}
			assert.Equal(t, e.MaxSpeed(), tc.want)
		})

	}
}

type FakeEngine struct{}

func (e *FakeEngine) MaxSpeed() int {
	return 5
}

func TestCar_Speed(t *testing.T) {
	type fields struct {
		Speeder Speeder
	}
	testCases := []struct {
		name   string
		fields fields
		want   int
	}{
		{name: "default engine car", fields: fields{Speeder: &DefaultEngine{}}, want: 80},
		{name: "turbo engine car", fields: fields{Speeder: &TurboEngine{}}, want: 100},
		{name: "fake engine car", fields: fields{Speeder: &FakeEngine{}}, want: 20}, //fake method
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c := Car{
				Speeder: tc.fields.Speeder,
			}
			if got := c.Speed(); got != tc.want {
				t.Errorf("Speed() = %v, want %v", got, tc.want)
			}
		})

	}
}

type MockEngine struct {
	mock.Mock
}

func (m *MockEngine) MaxSpeed() int {
	args := m.Called()
	return args.Get(0).(int)
}

func TestCar_Mock(t *testing.T) {
	mock := new(MockEngine)
	car := Car{
		Speeder: mock,
	}

	//return mocked function
	mock.On("MaxSpeed").Return(9)

	v := car.Speed()
	assert.Equal(t, 20, v)
}
