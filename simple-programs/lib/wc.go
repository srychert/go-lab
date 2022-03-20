package lib

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func WordCount() {
	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('.')

	m := make(map[string]int)
	for _, value := range strings.Fields(s) {
		m[value] += 1
	}

	fmt.Println(m)
}
