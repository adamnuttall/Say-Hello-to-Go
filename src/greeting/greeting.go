/**
 * Created by IntelliJ IDEA.
 * User: Adam
 * Date: 28/11/13
 * Time: 14:00
 * To change this template use File | Settings | File Templates.
 */
package greeting
import "fmt"

type Salutation struct {
    Name string
    Greeting string
}

type Printer func(string) ()

func Greet(salutation []Salutation, do Printer, isFormal bool, times int) {
    for _, s := range salutation {
        message, alternate := CreateMessage(s.Name, s.Greeting)

        if prefix := GetPrefix(s.Name); isFormal {
            do(prefix + message)
        } else {
            do(alternate)
        }
    }
}

func GetPrefix(name string) (prefix string) {

    prefixMap := map[string]string {
        "Bob" : "Mr ",
        "Joe" : "Dr ",
        "Amy" : "Dr ",
        "Mary" : "Mrs ",
    }

    prefixMap["Joe"] = "Jr "

    prefixMap["Mary"] = "", false
    //delete(prefixMap, "Mary")

    return prefixMap[name]
}

func CreateMessage(name, greeting string) (message string, alternate string) {
    message = greeting + " " + name
    alternate = "HEY!" + name
    return
}


func Print(s string) {
    fmt.Print(s)
}

func PrintLine(s string) {
    fmt.Println(s)
}

func CreatePrintFunction(custom string) Printer {
    return func(s string) {
        fmt.Println(s + custom)
    }
}
