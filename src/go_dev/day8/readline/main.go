package readline

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	path, err := os.Getwd()
	file, err := os.Open(path+"/src/go_dev/day8/readline/main.go")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	var line []byte
	for {
		// the different between readline and ....
		data, prefix, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		line = append(line, data...)
		if !prefix {
			fmt.Printf("data: %s\n", string(line))
			line = line[:]
		}
	}
}