## Overall Idea
1. Go has a single workspace for all the Go projects on the machine, which is totally different from almost all other languages. $HOME/go for example. Go tools will build based on the source codes at $HOME/go/src and generate corresponding command executable or package files in $HOME/go/bin or $HOME/go/pkg. 

2. Go is invented for building huge system like servers. It make the build process easier.

## Control flow

1. C's while loop is spelled for loop in Go, i.e. only the middle part of the three expressions remains. 


3. In Go's for and if statement, () is not needed but {} is always required and cannot be omitted.

4. 	Switch without a condition is the same as switch true. This construct can be a clean way to write long if-then-else chains.

5. **deferred function**. They are pushed into a stack.

## Methods

1. Go does not have classes. However, we can define methods on types.

2. A method is just a function with a receiver argument. **Note: for normal
   receiver, the method gets a copy of the receiver (i.e. pass by value), just as those for ordinary
   functions**.

3. You can only declare a method with a receiver whose type is defined in the
   same package as the method. You cannot declare a method with a receiver whose
   type is defined in another package (which includes the built-in types such as
   int).

4. If we need to modify the value of receiver itself (actually we do this in
   most cases), we must delcare methods with pointer receivers.

5. functions with a pointer argument must take a pointer. However, methods with
   pointer receivers take either a value or a pointer as the receiver when they
   are called. Go can automatically treat v.Scale(5) as (&v).Scale(5) since the
   Scale method has a pointer receiver. For function/methods with normal value
   parameter/receiver, the way Go treats them is equivalent to the above
   situation.

6. Reason for using pointer receiver: modifying value and space efficiency.

## Interface

1. An interface type if defined as a set of method signatures. All types
   implementing those methods can be held by an object of that interface.

2. Interfaces are implemented implicitly, a type implements an interface by
   implementing all the methods declared in the interface.

3. Under the hood, interface contains a tuple of a value and the concrete type.
   We can use the function describe to peek the content of interface.
    
    '''go
    func describe(i I) {
        fmt.Printf("(%v, %T)\n", i, i) // %v for value, %T for type
    }
    
    '''
4. The empty inferface, interface{}, is rather important in Go. The value of
   empty inferface can hold value of any other types since every type implement
   at least 0 method.

5. Stringer interface is a type that can describe itself with a string. And it's
   built-in in the package fmt.

6. Type can be checked with Go's reflect package.

## Errors

1. error type is a built-in interface similar to fmt.Stringer.

2. Functions often return an error value, and calling code should handle errors
   by testing whether the error equals nil.

3. See ./error.go for example. 


## Arrays and Slices

## Maps

