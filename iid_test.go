package imagid

import (
	"fmt"
)

func ExampleIID1() {
	var v uint64 = 149495437762496513 // 14vzpk09yxoh
	id := IID(v)
	fmt.Printf("%s", id)
	// Output:
	// hoxy90kpzv41
}

func ExampleIID2() {
	var v uint64 = 149497847983638530 // 14vzpk09yxoh
	id := IID(v)
	fmt.Printf("%s", id)
	// Output:
	// q6pex8bk0w41
}
