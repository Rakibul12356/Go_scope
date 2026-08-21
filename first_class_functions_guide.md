# 🧠 Go Language: First-Class & First-Order Functions (ফার্স্ট-ক্লাস ও ফার্স্ট-অর্ডার ফাংশন)

Go প্রোগ্রামিং ল্যাঙ্গুয়েজে ফাংশন একটি অত্যন্ত শক্তিশালী কনসেপ্ট। গো-তে ফাংশন হলো **First-Class Citizen (ফার্স্ট-ক্লাস সিটিজেন)**। 

এই গাইডে **First-Class Function**, **First-Order Function**, **Higher-Order Function** এবং **Anonymous Function (Closure)** নিয়ে বিস্তারিত আলোচনা করা হলো।

---

## 📌 সংক্ষেপে ৩টি মূল ধারণার পার্থক্য (Quick Comparison)

| ধারণার নাম (Concept) | বিবরণ (Description) | উদাহরণ (Example) |
|---|---|---|
| **First-Class Functions** | গো ল্যাঙ্গুয়েজের একটি বৈশিষ্ট্য। ফাংশনকে সাধারণ ভ্যারিয়েবলের মতো মান অ্যাসাইন, আর্গুমেন্ট পাস এবং রিটার্ন করা যায়। | `fn := func() { fmt.Println("Hi") }` |
| **First-Order Function** | যে সাধারণ ফাংশন অন্য কোনো ফাংশনকে আর্গুমেন্ট হিসেবে নেয় না এবং অন্য ফাংশন রিটার্নও করে না। | `func add(a int, b int) int` |
| **Higher-Order Function** | যে ফাংশন অন্য কোনো ফাংশনকে আর্গুমেন্ট হিসেবে গ্রহণ করে **অথবা** অন্য ফাংশন রিটার্ন করে। | `func calc(a, b int, op func(int,int)int)` |

---

## 🛠️ গো কোড উদাহরণ (Go Code Example)

[main.go](file:///c:/Users/User/Go/first_class_functions/main.go) ফাইলে নিচের কোডটি সংরক্ষিত আছে:

```go
package main

import "fmt"

// ১. কাস্টম ফাংশন টাইপ (Custom Function Type Definition)
type Operation func(int, int) int

// ২. First-Order Function (সাধারণ ফাংশন)
func simpleAdd(a int, b int) int {
	return a + b
}

// ৩. Higher-Order Function (ফাংশনকে Argument হিসেবে গ্রহণ করা)
func calculate(a int, b int, op Operation) int {
	return op(a, b) // পাস হওয়া ফাংশনটিকে এক্সিকিউট করা হচ্ছে
}

// ৪. Higher-Order Function (ফাংশন রিটার্ন করা / Function Factory)
func createMultiplier(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}

func main() {
	// A. Variable-এ Anonymous Function রাখা
	greeting := func(name string) string {
		return "Hello, " + name
	}
	fmt.Println(greeting("Karim"))

	// B. Higher-Order Function-এ আর্গুমেন্ট হিসেবে পাস করা
	res1 := calculate(10, 5, simpleAdd)
	fmt.Println("Calculate with simpleAdd:", res1)

	// C. Function Factory (ফাংশন রিটার্ন থেকে তৈরি)
	triple := createMultiplier(3)
	fmt.Println("Triple of 4:", triple(4)) // 12
}
```

---

## 🔍 ১. First-Class Functions (ফার্স্ট-ক্লাস ফাংশন) কী?

কোনো ল্যাঙ্গুয়েজে ফাংশন যদি সাধারণ ভ্যারিয়েবল (যেমন `int`, `string` বা `struct`)-এর মতো সমান সুবিধা পায়, তবে তাকে **First-Class Functions** বলে।

Go-তে ফার্স্ট-ক্লাস ফাংশন ব্যবহারের ৪টি প্রধান বৈশিষ্ট্য:

1. **ভ্যারিয়েবলে ফাংশন রাখা (Assigning to Variables):**
   ```go
   square := func(n int) int {
       return n * n
   }
   fmt.Println(square(4)) // ১৬
   ```

2. **ফাংশনকে Parameter/Argument হিসেবে পাস করা (Passing as Arguments):**
   ```go
   func process(x int, f func(int) int) {
       fmt.Println("Processed:", f(x))
   }
   ```

3. **ফাংশন থেকে অন্য ফাংশন রিটার্ন করা (Returning Functions):**
   ```go
   func adder() func(int) int {
       return func(x int) int { return x + 1 }
   }
   ```

4. **ডাটা স্ট্রাকচারে ফাংশন রাখা (Storing in Maps/Slices):**
   ```go
   operations := map[string]func(int, int) int {
       "add": func(a, b int) int { return a + b },
       "sub": func(a, b int) int { return a - b },
   }
   ```

---

## 🔍 ২. First-Order Function (ফার্স্ট-অর্ডার ফাংশন) কী?

**First-Order Function** হলো এমন একটি ফাংশন যা:
- কোনো ফাংশনকে ইনপুট আর্গুমেন্ট হিসেবে গ্রহণ করে না।
- কোনো ফাংশনকে আউটপুট রিটার্ন হিসেবে প্রদান করে না।
- এটি শুধুমাত্র সাধারণ ডাটা (যেমন `int`, `string`, `bool`, `struct` ইত্যাদি) নিয়ে কাজ করে।

### উদাহরণ:
```go
// এটি একটি First-Order Function
func multiply(x int, y int) int {
    return x * y
}
```

---

## 🔍 ৩. Higher-Order Function (হায়ার-অর্ডার ফাংশন) কী?

**Higher-Order Function** হলো ফার্স্ট-অর্ডার ফাংশনের বিপরীত। এটি:
- অন্তত একটি ফাংশনকে আর্গুমেন্ট হিসেবে গ্রহণ করে, **অথবা**
- একটি ফাংশনকে রিটার্ন ভ্যালু হিসেবে প্রদান করে।

### উদাহরণ (Argument হিসেবে ফাংশন গ্রহণ):
```go
func executeTwice(f func()) {
    f()
    f()
}
```

---

## 🧠 RAM মেমোরি ও ক্লোজার এক্সিকিউশন (Closure & Memory Simulation)

যখন কোনো অ্যানোনিমাস ফাংশন তার বাইরের স্কোপের ভ্যারিয়েবল ক্যাপচার করে, তখন তাকে **Closure** বলা হয়।

```text
createMultiplier(factor = 3)
       │
       ▼
┌──────────────────────────────────────────────┐
│ Heap / Closure Memory Space                  │
│  Captured Variable: factor = 3               │
│  Returned Function Signature: func(n int) int│
└──────────────────────────────────────────────┘
       │
       ▼
 triple(4) ───> 4 * (Captured factor = 3) ───> 12
```

1. **Step 1:** `createMultiplier(3)` রান হলে RAM-এ `factor = 3` মেমোরিতে থেকে যায়।
2. **Step 2:** রিটার্ন হওয়া অ্যানোনিমাস ফাংশনটি `factor`-এর মান মনে রাখে (Enclose করে)।
3. **Step 3:** পরবর্তীতে `triple(4)` কল করলে এটি বাইরের মেমোরির `factor` (3) ব্যবহার করে $4 \times 3 = 12$ রিটার্ন করে।

---

## 📁 তৈরি করা ফাইলসমূহ (Created Files):

1. **[first_class_functions/main.go](file:///c:/Users/User/Go/first_class_functions/main.go)** - সম্পূর্ণ Go কোড উদাহরণ।
2. **[first_class_functions/first_class_functions_guide.md](file:///c:/Users/User/Go/first_class_functions/first_class_functions_guide.md)** - বিস্তারিত গাইড।
3. **[first_class_functions_guide.md](file:///c:/Users/User/Go/first_class_functions_guide.md)** - Root নির্দেশিকা।
