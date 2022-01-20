package main

// importing the packages
import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func CreateFile() {

	fmt.Println("Enter text: ")
	ipReader := bufio.NewReader(os.Stdin)
	ip, _ := ipReader.ReadString('\n')

	file, err := os.Create("test1.txt")

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	f, err := os.OpenFile("deneme.txt", os.O_RDONLY, 0755)
	if err != nil {
		panic(err)
	}

	// Reads the file input in first way
	buff := make([]byte, 100)
	for no, err := f.Read(buff); err == nil; no, err = f.Read(buff) {
		if no > 0 {
			os.Stdout.Write(buff[0:no])
		}
	}

	// In second way
	// scanner := bufio.NewScanner(f)
	// for scanner.Scan() {
	// 	str := scanner.Text()
	// 	fmt.Println(str)
	// }
	// if err := scanner.Err(); err != nil {
	// 	log.Fatal(err)
	// }

	f, err = os.OpenFile("deneme.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	defer file.Close()

	f.Write([]byte(ip))

	file.WriteString("Let's Go, making my day prudential every day")
}

func ReadFile() {

	fName := "test1.txt"

	data, err := ioutil.ReadFile(fName)
	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}
	fmt.Printf("\nFile Name: %s", fName)
	fmt.Printf("\nSize: %d bytes", len(data))
	fmt.Printf("\nData: %s", data)

}

// main function
func main() {

	CreateFile()
	ReadFile()
}
