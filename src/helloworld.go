/**
 * Created by IntelliJ IDEA.
 * User: Adam
 * Date: 26/11/13
 * Time: 16:20
 * To change this template use File | Settings | File Templates.
 */
package main

import "./greeting"

func main() {
    //var s = greeting.Salutation{"Bob", "Hello"}

    slice := []greeting.Salutation{
        { "Bob", "Hello"},
        { "Joe", "Hi"},
        { "Mary", "What is up?" },
    }

    greeting.Greet(slice, greeting.CreatePrintFunction("?"), true, 6)
}
