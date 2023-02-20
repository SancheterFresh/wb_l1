package main

type value interface {
	printValue()
}

type client struct {
}

func (c *client) getValue(val value) {
	val.printValue()
}

type valueX struct {
	x int
}

func (vx *valueX) printValue() {
	println(vx.x)
}

type valueY struct {
	y int
}

func (vy *valueY) outputValue() {
	println(vy.y)
}

type yAdapter struct {
	valY *valueY
}

func (yadpt *yAdapter) printValue() {
	yadpt.valY.outputValue()
}

func main() {
	client := &client{}
	valueX := &valueX{x: 5}

	client.getValue(valueX)

	valueY := &valueY{y: 7}

	yAdapter := &yAdapter{
		valY: valueY,
	}

	client.getValue(yAdapter)

}
