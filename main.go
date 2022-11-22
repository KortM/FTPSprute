package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/jlaffaye/ftp"
)

func main() {
	fmt.Println("Start")
	//connect("213.189.208.62", "21", "anonymous", "anonymous")
	scanner, err := load_username_list()
	if err != nil {
		log.Fatal(err, scanner.Text())
	}
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	fmt.Println("end")

}

func connect(host string, port string, user string, passwd string) ([]string, error) {
	//Func to connect remote ftp host
	var result []string
	c, err := ftp.Dial(strings.Join([]string{host, port}, ":"), ftp.DialWithTimeout(time.Second*5))
	if err != nil {
		log.Fatal(err)
	}
	err = c.Login(user, passwd)
	if err != nil {
		//log.Fatal(err)
		return result, err
	} else {
		result[0] = user
		result[1] = passwd
		return result, nil
	}
}

func load_passwd_list() (*bufio.Scanner, error) {
	//Load password list
	passwd_list := os.Args[2]
	input, err := os.Open(passwd_list)
	if err != nil {
		//log.Fatal(err)
		return bufio.NewScanner(strings.NewReader("File read failed")), err
	}
	scanner := bufio.NewScanner(input)
	return scanner, nil
}

func load_username_list() (*bufio.Scanner, error) {
	//Load username file
	user_list := os.Args[1]
	input, err := os.Open(user_list)
	if err != nil {
		return bufio.NewScanner(strings.NewReader("File read failed")), err
	}
	scanner := bufio.NewScanner(input)
	return scanner, nil
}
func load_proxy() error {
	//TODO
	return nil
}
