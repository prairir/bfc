package b2c

import (
	"fmt"
	"strings"
)

const DefaultSize = 30000 // default brainfuck memory array size

func Transpile(bf string, arrSize uint) string {
	var out strings.Builder

	preamble := fmt.Sprintf(`#include <stdio.h>
int main(){char array[%d]={0};char* ptr = array;`, arrSize)

	out.WriteString(preamble)

	for _, val := range bf {
		switch val {
		case '>':
			out.WriteString("++ptr;")
		case '<':
			out.WriteString("--ptr;")
		case '+':
			out.WriteString("++*ptr;")
		case '-':
			out.WriteString("--*ptr;")
		case '.':
			out.WriteString("putchar(*ptr);")
		case ',':
			out.WriteString("*ptr=getchar();")
		case '[':
			out.WriteString("while(*ptr){")
		case ']':
			out.WriteString("}")
		default:
			continue
		}
	}

	out.WriteString("return 0;}")
	return out.String()
}
