package main

//------------查找重复的行------------
//方式一 
/*import(
	"fmt"
	"bufio"
	"os"
)

func main(){
	count := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	//length := 0
	for input.Scan(){
		count[input.Text()]++
		length++
		if length > 3{
			break
		}
	}
	for line,n := range(count){
		if n>1{
			fmt.Printf("%d\t%s\n",n,line)
		}
	}
}*/

//方式二,文件流的方式读
/*import(
	"fmt"
	"bufio"
	"os"
)

func main(){
	count := make(map[string]int)
	file := os.Args[1:]
	if len(file)==0{
		countlines(os.Stdin,count)
	}else{
		for _,arg := range file{
			f,err := os.Open(arg)
			if err != nil{
				fmt.Fprintf(os.Stderr,"dup2:%v",err)
				continue
			}
			countlines(f,count)
			f.Close()
		}
		for line,n := range(count){
			fmt.Printf("%d\t%s\n",n,line)
		}
	}
}

func countlines(f *os.File,count map[string]int){
	input:=bufio.NewScanner(f)
	for input.Scan(){
		count[input.Text()]++
	}
}*/


//把整个输入到内存中，然后分割成行读取
import(
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main(){
	count := make(map[string]int)
	for _,val := range os.Args[1:]{
		data,err := ioutil.ReadFile(val)
		if err!=nil{
			fmt.Fprintf(os.Stderr,"dup3:%v",err)
			continue
		}
		for _,line:=range strings.Split(string(data),"\n"){
			count[line]++
		}
	}
	for line,n := range count{
		if n>1{
			fmt.Printf("%d\t%s\n",n,line)
		}
	}
}
