package animal

import (
	"fmt"
)

/*
class Animal
   virtual abstract Speak() string
*/
type Animal interface {
	Speak() string
}

// private struct
type cat struct {
	Name string // public
	Age  int    // public
}

func (c cat) Meow() {
	println("Meow !!")
}

// NewCat public factory method
func NewCat(name string, age int) cat {
	cat := cat{name, age}
	return cat
}

func (c cat) String() string {
	return fmt.Sprintf("Cat Name: %s \nCat Age: %d", c.Name, c.Age)
}

func (c cat) Speak() string {
	return "Meow!"
}

// Dog Public struct
type Dog struct {
	name string
	age  int
}

// Bark public method of class Dog
func (dog Dog) Bark() {
	println("Bark !")
}

// SetName Public Setter (dog must be Pointer!)
func (dog *Dog) SetName(name string) {
	dog.name = name
}

// SetAge Public Setter (dog must be Pointer!)
func (dog *Dog) SetAge(age int) {
	dog.age = age
}

// NewDog Public factory method "Constractur"
func NewDog(name string, age int) Dog {
	dog := Dog{name, age}
	return dog
}

func (dog Dog) String() string {
	return fmt.Sprintf("Dog Name: %s \nDog Age: %d", dog.name, dog.age)
}

func (d Dog) Speak() string {
	return "Woof!"
}

/*
class Llama
  method Speak() string //non-virtual
     return "LaLLamaQueLLama!"
*/
type Llama struct {
}

func (l Llama) Speak() string {
	return "LaLLamaQueLLama!"
}
