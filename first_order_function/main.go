package main

import "fmt"

// ====================================================
// First-Order Function (ফার্স্ট-অর্ডার ফাংশন)
// ====================================================
// এটি একটি ফার্স্ট-অর্ডার ফাংশন কারণ:
// ১. এটি ইনপুট প্যারামিটার হিসেবে কেবল সাধারণ মান বা ডাটা (int) নেয়।
// ২. কোনো ফাংশনকে আর্গুমেন্ট হিসেবে গ্রহণ করে না।
// ৩. কোনো ফাংশন রিটার্ন করে না।
// ====================================================

// উদাহরণ ১: সাধারণ যোগ করার ফার্স্ট-অর্ডার ফাংশন
func add(a int, b int) int {
	return a + b
}

// উদাহরণ ২: কাস্টম স্ট্রাক্ট গ্রহণকারী ফার্স্ট-অর্ডার ফাংশন
type Product struct {
	Name  string
	Price float64
}

func calculateTotalPrice(p Product, quantity int) float64 {
	return p.Price * float64(quantity)
}

// উদাহরণ ৩: এরর হ্যান্ডলিং সহ ফার্স্ট-অর্ডার ফাংশন
func divide(a float64, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

func main() {
	fmt.Println("=== First-Order Function Demonstration in Go ===")
	fmt.Println()

	// ১. add() ফার্স্ট-অর্ডার ফাংশন কল
	sum := add(15, 25)
	fmt.Printf("1. add(15, 25) = %d\n\n", sum)

	// ২. calculateTotalPrice() ফার্স্ট-অর্ডার ফাংশন কল
	item := Product{Name: "Laptop", Price: 1200.50}
	total := calculateTotalPrice(item, 2)
	fmt.Printf("2. Product: %s | Total Price (2 units): $%.2f\n\n", item.Name, total)

	// ৩. divide() ফার্স্ট-অর্ডার ফাংশন কল
	result, err := divide(100.0, 4.0)
	if err == nil {
		fmt.Printf("3. divide(100, 4) = %.2f\n", result)
	}
}
