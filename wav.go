package WAV

import (
	"os"
	"fmt"
	"encoding/binary"
)

type Wav struct {
	dsize uint32
	fsize uint32
	head []byte
	data []byte
}

func NewWav() *Wav {
	w := &Wav {}
	w.fsize = 0
	w.dsize = 0
	w.head = make([]byte, 44)
	return w
}

func (self *Wav) Load(file string) {

	var fsz []byte = make([]byte, 4)
	var dsz []byte = make([]byte, 4)

	f, _ := os.Open(file)
	defer f.Close()
	f.Read(self.head)

	for i := 0; i < 4; i++ {
		fsz[i] = self.head[i + 4]
		dsz[i] = self.head[i + 40]
	}

	self.fsize = binary.LittleEndian.Uint32(fsz)
	self.dsize = binary.LittleEndian.Uint32(dsz)

	fmt.Println(self.fsize, self.dsize)

	self.data = make([]byte, self.dsize)
	f.ReadAt(self.data, 44)

}

func (self *Wav) Save(wfilen string) {
	wfile, _ := os.Create(wfilen)
	defer wfile.Close()
	wfile.Write(self.head)
	wfile.WriteAt(self.data, 44)
}

func (self *Wav) Mix(other *Wav) {

	var i uint32

	if other.dsize > self.dsize {

		for i = 0; i < self.dsize; i++ {
			self.data[i] = uint8(self.data[i] + other.data[i])
		}

	} else {

		for i = 0; i < other.dsize; i++ {
			self.data[i] = uint8(self.data[i] + other.data[i])
		}

	}

}

func (self *Wav) Splice(other *Wav) {
	newfsz := make([]byte, 4)
	newdsz := make([]byte, 4)

	self.fsize += (other.fsize + 8)
	self.dsize += other.dsize

	binary.LittleEndian.PutUint32(newfsz, self.fsize)
	binary.LittleEndian.PutUint32(newdsz, self.dsize)

	for i := 0; i < 4; i++ {
		self.head[i + 4] = newfsz[i]
		self.head[i + 40] = newdsz[i]
	}

	self.data = append(self.data, other.data...)
}
