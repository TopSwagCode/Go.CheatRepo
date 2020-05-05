# Array

## Working with basic arrays
Create and print a Array: 

```Go
primes := [6]int{2, 3, 5, 7, 11, 13}
fmt.Println(primes)

// Result in 
[2 3 5 7 11 13]
```

A way to work with arrays is getting an item from it at index.

```Go
var number int = primes[0]
fmt.Println(number)

// Result in 
2
```

Similar you would be able to get a slice / range of items.

```Go
var slice []int = primes[1:4]
fmt.Println(slice)

// Results in Item at index 1,2,3
[3 5 7]
```

If you would like to update an item in the array

```Go
primes[2] = 1337
fmt.Println(primes)

// Result in 
[2 3 1337 7 11 13]
```

Updating the array would also update the slice. So if you printed  out the same slice again. You would get an updated version of it.

```Go
var slice []int = primes[1:4]
fmt.Println(slice)

// Result in Item at index 1,2,3
[3 1337 7]
```

If you ever needed a total number of elements or the max capacity of it you could use the functions ```len()``` and ```cap()```

```Go
fmt.Println("Length of primes:",len(primes))
fmt.Println("Cap of primes:",cap(primes))

// Result in
Length of primes: 6
Cap of primes: 6
```

## Working with slices
In Go, arrays have a fixed size. The size is even part of the definition of an array, so the two arrays [10]int and [20]int are not just two int arrays of different size but are in fact different types.

Slices add a dynamic layer on top of arrays. Creating a slice from an array neither allocates new memory nor copies anything. A slice is nothing but a “window” to some part of the array. This also explains why changing the array before updated the slice.

In Go we can use ```make()``` for creating slices. Below we can see how creating it empty with room for 5 elements.

```Go
slice := make([]int, 5)
fmt.Printf("len=%d cap=%d %v\n", len(slice), cap(slice), slice)

// Result in
len=5 cap=5 [0 0 0 0 0]
```

With a slice in place we can now start appending items to it. Slice has a nice feature if you try to append beyound capacity it will control all the magic of handling the underlaying array and exstend it for you. Below you can see the result of such actions.

```Go
// Create slice with no elements in it
slice := make([]int, 0, 5)
fmt.Printf("len=%d cap=%d %v\n", len(slice), cap(slice), slice)

// For loop
for i := 0; i < 10; i++ {
    slice = append(slice, i*i)
    fmt.Printf("len=%d cap=%d %v\n", len(slice), cap(slice), slice)
}

// Result in
len=0 cap=5 []
len=1 cap=5 [0]
len=2 cap=5 [0 1]
len=3 cap=5 [0 1 4]
len=4 cap=5 [0 1 4 9]
len=5 cap=5 [0 1 4 9 16]
len=6 cap=10 [0 1 4 9 16 25]
len=7 cap=10 [0 1 4 9 16 25 36]
len=8 cap=10 [0 1 4 9 16 25 36 49]
len=9 cap=10 [0 1 4 9 16 25 36 49 64]
len=10 cap=10 [0 1 4 9 16 25 36 49 64 81]
```

## Going through all the elements in a array

Besides the For loop, there are 3 other ways running through an array / slice.

Getting index and value:
```Go
for i,v := range slice {
    fmt.Printf("slice**%d = %d\n", i, v)
}

// Result in
slice**0 = 0
slice**1 = 1
slice**2 = 4
slice**3 = 9
slice**4 = 16
slice**5 = 25
slice**6 = 36
slice**7 = 49
slice**8 = 64
slice**9 = 81

```

getting index only:
```Go
for i := range slice {
    fmt.Println("slice**index: ", i)
}

// Result in
slice**index:  0
slice**index:  1
slice**index:  2
slice**index:  3
slice**index:  4
slice**index:  5
slice**index:  6
slice**index:  7
slice**index:  8
slice**index:  9
```
 
getting value only:
```Go
for _, v := range slice {
    fmt.Println("slice**Value: ", v)
}

// Result in
slice**Value:  0
slice**Value:  1
slice**Value:  4
slice**Value:  9
slice**Value:  16
slice**Value:  25
slice**Value:  36
slice**Value:  49
slice**Value:  64
slice**Value:  81

```
Ofcourse you can combine it with only getting a subset of the values like so:
```Go
for _, v := range slice[1:5] {
    fmt.Println("slice**Value: ", v)
}

// Result in
slice**Value:  1
slice**Value:  4
slice**Value:  9
slice**Value:  16
```