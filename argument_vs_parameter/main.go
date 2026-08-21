package main

import "fmt"

// ----------------------------------------------------
// 1. Function Definition (ফাংশন ডেফিনেশন)
// এখানে 'a' এবং 'b' হলো PARAMETERS (প্যারামিটার)
// এরা হলো প্লেসহোল্ডার ভ্যারিয়েবল যা ডেটা গ্রহণ করতে প্রস্তুত থাকে।
// ----------------------------------------------------
func addNumbers(a int, b int) int { // <--- 'a' এবং 'b' হলো PARAMETERS
	sum := a + b
	return sum
}

// আরও একটি উদাহরণ:
// 'name' এবং 'age' হলো PARAMETERS
func printUserDetails(name string, age int) { // <--- 'name' এবং 'age' হলো PARAMETERS
	fmt.Printf("User Name: %s, Age: %d\n", name, age)
}

func main() {
	fmt.Println("=== Parameter vs Argument Demonstration in Go ===")

	// ----------------------------------------------------
	// 2. Function Call / Invocation (ফাংশন কল)
	// এখানে '10' এবং '20' হলো ARGUMENTS (আর্গুমেন্ট)
	// ফাংশন কল করার সময় যে আসল ভ্যালু বা ডেটা পাঠানো হয়।
	// ----------------------------------------------------
	result := addNumbers(10, 20) // <--- '10' এবং '20' হলো ARGUMENTS
	fmt.Printf("Result of addNumbers(10, 20): %d\n\n", result)

	// ভ্যারিয়েবল বানিয়ে সেই ভ্যারিয়েবলের মান পাঠানো:
	x := 50
	y := 75
	// এখানে 'x' এবং 'y'-এর ভেতরে থাকা মান (50 এবং 75) হলো ARGUMENTS
	total := addNumbers(x, y) // <--- x, y (যাদের মান 50, 75) হলো ARGUMENTS
	fmt.Printf("Result of addNumbers(%d, %d): %d\n\n", x, y, total)

	userName := "Rahim"
	userAge := 25
	// এখানে 'userName' ("Rahim") এবং 'userAge' (25) হলো ARGUMENTS
	printUserDetails(userName, userAge) // <--- ARGUMENTS
}
