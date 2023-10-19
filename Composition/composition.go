package main

import "fmt"

type author struct {
	firstname string
	lastname  string
	bio       string
}

type blogpost struct {
	title   string
	content string
	author  // struct reusage here
}

type website struct {
	blogposts []blogpost // embedding slice with structs
}

func (a *author) name() string {
	return fmt.Sprintf("%s %s", a.firstname, a.lastname) // everything gel
}

func (b *blogpost) details() {
	fmt.Println("Title:", b.title)
	fmt.Println(
		"Name:",
		b.name(),
	) // from our borrrowed author we got to use our method name as a whole
	fmt.Println(
		"Bio:",
		b.bio,
	) // we can use these because we already embedded out struct already
	fmt.Println("Content:", b.content) // norms
}

func (w *website) contents() {
	fmt.Println("Contents of the website")
	for _, v := range w.blogposts { // range into our different blogpost cause it is a slice
		v.details() // everything all together
		fmt.Println()
	}
}

func main() {
	author := author{"Dude", "Excel", "Your random script kiddie"}

	blogpost1 := blogpost{
		"Let talk Go",
		"Go is a language that do not supports inheritance but it supports somthing like these called composition in struct for code reuse",
		author, // remember right here our struct must be complete
	}
	blogpost2 := blogpost{
		"Go module",
		"Go module is the one that help up to keep tracks of all the packages we import in our program and it help us with the version installed also, We have go work also which i am going to explain latter in the work",
		author,
	}
	w := website{blogposts: []blogpost{blogpost1, blogpost2}} // and then our slice embedded struct
	w.contents()
}
