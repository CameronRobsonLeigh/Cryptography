package main

import (
    "encoding/hex"
    "encoding/base64"
    "fmt"
    "log"
)

// hexadecimal string -> each byte represented as 2 characters
// Base 64 encoding converts binary data -> ASCII string format for easier handling and transmission
// Conversion converts Hex -> bytes -> Base 64 (encoded)
func convertHexToBase64(hexStr string) (string, error) {
    // Decode the hex string into bytes
    bytes, err := hex.DecodeString(hexStr)
    if err != nil {
        return "", err
    }

    // Encode the byte slice into Base64
    base64Str := base64.StdEncoding.EncodeToString(bytes)
    
    return base64Str, nil
}

func main() {  
    base64Str, err := convertHexToBase64("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Base64:", base64Str)
}
