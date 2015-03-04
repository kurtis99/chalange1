package drum

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"os"
)

// DecodeFile decodes the drum machine file found at the provided path
// and returns a pointer to a parsed pattern which is the entry point to the
// rest of the data.
// TODO: implement
func DecodeFile(path string) (*Pattern, error) {

	fmt.Println("Reading file", path)

	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	stats, _ := f.Stat()
	size := stats.Size()

	data := make([]byte, size)

	defer f.Close()

	err = binary.Read(f, binary.LittleEndian, &data)
	if err != nil {
		fmt.Println("binary.Read failed:", err)
	}

	fmt.Println(hex.Dump(data))

	p := &Pattern{}
	return p, nil
}

// Pattern is the high level representation of the
// drum pattern contained in a .splice file.
// TODO: implement
type Pattern struct{}
