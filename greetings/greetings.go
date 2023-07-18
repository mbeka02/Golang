package greetings

import "fmt"



func Hello ( name string) string{
//:= declaring and intitalizing in one line
 message:=fmt.Sprintf(" hey there , %v",name)
 return message
}