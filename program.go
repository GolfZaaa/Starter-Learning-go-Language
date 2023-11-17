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

// func main() {
// 	name := "Avirut Chaisongkram"
// 	age := 25
// 	score := 95.5
// 	pass := true
// 	fmt.Printf("My Name is %v\n", name)
// 	fmt.Printf("Type name = %T\n", name)
// 	fmt.Printf("Age is %v\n", age)
// 	fmt.Printf(" Type Age = %T\n", age)
// 	fmt.Printf(" Score is %v\n", score)
// 	fmt.Printf(" Type Score = %T\n", score)
// 	fmt.Printf(" Pass is = %v\n", pass)
// 	fmt.Printf(" Type pass = %T\n", pass)
// }

// ตัวดำเนินการทางคณิตศาสตร์
// func main() {
// 	number1, number2 := 10, 3
// 	fmt.Println("ผลบวก = ", number1+number2)
// 	fmt.Println("ผลลบ = ", number1-number2)
// 	fmt.Println("ผลคูณ = ", number1*number2)
// 	fmt.Println("ผลหาร = ", number1/number2)
// 	fmt.Println("เศษ = ", number1%number2)
// }

// ตัวดำเนินการเปรี่ยบเทียบ
func main() {
	number1, number2 := 10, 3

	fmt.Println("เท่ากันหรือไม่ = ", number1 == number2)
	fmt.Println("ไม่เท่ากันหรือไม่ =", number1 != number2)
	fmt.Println(number1, " มีค่ามากกว่า ", number2, "=", number1 > number2)
	fmt.Println(number1, " มีค่าน้อยกว่า ", number2, "=", number1 < number2)
	fmt.Println(number1, " มีค่ามากกว่าหรือเท่ากับ ", number2, "=", number1 >= number2)
	fmt.Println(number1, " มีค่าน้อยกว่าหรือเท่ากับ ", number2, "=", number1 <= number2)

}
