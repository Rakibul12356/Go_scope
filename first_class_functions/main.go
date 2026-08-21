package main

import "fmt"

// ----------------------------------------------------
// ১. কাস্টম ফাংশন টাইপ (Custom Function Type)
// Go-তে ফাংশন টাইপকে সংজ্ঞায়িত করা যায়
// ----------------------------------------------------
type Operation func(int, int) int

// ----------------------------------------------------
// ২. First-Order Function (ফার্স্ট-অর্ডার ফাংশন)
// এটি একটি সাধারণ ফাংশন যা অন্য কোন ফাংশনকে আর্গুমেন্ট হিসেবে নেয় না
// এবং কোন ফাংশন রিটার্নও করে না।
// ----------------------------------------------------
func simpleAdd(a int, b int) int {
	return a + b
}

// ----------------------------------------------------
// ৩. Higher-Order Function (হায়ার-অর্ডার ফাংশন) - Case A:
// ফাংশনকে Argument হিসেবে গ্রহণ করা
// ----------------------------------------------------
func calculate(a int, b int, op Operation) int {
	return op(a, b) // পাস হওয়া ফাংশনটিকে এক্সিকিউট করা হচ্ছে
}

// ----------------------------------------------------
// ৪. Higher-Order Function - Case B:
// ফাংশন রিটার্ন করা (Function Factory / Closure)
// ----------------------------------------------------
func createMultiplier(factor int) func(int) int {
	// একটি Anonymous Function রিটার্ন করা হচ্ছে
	return func(number int) int {
		return number * factor
	}
}

func main() {
	fmt.Println("=== Go First-Class & First-Order Functions Demo ===")
	fmt.Println()

	// ====================================================
	// ক্যানভাস ১: Function Assigned to Variable (ভ্যারিয়েবলে ফাংশন রাখা)
	// ====================================================
	fmt.Println("--- 1. Assigning Anonymous Function to Variable ---")
	
	// বেনামী (Anonymous) ফাংশন ভ্যারিয়েবলে অ্যাসাইন করা
	greeting := func(name string) string {
		return "Hello, " + name + "! Welcome to Go."
	}
	
	// ভ্যারিয়েবলটি ব্যবহার করে ফাংশন কল করা
	fmt.Println(greeting("Karim"))
	fmt.Println()

	// ====================================================
	// ক্যানভাস ২: Passing Function as Argument (Callback)
	// ====================================================
	fmt.Println("--- 2. Passing Function as Argument (Higher-Order Function) ---")
	
	// বিয়োগ করার জন্য একটি অ্যানোনিমাস ফাংশন তৈরি
	subtract := func(x int, y int) int {
		return x - y
	}

	// calculate ফাংশনে সাধারণ ফাংশন এবং অ্যানোনিমাস ফাংশন পাস করা
	res1 := calculate(10, 5, simpleAdd) // simpleAdd পাস করা হলো
	res2 := calculate(10, 5, subtract)  // subtract ভ্যারিয়েবল পাস করা হলো
	res3 := calculate(10, 5, func(x int, y int) int { // ইনলাইন ফাংশন পাস
		return x * y
	})

	fmt.Printf("Addition result (simpleAdd): %d\n", res1)
	fmt.Printf("Subtraction result (subtract): %d\n", res2)
	fmt.Printf("Inline Multiplication result: %d\n\n", res3)

	// ====================================================
	// ক্যানভাস ৩: Returning Function from Function (Closure / Factory)
	// ====================================================
	fmt.Println("--- 3. Returning Function from Function (Function Factory) ---")
	
	// ৩ গুণ ও ৫ গুণ করার জন্য ফাংশন জেনারেট করা
	triple := createMultiplier(3)
	quintuple := createMultiplier(5)

	fmt.Printf("Triple of 4: %d\n", triple(4))         // 4 * 3 = 12
	fmt.Printf("Quintuple of 4: %d\n", quintuple(4))   // 4 * 5 = 20
	fmt.Println()

	// ====================================================
	// ক্যানভাস ৪: Immediately Invoked Function Expression (IIFE)
	// ====================================================
	fmt.Println("--- 4. Immediately Invoked Function (IIFE) ---")
	
	// তৈরি হওয়ার সাথে সাথেই এক্সিকিউট হওয়া ফাংশন
	func(msg string) {
		fmt.Println("IIFE Executed Immediately with message:", msg)
	}("Go is awesome!")
}
