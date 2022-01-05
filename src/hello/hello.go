package hello

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Hello, World!")
	resp, err := http.Get("https://www.baidu.com")
	check(err)
	body, err := ioutil.ReadAll(resp.Body)
	check(err)
	fmt.Println(len(body))
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
