package hashtable

import (
	"fmt"
	"testing"
)

func TestHash(t *testing.T) {
	h := hash("hello")
	fmt.Println(h)
}
