package mix

import (
	"fmt"
	"log"
	"path/filepath"
)

func main() {
	files, err := filepath.Glob("*")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(files)
}
