package main

import (
	"fmt"
)

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
// func main() {
// 	var number int
// 	fmt.Print("กรุณาป้อนหมายเลข 1กับ2 \n 1.เปิดบัญชีธนาคาร \n 2.ถอน-ฝากเงิน \n ถ้าเป็นตัวเลขอื่นจะ ข้อมูลไม่ถูกต้อง \n")
// 	fmt.Print("กรุณาป้อนหมายเลข = ")
// 	fmt.Scanf("%d", &number)
// 	switch number {
// 	case 1:
// 		fmt.Println("เปิดบัญชีธนาคาร")
// 	case 2:
// 		fmt.Println("ถอน-ฝากเงิน")
// 	default:
// 		fmt.Println("ข้อมูลไม่ถูกต้อง")
// 	}
// }

// Array
// func main() {
// 	// number1 := 100
// 	// number2 := 200
// 	// number3 := 300

// 	//สร้าง Array แบบกำหนดจำนวน
// 	// var numbers [3]int = [3]int{100, 200, 300} // 0 = 100 ,1 = 200 ,2 = 300
// 	// show := len(numbers)
// 	// fmt.Println("Show Array =", show)

// 	//สร้าง Array แบบไม่กำหนดจำนวน
// 	numbers := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
// 	show := len(numbers)
// 	fmt.Println("Show Array ... =", show)
// }

// Slice
// func main() {
// 	numbers := []int{100, 200, 300} //มีค่าเริ่มต้น
// 	numbers = append(numbers, 400)
// 	numbers = append(numbers, 500)
// 	fmt.Println(numbers)
// 	fmt.Println(numbers[1:])       // คือเอาตั้งแต่ตำแหน่งที่ 1 ของ Array ถึง สุดท้าย 1-4
// 	fmt.Println("EE", numbers[:3]) // คือเอาตั้งแต่ตำแหน่งที่ 0 - 3
// }

// map
// func main() {
// 	// แบบรวดรัด
// 	// country := map[string]string{"TH": "Thailand", "JP": "Japan"}

// 	// แบบไม่รวดรัด
// 	country := map[string]string{}
// 	country["TH"] = "ไทย"
// 	country["JP"] = "ญี่ปุ่น"

// 	value, check := country["JP"]

// 	if check {
// 		fmt.Println(value)
// 	} else {
// 		fmt.Print("ไม่พบข้อมูล")
// 	}
// 	// fmt.Println(country["JP"])
// }

// for...loop
// func main() {
// 	// for i := 1; i <= 3; i++ {
// 	// 	fmt.Println("ดีครับ", i)
// 	// }

// 	for i := 10; i >= 0; i-- {
// 		fmt.Println(i)
// 	}
// }

//Break & Continue

// func main() {
// 	for i := 1; i <= 10; i++ {
// 		if i == 5 {
// 			continue
// 		}
// 		fmt.Println(i)
// 	}
// 	fmt.Println("จบโปรแกรม")
// }

func main() {
	// number := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	// for i := 0; i < len(number); i++ {
	// 	fmt.Println("Index = ", i, "Value = ", number[i])
	// }

	// for index, value := range number {
	// 	fmt.Println("Index = ", index, "Value = ", value)
	// }

	//กรณีไม่เอา index
	// for _, value := range number {
	// 	fmt.Println("Value = ", value)
	// }

	language := map[string]string{"TH": "Thailand", "EN": "English","JP":"Japanese"}
	// for key, value := range language {
	// 	fmt.Println("Key = ", key, "Value = ", value)
	// }

	//กรณีไม่เอา Key ให้ใส่ _
	for _, value := range language {
		fmt.Println("Value = ", value)
	}

}
