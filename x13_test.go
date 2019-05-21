package main

import (
	"encoding/hex"
	"fmt"
	"github.com/ccconnor/go-x13bcd/x13"
	"testing"
)

func TestX13(t *testing.T) {
	blockHeader, _ := hex.DecodeString("00000060d7954b565e2ede08fadd961e1bc97e0dd4b4f2ab36f8fb280a673862049dd97e42aae4081ff493c65c7b5b8d190a9b2b5332c3a0feaf2bc9a02fa7b0559c35a7316de35c607a431a058c4114")
	x13hash := x13.HashX13sm3(blockHeader)
	hash := hex.EncodeToString(x13hash[:])
	fmt.Println(hash)
}
