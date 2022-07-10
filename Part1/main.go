package main

import "fmt"

type getAnimaWaight interface {
	getWaight() float64
}

func foderPerMontCat(wgh getAnimaWaight, numMontes int) float64 {
	return wgh.getWaight() * float64(numMontes)
}

func main() {

	//var afm getAnimaWaight
	//
	//cat1 := cat{
	//	name:  "Mike",
	//	weight: 15,
	//}
	//cow1 := cow{
	//	name:  "Mike",
	//	weight: 15,
	//}
	//dog1 := dog{
	//	name:  "Mike",
	//	weight: 15,
	//}

	ferm := map[string]getAnimaWaight{
		"cat Noran": cat{weight: 20},
		"cow Benni": cow{weight: 120},
		"dog Jack":  dog{weight: 32},
	}

	for animalName, animalWaight := range ferm {

		monthes := 6
		fodderNeded := foderPerMontCat(animalWaight, monthes)
		fmt.Printf("\n Foderr neaded for %s : %f", animalName, fodderNeded)

	}

	//afm = cat1

}
