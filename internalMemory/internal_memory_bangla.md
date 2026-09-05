# Go-তে মেমোরি সিমুলেশন ও Internal Memory

এই নোটে `memory.go` প্রোগ্রামটি চালানোর সময় Go প্রোগ্রামের বিভিন্ন value কোথায় থাকে এবং কোন ক্রমে কাজ করে তা সহজভাবে দেখানো হয়েছে।

## ১. Internal Memory কী?

প্রোগ্রাম চালানোর সময় অপারেটিং সিস্টেম ও Go runtime মিলে process-এর জন্য মেমোরির একটি address space তৈরি করে। শেখার সুবিধার জন্য এই মেমোরিকে কয়েকটি অংশে ভাগ করে ভাবা যায়:

- **Code segment**: কম্পাইল হওয়া instruction বা machine code থাকে।
- **Data বা Global segment**: package-level variable এবং স্থির data থাকে।
- **Stack**: function call-এর সময় তৈরি হওয়া local variable, parameter এবং return information-এর জন্য ব্যবহৃত হয়।
- **Heap**: runtime-এ প্রয়োজন অনুযায়ী তৈরি হওয়া data এখানে থাকতে পারে।

এটি একটি conceptual model। Go compiler এবং garbage collector implementation অনুযায়ী কোনো value stack বা heap-এ রাখবে তা নির্ধারণ করে। তাই source code দেখে কোনো variable-এর নির্দিষ্ট memory address ধরে নেওয়া উচিত নয়।

## ২. `memory.go`-এর গুরুত্বপূর্ণ অংশ

```go
const a = 10
var p = 100
```

- `a` হলো package-level constant। এর value পরিবর্তন করা যায় না।
- `p` হলো package-level variable। এটি package scope-এ আছে এবং `call()` থেকে ব্যবহার করা যায়।
- দুটিই `main()`-এর local variable নয়। শেখার মডেলে এগুলোকে global/data অংশে ভাবা যায়।

```go
func call() {
    add := func(x int, y int) int {
        z := x + y
        fmt.Println("The sum is:", z)
        return z
    }

    add(5, 10)
    add(a, p)
}
```

- `call()` চালু হলে তার জন্য একটি function call frame তৈরি হয়।
- `x`, `y` এবং `z` function-এর local value।
- `add` একটি anonymous function-এর reference। এটি `call()`-এর ভিতরে তৈরি হয়েছে।
- `add(a, p)`-এ package-level `a` ও `p` ব্যবহার করা হয়েছে।

```go
func main() {
    a := 20
    call()
    fmt.Println(a)
}
```

এখানে `a := 20` package-level `const a = 10`-কে পরিবর্তন করে না। এটি `main()`-এর আলাদা local variable। একই নাম হলেও scope আলাদা:

- `call()`-এর বাইরে থাকা `a` হলো package-level constant, value `10`।
- `main()`-এর ভিতরের `a` হলো local variable, value `20`।

## ৩. Program execution order

Go প্রোগ্রাম চালু হলে সাধারণত এই ক্রমটি দেখা যায়:

```text
1. package-level declaration প্রস্তুত হয়
2. package-level variable initialize হয়
3. init() function চলে
4. main() function চলে
5. main() থেকে call() চলে
6. call() শেষ হলে main() আবার পরের statement চালায়
```

এই উদাহরণে সম্ভাব্য output:

```text
init function called
The sum is: 15
The sum is: 110
20
```

`init()` আগে চলে, কারণ `main()` শুরু হওয়ার আগে package initialization সম্পন্ন হয়। এরপর `main()`-এর local `a`-এর value `20` print হয়।

## ৪. Stack frame-এর সহজ simulation

`main()` শুরু হলে একটি frame কল্পনা করা যায়:

```text
main frame
└── a = 20
```

`main()` যখন `call()` চালায়, তখন নতুন frame তৈরি হয়:

```text
call frame
├── add = anonymous function
├── x, y = প্রথম call-এর parameter
└── z = x + y
```

`add(5, 10)` শেষ হলে সেই call-এর frame আর প্রয়োজন থাকে না। পরে `add(a, p)`-এর জন্য আবার parameter ও `z`-এর নতুন values ব্যবহৃত হয়।

বাস্তব Go runtime এই বিষয়গুলো optimize করতে পারে। তাই এই diagram শেখার সুবিধার জন্য, debugger-এর exact memory map নয়।

## ৫. Scope এবং shadowing

একই নাম ভিন্ন scope-এ থাকলে inner scope-এর declaration outer scope-কে shadow করতে পারে। উদাহরণ:

```go
const a = 10

func main() {
    a := 20
    fmt.Println(a) // 20
}
```

`main()`-এর ভিতরে `a` খোঁজার সময় Go আগে local scope দেখে। তাই সেখানে `20` পাওয়া যায়। `main()`-এর বাইরে `a` ব্যবহার করলে package-level `const a = 10` পাওয়া যাবে।

## ৬. Function এবং closure-এর সম্পর্ক

`add` হলো একটি anonymous function। এটি `call()`-এর ভিতরে তৈরি হলেও এই উদাহরণে বাইরের কোনো local variable capture করছে না। কোনো anonymous function যদি বাইরের variable ব্যবহার করে, তখন সেটিকে closure বলা হয়। যেমন:

```go
func makeAdder() func(int) int {
    base := 10

    return func(value int) int {
        return base + value
    }
}
```

এখানে returned function-এর প্রয়োজন শেষ না হওয়া পর্যন্ত `base` value-টি runtime-কে ধরে রাখতে হতে পারে। এই ধরনের value stack থেকে heap-এ চলে যেতে পারে। একে **escape to heap** বলা হয়। সিদ্ধান্তটি compiler নেয়।

## ৭. সংক্ষেপে মনে রাখার বিষয়

- `init()` আগে, `main()` পরে চলে।
- package-level declaration-এর scope পুরো package জুড়ে।
- local variable function বা block-এর scope-এ থাকে।
- একই নামের local variable outer variable-কে shadow করতে পারে।
- function call-এর জন্য stack frame-এর ধারণা ব্যবহার করা যায়।
- Go-তে stack বনাম heap compiler নির্ধারণ করে; programmer সরাসরি তা নিয়ন্ত্রণ করে না।
- garbage collector heap-এর আর প্রয়োজন নেই এমন object পরিষ্কার করতে সাহায্য করে।

## ৮. উদাহরণ চালানো

`internalMemory` directory থেকে চালান:

```bash
go run memory.go
```

আরও পরিষ্কারভাবে execution দেখতে চাইলে:

```bash
gofmt -w memory.go
go run memory.go
```
