package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

func main() {
	s := "Love is but a song to sing Fear's the way we die You can make the mountains ring Or make the angels cry Though the bird is on the wing And you may not know why Come on people now Smile on your brother Everybody get together Try to love one another Right now"

	encodeStd := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	s64 := base64.NewEncoding(encodeStd).EncodeToString([]byte(s))
	s64_2 := base64.StdEncoding.EncodeToString([]byte(s))

	fmt.Println(len(s))
	fmt.Println(len(s64))
	fmt.Println(s)
	fmt.Println(s64)
	fmt.Println(s64_2)
	fmt.Println()

	bs, err := base64.StdEncoding.DecodeString(s64_2)
	if err != nil {
		log.Fatalln("Im giving her all she's got captain", err)
	}

	fmt.Println(string(bs))
}
