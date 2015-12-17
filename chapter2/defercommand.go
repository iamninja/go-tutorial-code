package main

import "fmt"

func CopyFile(dstName, srcName string) error {
	src, err := os.Open(srcName)
	if err != nil {
		return err
	}
	defer src.Close()			// runs at the end of function

	dst, err := os.Open(dstName)
	if err != nil {
		return err
	}
	defer dst.Close()			// runs at the end of function

	_, err := io.Copy(dst, src)
	return err
}

func main() {

}