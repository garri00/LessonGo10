package main

type cat struct {
	name   string
	weight float64
	meaw   string
	//catFeeder := 7
}

func (ct cat) getWaight() float64 {
	return ct.weight * 7
}

type cow struct {
	name   string
	weight float64
	milk   float64
}

func (cw cow) getWaight() float64 {
	return cw.weight * 25
}

type dog struct {
	name   string
	weight float64
}

func (dg dog) getWaight() float64 {
	return dg.weight * 7
}
