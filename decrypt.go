package main

import (
	"crypto/aes"
	"crypto/cipher"
	"io"
	"log"
	"os"
	"strings"
)

func decrypt(file string) {
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

	fi, err := infile.Stat()
	if err != nil {
		log.Fatal(err)
	}

	iv := make([]byte, block.BlockSize())
	msgLen := fi.Size() - int64(len(iv))
	_, err = infile.ReadAt(iv, msgLen)
	if err != nil {
		log.Fatal(err)
	}

	outfile, err := os.OpenFile(strings.ReplaceAll(file, ".davidencoded", ""), os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		log.Fatal(err)
	}
	defer outfile.Close()

	buf := make([]byte, 1024)
	stream := cipher.NewCTR(block, iv)
	for {
		n, err := infile.Read(buf)
		if n > 0 {

			if n > int(msgLen) {
				n = int(msgLen)
			}
			msgLen -= int64(n)
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
	log.Printf("successfully decrypted %s to %s", infile.Name(), outfile.Name())
}
