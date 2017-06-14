package a1

import "testing"

func TestLinearSearch(t *testing.T) {

	//Shouldn't use slice of interface{} for an slice of generic type
	// **Simply use interface{} and use reflect.Valueof() to get the underlying slice.**

	// a := []int{4, 2, -1, 5, 0}
	// case1 := make([]interface{}, len(a))
	// for i := range a {
	// 	case1[i] = a[i]
	// }

	// b := []string{"cat", "nose", "egg"}
	// case2 := make([]interface{}, len(b))
	// for i := range b {
	// 	case2[i] = b[i]
	// }

	cases := []struct {
		x    interface{}
		lst  interface{}
		want int
	}{
		{x: 5, lst: []int{4, 2, -1, 5, 0}, want: 3},
		{x: 3, lst: []int{4, 2, -1, 5, 0}, want: -1},
		{x: "egg", lst: []string{"cat", "nose", "egg"}, want: 2},
		{x: "up", lst: []string{"cat", "nose", "egg"}, want: -1},
	}

	for _, c := range cases {
		got := linearSearch(c.x, c.lst)
		if got != c.want {
			t.Errorf("linearSearch(%v, %v) got %v, should get %v", c.x, c.lst, got, c.want)
		}
	}

}
