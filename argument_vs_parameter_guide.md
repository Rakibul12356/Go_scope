# 🧠 Go Language: Parameter vs Argument (প্যারামিটার বনাম আর্গুমেন্ট)

এই গাইডে **Parameter (প্যারামিটার)** এবং **Argument (আর্গুমেন্ট)**-এর মধ্যে পার্থক্য, উদাহরণ, রাম (RAM) মেমোরি সিমুলেশন এবং বাস্তব জীবনের অ্যানালজি সহ বিস্তারিত আলোচনা করা হলো।

---

## 📌 সংক্ষেপে মূল পার্থক্য (Quick Summary)

* **Parameter (প্যারামিটার):** ফাংশন **ডিক্লেয়ারেশন বা ডেফিনেশন (Definition)** করার সময় ব্র্যাকেটের `()` ভেতরে যে ভ্যারিয়েবল লেখা হয়। এটি একটি প্লেসহোল্ডার (Placeholder)।
* **Argument (আর্গুমেন্ট):** ফাংশন **কল (Call/Invocation)** করার সময় ব্র্যাকেটের `()` ভেতরে যে আসল মান (Actual Value) বা ডেটা পাঠানো হয়।

---

## 🛠️ গো কোড উদাহরণ (Go Code Example)

[main.go](file:///c:/Users/User/Go/argument_vs_parameter/main.go) ফাইলে নিচের কোডটি রাখা আছে:

```go
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
```

---

## 🔍 বিস্তারিত তুলনা (Detailed Comparison)

| বিষয় (Feature) | Parameter (প্যারামিটার) | Argument (আর্গুমেন্ট) |
|---|---|---|
| **সংজ্ঞা (Definition)** | ফাংশন সংজ্ঞায়িত (Declare) করার সময় যে ভ্যারিয়েবল ডিক্লেয়ার করা হয়। | ফাংশন চালনা (Call) করার সময় যে বাস্তব ডেটা/ভ্যালু পাস করা হয়। |
| **পর্যায় (Stage)** | Function Definition / Signature Stage | Function Invocation / Call Stage |
| **প্রকৃতি (Nature)** | এটি এক ধরনের প্লেসহোল্ডার ভ্যারিয়েবল (Variable Placeholder)। | এটি আসল মান (Literal Value, Variable Value, or Expression)। |
| **মেমোরি (Memory Scope)** | ফাংশনের নিজস্ব লোকাল স্কোপে (Stack Frame) তৈরি হয়। | কলার (Caller) ফাংশন থেকে প্যারামিটারে ডেটা হিসেবে কপি বা রেফারেন্স হয়। |
| **পরিবর্তনশীলতা** | এটি স্থির ডিক্লেয়ারেশন (Static Signature)। | বিভিন্ন ফাংশন কলে বিভিন্ন মান পাঠানো যেতে পারে (Dynamic Values)। |

---

## 🧠 RAM মেমোরি ও এক্সিকিউশন সিমুলেশন (RAM & Stack Frame Simulation)

যখন কোডে `addNumbers(10, 20)` রান হয়:

```text
[ Function Call: main() ]
       │
       │  Arguments: (10, 20) পাস করা হচ্ছে
       ▼
[ RAM Call Stack Frame: addNumbers() ]
┌────────────────────────────────────────┐
│ Parameter 'a' = 10 (Argument থেকে পাওয়া)│
│ Parameter 'b' = 20 (Argument থেকে পাওয়া)│
│ Local Var 'sum' = 30                   │
└────────────────────────────────────────┘
```

1. **Step 1:** Go Runtime মেমোরির Stack-এ `addNumbers()` ফাংশনের জন্য একটি নতুন Call Stack Frame বরাদ্দ করে।
2. **Step 2:** `main()` থেকে পাঠানো Arguments (`10`, `20`) প্যারামিটার `a` এবং `b`-তে কপি হয় (`a = 10`, `b = 20`)।
3. **Step 3:** `addNumbers()` ফাংশন তার লোকাল স্কোপ থেকে `a + b` হিসাব সম্পন্ন করে রিটার্ন করে এবং Stack Frame খালি (pop) হয়ে যায়।

---

## 💡 বাস্তব জীবনের সহজ উদাহরণ (Real-World Analogies)

### ১. জুস ব্লেন্ডার মেশিন (Juice Blender Analogy)
* **Parameter:** ব্লেন্ডার মেশিনের ইনপুট স্লট (যেমন: `Fruit Slot`)। এটি ডিফাইন করে যে মেশিনে ফল দিতে হবে।
* **Argument:** আপনি সত্যি সত্যি যে ফলটি স্লটে দিলেন (যেমন: `আম` বা `আপেল`)।

### ২. ফর্মে খালি ঘর (Form Input Field Analogy)
* **Parameter:** রেজিস্ট্রেশন ফর্মে লেখা ঘর: `First Name: [ _________ ]` -> `First Name` হলো প্যারামিটার।
* **Argument:** আপনি ফর্ম পূরণ করার সময় লিখলেন `Rahim` -> `Rahim` হলো আর্গুমেন্ট (প্রকৃত ডেটা)।

---

## 📁 তৈরি করা ফাইলসমূহ (Created Files):

1. **[argument_vs_parameter/main.go](file:///c:/Users/User/Go/argument_vs_parameter/main.go)** - Go কোড ইমপ্লিমেন্টেশন।
2. **[argument_vs_parameter/argument_vs_parameter_guide.md](file:///c:/Users/User/Go/argument_vs_parameter/argument_vs_parameter_guide.md)** - বিস্তারিত গাইড।
