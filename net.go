package main

/*import (
	"fmt"
	"net/http"
	"os"
	"io/ioutil"
)

func main(){
	for _,url := range os.Args[1:]{
		resp,err:=http.Get(url)
		if err != nil{
			fmt.Fprintf(os.Stderr,"fetch:%v\n",err)
			os.Exit(1)
		}
		b,err1 := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err1 != nil{
			fmt.Fprintf(os.Stderr,"fetch:%v\n",err1)
			os.Exit(1)
		}
		fmt.Printf("%s",b)
	}
}*/

//并发读取多个url
/*import(
	"fmt"
	"os"
	"io/ioutil"
	"net/http"
	"time"
	"io"
)

func main(){
	start := time.Now()
	ch := make(chan string)
	for _,url := range os.Args[1:]{
		go fetch(url,ch)
	}
	for range os.Args[1:]{
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs,elapsed\n",time.Since(start).Seconds())
}

func fetch(url string,ch chan<- string){
	start := time.Now()
	resq,err := http.Get(url)
	if err !=nil{
		ch <- fmt.Sprint(err)
		return
	}
	nbytes,err1 := io.Copy(ioutil.Discard,resq.Body)
	resq.Body.Close()
	if err1 != nil{
		ch <- fmt.Sprintf("while reading %s:%v",url,start)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s",secs,nbytes,url)
}*/



//web服务         
/*
import(
	"fmt"
	"log"
	"net/http"
)

func main(){
	http.HandleFunc("/",handler)
	log.Fatal(http.ListenAndServe("localhost:8000",nil))
}

func handler(resp http.ResponseWriter,req *http.Request){
	fmt.Fprintf(resp,"url path = %q\n",req.URL.Path)
}
*/
//计数
import(
	"fmt"
	"log"
	"net/http"
	"sync"
)
var mu sync.Mutex
var count int

func main(){
	http.HandleFunc("/",handler)
	http.HandleFunc("/count",counter)
	log.Fatal(http.ListenAndServe("localhost:8000",nil))
}

func handler(w http.ResponseWriter,r *http.Request){
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w,"URL.Path:%q\n",r.URL.Path)
}

func counter(w http.ResponseWriter,r *http.Request){
	mu.Lock()
	fmt.Fprintf(w,"Count:%d\n",count)
	mu.Unlock()
}
