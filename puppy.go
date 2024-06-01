package puppy

import (
	dog "github.com/praveenrvkl/dogs"
)

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof Woof Woof!!!"
}

func BigBark() string {
	return dog.WhenGrownUp(Bark())
}

func BigBarks() string {
	return dog.WhenGrownUp(Barks())
}

func BigBarks2() string {
	return dog.WhenGrownUp(Barks())
}
