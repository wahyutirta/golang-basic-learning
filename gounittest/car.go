package main

type Speeder interface {
	MaxSpeed() int
}

type DefaultEngine struct {
}

func (e *DefaultEngine) MaxSpeed() int {
	return 50
}

type TurboEngine struct {
}

func (e *TurboEngine) MaxSpeed() int {
	return 100
}

type Car struct {
	Speeder Speeder
}

func (c *Car) Speed() int {
	defaultSpeed := 80
	v := c.Speeder.MaxSpeed()
	if v >= defaultSpeed {
		return c.Speeder.MaxSpeed()
	}
	if v <= 20 {
		return 20
	}
	return defaultSpeed
}

func NewCar(speeder Speeder) *Car {
	return &Car{
		Speeder: speeder,
	}
}
