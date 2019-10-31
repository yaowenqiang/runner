package main

import (
	"fmt"
	"pack"
	"math/rand"
	"reflect"
	"testing/quick"
	"time"
	"testing"
)

func main() {
	/*
	pi := pack.PolyIntegrator{}

	fmt.Println(pi.Integrate(0, 10, 3))
	fmt.Println(pi.Integrate(0, 10, 1, 0))

	ri := pack.RiemannIntegrator{}

	fmt.Println(ri.Integrate(0, 10, 3))
	fmt.Println(ri.Integrate(0, 10, 1, 0))

	fmt.Println(pack.QuickSort(1,2,3,4))
	fmt.Println(pack.QuickSort(10, 8, 2, 10, 1000, 20001,2,3,4))
	*/

	//val, ok := quick.Value(reflect.TypeOf(1),
	val, ok := quick.Value(reflect.TypeOf(MyInt(1)),
		rand.New(rand.NewSource(time.Now().Unix())))

	if ok {
		fmt.Println(val.Int())
	}


//	pack.PrintWeather(101010100)

	br := testing.Benchmark(func(b *testing.B) {
		for i := 0; i < b.N; i ++ {
			pack.PrintWeather(101010100)
		}
	})
	fmt.Println(br)
}
type MyInt int

func (mi MyInt) Generate(rand *rand.Rand, size int) reflect.Value {
	result  := rand.Float32() * 20. - 10.

	return reflect.ValueOf(int(result))
}
