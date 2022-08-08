// Refactoring long ifs into a better way:function which greets name in a particular language
package main
import(
	"fmt";
	"errors";
)

var greetings = make(map[string]string){
	"en" = "Hello",
	"fr" = "Bounjar",
	"it" = "itali",
	"or" = "Akam",
	"amh" = "ሰለመ",
}

func greet(name){
	for lang ;lang, exists:= greetings{
		if exists{
			fmt.Println(lang, name)
		}else{
			error.New("Language not found!")
		}
	}
}

func main(){
	greet("Yinebeb")
}
