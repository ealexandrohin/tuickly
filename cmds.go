package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	helix "github.com/nicklaw5/helix"
)

type ContentMsg string

func start(m *Model) tea.Cmd {
	return func() tea.Msg {
		live, err := paginate(func(after string) ([]helix.Stream, string, error) {
			params := &helix.FollowedStreamsParams{
				UserID: m.Auth.User.ID,
				First:  100,
				After:  after,
			}

			resp, err := client.GetFollowedStream(params)
			if err != nil {
				return nil, "", err
			}

			return resp.Data.Streams, resp.Data.Pagination.Cursor, nil
		})
		if err != nil {
			panic(err)
		}

		var s string

		s += fmt.Sprintf("len: %v\n", len(live))

		// for _, stream := range live {
		// 	s += fmt.Sprintf("%v %v\n", stream.UserName, stream.ViewerCount)
		// }

		return ContentMsg(s)
	}
}
