package main

// Links para estudar depois:
// https://blog.boot.dev/golang/how-to-use-golangs-generics/
// https://go.dev/doc/tutorial/generics#add_generic_function
// https://bitfieldconsulting.com/golang/generics

func SomaInteiros(mp map[string]int64) (soma int64) {
	for _, x := range mp {
		soma += x
	}
	return
}

// func SomaGenerica[T int64 | float64](mp map[string]T) (soma T) {
// 	for _, x := range mp {
// 		soma += x
// 	}
// 	return
// }

type Number interface {
	~int | ~int64 | ~float64
}

type MyNumber int

func SomaGenerica[T Number](mp map[string]T) (soma T) {
	for _, x := range mp {
		soma += x
	}
	return
}

// func main() {
// 	somaInteiros := map[string]int64{"a": 1, "b": 2, "c": 3, "d": 4}
// 	fmt.Println(SomaInteiros(somaInteiros))

// 	// somaGenericos := map[string]float64{"a": 1, "b": 2.4, "c": 3, "d": 4.4}
// 	// fmt.Println(SomaGenerica(somaGenericos))

// 	// var x, y, z Number

// 	somaGenerica := map[string]MyNumber{"a": 1, "b": 2, "c": 3, "d": 4}
// 	fmt.Println(SomaGenerica(somaGenerica))
// }
