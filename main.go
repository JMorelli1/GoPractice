package main

import(
	"fmt"
	"errors"
	// "github.com/JMorelli1/GoPractice/models"
)

func main(){
	portUsed, errMsg := startServer(3000)
	fmt.Println(portUsed, errMsg)
}

func startServer(port int) (int, error){
	fmt.Println("Server starting.....")
	
	fmt.Println("Server started successfully")
	return 3000, errors.New("Something went wrong")
}