package utils

import (
	"fmt"
	"image"
	"net/http"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	"github.com/muesli/termenv"
	"github.com/nfnt/resize"
	"github.com/nicklaw5/helix"
)

// func Max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// func Gradient(base lipgloss.Style, s string) string {
// 	startColor, _ := colorful.Hex("#ffbf00")
// 	endColor, _ := colorful.Hex("#9146ff")
// 	n := len([]rune(s))
// 	var str string
// 	for i, ss := range []rune(s) {
// 		t := float64(i) / float64(max(1, n-1))
// 		c := startColor.BlendLab(endColor, t).Clamped()
// 		str += base.Foreground(lipgloss.Color(c.Hex())).Render(string(ss))
// 	}
// 	return str
// }

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

func ConvertToItems[S any](sourceSlice []S, mapper func(item S) list.Item) []list.Item {
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

func GetStreamPreview(stream helix.Stream, previewSize int) (string, error) {
	thumbnailURL := strings.NewReplacer(
		"{width}", "320",
		"{height}", "180",
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
