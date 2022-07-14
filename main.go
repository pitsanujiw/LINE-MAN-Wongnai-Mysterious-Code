package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	secret := "aWFuZ25vVzpOQU06RU5JTDp0YTpzdTpuaW9K"

	sd, _ := base64.StdEncoding.DecodeString(secret)

	fmt.Println(reserveStr(string(sd)))
	// Result: Join:us:at:LINE:MAN:Wongnai
}

func reserveStr(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}
