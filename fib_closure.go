package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
        prev := 0
            next := 1
            return func() int {
                        tmp := prev
                                prev = next
                                        next = tmp + next
                                                return tmp
                                                    
            }

}

func main() {
        f := fibonacci()
        for i := 0; i < 10; i++ {
                    fmt.Println(f())
                        
        }

}

