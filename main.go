package g0ng

import (
	"errors"
)

//Delimiter to denote the end of a string sequence
const (
	Seed      = rune('\u0002')
	Delimiter = rune('\u0003')
)

//inppend, inserts el into the location i
//and shifts all proceding elements over
func inppend(arr []*Tree, el *Tree, i int) []*Tree {

	if len(arr) == 0 {
		return []*Tree{el}
	}

	ret := append(arr, &Tree{})
	copy(ret[i+1:], ret[i:])
	ret[i] = el

	return ret
}

//Tree is a recurent struct which stores ngrams.
//A tree holds a value and pointers to other trees
//which branch from that tree.
//branches holds the value pointers and val holds the
//value
type Tree struct {
	branches []*Tree
	val      rune
}

//Traverse returns all whole words present in the tree
func (t *Tree) Traverse() [][]rune {
	if t.val == Delimiter {
		return nil
	}

	//return, list of words, each word is an array of runes
	ret := [][]rune{}
	//own value will be appended(tacked on) to all the
	//results from this node's branches
	tack := t.val

	//replace tack with invisible character if seed
	if tack == Seed {
		tack = 0
	}

	//range over branches and execute DeepToString
	for _, v := range t.branches {
		set := v.Traverse()

		if set == nil {
			word := []rune{tack}
			ret = append(ret, word)
		}

		for _, val := range set {
			//prepend
			word := append([]rune{tack}, val...)
			ret = append(ret, word)
		}

	}

	return ret
}

//String for use with fmt
func (t *Tree) String() string {
	if t == nil {
		return "<nil>"
	}

	slice := t.Traverse()
	var ret string

	for _, v := range slice {
		ret += string(v) + "\n"
	}

	//remove last newline
	return ret[:len(ret)-1]
}

//binarySearch
//Searches the tree's branches for the depicted rune
//using the binary search algorithm
//https://en.wikipedia.org/wiki/Binary_search_algorithm#Procedure
func (t *Tree) binarySearch(T rune) (int, *Tree) {
	if len(t.branches) == 0 {
		return 0, nil
	}

	n := len(t.branches)

	L, R := 0, n-1
	m := (L + R) / 2

	for L <= R {

		selItem := t.branches[m]

		//search complete
		if selItem.val == T {
			return m, selItem
		} else if selItem.val < T {
			L = m + 1
		} else if selItem.val > T {
			R = m - 1
		}

		m = (L + R) / 2
	}

	return L, nil
}

//Insert is used to insert a word into the ngram.
//The word delimiter is automatically added
//
//Note: the input is of type []rune
func (t *Tree) Insert(r []rune) {
	r = append(r, Delimiter)
	t.rInsert(r)
}

//recursive insert
func (t *Tree) rInsert(r []rune) {
	i, n := t.binarySearch(r[0])
	if n == nil {
		n = &Tree{
			branches: []*Tree{},
			val:      r[0]}
		t.branches = inppend(t.branches, n, i)
	}

	if len(r) <= 1 {
		return
	}

	n.rInsert(r[1:])
}

//Del
func (t *Tree) Del(r []rune) {
	//r = append(r, Delimiter)
	t.rDel(r)
}

//rDel
func (t *Tree) rDel(r []rune) bool {
	//if there are no more runes to check
	//then the algorithm has successfully
	//reached the node to be deleted
	if len(r) == 0 {
		return true
	}

	//find location and value of r[0] in
	//the branches of the current tree
	k, v := t.binarySearch(r[0])
	if v == nil {
		//if there is no matching branch
		//nothing will be deleted
		return false
	}

	//if a branch must be deleted
	//delete it
	if t.branches[k].rDel(r[1:]) {
		copy(t.branches[k:], t.branches[k+1:])
		t.branches = t.branches[:len(t.branches)-1]
	}

	return false
}

//Find is used to find a subset of the main ngram.
//It returns a tree and an error, the error means that
//no match exists.
//The returned Tree contains branches with the input string.
//This means that the input values will be present in the
//results of Traverse() and String()
//
//Note: the input is of type []rune
func (t *Tree) Find(r []rune) (*Tree, error) {
	if len(r) < 1 {
		return t, nil
	}

	ret := &Tree{
		val:      t.val,
		branches: []*Tree{}}

	_, v := t.binarySearch(r[0])
	if v == nil {
		return nil, errors.New("gong: rune sequence not found")
	}

	ap, err := v.Find(r[1:])
	if err != nil {
		return nil, err
	}

	ret.branches = append(ret.branches, ap)

	return ret, nil
}

//New returns a new tree root
func New() *Tree {
	return &Tree{
		val: Seed}
}
