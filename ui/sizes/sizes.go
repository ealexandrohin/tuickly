package sizes

type Sizes struct {
	StreamList struct {
		Height  int
		Width   int
		Spacing int

		Preview struct {
			Width int
		}
	}
	SideList struct {
		Height  int
		Width   int
		Spacing int
	}
}

func New() Sizes {
	s := Sizes{}

	// streamlist

	s.StreamList.Height = 15
	s.StreamList.Width = 41
	s.StreamList.Spacing = 1

	// s.StreamList.Height = 13
	// s.StreamList.Width = 39
	// s.StreamList.Spacing = 2

	s.StreamList.Preview.Width = 39

	// sidelist

	s.SideList.Height = 2
	s.SideList.Width = 16
	s.SideList.Spacing = 0

	return s
}
