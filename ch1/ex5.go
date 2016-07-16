package main import (
    "image"
    "image/color"
    "image/gif"
    "io"
    "math"
    "math/rand"
    "os"
)

//var palette = []color.Color{color.White, color.Black}
//var palette = []color.Color{color.Black, color.RGBA{0, 255, 0, 0}}
//var palette = []color.Color{color.RGBA{0, 255, 0, 0}, color.Black}
//var palette = []color.Color{color.RGBA{0, 255, 0, 1}, color.RGBA{255, 0, 0, 1}} // red on green
var palette = []color.Color{color.RGBA{0, 0, 0, 1}, color.RGBA{0, 255, 0, 1}}

const (
//    whiteIndex = 0
//    blackIndex = 1
    //redIndex = 1
//    greenIndex = 0
    blackIndex = 0 // black is background color
    greenIndex = 1 // green is changing color
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
            img.SetColorIndex(size + int(x * size + 0.5), size + int(y * size + 0.5), greenIndex) //green is the color changing as frames
        }
        phase += 0.1
        anim.Delay = append(anim.Delay, delay)
        anim.Image = append(anim.Image, img)
    }
    gif.EncodeAll(out, &anim)


}
