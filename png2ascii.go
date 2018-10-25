package main

import (
    "github.com/nfnt/resize"

    "image"
    "image/color"
    "bytes"
    _ "image/png"
    "reflect"
    "fmt"
    "log"
    "os"
)

func Turn2Ascii(img image.Image, w uint) []byte {
    var ASCIISTR = " .,:-=+!?I7Z$8#%@"
    table := []byte(ASCIISTR)
    buf := new(bytes.Buffer)
    smolImg := resize.Resize(w, 0, img, resize.Lanczos3)

    bounds := smolImg.Bounds()
    /*w, h := bounds.Max.X, bounds.Max.Y*/
    // grayScale := image.NewGray(bounds)
    for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
        for x := bounds.Min.X; x < bounds.Max.X; x++ {
            //grayScale.Set(x, y, img.At(x, y))
            g := color.GrayModel.Convert(smolImg.At(x,y))
            y := reflect.ValueOf(g).FieldByName("Y").Uint()
            pos := int(y * 16 / 255)
            _ = buf.WriteByte(table[pos])
        }
        _ = buf.WriteByte('\n')
    }

    return buf.Bytes()
}

func main() {
    if len(os.Args[1:]) < 1 {
        log.Printf("useage ./asciify [png file]\n")
        return
    }
    filename := os.Args[1]
    // fmt.Println("filname = %s", filenameFoo)
    // filename := "kirby.png"
    infile, err := os.Open(filename)

    if err != nil {
        log.Printf("failed opening %s: %s", filename, err)
        panic(err.Error())
    }
    defer infile.Close()

    imgSrc, _, err := image.Decode(infile)
    if err != nil {
        panic(err.Error())
    }

    ascii := Turn2Ascii(imgSrc, uint(80))
    fmt.Print(string(ascii))
}
