package main
import (
    "image"
    "image/color"
    "image/gif"
    "io"
    "math"
    "math/rand"
    "os"
)

var palette = []color.Color{color.RGBA{255, 255, 255, 1}, color.RGBA{0, 255, 0, 1}, color.RGBA{255, 0, 0, 1}, color.RGBA{0, 0, 255, 1}}

const (
    whiteIndex = 0 // whit is background color
    greenIndex = 1 // green is changing color
    redIndex = 2
    blueIndex = 3
)

func main () {
    lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
    const (
        cycles = 5
        res = 0.001
        size = 100
        nframes = 64
        delay = 8
    )

    freq := rand.Float64() * 3.0
    anim := gif.GIF{LoopCount: nframes}
    phase := 0.0
    for i := 0; i < nframes; i ++ {
        rect := image.Rect(0, 0, 2 * size + 1, 2 * size + 1)
        img := image.NewPaletted(rect, palette)
        for t := 0.0; t < cycles * 2 * math.Pi; t+= res {
            x := math.Sin(t)
            y := math.Sin(t * freq + phase)
            if x > 0.5 {
                img.SetColorIndex(size + int(x * size + 0.5), size + int(y * size + 0.5), greenIndex) //green is the color changing as frames
            } else if x < 0.5 && x > 0 {
                img.SetColorIndex(size + int(x * size + 0.5), size + int(y * size + 0.5), blueIndex) //blue is the color changing as frames
            } else {
                img.SetColorIndex(size + int(x * size + 0.5), size + int(y * size + 0.5), redIndex) //red is the color changing as frames
            }
        }
        phase += 0.1
        anim.Delay = append(anim.Delay, delay)
        anim.Image = append(anim.Image, img)
    }
    gif.EncodeAll(out, &anim)


}
