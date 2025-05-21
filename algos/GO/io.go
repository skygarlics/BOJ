package algos

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	reader = bufio.NewReaderSize(os.Stdin, 1<<20)
	scanner = bufio.NewScanner(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func init() {
	scanner.Split(bufio.ScanWords)
	scanner.Buffer(make([]byte, 1024), 1<<20)
}
func scanString() string {
	scanner.Scan()
	return scanner.Text()
}
func scanInt() int {
	ret, _ := strconv.Atoi(scanString())
	return ret
}

func fscanString() string {
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
func fscanInt() int {
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