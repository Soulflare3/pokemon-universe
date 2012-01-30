package main

import "os"
import "bufio"
import "image"
import "image/png"
import "fmt"
import "image/draw"
import "strings"
import "strconv"
import "io/ioutil"

const (
	TILE_WIDTH  = 48
	TILE_HEIGHT = 48
)

var ImagesMap map[int]*ImageFile = make(map[int]*ImageFile)

type ImageFile struct {
	image         image.Image
	width, height int
}

func LoadImages() {
	files, err := ioutil.ReadDir("tiles")
	if err != nil {
		fmt.Printf("Couldn't open tile directory. Error: %v\n", err)
		return
	}

	for i := 0; i < len(files); i++ {
		img := NewImageFile("tiles/" + files[i].Name())
		if img != nil {
			name := strings.Replace(files[i].Name(), ".png", "", -1)
			id, err := strconv.Atoi(name)
			if err == nil {
				ImagesMap[id] = img
			}
		}
	}
}

func NewImageFile(_path string) *ImageFile {
	img := &ImageFile{}
	file, err := os.OpenFile(_path, 0, 0)
	if file == nil {
		fmt.Printf("Error opening image file: %s\n", err)
		return nil
	}
	defer file.Close()
	data := bufio.NewReader(file)
	img.image, err = png.Decode(data)
	if err != nil {
		fmt.Printf("Error decoding png: %s\n", err)
		return nil
	}
	img.width = img.image.Bounds().Size().X
	img.height = img.image.Bounds().Size().Y
	return img
}

func NewImage(_width int, _height int) *ImageFile {
	img := &ImageFile{}
	img.image = image.NewNRGBA(image.Rect(0, 0, _width, _height))
	img.width = _width
	img.height = _height
	return img
}

func (i *ImageFile) WriteToFile(_file string) {
	file, err := os.OpenFile(_file, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Printf("Error writing image file: %s\n", err)
		return
	}
	defer file.Close()
	data := bufio.NewWriter(file)
	err = png.Encode(data, i.image)
	if err != nil {
		fmt.Printf("Error encoding png: %s\n", err)
		return
	}
	data.Flush()
}

const (
	IMAGE_RGBA = iota
	IMAGE_NRGBA
)

func (i *ImageFile) DrawOn(_dst *ImageFile, _x int, _y int) {
	var dst draw.Image
	switch t := _dst.image.(type) {
	case *image.RGBA:
		dst = draw.Image(_dst.image.(*image.RGBA))
	case *image.NRGBA:
		dst = draw.Image(_dst.image.(*image.NRGBA))
	default:
		fmt.Printf("[ImageFile] ERROR GET dst type: %T\n", t)
	}

	draw.Draw(dst, image.Rect(_x, _y, _x+_dst.width, _y+_dst.height), i.image, image.Pt(0, 0), draw.Src)
}
