package main

import "log"

// a struct can have its own function

type someStruct struct {
	firstName string
}

func main() {
	var1 := someStruct{
		firstName: "Hello",
	}
	log.Println("Actual value : ", var1.firstName)
	var1.returnName()
	log.Println("Value after temporary change: ", var1.firstName)
	var1.returnChangedName()
	log.Println("Permanently Changed value : ", var1.firstName)
}

// function associated with someStruct and can access its variables though it wa snot passed explicitly
// it modifies as it wors on the pointer
func (s *someStruct) returnChangedName() string {
	s.firstName = "Bye"
	return s.firstName
}

// function associated with someStruct and can access its variables though it wa snot passed explicitly
// it doesn modifiy, it only copies
func (s someStruct) returnName() string {
	s.firstName = "Bye"
	log.Println("Temporarily changed copy : ", s.firstName)
	return s.firstName
}

// func (s someStruct) returnName() string {
// 	return s.firstName
// }

/*OUTPUT
2026/04/05 19:11:29 Hello
*/

/*EXPLANTION
Both will work for **reading** a value, but they behave differently in important ways.

## The Difference: Pointer vs Value Receiver

| | Pointer Receiver `*someStruct` | Value Receiver `someStruct` |
|---|---|---|
| **Modifies original?** | ✅ Yes | ❌ No (works on a copy) |
| **Nil-safe?** | ❌ Can panic if nil | ✅ Safer |
| **Memory** | More efficient for large structs | Copies the whole struct |

---

### Value Receiver — gets a **copy**
```go
func (s someStruct) returnName() string {
    return s.firstName // reading from a COPY
}
```
Any changes to `s` inside this method **won't affect** the original struct.

---

### Pointer Receiver — gets the **original**
```go
func (s *someStruct) returnName() string {
    return s.firstName // reading from the ORIGINAL
}
```
Changes to `s` here **will affect** the original struct.

---

## Why you get the same result

Since you're only **reading** `firstName` and not modifying anything, both produce the same output. The difference only becomes visible when you **mutate** the struct:

```go
func (s someStruct) setName() {
    s.firstName = "John" // ❌ only changes the copy, original unchanged
}

func (s *someStruct) setName() {
    s.firstName = "John" // ✅ changes the original
}
```

---

## Which one to use?

- Use **pointer receiver `*someStruct`** if:
  - You need to modify the struct
  - The struct is large (avoids copying)
  - You want **consistency** (if any method uses a pointer receiver, use it for all)

- Use **value receiver `someStruct`** if:
  - The struct is small and you never mutate it
  - You intentionally want to work on a copy

> **General rule of thumb:** Default to pointer receivers. It's the most common pattern in Go and avoids subtle bugs.
*/
