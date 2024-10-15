package main 
import (
	"errors"
	"io/fs"
	"os"
	"fmt"
)
func main(){
	if _, err := os.Open("non-existing"); err !=nil{
		if errors.Is(err, fs.ErrNotExist){
			fmt.Println("file does not exist")
		}else{
			fmt.Println(err)
		}
	}
}