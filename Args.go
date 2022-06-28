package main


//---------------------命令行参数-----------------------
//第一种方式
/*import(
	"os"
	"fmt"
)

func main(){
	var s,sep string
	for i:=1;i<len(os.Args);i++{
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}*/
//第二种方式是如果args是string类型或者slice类型，则可这样遍历
import(
	"fmt"
	"os"
	"strings"
)

func main(){
	//var sep string
	/*for _,val := range(os.Args[1:]){
		sep += val + " "
	}
	fmt.Println(sep)*/
	//输出为slice带中括号的数据
	fmt.Println(os.Args[1:])
	//strings包内的join方法
	fmt.Println(strings.Join(os.Args[1:]," "))
}

