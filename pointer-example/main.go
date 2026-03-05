package main

import "fmt"

// By default, Go passes arguments by value (makes a copy).
// If you want a function to change the original variable,
// you pass the address (&) and dereference (*) inside.

// In Go, you don't need pointers as often as in C++ because
// slices and maps are already reference types.

// Use pointers primarily when you need to modify a variable
// or avoid copying large structs
// Use pointers/references when working with large structs

// Value = Copy
// Reference = Not copy. Its the actual data in memory

// EXAMPLE 1

func example1() {
	// Basic Pointer Example
	y := 10
	p := &y
	*p = 5
	fmt.Println("Example1:", y)
}

// EXAMPLE 2
// Passing a reference or address of a pass-by-value type (int, string)

func addOne(n *int) {
	*n = *n + 1 // Modify the value at the address
}

func example2() {
	x := 10
	addOne(&x)                  // Pass the address
	fmt.Println("Example2:", x) // Output: 11 (Original x was modified)
}

// EXAMPLE 3: Receivers

// Value Receiver for read only operations on small structs?
// Pointer Receiver for mutating operations, state changes, delete, or large structs

// func (cache *Cache) Set(key string, value []byte) {}
// The *Cache in the receiver means Go is going to
// automatically pass the address and then automatically derefernce it.

type largeData struct {
	data [1000000]int
	name string
}

func (ld *largeData) GetData(size int) []int {
	return ld.data[:size]
}

func example3() {
	ld := largeData{name: "test"}
	data := ld.GetData(20)
	fmt.Println("Example3:", data)
}

// EXAMPLE 4: Nullable/Optional

// When Representing "Nullable" or Optional Values
// The Rule: Basic types in Go cannot be nil (they have zero values: 0, "", false).
// If you need to distinguish between "zero value" and "not set"/nil use a pointer.

type UserUpdate struct {
	Name *string `json:"name"`
	Age  *int    `json:"age"`
}

func example4() {
	age := 25
	// Only Age is provided; Name is nil (omitted)
	payload := UserUpdate{
		Age:  &age,
		Name: nil, // nil is allowed here because of the string pointer type in the struct
	}

	if payload.Name != nil { // We can check for nil here because of the string pointer type in the struct
		// Update name
	}
	if payload.Age != nil { // We can check for nil here because of the string pointer type in the struct
		fmt.Println("Example4:", *payload.Age) // Output: 25
	}
}

// EXAMPLE 5: Sharing and accessing state across Goroutines

func worker(id int, counter *int) {
	for i := 0; i < 100; i++ {
		(*counter)++ // All workers increment the SAME variable
	}
}

func example5() {
	count := 0
	// Pass pointer so all goroutines see the same 'count'
	go worker(1, &count)
	go worker(2, &count)

	// Note: In real code, you'd need a Mutex here to prevent race conditions!
}

func main() {
	example1()
	example2()
	example3()
	example4()
	example5()
}
