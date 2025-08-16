package main

import (
	tea "github.com/charmbracelet/bubbletea"
	helix "github.com/nicklaw5/helix"
)

type LiveMsg []helix.Stream

func Live(m *Model) tea.Cmd {
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
			return ErrorMsg{Err: err}
		}

		return LiveMsg(live)
	}
}
