// user represents an user in the system.
type user struct {
	name string
	email string
}

func stayOnStack () user {
	// In the stack frame, create a value and initialize it.
	u := user{
	name: "Hoanh An" ,
	email: "hoanhan101@gmail.com" ,
}
	// Return the value, pass back up to the main stack frame.
	return u
}

func escapeToHeap () * user {
	u := user{
	name: "Hoanh An" ,
	email: "hoanhan101@gmail.com" ,
	}
	return &u
}

// The only difference between stayOnStack and escapeToHeap is that compiler understands that in 
// escapeToHeap we use pointer, so main will end up having a pointer to the heap, because stack is
// self cleaning and we are going to loose data on the function escape. On the heap it will stay.  