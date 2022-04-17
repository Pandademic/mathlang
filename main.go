package main

import (
    "fmt"
    "strconv"
    "strings"
)

func main() {
    code := "1 + 1"                  // the MATH
    selector := strings.Fields(code) // get all the different parts
    oFirstNum := selector[0]         // select the first part
    operator := selector[1]          // the middle part(the operator)
    oSecondNum := selector[2]        // last part the second num // for conversion purposes it's an alias
    // Great but we can't add string so we Convert

    firstNum, _ := strconv.Atoi(oFirstNum) // it can't work it's magic on itself! 

    secondNum, _ := strconv.Atoi(oSecondNum)

    /*
       We have to do a different operation based on what we got as our operator...
    */
    switch operator {
    case "+": // Its a 'plus' sign ...
        result := firstNum + secondNum
        fmt.Println("Result", result, "\n")
    case "*": // So on!
        result := firstNum * secondNum
        fmt.Println("Result", result, "\n")
    case "/":
        result := firstNum / secondNum
        fmt.Println("Result", result, "\n")
    case "-":
        result := firstNum - secondNum
        fmt.Println("Result", result, "\n")
    }
}
