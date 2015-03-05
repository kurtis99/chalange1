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

	p := &Pattern{}

	err = binary.Read(f, binary.LittleEndian, p)
	if err != nil {
		fmt.Println("binary.Read failed:", err)
	}

	fmt.Println(hex.Dump(data))

	return p, nil
}

/*
Dump:
00000000  53 50 4c 49 43 45 00 00  00 00 00 00 00 c5 30 2e  |SPLICE........0.|
00000010  38 30 38 2d 61 6c 70 68  61 00 00 00 00 00 00 00  |808-alpha.......|
00000020  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
00000030  f0 42 00 00 00 00 04 6b  69 63 6b 01 00 00 00 01  |.B.....kick.....|
00000040  00 00 00 01 00 00 00 01  00 00 00 01 00 00 00 05  |................|
00000050  73 6e 61 72 65 00 00 00  00 01 00 00 00 00 00 00  |snare...........|
00000060  00 01 00 00 00 02 00 00  00 04 63 6c 61 70 00 00  |..........clap..|
00000070  00 00 01 00 01 00 00 00  00 00 00 00 00 00 03 00  |................|
00000080  00 00 07 68 68 2d 6f 70  65 6e 00 00 01 00 00 00  |...hh-open......|
00000090  01 00 01 00 01 00 00 00  01 00 04 00 00 00 08 68  |...............h|
000000a0  68 2d 63 6c 6f 73 65 01  00 00 00 01 00 00 00 00  |h-close.........|
000000b0  00 00 00 01 00 00 01 05  00 00 00 07 63 6f 77 62  |............cowb|
000000c0  65 6c 6c 00 00 00 00 00  00 00 00 00 00 01 00 00  |ell.............|
000000d0  00 00 00                                          |...|

Binary form:
00 00 00 00  04  6b 69 63 6b              01 00 00 00  01 00 00 00  01 00 00 00  01 00 00 00
01 00 00 00  05  73 6e 61 72 65           00 00 00 00  01 00 00 00  00 00 00 00  01 00 00 00
02 00 00 00  04  63 6c 61 70              00 00 00 00  01 00 01 00  00 00 00 00  00 00 00 00
03 00 00 00  07  68 68 2d 6f 70 65 6e     00 00 01 00  00 00 01 00  01 00 01 00  00 00 01 00
04 00 00 00  08  68 68 2d 63 6c 6f 73 65  01 00 00 00  01 00 00 00  00 00 00 00  01 00 00 01
05 00 00 00  07  63 6f 77 62 65 6c 6c     00 00 00 00  00 00 00 00  00 00 01 00  00 00 00 00

expected form:
(0) kick	|x---|x---|x---|x---|
(1) snare	|----|x---|----|x---|
(2) clap	|----|x-x-|----|----|
(3) hh-open	|--x-|--x-|x-x-|--x-|
(4) hh-close	|x---|x---|----|x--x|
(5) cowbell	|----|----|--x-|----|
*/

// Pattern is the high level representation of the
// drum pattern contained in a .splice file.
// TODO: implement
type Pattern struct {
	Magic   [13]byte
	Size    byte
	Version [32]byte
	Tempo   float32
	//	Tracks  []drum.Track
}

type Track struct {
}

// stringify Pattern
// This code printfs first proper strings
func (t Pattern) String() string {
	return fmt.Sprintf("Saved with HW Version: %s\nTempo: %f",
		string(t.Version[:]), t.Tempo)
}
