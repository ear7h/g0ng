g0ng
====

A Golang N-Gram package

# Install

In your command line run:

```sh
go get github.com/ear7h/gong
```

# Quick reference API 

```go
r := g0ng.New()                   //new ngram
r.Insert([]rune("a string"))      //insert a string
sub, err := r.Find([]rune("a s")) //get subtree where all items match the input
if err == nil {
	//called when subtree not found
}
arr2d := r.Traverse() //returns 2d slice, can be thought of as word list
r.Del([]rune("a s"))  //deletes matching subtree and non-Deliminator trees with not dependents
```


# Usage

### Import: 

```go
import (
	"github.com/ear7h/g0ng"
)
```

### Create a root for the n-gram tree:

```go
r := g0ng.New()
```

### Adding strings to root:

```go
//the Insert() method only accepts type []rune
r.Insert([]rune("hello world"))
```

### Printing elements of the tree:

```go
println(r.String())

//fmt automatically calls the String() method
fmt.Print(r)
```

### Searching for elements from a tree:

```go
//once again this method accepts type []rune
sub, err := r.Find([]rune("hello"))
if err != nil {
	//not called
}

//sub is of type *Tree
println(sub.String()) //will print "hello world"

//err != nil when no elements are found

sub, err = r.Find([]rune("hell0"))
if err != nil {
	//called
}
```

### Retrieving all elements from a tree:

```go
//the Traverse() method returns a [][]rune
//which essentially represents a list of words
//
//note: the first element of each word is
//the Seed character
els := r.Traverse()

println(string(els[0])) //prints hello world
println(string(els[0][0])) //prints 
println(string(els[0][1])) //prints h


```

### Deleting elements from a tree:

```go
//once again this method accepts type []rune
//del deletes the nodes specified by the value
//AND nodes without dependents
r.Insert([]rune("hello"))

r.Del([]rune("hello "))

println(r.String()) //prints "hello"
```


