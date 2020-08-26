// pointers are used for sharing

// when program starts it creates Goroutine (like lightweight thread) with its own block of memory
// called stack

// the stack mem in Go starts at 2K and is very small

// every function is given a stack frame

// no value can be put in stack if its size is unknown ahead of time. If it is unknown at compile time
// then it has to be on heap

// stack are self-cleaning on every creation of function

// example of passing by value:

count := 10
// To get the address of a value, we use &.
println ( "count:\tValue Of[" , count, "]\tAddr Of[" , &count, "]" )
// Pass the "value of" count.
increment1(count)
// Printing out the result of count. Nothing has changed.
println ( "count:\tValue Of[" , count, "]\tAddr Of[" , &count, "]" )
// Pass the "address of" count. This is still considered pass by value,
// not by reference because the address itself is a value.
increment2(&count)
// Printing out the result of count. count is updated.
println ( "count:\tValue Of[" , count, "]\tAddr Of[" , &count, "]" )

func increment1 (inc int ) {
	// Increment the "value of" inc.
	inc++
	println ( "inc1:\tValue Of[" , inc, "]\tAddr Of[" , &inc, "]" )
	}

// increment2 declares count as a pointer variable whose value is always
// an address and points to values of type int. The * here is not an
// operator. It is part of the type name. Every type that is declared,
// whether you declare or it is predeclared, you get for free a pointer.
func increment2 (inc * int ) {
// Increment the "value of" count that the "pointer points to".
//The * is an operator. It tells us the value of the pointer points to.
    *inc++
    println ( "inc2:\tValue Of[" , inc, "]\tAddr Of[" , &inc, "]\tValue Points To[" , *inc, "]" )
}

// output
count: Value Of[ 10 ] Addr Of[ 0xc000050738 ]
inc1: Value Of[ 11 ] Addr Of[ 0xc000050730 ]
count: Value Of[ 10 ] Addr Of[ 0xc000050738 ]
inc2: Value Of[ 0xc000050738 ] Addr Of[ 0xc000050748 ] Value
Points To[ 11 ]