package main

import "fmt"

// นิยามตัวแปร
// func main() {
// 	name := "Avirut Chaisongkram"
// 	age := 25
// 	score := 95.5
// 	pass := true
// 	fmt.Println("My Name is", name)
// 	fmt.Println("Age is", age)
// 	fmt.Println("Score is", score)
// 	fmt.Println("Pass Excemple is", pass)
// }

// นิยามค่าคงที่
// func main() {
// 	const name string = "Avirut Chai"
// 	fmt.Println("My Name is", name)
// }

func main() {
	name := "Avirut Chaisongkram"
	age := 25
	score := 95.5
	pass := true
	fmt.Printf("My Name is %v\n", name)
	fmt.Printf("Type name = %T\n", name)
	fmt.Printf("Age is %v\n", age)
	fmt.Printf(" Type Age = %T\n", age)
	fmt.Printf(" Score is %v\n", score)
	fmt.Printf(" Type Score = %T\n", score)
	fmt.Printf(" Pass is = %v\n", pass)
	fmt.Printf(" Type pass = %T\n", pass)

}
