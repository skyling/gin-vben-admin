package avatar

import (
	"errors"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
	"image"
	"image/color"
	"image/draw"
)

var (
	errFontRequired = errors.New("font file is required")
	errInvalidFont  = errors.New("invalid font")
)

// drawer draws an image.Image
type drawer struct {
	fontSize    float64
	dpi         float64
	fontHinting font.Hinting
	face        font.Face
	font        *truetype.Font
}

func newDrawer(fontBytes []byte, fontSize float64) (*drawer, error) {
	if fontBytes == nil {
		return nil, errFontRequired
	}
	g := new(drawer)
	g.fontSize = fontSize
	g.dpi = 72.0
	g.fontHinting = font.HintingNone

	font, err := parseFont(fontBytes)
	if err != nil {
		return nil, errInvalidFont
	}
	g.face = truetype.NewFace(font, &truetype.Options{
		Size:    g.fontSize,
		DPI:     g.dpi,
		Hinting: g.fontHinting,
	})

	g.font = font
	return g, nil
}

// our avatar image is square
func (g *drawer) Draw(s string, size int, bg *color.RGBA) image.Image {
	// draw the background
	dst := image.NewRGBA(image.Rect(0, 0, size, size))
	draw.Draw(dst, dst.Bounds(), &image.Uniform{bg}, image.ZP, draw.Src)
	face := g.face
	fontSize := g.fontSize
	if g.fontSize == 0 {
		fontSize = float64(size) * 0.55
		face = truetype.NewFace(g.font, &truetype.Options{
			Size:    fontSize,
			DPI:     g.dpi,
			Hinting: g.fontHinting,
		})
	}
	// draw the text
	drawer := &font.Drawer{
		Dst:  dst,
		Src:  image.White,
		Face: face,
	}
	// font index
	fi := g.font.Index([]rune(s)[0])

	// glyph example: http://www.freetype.org/freetype2/docs/tutorial/metrics.png
	var gbuf truetype.GlyphBuf
	var err error
	fsize := fixed.Int26_6(fontSize * g.dpi * (64.0 / 72.0))
	err = gbuf.Load(g.font, fsize, fi, font.HintingFull)
	if err != nil {
		drawer.DrawString("O")
		return dst
	}

	// center
	dY := int((size - int(gbuf.Bounds.Max.Y-gbuf.Bounds.Min.Y)>>6) / 2)
	dX := int((size - int(gbuf.Bounds.Max.X-gbuf.Bounds.Min.X)>>6) / 2)
	y := int(gbuf.Bounds.Max.Y>>6) + dY
	x := 0 - int(gbuf.Bounds.Min.X>>6) + dX

	drawer.Dot = fixed.Point26_6{
		X: fixed.I(x),
		Y: fixed.I(y),
	}
	drawer.DrawString(s)

	return dst
}

// parseFont parse the font file as *truetype.Font (TTF)
func parseFont(fontBytes []byte) (*truetype.Font, error) {
	font, err := truetype.Parse(fontBytes)
	if err != nil {
		return nil, err
	}

	return font, nil
}
