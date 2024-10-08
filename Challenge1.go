package main

import (
    "encoding/hex"
    "encoding/base64"
    "fmt"
    "log"
    "errors"
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

func fixedXor(bufferOne string, bufferTwo string) (string, error) {
    // Decode first string into bytes
    bytesOne, err := hex.DecodeString(bufferOne)
    if err != nil {
        return "", err
    }
    // Same with second string
    bytesTwo, err := hex.DecodeString(bufferTwo)
    if err != nil {
        return "", err
    }

    // length check, requirement is of them to be equal length
    // e.g. matching a key and plaintext
    if len(bytesOne) != len(bytesTwo) {
        return "", errors.New("buffers must be of equal length")
    }

    // create empty byte array of length of bytes one (we know will be equal)
    // more efficient that declaration using Go make function as it pre-allocates the memory.
    xorResult := make([]byte, len(bytesOne))
    for i := 0; i < len(bytesOne); i++ {
        // xor operation on both hex strings 
        xorResult[i] = bytesOne[i] ^ bytesTwo[i]
    }

    // finally, encode xor result to string (not base 64)
    return hex.EncodeToString(xorResult), nil

}

func main() {  
    base64Str, err := convertHexToBase64("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Base64:", base64Str)

    str, err := fixedXor("1c0111001f010100061a024b53535009181c", "686974207468652062756c6c277320657965");
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("XOR:", str)


}
