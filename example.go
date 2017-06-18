package g0ng

import (
//"github.com/ear7h/g0ng"
)

func main() {
	//r := g0ng.New() //
	r := New()

	r.Insert([]rune("hello world"))

	//the Insert() method only accepts type []rune
	println(r.String()) //prints "hello world"

	//once again this method accepts type []rune
	sub, err := r.Find([]rune("hello"))
	if err != nil {
		//not called
	}

	//sub is of type *Tree
	println(sub.String()) //prints "hello world"

	//err != nil when no elements are found

	sub, err = r.Find([]rune("hell0"))
	if err != nil {
		//called
	}

	//the Traverse() method returns a [][]rune
	//which essentially represents a list of words
	//
	//note: the first element of each word is
	//the Seed character
	els := r.Traverse()

	println(string(els[0]))    //prints hello world
	println(string(els[0][0])) //prints
	println(string(els[0][1])) //prints h

	//once again this method accepts type []rune
	//del deletes the nodes specified by the value
	//AND nodes without dependents
	r.Insert([]rune("hello"))

	r.Del([]rune("hello "))

	println(r.String()) //prints "hello"

}
