package main

import (
	"fmt"
	"github.com/gopherjs/vecty"
	. "github.com/kooksee/bulma"
)

func main() {
	fmt.Println(
		Box{
			Slot: CpsOf(&Buttons{
				Size: Size.IsMedium,
				Slot: CpsOf(&Button{
					Slot:  CpsOf(vecty.Text("Loading")),
					Size:  Size.IsMedium,
					State: State.IsLoading,
				}),
			}, &Buttons{
				Size: Size.IsLarge,
			}),
		},
	)
}
