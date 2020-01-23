package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// handles errors
func checkerr(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {
	// set file root directory
	root := `D:\My Documents\vincent\` //---> change to root dir of the files

	// read all files in the root directory
	// returns a slice of file info of type []os.FileInfo
	files, err := ioutil.ReadDir(root)
	checkerr(err)

	// loop through files in the root directory
	for _, file := range files {
		fsize := make([]byte, file.Size())
		fpath := root + file.Name()
		f, err := os.OpenFile(fpath, os.O_RDWR, 0755)
		checkerr(err)
		line, err := f.Read(fsize)
		checkerr(err)
		log.Println("file currently being read: ", fpath)

		var lnCounter int // new line number Counter

		for i := 1; i < line; i++ {
			// time.Sleep(time.Duration(10000))
			if string(fsize[i]) == "\n" {
				// fmt.Println("Line data is", string(fsize[lnCounter:i]))
				// check if line contains <<"SEGMENT CREATION DEFERRED">>
				if strings.Contains(string(fsize[lnCounter:i]), "SEGMENT CREATION DEFERRED") {
					log.Println("Defect found in ", file.Name())
					wData := []byte("    ) --SEGMENT CREATION DEFERRED\n") //---> Data to be written to file

					//replace line with wdata
					l, err := f.WriteAt(wData, int64(lnCounter))
					checkerr(err)
					log.Println("data replaced in ", file.Name(), l)
				}
				lnCounter = i + 1 // ---> move to next line
			}
		}
		f.Close()
	}
	fmt.Println("done")
}
