package main

import (
	"fmt"
	"os"
)

const s = `
             ,_---~~~~~----._
      _,,_,*^____      _____''*g*\"*,
     / __/ /'     ^.  /      \ ^@q   f
    [  @f | @))    |  | @))   l  0 _/
     \'/   \~____ / __ \_____/    \
      |           _l__l_           I
      }          [______]           I
      ]            | | |            |
      ]             ~ ~             |
      |                            |
       |                           |
    
        Im in ur UEFI, running rt0
`

func main() {
	if len(os.Args) == 3 {
		str2asm(os.Args[1], os.Args[2])
	} else {
		str2asm("oktext<>", s)
	}
}

func str2asm(n, v string) {
	extra := 0
	for i, r := range v {
		fmt.Printf("DATA %s+0x%x(SB)/2, $0x%04x\n", n, (i+extra)*2, r)
		if r == '\n' {
			extra++
			fmt.Printf("DATA %s+0x%x(SB)/2, $0x%04x\n", n, (i+extra)*2, '\r')
		}
	}
	fmt.Printf("DATA %s+0x%x(SB)/2, $0\n", n, (len(v)+extra)*2)
	fmt.Printf("GLOBL %s(SB),NOPTR,$%d\n", n, (len(v)+extra)*2+2)
}
