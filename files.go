package main

import (
    "os"
	"path/filepath"
    "image/jpeg"
	"github.com/nfnt/resize"
	"log"
)

func getPaths(root string) []string {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
        files = append(files, path)
        return nil
	})
	if err != nil {
        panic(err)
    }
	return files
}

func resizeImage(path string)  {

}  

func main() {
    // files := getPaths("/workspace/test-go/")

    // for _, file := range files {
    //     fmt.Println(file)
	// }
	file, err := os.Open("/workspace/test-go/a.jpg")
	if err != nil {
		log.Fatal(err)
	}
	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	m := resize.Resize(100, 0, img, resize.Lanczos3)

	out, err := os.Create("test_resized.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	jpeg.Encode(out, m, nil)

}

