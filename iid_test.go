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
	var v uint64 = 149497847983638530 // 14w0kb8xep6q
	id := IID(v)
	fmt.Printf("%s", id)
	// Output:
	// aq6pex8bk0w41
}
