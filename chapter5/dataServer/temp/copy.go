package temp

import (
	"io"
	"os"
)

func copyTempToObjects(f *os.File, name string) {
	w, _ := os.Create(name)
	io.Copy(w, f)
}
