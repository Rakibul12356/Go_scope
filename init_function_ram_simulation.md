# 🧠 Go `init()` Function & RAM Execution Simulation (মেমোরি ও রাম সিমুলেশন)

এই গাইডে আপনার আপডেট করা `initFunc.go` কোডটি কীভাবে রাম (RAM)-এ এক্সিকিউট হয়ে লাইন-বাই-লাইন আউটপুট প্রদান করেছে তার স্টেপ-বাই-স্টেপ সিমুলেশন দেওয়া হলো।

---

## 📄 বিশ্লেষণকৃত কোড:

```go
package main

import "fmt"

var a int = 10 // ১. গ্লোবাল ভ্যারিয়েবল ইনিশিয়ালাইজেশন

func main() {
	fmt.Println("I am main function") // ৫. main এর প্রথম লাইন
	fmt.Println(a)                   // ৬. a এর নতুন মান প্রিন্ট (২০)
}

func init() {
	fmt.Println("I am init function", a) // ৩. a এর প্রাথমিক মান সহ প্রিন্ট (১০)
	a = 20                              // ৪. a এর মান পরিবর্তন করে ২০ করা হলো
}
```

---

## 🖥️ টার্মিনাল আউটপুট (Output):
```text
I am init function 10
I am main function
20
```

---

## 🧠 RAM মেমোরি ও এক্সিকিউশন স্টেপস (Step-by-Step Simulation):

### 📍 Step 1: Data Segment-এ ভ্যারিয়েবল ডিক্লেয়ারেশন
* Go Runtime চালু হওয়ার পরই RAM-এর **Data Segment**-এ গ্লোবাল ভ্যারিয়েবল `a`-এর জন্য মেমোরি অ্যালকেট করে এবং মান সেট করে `10`।

```
RAM Data Segment:
+-------------------+
|  a = 10           |
+-------------------+
```

---

### 📍 Step 2: `init()` Call Stack-এ Push ও Execution
`main()` শুরু হওয়ার আগেই Go প্রথমে `init()` ফাংশন রান করায়।

1. **Line 13 (`fmt.Println("I am init function", a)`):**
   - RAM থেকে `a`-এর মান পড়া হয় (যা ছিল `10`)।
   - **আউটপুট প্রিন্ট হয়:** `I am init function 10`

2. **Line 14 (`a = 20`):**
   - RAM-এর Data Segment-এ `a`-এর মান ১০ থেকে বদলে **২০** হয়ে যায়।

```
RAM Data Segment (Updated):
+-------------------+
|  a = 20           |  <-- a এর মান ওভাররাইট হয়ে ২০ হলো
+-------------------+
```

3. `init()` শেষ হয় এবং RAM Call Stack থেকে পপ (ক্লিয়ার) হয়ে যায়।

---

### 📍 Step 3: `main()` Call Stack-এ Push ও Execution
`init()` সম্পূর্ণ শেষ হওয়ার পর GoRuntime `main()` ফাংশন চালু করে।

1. **Line 8 (`fmt.Println("I am main function")`):**
   - **আউটপুট প্রিন্ট হয়:** `I am main function`

2. **Line 9 (`fmt.Println(a)`):**
   - RAM থেকে `a`-এর বর্তমান মান পড়ে। (যেহেতু `init()` এর ভেতরে `a = 20` করা হয়েছিল, তাই এখন মান ২০)।
   - **আউটপুট প্রিন্ট হয়:** `20`

---

## ⏱️ এক্সিকিউশন টাইমলাইন সামারি (Timeline Summary)

```
[Time 0: Memory Alloc] ----> var a = 10 in RAM
[Time 1: init() Runs]  ----> Print "I am init function 10" -> Set a = 20 in RAM
[Time 2: init() Exits] ----> Stack Cleared
[Time 3: main() Runs]  ----> Print "I am main function" -> Print a (20)
```
