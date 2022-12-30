package encoding

import (
	"encoding/binary"
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	log.Println("Do stuff BEFORE the tests!")
	exitVal := m.Run()
	log.Println("Do stuff AFTER the tests!")

	os.Exit(exitVal)
}

func TestBinary(t *testing.T) {
	buf := make([]byte, binary.MaxVarintLen64)
	for _, x := range []uint64{1, 2, 127, 128, 255, 256} {
		n := binary.PutUvarint(buf, x)
		t.Logf("%x\n", buf[:n])
	}

}
