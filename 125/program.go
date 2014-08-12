
package main

import "fmt"
import "log"
import "bufio"
import "os"
import "strconv"

func main() {
    file, err := os.Open(os.Args[1])
    if err != nil {
        log.Fatal(err)
    }   
    defer file.Close()
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        //'scanner.Text()' represents the test case, do something with it
        atPosition, _ := strconv.Atoi(scanner.Text())
        numberOfHalves := 0
	    for i := atPosition; i >= 1; i = fuzzyHalfer(i) {
		    numberOfHalves++
	    }
	    
	     fmt.Println(determineNumber(numberOfHalves))
       
    }   
}



func fuzzyHalfer(someIndex int) (nextIndex int) {

	for nextIndex = someIndex; nextIndex != 0 && (nextIndex%2 == 0 || nextIndex == someIndex); {
		nextIndex /= 2
	}

	// if we initially started on an even number we'll need to take one more half
	if someIndex%2 == 0 {
		nextIndex /= 2
	}

	return nextIndex
}

func determineNumber(numberOfHalves int) (theNumber int) {

	switch numberOfHalves % 3 {
	case 0:
		theNumber = 0
	case 2:
		theNumber = 2
	case 1:
		theNumber = 1
	}

	return
}

func generate(toLength int) (gennedString string) {
	gennedString = "0"
	for len(gennedString) < toLength {
		gennedString += translate(gennedString)
	}

	return
}

func translate(someString string) string {
	trans := func(r rune) rune {
		switch {
		case r == '0':
			return '1'
		case r == '1':
			return '2'
		case r == '2':
			return '0'
		}
		return r
	}

	return strings.Map(trans, someString)
}