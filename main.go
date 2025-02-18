Ось приклад базової обробки даних на Go. Цей код включає роботу з вхідними даними, їхню перевірку, сортування та вивід результатів:

```go
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sort"
)

// Create a struct to hold our data
type Person struct {
	Name string
	Age  int
}

// Implement the sort.Interface for our struct
type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

// Function to check and handle errors
func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Function to process our data
func processData(input []string) []Person {
	var people []Person

	for _, line := range input {
		split := strings.Split(line, ", ")

		name := split[0]
		age, err := strconv.Atoi(split[1])
		handleErr(err)

		// Create a new Person and append it to our slice
		people = append(people, Person{Name: name, Age: age})
	}

	return people
}

// Function to read data from a file
func readData(file string) []string {
	f, err := os.Open(file)
	handleErr(err)
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	handleErr(scanner.Err())

	return lines
}

// Function to print our data
func printData(people []Person) {
	for _, p := range people {
		fmt.Printf("%s is %d years old.\n", p.Name, p.Age)
	}
}

func main() {
	// Read data from a file
	lines := readData("data.txt")

	// Process the data
	people := processData(lines)

	// Sort the data by age
	sort.Sort(ByAge(people))

	// Print the data
	printData(people)
}
```

Цей код очікує файл `data.txt` у тому ж каталозі, де виконується програма. Кожен рядок файлу має бути у форматі `Name, Age`.