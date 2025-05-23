package main

import (
	"bufio"
	"fmt"
	"os"
)


var (
	reader = bufio.NewReaderSize(os.Stdin, 1<<20)
	writer = bufio.NewWriter(os.Stdout)
)

func scanString() string {
	var b byte
	for {
		b, _ = reader.ReadByte()
		if b > ' ' {
			break
		}
	}
	buf := make([]byte, 0, 20)
	for {
		buf = append(buf, b)
		b, _ = reader.ReadByte()
		if b <= ' ' {
			break
		}
	}
	return string(buf)
}
func scanInt() int {
	var n int
	sign := 1
	b, _ := reader.ReadByte()
	for (b < '0' || b > '9') && b != '-' {
		b, _ = reader.ReadByte()
	}
	if b == '-' {
		sign = -1
		b, _ = reader.ReadByte()
	}
	for '0' <= b && b <= '9' {
		n = n*10 + int(b-'0')
		b, _ = reader.ReadByte()
	}
	return n * sign
}
func print(a ...any) { fmt.Fprint(writer, a...) }
func println(a ...any) { fmt.Fprintln(writer, a...) }


type Booth struct {
	id    int
	sweet int
	salty int
	hinting int
	neighbors []int
	values map[int]int  // index로 갈때의 value
}
func (b *Booth) visited(bitmask int) bool {
	return (bitmask & (1 << b.id)) != 0
}

func main() {
	defer writer.Flush()

	a := scanInt()
	print(a)
}