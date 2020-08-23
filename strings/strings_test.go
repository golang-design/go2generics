// Code generated by go2go; DO NOT EDIT.


//line strings_test.go2:1
package strings

//line strings_test.go2:1
import (
//line strings_test.go2:1
 "fmt"
//line strings_test.go2:1
 "reflect"
//line strings_test.go2:1
 "testing"
//line strings_test.go2:1
)

//line strings_test.go2:9
type Int int

func (i Int) String() string {
	return fmt.Sprintf("%d", i)
}

func TestStringifiable(t *testing.T) {
	ii := []Int{1, 2, 3, 4, 5}
	ss := instantiate୦୦Stringify୦strings୮aInt(ii)
	if !reflect.DeepEqual(ss, []string{"1", "2", "3", "4", "5"}) {
		t.Errorf("Stringify wrong, got: %v", ss)
	}
}
//line strings.go2:10
func instantiate୦୦Stringify୦strings୮aInt(s []Int,) []string {
	ret := []string{}
	for _, v := range s {
		ret = append(ret, v.String())
	}
	return ret
}

//line strings.go2:16
var _ = fmt.Errorf
//line strings.go2:16
var _ = reflect.Append
//line strings.go2:16
var _ = testing.AllocsPerRun