# Go LeetCode Framework

A Go-based framework to make solving LeetCode problems easier.

Adapted from [Tom on the Internet](https://www.youtube.com/@tomontheinternet) — check out his channel for great developer content and tutorials.

---

## 🚀 Features

- ⚙️ Auto-generated test files using [`gotests`](https://github.com/cweill/gotests)
- 🔁 Auto-reloading test runner using [`gow`](https://github.com/mitranim/gow)
- 🧪 Write, test, and iterate quickly on Go solutions to coding problems
- 🗂 Clean structure, minimal setup

---

## 🚧 Getting Started

First, install the required tools:

```bash
go install github.com/cweill/gotests/...@latest
go install github.com/mitranim/gow@latest
```

Then:

```bash
# 1. Clone the repo
git clone https://github.com/alex-r89/go-leetcode-framework.git
cd go-leetcode-framework

# 2. Create a new folder for the problem
mkdir 1-two-sum
cd 1-two-sum

# 3. Generate a test file (also creates the solution file if missing)
gotests -all -w -parallel .

# 4. Open the test file and:
#    - Use your editor to expand the generated test struct (VSCode: cmd + .)
#    - Fill in example inputs from LeetCode into `args`
#    - Fill in expected output into `want`
#    - Repeat for as many test cases as needed

# 5. Run the tests using gow
gow test 1-two-sum/1-two-sum.go 1-two-sum/1-two-sum_test.go
```
Now you can iteratively fill in your function logic and instantly see if the test cases pass.

---

## 📁 Structure

Each problem lives in its own folder, containing the solution file and its corresponding `_test.go` file.

```
.
├── 1-two-sum/
│   ├── 1-two-sum.go
│   └── 1-two-sum_test.go
├── 2-longest-substring/
│   ├── 2-longest-substring.go
│   └── 2-longest-substring_test.go
...
```

---

## 🧪 Example

See [`1-two-sum/1-two-sum.go`](1-two-sum/1-two-sum.go) and [`1-two-sum/1-two-sum_test.go`](1-two-sum/1-two-sum_test.go) for a basic example of how a solution and its tests are structured.


---

## 📣 Credits

This project was inspired by [Tom on the Internet](https://www.youtube.com/@tomontheinternet).

Tools used:

- [gotests](https://github.com/cweill/gotests)
- [gow](https://github.com/mitranim/gow)
