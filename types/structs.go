package main

import "fmt"

const usixteenbitmax float64 = 65535
const kmh_multiple float64 = 1.60934

type car struct {
	gas_pedal      uint16
	brake_pedal    uint16
	steering_wheel int16
	top_speed_kmh  float64
}

// Value receiver methods for structs
// Doesn't modify any of the values
func (c car) kmh() float64 {
	return float64(c.gas_pedal) * (c.top_speed_kmh / usixteenbitmax)
}

func (c car) mph() float64 {
	return c.kmh() / kmh_multiple
}

// Pointer receiver methods
// Modifies values inside the structs
func (c *car) set_top_speed(newspeed float64) {
	c.top_speed_kmh = newspeed
}

// The difference is basically passing by value x passing by reference

// Also possible to write it as a basic function instead of a method
func set_top_speed_func(c car, speed float64) car {
	c.top_speed_kmh = speed
	return c
}

func main() {
	a_car := car{
		gas_pedal:      1234,
		brake_pedal:    0,
		steering_wheel: 10,
		top_speed_kmh:  200.0,
	}

	fmt.Println(a_car.gas_pedal)
	fmt.Println(a_car.kmh())
	fmt.Println(a_car.mph())
	a_car.set_top_speed(500)
	fmt.Println(a_car.kmh())
	fmt.Println(a_car.mph())

	// Golang also has anonymous structs

	aDoctor := struct{ name string }{name: "John Pertwee"}
	fmt.Println(aDoctor)

	// Go passes structs by value, so watch out for big struct copying
	anotherDoctor := aDoctor
	anotherDoctor.name = "Tom Baker"
	fmt.Println(aDoctor)
	fmt.Println(anotherDoctor)
}
