package sizes

type Sizes struct {
	Padding int

	Tabs struct {
		Height  int
		Spacing int
	}

	StreamList struct {
		Height  int
		Width   int
		Spacing int

		Inner struct {
			Height int
			Width  int
		}
	}

	SideList struct {
		Height  int
		Width   int
		Spacing int

		Inner struct {
			Height int
			Width  int
		}
	}
}

func New() Sizes {
	s := Sizes{}

	s.Padding = 2

	// tabs

	s.Tabs.Height = 1
	s.Tabs.Spacing = 1

	// streamlist

	s.StreamList.Height = 15
	s.StreamList.Width = 41
	s.StreamList.Spacing = 1

	s.StreamList.Inner.Height = 2
	s.StreamList.Inner.Width = 39

	// sidelist

	s.SideList.Height = 2
	s.SideList.Width = 18
	s.SideList.Spacing = 0

	s.SideList.Inner.Height = 2
	s.SideList.Inner.Width = 16

	return s
}
