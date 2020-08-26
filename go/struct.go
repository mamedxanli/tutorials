// example
type example struct {
	flag    bool
	counter int16
	pi      float32
}

// initialize the struct
var e1 example

ftm.Printf("%+v\n", e1)
// output: {flag:false counter:0 pi:0}

// initialize with non default values
e2 := example{
	flag:    true,
	counter: 10,
	pi:      3.141592,
}
fmt.Println( "Flag" , e2.flag)
fmt.Println( "Counter" , e2.counter)
fmt.Println( "Pi" , e2.pi)
// output
// Counter 10
// Pi 3.141592
// Flag true


// it is possible to declare a variable of anonymous type
e3 := struct {
	flag bool
	counter int16
	pi float32
}{
	flag: true ,
	counter: 10 ,
	pi: 3.141592 ,
	}
fmt.Println( "Flag" , e3.flag)
fmt.Println( "Counter" , e3.counter)
fmt.Println( "Pi" , e3.pi)
// output
// Counter 10
// Pi 3.141592
// Flag true

// two name type identical struct cannot be be assigned to each other. But if one is anonymous 
// like e3 above, then it is possible to assign.
var e4 example
e4 = e3
fmt.Printf("%+v\n", e4)

// output: 
{flag:true counter:10 pi:3.141592}

