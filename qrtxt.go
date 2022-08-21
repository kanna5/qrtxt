package qrtxt

import (
	"strings"

	"github.com/skip2/go-qrcode"
)

type RecoveryLevel int

const (
	Low     = RecoveryLevel(qrcode.Low)
	Medium  = RecoveryLevel(qrcode.Medium)
	High    = RecoveryLevel(qrcode.High)
	Highest = RecoveryLevel(qrcode.Highest)
)

var (
	pixels = []rune{'█', '▀', '▄', ' '}

	b2i = map[bool]int{false: 0, true: 1}
)

type Encoded []string

func (e *Encoded) String() string {
	if e == nil {
		return ""
	}
	return strings.Join([]string(*e), "\n")
}

func Encode(text string, level RecoveryLevel) (*Encoded, error) {
	qr, err := qrcode.New(text, qrcode.RecoveryLevel(level))
	if err != nil {
		return nil, err
	}

	bitmap := qr.Bitmap()

	textWidth := len(bitmap[0])
	textLines := len(bitmap) / 2
	if len(bitmap)%2 != 0 {
		textLines += 1
		bitmap = append(bitmap, make([]bool, textWidth))
	}

	buf := make([]rune, textLines*textWidth)
	for y := 0; y < textLines; y++ {
		offset := y * textWidth
		for x := 0; x < textWidth; x++ {
			upper := bitmap[y*2][x]
			lower := bitmap[y*2+1][x]
			buf[offset+x] = pixels[b2i[upper]<<1|b2i[lower]]
		}
	}
	ret := Encoded(make([]string, textLines))
	for y := 0; y < textLines; y++ {
		ret[y] = string(buf[y*textWidth : (y+1)*textWidth])
	}
	return &ret, nil
}
