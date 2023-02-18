package binary_test

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"testing"
)

func TestBinary(t *testing.T) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, "b")
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
	fmt.Printf("%s", 100)

}
