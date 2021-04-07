package main

import (
	"encoding/base64"
	"flag"
	"log"

	"github.com/atotto/clipboard"
)

var decode = flag.Bool("d", false, "decode string in clipboard")

func main() {
	flag.Parse()

	str, err := clipboard.ReadAll()
	if err != nil {
		log.Fatalln(err)
	}

	if *decode {
		b, err := base64.RawURLEncoding.DecodeString(str)
		if err != nil {
			log.Fatalln(err)
		}
		str = string(b)
	} else {
		str = base64.RawURLEncoding.EncodeToString([]byte(str))
	}

	err = clipboard.WriteAll(str)
	if err != nil {
		log.Fatalln(err)
	}
}
