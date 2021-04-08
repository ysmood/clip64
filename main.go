package main

import (
	"encoding/base64"
	"flag"
	"log"

	"github.com/atotto/clipboard"
)

var decode = flag.Bool("d", false, "decode string in clipboard")
var urlSafe = flag.Bool("u", false, "whether to use url-safe encoding or not")
var noPadding = flag.Bool("n", false, "whether to use padding or not")

func main() {
	flag.Parse()

	str, err := clipboard.ReadAll()
	if err != nil {
		log.Fatalln(err)
	}

	var enc *base64.Encoding
	if *urlSafe {
		enc = base64.URLEncoding
	} else {
		enc = base64.StdEncoding
	}

	if *noPadding {
		enc = enc.WithPadding(base64.NoPadding)
	}

	if *decode {
		b, err := enc.DecodeString(str)
		if err != nil {
			log.Fatalln(err)
		}
		str = string(b)
	} else {
		str = enc.EncodeToString([]byte(str))
	}

	err = clipboard.WriteAll(str)
	if err != nil {
		log.Fatalln(err)
	}
}
