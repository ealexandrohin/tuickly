// Package utils are utils, duh
package utils

import (
	"fmt"
	"image"
	"math"
	"net/http"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	"github.com/muesli/termenv"
	"github.com/nfnt/resize"
	"github.com/nicklaw5/helix"
)

// Humanize converts a large integer into a concise, human-readable string
// representation using suffixes:
//
// K for thousands (1,000)
//
// M for millions (1,000,000)
//
// B for billions (1,000,000,000)
//
// The value is rounded to the nearest whole number before applying the suffix.
// For numbers less than 1000, the original number is returned as a string.
//
// Examples:
//
// 999 => 999
// 1000 => 1K
// 1500 => 2K
// 50000 => 50K
// 500000 => 500K
// 1000000 => 1M
// 1500000 => 2M
// 500000000 => 500M
// 1000000000 => 1B
// 1500000000 => 2B
func Humanize(number int) string {
	f := float64(number)
	result := strings.Builder{}

	switch {
	case f >= 1000000000.0:
		fmt.Fprintf(&result, "%dB", int(math.Round(f/1000000000.0)))
	case f >= 1000000.0:
		fmt.Fprintf(&result, "%dM", int(math.Round(f/1000000.0)))
	case f >= 1000.0:
		fmt.Fprintf(&result, "%dK", int(math.Round(f/1000.0)))
	// takes too much space
	// case f >= 10000.0:
	// 	fmt.Fprintf(&result, "%dK", int(math.Round(f/1000.0)))
	// case f >= 1000.0:
	// 	fmt.Fprintf(&result, "%.1fK", math.Round(f/100.0)/10.0)
	default:
		fmt.Fprint(&result, number)
	}

	return result.String()
}

// Paginate fetches all pages of resources of type T by repeatedly calling the
// provided request function until no more pages are available.
//
// The type parameter T represents the type of the items being paginated.
//
// The request function takes a string, 'after', which is the cursor for the
// next page, and returns a slice of items ([]T), the next cursor string, and
// an error. The initial call to request should receive an empty string ("")
// for 'after'. Pagination stops when the request function returns an empty
// string for the next cursor.
//
// It returns a single slice containing all fetched items and an error if any
// request fails.
func Paginate[T any](
	request func(after string) ([]T, string, error),
) ([]T, error) {
	var (
		all    []T
		cursor string
	)

	for {
		items, nextCursor, err := request(cursor)
		if err != nil {
			return nil, err
		}

		all = append(all, items...)
		if nextCursor == "" {
			break
		}

		cursor = nextCursor
	}

	return all, nil
}

// ConvertToItems converts a slice of source items (sourceSlice) of any type S
// into a slice of list.Item using a custom mapping function.
//
// S represents the type of the source items in the input slice.
//
// The mapper function is responsible for transforming a single source item (S)
// into a list.Item.
//
// It returns a new slice of list.Item with the same length as the source slice.
func ConvertToItems[S any](
	sourceSlice []S,
	mapper func(item S) list.Item,
) []list.Item {
	items := make([]list.Item, len(sourceSlice))

	for i, s := range sourceSlice {
		items[i] = mapper(s)
	}

	return items
}

// GetImagePreview was shamelessly copied from https://github.com/savedra1/clipse
func GetImagePreview(img image.Image, imageSize int) string {
	img = resize.Resize(uint(imageSize), 0, img, resize.Lanczos3)

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	p := termenv.ColorProfile()

	var sb strings.Builder

	for y := 0; y < height; y += 2 {
		for x := 0; x < width; x++ {
			upperR, upperG, upperB, _ := img.At(x, y).RGBA()
			lowerR, lowerG, lowerB, _ := img.At(x, y+1).RGBA()

			upperColor := p.Color(
				fmt.Sprintf(
					"#%02x%02x%02x", uint8(upperR>>8), uint8(upperG>>8), uint8(upperB>>8),
				),
			)
			lowerColor := p.Color(
				fmt.Sprintf(
					"#%02x%02x%02x", uint8(lowerR>>8), uint8(lowerG>>8), uint8(lowerB>>8),
				),
			)
			sb.WriteString(termenv.String("▀").Foreground(lowerColor).Background(upperColor).String())
		}
		sb.WriteString("\n")
	}

	return sb.String()
}

// GetStreamPreview fetches a thumbnail image from a stream and generates
// a terminal-based preview of it.
//
// It takes a [helix.Stream] struct, which contains the stream details including
// the 'ThumbnailURL' template, and a desired 'previewSize' (the target width
// for the terminal preview).
//
// The 'ThumbnailURL' template is modified to request a 320x180 image size before
// fetching. The image data is then decoded and passed to GetImagePreview for
// terminal rendering.
//
// It returns the ANSI-encoded string for the terminal preview or an error if
// fetching or decoding the image fails, or if the HTTP request is unsuccessful.
func GetStreamPreview(stream helix.Stream, previewSize int) (string, error) {
	thumbnailURL := strings.NewReplacer(
		"{width}", "320",
		"{height}", "180",
		// twitch's default preview size
	).Replace(stream.ThumbnailURL)

	resp, err := http.Get(thumbnailURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", err
	}

	img, _, err := image.Decode(resp.Body)
	if err != nil {
		return "", err
	}

	return GetImagePreview(img, previewSize), nil
}
