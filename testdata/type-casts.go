package fixtures

import "fmt"

func main() {
	type A struct {}
	type B struct {}

	type someInterfaceA interface {}
	type someInterfaceB interface {}

	var a someInterfaceA = &A{}
	var b someInterfaceB = &B{}

	if _, ok := a.(*A); ok {
		fmt.Println("SUCCESS")

		bCasted := b.(*B) // MATCH /type assert can panic/
		fmt.Printf("%v", bCasted)
	}

	switch a.(type) {
	case someInterfaceA:
		if _, ok := a.(*A); ok {
			fmt.Println("SUCCESS")

			bCasted := b.(*B) // MATCH /type assert can panic/
			fmt.Printf("%v", bCasted)
		}
	default:
		bCasted := b.(*B) // MATCH /type assert can panic/
		fmt.Printf("%v", bCasted)

		switch t := b.(type) {
		case someInterfaceB:
			aCasted := a.(*A) // MATCH /type assert can panic/
			fmt.Printf("%v", aCasted)

			if _, ok := a.(*A); ok {
				fmt.Println("SUCCESS")

				bCasted = b.(*B) // MATCH /type assert can panic/
				fmt.Printf("%v", t)
			}
		}
	}
}
