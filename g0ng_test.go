package g0ng

import (
	"fmt"
	"testing"
)

func TestTree(test *testing.T) {
	t := &Tree{}

	for i := 0; i < 900; i++ {
		s := fmt.Sprintf("%03d", i)
		t.Insert([]rune(s))
	}

	t.String()

	stree, err := t.Find([]rune("91"))
	if err != nil {
		s := "error caught"
		s = s + "!"
	}

	n := stree.String()

	if n != "<nil>" {
		test.Fatal()
	}

	t.Traverse()
	t.String()
	t.Del([]rune("91"))
	t.Del([]rune("80"))
}
