package main

import "fmt"
import "log"
import "bufio"
import "os"
import "strings"
import "regexp"
import "runtime"

type Node struct {
  left *Node
  right *Node
  data string
}

func main() {
    file, err := os.Open(os.Args[1])
    if err != nil {
        log.Fatal(err)
    }   
    defer file.Close()
    scanner := bufio.NewScanner(file)

    tree := new(Node)
    
    for scanner.Scan() {
       
        inputArray := strings.Split(scanner.Text(), " ")
        
        pattern := strings.Split(inputArray[0], "") //convert pattern string to slice
        result := inputArray[1]

        tree = buildTree(pattern, "")

        possibilities := findLeaves(tree)
        
        if testResults(possibilities, result) {
      fmt.Println("Yes")
      } else {
          fmt.Println("No") 
      }
    }   
}

func findLeaves(forRoot *Node) (leaves []*Node) {
    
  if forRoot.left == nil && forRoot.right == nil {
    return append(leaves, forRoot)
  }

  if forRoot.left != nil {
    for _, leaf := range findLeaves(forRoot.left) {
      leaves = append(leaves, leaf)
    }
  }

  if forRoot.right != nil {
    for _, leaf := range findLeaves(forRoot.right) {
      leaves = append(leaves, leaf)
    }
  }

  return leaves
}

func buildTree(pattern []string, accumData string) *Node {
  if len(pattern) == 0 {
     return &Node{ data: accumData }
  }

  //pop a token off the front
  token, pattern := pattern[0], pattern[1 :len(pattern)]

  // we'll always have a left / A branch since both 0 & 1 can translate to that
  newNode := &Node{ data: accumData }
  newNode.left = buildTree(pattern, newNode.data + "A")

  // if the token is a 1 we'll need a "B" branch as well
  if token == "1" {
    newNode.right =  buildTree(pattern, newNode.data + "B")
  }

  runtime.GC();

  return newNode
}

func testResults(possibleResults []*Node, resultToMatch string) (matched bool) {
  
  for _, testCase := range possibleResults {

    // Build the greedy regex from the possibility: AABA translates to ^A+A+B+A+$
    regex := "^" + strings.Join(strings.Split(testCase.data, ""), "+") + "+$" 

    matched, _ = regexp.MatchString(regex , resultToMatch)
      
      if matched  {
        break
      }
  }

  return 
}

