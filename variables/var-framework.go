package main

import (
	"fmt"
	"os"
)

func main() {
	name := os.Getenv("USER")
	course := "Go Fundamentals"

	fmt.Println("Hi", name, "your current course is", course)
	updateCourse(&course)
	fmt.Println("Your current course is", course)
}

func updateCourse(course *string) string {
	*course = "Some new course"
	fmt.Println("Updated course to", *course)
	return *course
}
