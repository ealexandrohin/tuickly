package streamlist

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/x/ansi"
	"github.com/eAlexandrohin/tuickly/ctx"
	"github.com/muesli/termenv"
	"github.com/nfnt/resize"
)

type Item struct {
	UserID      string
	UserLogin   string
	UserName    string
	GameName    string
	Title       string
	ViewerCount int
	StartedAt   time.Time
}

func (s Item) FilterValue() string {
	return fmt.Sprintf("%s %s %s", s.UserLogin, s.GameName, s.Title)
}

func (s Item) Get() Item {
	return s
}

type Delegate struct {
	Ctx *ctx.Ctx
	// UpdateFunc      func(tea.Msg, *Model) tea.Cmd
	// ShortHelpFunc   func() []key.Binding
	// FullHelpFunc    func() [][]key.Binding
}

func New(ctx *ctx.Ctx) Delegate {
	return Delegate{
		Ctx: ctx,
	}
}

// func (d StreamDelegate) SetHeight(h int) {
// 	d.height = h
// }

func (d Delegate) Height() int {
	// return d.height
	return 14
}

// func (d StreamDelegate) SetSpacing(s int) {
// 	d.spacing = s
// }

func (d Delegate) Spacing() int {
	// return d.spacing
	return 1
}

func (d Delegate) Update(msg tea.Msg, m *list.Model) tea.Cmd {
	return nil
}

func (d Delegate) Render(w io.Writer, m list.Model, index int, item list.Item) {
	if m.Width() <= 0 {
		return
	}

	// log.Printf("%+v", m)

	var (
		userName string
		gameName string
		// uptime      string
		viewerCount string
		title       string
		s           = &d.Ctx.Styles
	)

	if i, ok := item.(Item); ok {
		userName = i.UserName
		gameName = ansi.Truncate(i.GameName, 16, "...")

		viewerCount = strconv.Itoa(i.ViewerCount)
		title = ansi.Truncate(i.Title, 40, "...")
	} else {
		return
	}

	isSelected := index == m.Index()

	var (
		topRow string

		bottomLeft  string
		bottomRight string
		bottomRow   string
	)

	if isSelected {
		topRow = s.StreamList.Selected.Top.Style.Render(title)

		bottomLeft = s.StreamList.Selected.Bottom.Left.Render(userName)
		bottomRight = s.StreamList.Selected.Bottom.Right.Render(viewerCount)

		bottomRow = lipgloss.JoinHorizontal(lipgloss.Left,
			bottomLeft,
			s.StreamList.Selected.Bottom.Middle.Width(40-lipgloss.Width(bottomLeft)-lipgloss.Width(bottomRight)).Render(gameName),
			bottomRight,
		)
	} else {
		topRow = s.StreamList.Normal.Top.Style.Render(title)

		bottomLeft = s.StreamList.Normal.Bottom.Left.Render(userName)
		bottomRight = s.StreamList.Normal.Bottom.Right.Render(viewerCount)

		bottomRow = lipgloss.JoinHorizontal(lipgloss.Left,
			bottomLeft,
			s.StreamList.Normal.Bottom.Middle.Width(40-lipgloss.Width(bottomLeft)-lipgloss.Width(bottomRight)).Render(gameName),
			bottomRight,
		)
	}

	img, err := getDecodedImg("/home/alex/Downloads/1920-x-1080-hd-1qq8r4pnn8cmcew4.jpg")
	if err != nil {
		panic(err)
	}

	preview := getBasicString(img, 40)

	preview = preview[:len(preview)-1]

	stream := lipgloss.JoinVertical(lipgloss.Left, preview, topRow, bottomRow)

	// log.Printf("%q", stream)

	fmt.Fprintf(w, "%v", stream)
}

// func (d Delegate) Render(w io.Writer, m list.Model, index int, item list.Item) {
// 	if m.Width() <= 0 {
// 		return
// 	}
//
// 	var (
// 		userName string
// 		gameName string
// 		// uptime      string
// 		viewerCount string
// 		title       string
// 		s           = &d.Ctx.Styles
// 	)
//
// 	if i, ok := item.(Item); ok {
// 		userName = i.UserName
// 		gameName = ansi.Truncate(i.GameName, 16, "...")
//
// 		// d := time.Since(i.StartedAt)
// 		// h := int(d / time.Hour)
// 		// d -= time.Duration(h) * time.Hour
// 		// m := int(d / time.Minute)
// 		// d -= time.Duration(m) * time.Minute
// 		// s := int(d / time.Second)
// 		// uptime = fmt.Sprintf("%02d:%02d:%02d", h, m, s)
// 		// god i wish for .toLocaleTimeString()
//
// 		viewerCount = strconv.Itoa(i.ViewerCount)
// 		title = ansi.Truncate(i.Title, 40, "...")
// 	} else {
// 		return
// 	}
//
// 	textwidth := m.Width() - s.StreamList.Normal.Title.GetPaddingLeft() - s.StreamList.Normal.Title.GetPaddingRight()
// 	title = ansi.Truncate(title, textwidth, "...")
//
// 	var (
// 		isSelected  = index == m.Index()
// 		emptyFilter = m.FilterState() == list.Filtering && m.FilterValue() == ""
// 		isFiltered  = m.FilterState() == list.Filtering || m.FilterState() == list.FilterApplied
// 	)
//
// 	if isFiltered {
// 		matchedRunes = m.MatchesForItem(index)
// 	}
//
// 	if emptyFilter {
// 		userName = s.StreamList.Dimmed.UserName.Render(userName)
// 		gameName = s.StreamList.Dimmed.GameName.Render(gameName)
// 		uptime = s.StreamList.Dimmed.Uptime.Render(uptime)
// 		viewerCount = s.StreamList.Dimmed.ViewerCount.Render(viewerCount)
// 		title = s.StreamList.Dimmed.Title.Render(title)
// 	} else if isSelected && m.FilterState() != list.Filtering {
// 		// if isFiltered {
// 		// 	unmatched := s.StreamList.Selected.Style.Inline(true)
// 		// 	matched := unmatched.Inherit(s.StreamList.FilterMatch)
// 		// 	title = lipgloss.StyleRunes(title, matchedRunes, matched, unmatched)
// 		// }
// 		userName = s.StreamList.Selected.UserName.Render(userName)
// 		gameName = s.StreamList.Selected.GameName.Render(gameName)
// 		uptime = s.StreamList.Selected.Uptime.Render(uptime)
// 		viewerCount = s.StreamList.Selected.ViewerCount.Render(viewerCount)
// 		title = s.StreamList.Selected.Title.Render(title)
// 	} else {
// 		// if isFiltered {
// 		// 	// Highlight matches
// 		// 	unmatched := s.NormalTitle.Inline(true)
// 		// 	matched := unmatched.Inherit(s.FilterMatch)
// 		// 	title = lipgloss.StyleRunes(title, matchedRunes, matched, unmatched)
// 		// }
// 		userName = s.StreamList.Normal.UserName.Render(userName)
// 		gameName = s.StreamList.Normal.GameName.Render(gameName)
// 		uptime = s.StreamList.Normal.Uptime.Render(uptime)
// 		viewerCount = s.StreamList.Normal.ViewerCount.Render(viewerCount)
// 		title = s.StreamList.Normal.Title.Render(title)
// 	}
//
// 	preview := lipgloss.JoinHorizontal(lipgloss.Top, getBasicString())
//
// 	fmt.Fprintf(w, "%v --- %v --- %v --- %v\n\t>>>%v", userName, gameName, uptime, viewerCount, title)
//
// 	i, ok := item.(StreamItem)
// 	if !ok {
// 		return
// 	}
//
// 	var style lipgloss.Style
// 	if index == m.Index() {
// 		style = lipgloss.NewStyle().
// 			Background(lipgloss.Color("205")).
// 			Foreground(lipgloss.Color("235"))
// 	} else {
// 		style = lipgloss.NewStyle().
// 			Foreground(lipgloss.Color("250"))
// 	}
//
// 	str := fmt.Sprintf("%v --- %v --- %v --- %v\n\t>>>%v", i.UserName, i.GameName, i.StartedAt, i.ViewerCount, i.Title)
//
// 	fmt.Fprint(w, style.Render(str))
// }

func getDecodedImg(fp string) (image.Image, error) {
	file, err := os.Open(fp)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	return img, nil
}

func getBasicString(img image.Image, windowSize int) string {
	img = resize.Resize(uint(windowSize), 0, img, resize.Lanczos3)

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
