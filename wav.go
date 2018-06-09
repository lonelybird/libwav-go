package main

import (
	"os"
	//"fmt"
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

func (w *Wav) load(file string) {

	var fsz []byte = make([]byte, 4)
	var dsz []byte = make([]byte, 4)

	f, _ := os.Open(file)
	defer f.Close()
	f.Read(w.head)

	for i := 0; i < 4; i++ {
		fsz[i] = w.head[i + 4]
		dsz[i] = w.head[i + 40]
	}

	w.fsize = binary.LittleEndian.Uint32(fsz)
	w.dsize = binary.LittleEndian.Uint32(dsz)

	w.data = make([]byte, w.dsize)
	f.ReadAt(w.data, 44)

}

func (self *Wav) save(file string) {
	
}

func (self *Wav) mix(other *Wav) {

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

func main() {
	w := NewWav()
	w.load("../heart of steel.wav")
	o := NewWav()
	o.load("../heart of steel.wav")
	w.mix(o)
}
