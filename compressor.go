package compressPngJpeg

import (
	"fmt"
	"github.com/nfnt/resize"
	"github.com/ultimate-guitar/go-imagequant"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"os"
)

func crushFile(file *os.File, destfile string, speed int, compression png.CompressionLevel) error {
	image, err := ioutil.ReadAll(file)
	if err != nil {
		return fmt.Errorf("ioutil.ReadAll: %s", err.Error())
	}
	optiImage, err := imagequant.Crush(image, speed, compression)
	if err != nil {
		return fmt.Errorf("imagequant.Crush: %s", err.Error())
	}

	destFh, err := os.OpenFile(destfile, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("os.OpenFile: %s", err.Error())
	}
	defer destFh.Close()

	destFh.Write(optiImage)
	return nil
}

//version  - default false
//in - Input filename
//out - Output filename
//speed - Speed (1 slowest, 10 fastest)
//compression - Compression level (DefaultCompression = 0, NoCompression = -1, BestSpeed = -2, BestCompression = -3)
func PNGQuant(version bool, file *os.File, out string, speed, compression int)  {
	if version {
		fmt.Printf("libimagequant '%s' (%d)\n", imagequant.GetLibraryVersionString(), imagequant.GetLibraryVersion())
		os.Exit(0)
	}

	var cLevel png.CompressionLevel
	switch compression {
	case 0:
		cLevel = png.DefaultCompression
	case -1:
		cLevel = png.NoCompression
	case -2:
		cLevel = png.BestSpeed
	case -3:
		cLevel = png.BestCompression
	default:
		cLevel = png.BestCompression
	}

	err := crushFile(file, out, speed, cLevel)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	os.Exit(0)

}

func JPEGQuant(file *os.File, outPath string, width, height uint)  {

	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	m := resize.Resize(width, height, img, resize.Lanczos3)

	out, err := os.Create(outPath)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	jpeg.Encode(out, m, nil)
}
