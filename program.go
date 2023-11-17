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
// func main() {
// 	number1, number2 := 10, 3
// 	fmt.Println("เท่ากันหรือไม่ = ", number1 == number2)
// 	fmt.Println("ไม่เท่ากันหรือไม่ =", number1 != number2)
// 	fmt.Println(number1, " มีค่ามากกว่า ", number2, "=", number1 > number2)
// 	fmt.Println(number1, " มีค่าน้อยกว่า ", number2, "=", number1 < number2)
// 	fmt.Println(number1, " มีค่ามากกว่าหรือเท่ากับ ", number2, "=", number1 >= number2)
// 	fmt.Println(number1, " มีค่าน้อยกว่าหรือเท่ากับ ", number2, "=", number1 <= number2)
// }

// scanf การป้อนค่า
// %s คือการป้อนค่าแบบตัวอักษร string
// %d คือการป้อนค่าแบบจำนวนเต็ม int
// %f คือการป้อนค่าแบบจำนวนมีเศษ float
// func main() {
// // var name string
// // fmt.Print("กรุณาป้อนชื่อนักเรียน = ")
// // fmt.Scanf("%s", &name)

// // fmt.Println("สวัสดี", name)

// // var score int
// // fmt.Print("กรุณาป้อนคะแนนนักเรียน = ")
// // fmt.Scanf("%d", &score)
// // fmt.Println("คะแนนสอบ + (จิตพิสัย)", score+10)

// 	var score float64
// 	fmt.Print("กรุณาป้อนคะแนนนักเรียน = ")
// 	fmt.Scanf("%f", &score)
// 	fmt.Println("คะแนนสอบ + (จิตพิสัย)", score+10)
// }

// if...else if...else
// 100 => >50 สอบผ่าน <50 ไม่ผ่าน
// func main () {
// 	var score int
// 	fmt.Print("กรุณาป้อนคะแนนนักเรียน = ")
// 	fmt.Scanf("%d", &score)
// 	fmt.Println("คะแนนสอบ + (จิตพิสัย)", score)

// 	//ประมวลผล
// 	if score >= 50 {
// 		fmt.Println("สอบผ่าน")
// 	}else{
// 		fmt.Println("สอบไม่ผ่าน")
// 	}
// }

// func main() {
// 	var number int
// 	fmt.Print("กรุณาป้อนตัวเลข = ")
// 	fmt.Scanf("%d", &number)

// 	if number%2 == 0 {
// 		fmt.Println("เลขคู่")
// 	} else {
// 		fmt.Println("เลขคี้")
// 	}
// }

//ลองทำโจทย์ปัญหา
// func main() {
// 	var number int
// 	fmt.Print("กรุณาป้อนหมายเลข 1กับ2 \n 1.เปิดบัญชีธนาคาร \n 2.ถอน-ฝากเงิน \n ถ้าเป็นตัวเลขอื่นจะ ข้อมูลไม่ถูกต้อง \n")
// 	fmt.Print("กรุณาป้อนหมายเลข = ")
// 	fmt.Scanf("%d", &number)
// 	if number == 1 {
// 		fmt.Print("เปิดบัญชีธนาคาร")
// 	} else if number == 2 {
// 		fmt.Print("ถอน-ฝากเงิน")
// 	} else {
// 		fmt.Print("ข้อมูลไม่ถูกต้อง")
// 	}
// }

// Switch...Case
func main() {
	var number int
	fmt.Print("กรุณาป้อนหมายเลข 1กับ2 \n 1.เปิดบัญชีธนาคาร \n 2.ถอน-ฝากเงิน \n ถ้าเป็นตัวเลขอื่นจะ ข้อมูลไม่ถูกต้อง \n")
	fmt.Print("กรุณาป้อนหมายเลข = ")
	fmt.Scanf("%d", &number)

	switch number {
	case 1:
		fmt.Println("เปิดบัญชีธนาคาร")
	case 2:
		fmt.Println("ถอน-ฝากเงิน")
	default:
		fmt.Println("ข้อมูลไม่ถูกต้อง")
	}

}
