package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"log"
	"os"
)

func encrypt(file string) {
	infile, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer infile.Close()

	key, err := os.ReadFile("key")
	if err != nil {
		log.Fatal(err)
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatal(err)
	}

	iv := make([]byte, block.BlockSize())
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		log.Fatal(err)
	}

	outfile, err := os.OpenFile(fmt.Sprintf("%s%s", file, ".davidencoded"), os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		log.Fatal(err)
	}
	defer outfile.Close()

	buf := make([]byte, 1024)
	stream := cipher.NewCTR(block, iv)
	for {
		n, err := infile.Read(buf)
		if n > 0 {
			stream.XORKeyStream(buf, buf[:n])

			outfile.Write(buf[:n])
		}

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Printf("Read %d bytes: %v", n, err)
			break
		}
	}

	outfile.Write(iv)
	log.Printf("successfully encrypted %s to %s", infile.Name(), outfile.Name())
}
