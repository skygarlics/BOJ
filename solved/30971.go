package main_30971

// https://www.acmicpc.net/problem/30971
// 30971. 부스

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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
	
	N, K := scanInt(), scanInt()
	booths := make([]Booth, N)
	for i := 0; i < N; i++ {
		booths[i].sweet = scanInt()

		booths[i].id = i
		booths[i].values = make(map[int]int)
	}
	for i := 0; i < N; i++ {
		booths[i].salty = scanInt()
	}
	for i := 0; i < N; i++ {
		booths[i].hinting = scanInt()
		for j:= 0; j < i; j++ {
			if booths[i].hinting * booths[j].hinting <= K {
				booths[i].neighbors = append(booths[i].neighbors, j)
				booths[i].values[j] = booths[i].sweet * booths[j].salty

				booths[j].neighbors = append(booths[j].neighbors, i)
				booths[j].values[i] = booths[j].sweet * booths[i].salty
			}
		}
	}

	// dp[visited][last] : total value of visited booth with last booth
	INF := -1
	FULL := (1 << N)
	dp := make([][]int, FULL)
	for i := 0; i < FULL; i++ {
		dp[i] = make([]int, N)
		for j := 0; j < N; j++ {
			dp[i][j] = INF
		}
	}
	for i := 0; i < N; i++ {
		// initial visit is 0 value
		dp[1<<i][i] = 0
	}

	for state := 0; state < FULL; state++ {
		for last := 0; last < N; last++ {
			if dp[state][last] == INF {
				continue
			}
			for _, next := range booths[last].neighbors {
				if booths[next].visited(state) {
					continue
				}
				newVisited := state | (1 << next)
				newValue := dp[state][last] + booths[last].values[next]
				if newValue > dp[newVisited][next] {
					dp[newVisited][next] = newValue
				}
			}
		}
	}
	ret := slices.Max(dp[FULL-1])
	if ret == INF {
		println(-1)
	} else {
		println(ret)
	}
}