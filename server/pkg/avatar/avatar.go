package avatar

import (
	"bytes"
	"errors"
	"fmt"
	svg "github.com/ajstarks/svgo"
	"github.com/duke-git/lancet/v2/cryptor"
	"image/color"
	"image/jpeg"
	"image/png"
	"io"
	"strings"
	"unicode"

	"github.com/dchest/lru"
	"stathat.com/c/consistent"
)

var (
	BgColors = []string{
		//"#323A49",
		//"#6A7EA3",
		//"#9145F2",
		//"#16BBA7",
		//"#0D65FF",
	}

	defaultColorKey = "#323A49"

	// ErrUnsupportChar is returned when the character is not supported
	ErrUnsupportChar = errors.New("unsupported character")

	// ErrUnsupportedEncoding is returned when the given image encoding is not supported
	ErrUnsupportedEncoding = errors.New("avatar: Unsuppored encoding")
	BGC                    = consistent.New()
)

// InitialsAvatar represents an initials avatar.
type InitialsAvatar struct {
	drawer *drawer
	cache  *lru.Cache
}

// New creates an instance of InitialsAvatar
func New(fontBytes []byte) *InitialsAvatar {
	avatar := NewWithConfig(Config{
		MaxItems:  1024, // default to 1024 items.
		FontBytes: fontBytes,
	})
	return avatar
}

// Config is the configuration object for caching avatar images.
// This is used in the caching algorithm implemented by  https://github.com/dchest/lru
type Config struct {
	// Maximum number of items the cache can contain (unlimited by default).
	MaxItems int

	// Maximum byte capacity of cache (unlimited by default).
	MaxBytes int64

	// TrueType Font file bytes
	FontBytes []byte

	// TrueType Font size
	FontSize float64
}

// NewWithConfig provides config for LRU Cache.
func NewWithConfig(cfg Config) *InitialsAvatar {
	var err error

	avatar := new(InitialsAvatar)
	avatar.drawer, err = newDrawer(cfg.FontBytes, cfg.FontSize)
	if err != nil {
		panic(err.Error())
	}
	avatar.cache = lru.New(lru.Config{
		MaxItems: cfg.MaxItems,
		MaxBytes: cfg.MaxBytes,
	})

	return avatar
}

// DrawToBytes draws an image base on the name and size.
// Only initials of name will be draw.
// The size is the side length of the square image. Image is encoded to bytes.
//
// You can optionaly specify the encoding of the file. the supported values are png and jpeg for
// png images and jpeg images respectively. if no encoding is specified then png is used.
func (a *InitialsAvatar) DrawToBytes(colors *consistent.Consistent, name string, size int, encoding ...string) ([]byte, error) {
	if size <= 0 {
		size = 48 // default size
	}
	bgcolor := getColorByName(colors, name)
	cackeKey := lru.Key(fmt.Sprintf("%s%d%v", name, size, encoding))
	name = strings.TrimSpace(name)
	firstRune := []rune(name)[0]
	if !isHan(firstRune) && !unicode.IsLetter(firstRune) && !unicode.IsNumber(firstRune) {
		name = "O"
	}
	initials := getInitials(name)
	// get from cache
	v, ok := a.cache.GetBytes(cackeKey)
	if ok {
		return v, nil
	}

	m := a.drawer.Draw(initials, size, bgcolor)

	// encode the image
	var buf bytes.Buffer
	enc := "png"
	if len(encoding) > 0 {
		enc = encoding[0]
	}
	switch enc {
	case "jpeg":
		err := jpeg.Encode(&buf, m, nil)
		if err != nil {
			return nil, err
		}
	case "png":
		err := png.Encode(&buf, m)
		if err != nil {
			return nil, err
		}
	default:
		return nil, ErrUnsupportedEncoding
	}

	// set cache
	a.cache.SetBytes(cackeKey, buf.Bytes())

	return buf.Bytes(), nil
}

// Is it Chinese characters?
func isHan(r rune) bool {
	if unicode.Is(unicode.Scripts["Han"], r) {
		return true
	}
	return false
}

// random color
func getColorByName(c *consistent.Consistent, name string) *color.RGBA {
	key, err := c.Get(name)
	if err != nil {
		key = "#" + cryptor.Md5String(name)[0:6]
		//key = defaultColorKey
	}
	cc, err := parseHexColor(key)
	if err != nil {
		cc = color.RGBA{69, 189, 243, 255}
	}
	return &cc
}

func getColorHexByName(c *consistent.Consistent, name string) string {
	key, err := c.Get(name)
	if err != nil {
		key = defaultColorKey
	}
	return key
}

func parseHexColor(s string) (c color.RGBA, err error) {
	c.A = 0xff
	switch len(s) {
	case 7:
		_, err = fmt.Sscanf(s, "#%02x%02x%02x", &c.R, &c.G, &c.B)
	case 4:
		_, err = fmt.Sscanf(s, "#%1x%1x%1x", &c.R, &c.G, &c.B)
		// Double the hex digits:
		c.R *= 17
		c.G *= 17
		c.B *= 17
	default:
		err = fmt.Errorf("invalid length, must be 7 or 4")

	}
	return
}

func getInitials(name string) string {
	if len(name) == 0 {
		return ""
	}
	o := opts{
		allowEmail: true,
		limit:      3,
	}
	i, _ := parseInitials(strings.NewReader(name), o)
	return i
}

func Init(fontBytes []byte) *InitialsAvatar {
	for _, key := range BgColors {
		BGC.Add(key)
	}
	return New(fontBytes)
}

func Svg(w io.Writer, name string, sz int) {
	bgcolor := getColorHexByName(BGC, name)
	name = strings.TrimSpace(name)
	firstRune := []rune(name)[0]
	if !isHan(firstRune) && !unicode.IsLetter(firstRune) && !unicode.IsNumber(firstRune) {
		name = "O"
	}
	initials := getInitials(name)

	canvas := svg.New(w)
	canvas.Start(sz, sz)
	canvas.Rect(0, 0, sz, sz, fmt.Sprintf("fill:%s", bgcolor))
	fontSize := int(float64(sz) * 0.6)
	y := (2*sz - fontSize) / 2
	if unicode.IsLower(firstRune) {
		y = sz / 2
	}
	y = sz / 2
	canvas.Text(sz/2, y, initials, fmt.Sprintf("dominant-baseline:middle;text-anchor:middle;font-size:%dpx;fill:white;", fontSize))
	canvas.End()
}
