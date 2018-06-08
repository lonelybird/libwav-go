package main

import (
	"os"
	"fmt"
)

type Wav struct {
	dsize uint32
	fsize uint32
	head []byte
	data []byte
}

func NewWav() *Wav {
	wav := &Wav {}
	wav.head = make([]byte, 44)
	return wav
}

func (w *Wav) load(file string) {
	f, _ := os.Open(file)
	defer f.Close()
	f.Read(w.head)
	fmt.Println("head:\n", w.head)
}

func main() {
	w := NewWav()
	w.load("../lua/pb/8b/8b_1.wav")
}
