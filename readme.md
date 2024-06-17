Capital letter methods means public.

Go compiler will give errors if the variables defined are not used.

there are default values which are there if the value is not assigned.

IF we dont define the type of the variable, then lexer will implicitly define the type of the variable.

No var style in Go language:
We use walrus assignment symbol
// no var
	numberofProjects := 30000000
	fmt.Println(numberofProjects)


WE CANNOT USE WALRUS ASSIGNMENT OUTSIDE FUNCTIONS

any variable starting with a Caps letter is a public variable


comma ok syntax

Time is what we give with the same numbers given below:
	fmt.Println(presentTime.Format("    "))
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

Memory allocation and Deallocation happens automatically
Two methods we use:
new() and make()

new()--> Allocate memory but no initialize
we will get a memory address
Zeroed storage

make()--> Allocate memory and initialize
we will get a memory address
non zeroed storage


Garbage collection(GC)
GC happens automatically
Anything out of scope or nil becomes part of GC, but the cleaning happens only when a certain number of parameters are met.

func NumCPU() int
returns the number of logical CPUs usable by the current process.
Based on this, we can think of 


Pointers
Pointers are defined in order to get get the exct reference or the address where the value is saved  in the memory.
In some functions, a reference to the variable is passes sometimes. However, when we use pointers, we can pass the exact value which is present at a given address in the memoory.

Pointers point to the value in the memory. 
Whenever we use the *pointer_variable, then we get the value of the variable to which it is pointing

Moral of the story: Operation is performed on the original value.



Array
In array, mentioing size is a MUST
even though you put less value here, the length and space in memory will be occupied as much as it is defined when the variable is initialized.

Slices


