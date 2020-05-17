package algo

import (
	"fmt"
	"testing"

	"github.com/gogf/gf/util/gutil"
)

func TestBinarySearch(t *testing.T) {
	l := []interface{}{
		1, 2, 3, 4, 5,
	}
	fmt.Println(BinarySearch(l, 1, gutil.ComparatorInt))
}
