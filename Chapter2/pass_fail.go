// pass_fail reports whether a grade is passing or failing.
package main

import (
	// "bufio"
	"fmt"
	"log"

	"github.com/HeadFirstGo/Chapter4/keyboard"
	// "os"
	// "strconv"
	// "strings"
)

func main() {
	fmt.Print("Enter a grade: ")
	// reader := bufio.NewReader(os.Stdin)
	// Option 1: blank identifier _
	//input, _ := reader.ReadString('\n')
	//Option 2: Catch the error
	// input, err := reader.ReadString('\n')
	grade, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}

	// input = strings.TrimSpace(input)
	// grade, err := strconv.ParseFloat(input, 64)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	var status string
	if grade >= 60 {
		status = "passing"
	} else {
		status = "failing"
	}

	fmt.Println("A grade of", grade, "is", status)
}
