# 🎈 Go Package Scope: The Toy Room Rules! 🧸

Imagine your Go code is a giant house full of **Rooms**. In Go, we call these rooms **Packages**!

Here are the 5 easy rules of how toys (variables, constants, and functions) work inside and outside of rooms!

---

## 🚪 What is a Package?
Every Go file starts by naming its room at the top:
```go
package main
```
All files that say `package main` live in the **same room**!

---

## 📜 Rule 1: No Quick Shortcuts `:=` Out in the Room!
Inside a function (like a pencil case), you can use the quick `:=` shortcut. 
But out in the open room (package scope), you **MUST** use full magic words (`var`, `const`, or `func`).

❌ **Wrong (The room gets confused!):**
```go
package mathlib

Money := 200 // 💥 ERROR! No := outside of functions!
```

✅ **Right (Clear declaration):**
```go
package mathlib

var Money = 200 // 👍 Good! The room knows what Money is!
```

---

## 🔠 Rule 2: Capital Letters are Magic Passports! ✈️
How do other rooms see your toys? Look at the **FIRST LETTER** of the name!

### 1. BIG Capital Letter = Superpower (Exported)
If the name starts with a **Capital Letter**, it gets a passport to travel to other packages!

```go
package mathlib

var Money = 200         // ✈️ Other packages can visit and see Money!
func Sum(a, b int) int  // ✈️ Other packages can use Sum!
```

### 2. small lowercase Letter = Room Secret (Unexported)
If the name starts with a **lowercase letter**, it must stay inside its own room!

```go
package mathlib

var secretToy = "Robot" // 🔒 Only code inside mathlib can see this!
```

---

## 🤝 Rule 3: Everyone in the Same Room Shares Everything!
If two files are in the same folder and both say `package main`:
* [main.go](file:///c:/Users/User/Go/main.go)
* [add.go](file:///c:/Users/User/Go/add.go)

They can play with each other's toys automatically—even lowercase ones! 

Example: `main.go` can call `add()` directly because `add.go` is in the same room (`package main`).

> 💡 **Tip:** Always run them together using:
> ```bash
> go run .
> ```

---

## 🚫 Rule 4: No Two Toys Can Have the Same Name!
You can't have two variables or functions with the exact same name at package level in the same room.

❌ **Playground Argument (Error):**
* `main.go` -> `var Score = 10`
* `add.go` -> `var Score = 50` 💥 *(Error: Score redeclared!)*

---

## 🏃 Rule 5: Action Statements Must Stay Inside Functions!
You can declare variables in the room, but **actions** (like printing or calculating) MUST happen inside a function `func`.

❌ **Wrong:**
```go
package main

var Age = 7
fmt.Println(Age) // 💥 ERROR! Actions can't sit outside functions!
```

✅ **Right:**
```go
package main

import "fmt"

var Age = 7

func main() {
    fmt.Println(Age) // 👍 Actions happen inside func main!
}
```

---

## 🎓 Summary Cheat Sheet

| Feature | Inside Function `func` | Outside Function (Package Scope) |
| :--- | :--- | :--- |
| **`:=` Short Declaration** | ✅ Allowed | ❌ Forbidden! Use `var` |
| **Capital Name (`Money`)** | Local to function | ✈️ Shared with ALL packages |
| **lowercase Name (`secret`)** | Local to function | 🏠 Shared ONLY within same package |
| **Actions (`fmt.Println`)** | ✅ Allowed | ❌ Forbidden |
