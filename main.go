package main

import (
	"fmt"
	"os"
	"sync"
)

// Made with help of https://levelup.gitconnected.com/a-short-guide-to-encryption-using-go-da97c928259f

var filesFull []string

func main() {
	fmt.Println("WARNING! USE WITH CAUTION. DON'T USE WITH IMPORTANT FILES")
	fmt.Println("WARNING! MAKE SURE NOT TO CHANGE KEY AFTER ENCRYPTION AS FILES WILL GET CORRUPTED WHEN DECRYPTING WITH WRONG KEY")
	initialize()
}
func initialize() {
	var wg sync.WaitGroup
	var action string
	var foldername string
	fmt.Printf("Type %q to encrypt, %q to decrypt: ", "e", "d")
	fmt.Scan(&action)
	if action != "e" && action != "d" {
		fmt.Println("Invalid action provided, please try again")
		initialize()
		return
	}
	fmt.Print("Enter folder name, (must exist in this path): ")
	fmt.Scan(&foldername)
	path := "./" + foldername
	if action == "e" {
		readFiles(path)
		for _, v := range filesFull {
			wg.Add(1)
			go func(file string) {
				defer wg.Done()
				encrypt(file)
				os.Remove(file)
			}(v)
		}
	} else if action == "d" {
		readFiles(path)
		for _, v := range filesFull {
			wg.Add(1)
			go func(file string) {
				defer wg.Done()
				decrypt(file)
				os.Remove(file)
			}(v)
		}

	}
	wg.Wait()
}
