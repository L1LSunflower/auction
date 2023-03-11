package webp

import (
	"github.com/L1LSunflower/auction/pkg/webp/decoder"
	"github.com/L1LSunflower/auction/pkg/webp/encoder"
	"image"
	"image/color"
	"io"
)

func init() {
	image.RegisterFormat("webp", "RIFF????WEBPVP8", quickDecode, quickDecodeConfig)
}

func quickDecode(r io.Reader) (image.Image, error) {
	return Decode(r, &decoder.Options{})
}

func quickDecodeConfig(r io.Reader) (image.Config, error) {
	return DecodeConfig(r, &decoder.Options{})
}

// Decode picture from reader
func Decode(r io.Reader, options *decoder.Options) (image.Image, error) {
	if dec, err := decoder.NewDecoder(r, options); err != nil {
		return nil, err
	} else {
		return dec.Decode()
	}
}

func DecodeConfig(r io.Reader, options *decoder.Options) (image.Config, error) {
	if dec, err := decoder.NewDecoder(r, options); err != nil {
		return image.Config{}, err
	} else {
		return image.Config{
			ColorModel: color.RGBAModel,
			Width:      dec.GetFeatures().Width,
			Height:     dec.GetFeatures().Height,
		}, nil
	}
}

// Encode picture and write to io.Writer
func Encode(w io.Writer, src image.Image, options *encoder.Options) error {
	if enc, err := encoder.NewEncoder(src, options); err != nil {
		return err
	} else {
		return enc.Encode(w)
	}
}
