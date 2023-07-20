Exercises
sum is a function that takes a slice of numbers and adds them together. What would its function signature look like in Go?

Write a function that takes an integer and halves it and returns true if it was even or false if it was odd. For example, half(1) should return (0, false) and half(2) should return (1, true).

Write a function with one variadic parameter that finds the greatest number in a list of numbers.

Using makeEvenGenerator as an example, write a makeOddGenerator function that generates odd numbers.

The Fibonacci sequence is defined as: fib(0) = 0, fib(1) = 1, fib(n) = fib(n-1) + fib(n-2). Write a recursive function that can find fib(n).

What are defer, panic, and recover? How do you recover from a runtime panic?

How do you get the memory address of a variable?

How do you assign a value to a pointer?

How do you create a new pointer?

What is the value of x after running this program:

func square(x *float64) {
    *x = *x * *x
}
func main() {
    x := 1.5
    square(&x)
}
Write a program that can swap two integers (x := 1; y := 2; swap(&x, &y) should give you x=2 and y=1).