package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"
)

func main() {

	val := 23

	fmt.Printf("hello, world\n")
	fmt.Println(" Test: ", val)

	numslice := []int{2, 3, 4, 5}

	// dynamically define slice with len and capacity
	// int slices inited to 0 for len provided
	numslice2 := make([]int, 6, 10)

	copy(numslice2, numslice)

	numslice2 = append(numslice2, 2, 6, 7)
	for _, val := range numslice2 {
		fmt.Printf("%d ", val)
	}
	fmt.Println()

	// Map use case string ==> int
	// var stocks map[string]int
	stocks := make(map[string]int)

	stocks["INTC"] = 50

	fmt.Println(stocks["INTC"])
	stocks["CRM"] = 200

	ok := delVal(stocks, "AMZ")

	if ok {
		fmt.Println(len(stocks))
	} else {
		fmt.Println("Not there")
	}

	addVals(stocks, numslice2, "XYZ", "PQR", "MSFT")

	fmt.Println(len(stocks), stocks)

	//defer test
	defer sysexit("Done TGIF")

	//closure test
	var a, b int = 12, 13
	//a = 12
	//b = 13
	//var sum, diff int
	prod := func() int {
		a += 2
		b += 2
		return a + b
	}
	fmt.Printf("closure func: %d, a=%d, b=%d\n", prod(), a, b)
	//fmt.Printf("%d\n",prod());

	fmt.Println(factorial(3))

	//math
	//fmt.Println(safeDiv(6,0))
	iptr := new(int)
	*iptr = 7
	fmt.Println("Value at ptr:", iptr, "=", *iptr)
	doubleIt(iptr)
	fmt.Println("Value at ptr:", iptr, "=", *iptr)

	fmt.Printf("\n---------------\n")

	//struct usage
	// var myself Person
	// myself.age=35
	myself := Person{age: 35, name: "Abhishek", addr: "Roosevelt Ave"}
	alexa := Robot{name: "Alexa", model: "Nexus", version: 6}
	//printPerson(myself)
	myself.printDetails()

	//enquire(myself);
	enquire(alexa)

	fmt.Printf("\n---------------\n")

	//string manipulation
	myStr := "Hasta Manana Maria"
	fmt.Println(strings.HasSuffix(myStr, "na"))
	s2 := strings.Replace(myStr, "na", "la", 5)
	fmt.Println(s2)
	fmt.Println(strings.Index(s2, "al"))
	fmt.Println(strings.SplitAfter(s2, "M"))

	//algo
	if !sort.IntsAreSorted(numslice2) {
		sort.Ints(numslice2)
		fmt.Println(numslice2)
	}

	fmt.Printf("\n---------------\n")

	//OS calls and file IO
	file, err := os.Create("sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	file.WriteString("Hola Senora!")
	file.Close()

	stream, err := ioutil.ReadFile("sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(stream))

	e := os.Remove("sample.txt")
	if e != nil {
		log.Fatal(e)
	}

	fmt.Printf("\n---------------\n")

	//Like maps and slices, channels must be created before use:
	c := make(chan int, 10)
	go fibn(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}

	fmt.Printf("\n---------------\n")

	//panic
	dontPanic()

	// iterator tests tests
	s := NewSliceIterator([]string{"a", "b", "c"})
	IteratorFun(s)

	r := NewRoundRobinSliceIterator([][]string{
		{"a1", "a2", "a3"},
		{"b1", "b2"},
		{"c1", "c2", "c3"},
	})

	IteratorFun2(r)
	IteratorFun(r)

	fmt.Printf("\n---------------\n")

}

//----STRUCTS/INTERFACES ------

// Person is a person
type Person struct {
	name, addr string
	// addr string
	age int
}

// Robot is a robot
type Robot struct {
	name, model string
	// model string
	version int
}

// Entity is an entity
type Entity interface {
	printDetails()
	isHuman() bool
}

func (r Robot) printDetails() {
	fmt.Println(r.name, "--", r.model, "-v", r.version)
}

func (p Person) printDetails() {
	fmt.Println(p.name, "--", p.addr, "--", p.age)
}

func (r Robot) isHuman() bool {
	return false
}

// func (p Person) isHuman() bool{
// 	return true
// }

// -----FUNCTIONS-------
func fibn(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	//close channel after n elements
	close(c)

}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func handler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola Taranika!")
}

func delVal(stk map[string]int, sym string) bool {
	var ret bool

	//check if value exists
	_, ret = stk[sym]

	delete(stk, sym)

	return ret
}

func addVals(stk map[string]int, vals []int, syms ...string) {
	var i int
	var str string
	for i, str = range syms {
		stk[str] = vals[i]
	}
}

func factorial(num int) int {
	if num == 0 {
		return 1
	}
	return (num * factorial(num-1))
}

func sysexit(str string) {
	fmt.Println("Exiting with string", str)
}

func safeDiv(num1, num2 int) float64 {
	defer func() {
		fmt.Println(recover())
	}()

	result := float64(num1 / num2)
	return result
}

func dontPanic() {
	defer func() {
		fmt.Println(recover())
	}()

	panic("OMG!!")
}

func doubleIt(num *int) {
	*num *= 2
}

func printPerson(p Person) {
	fmt.Println(p.name, "--", p.addr, "--", p.age)
}

func enquire(e Entity) {
	fmt.Println("The given entity:")
	e.printDetails()
	if e.isHuman() {
		fmt.Println("It is a Human Being")
	} else {
		fmt.Println("It is NOT a Human Being")
	}

}
