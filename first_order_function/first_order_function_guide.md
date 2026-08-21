# 🧠 Go Language: First-Order Function (ফার্স্ট-অর্ডার ফাংশন)

প্রোগ্রামিংয়ে **First-Order Function (ফার্স্ট-অর্ডার ফাংশন)** বলতে এমন সাধারণ ফাংশনকে বোঝায় যা কেবল সাধারণ ডাটা বা ভ্যালু নিয়ে কাজ করে। 

এই গাইডে ফার্স্ট-অর্ডার ফাংশনের সংজ্ঞা, বৈশিষ্ট্য, গো (Go) ভাষার উদাহরণ এবং মেমোরি এক্সিকিউশন নিয়ে বিস্তারিত আলোচনা করা হলো।

---

## 📌 ফার্স্ট-অর্ডার ফাংশন কী? (What is a First-Order Function?)

**First-Order Function** হলো এমন একটি স্ট্যান্ডার্ড ফাংশন যা:
1. ইনপুট হিসেবে কেবল **সাধারণ ডাটা/ভ্যালু** (যেমন: `int`, `string`, `struct`, `slice`, `pointer` ইত্যাদি) গ্রহণ করে।
2. কোনো **ফাংশনকে** প্যারামিটার বা আর্গুমেন্ট হিসেবে **গ্রহণ করে না**।
3. কোনো **ফাংশনকে** আউটপুট হিসেবে **রিটার্ন করে না**।

> 💡 **এক কথায়:** এটি এমন এক সাধারণ ফাংশন যা ডাটা গ্রহণ করে এবং ডাটা প্রসেস করে রেজাল্ট প্রদান করে, অন্য কোনো ফাংশনকে ইনপুট বা আউটপুট হিসেবে হ্যান্ডেল করে না।

---

## 🛠️ গো কোড উদাহরণ (Go Code Example)

[main.go](file:///c:/Users/User/Go/first_order_function/main.go) ফাইলে থাকা ফার্স্ট-অর্ডার ফাংশনের উদাহরণ:

```go
package main

import "fmt"

// ----------------------------------------------------
// 1. First-Order Function (সাধারণ যোগ করার ফাংশন)
// এটি ইনপুট হিসেবে দুটি int গ্রহণ করে এবং একটি int রিটার্ন করে।
// ----------------------------------------------------
func add(a int, b int) int {
	return a + b
}

// ----------------------------------------------------
// 2. First-Order Function (কাস্টম টাইপ/স্ট্রাক্ট গ্রহণকারী)
// ----------------------------------------------------
type Product struct {
	Name  string
	Price float64
}

func calculateTotalPrice(p Product, quantity int) float64 {
	return p.Price * float64(quantity)
}

func main() {
	// add() ফাংশন কল (First-Order Function)
	sum := add(15, 25)
	fmt.Println("Sum:", sum)

	// calculateTotalPrice() ফাংশন কল (First-Order Function)
	laptop := Product{Name: "Laptop", Price: 1200.00}
	total := calculateTotalPrice(laptop, 2)
	fmt.Printf("Total price for 2 laptops: $%.2f\n", total)
}
```

---

## 🔍 First-Order Function বনাম Higher-Order Function

| বৈশিষ্ট্য (Feature) | First-Order Function | Higher-Order Function |
|---|---|---|
| **প্যারামিটার (Parameter)** | কেবল সাধারণ ডাটা (int, string, struct) নেয়। | অন্য একটি ফাংশনকে প্যারামিটার হিসেবে নিতে পারে। |
| **রিটার্ন ভ্যালু (Return Value)** | কেবল সাধারণ মান বা ডাটা রিটার্ন করে। | অন্য কোনো ফাংশন রিটার্ন করতে পারে। |
| **জটিলতা (Complexity)** | সহজ, সোজা এবং প্রাতিষ্ঠানিক ফাংশন স্ট্রাকচার। | ডাইনামিক আচরণ (Callback / Closure)। |
| **উদাহরণ** | `func add(a int, b int) int` | `func calc(a int, b int, f func(int,int)int) int` |

---

## 🧠 RAM মেমোরি ও স্ট্যাক ফ্রেম এক্সিকিউশন (RAM Stack Frame)

যখন `add(15, 25)` কল করা হয়:

```text
[ Function Call: main() ]
       │
       │  Arguments: (15, 25) পাস করা হলো
       ▼
[ RAM Stack Frame: add() ]
┌────────────────────────────────────────┐
│ Parameter 'a' = 15                     │
│ Parameter 'b' = 25                     │
│ Return Value   = 40                    │
└────────────────────────────────────────┘
       │
       ▼ Stack Frame Pop/Cleared
```

1. **Step 1:** Call Stack-এ `add()` এর জন্য একটি নির্দিষ্ট ফ্রেম তৈরি হয়।
2. **Step 2:** `a = 15` এবং `b = 25` দিয়ে প্রসেস হয়ে ফলাফল `40` রিটার্ন হয়।
3. **Step 3:** মেমোরি থেকে `add()` ফাংশনের ফ্রেমটি ক্লিয়ার (pop) হয়ে যায়।

---

## 💡 বাস্তব জীবনের সহজ উদাহরণ (Real-World Analogy)

### পকেট ক্যালকুলেটর (Pocket Calculator Analogy)
* **First-Order Function:** একটি পকেট ক্যালকুলেটরের `+` বা `-` বাটন। আপনি দুটি সংখ্যা দেবেন (যেমন $15 + 25$), ক্যালকুলেটর আপনাকে একটি সংখ্যা ($40$) ফেরত দেবে। এটি কোনো নতুন বাটন বা প্রোগ্রাম তৈরি করে না বা নেয় না।

---

## 📁 তৈরি করা ফাইলসমূহ (Created Files):

1. **[first_order_function/main.go](file:///c:/Users/User/Go/first_order_function/main.go)** - First-Order Function-এর Go কোড।
2. **[first_order_function/first_order_function_guide.md](file:///c:/Users/User/Go/first_order_function/first_order_function_guide.md)** - বিস্তারিত গাইড।
3. **[first_order_function_guide.md](file:///c:/Users/User/Go/first_order_function_guide.md)** - Root গাইড।
