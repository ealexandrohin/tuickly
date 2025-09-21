package cmds

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/eAlexandrohin/tuickly/ctx"
	"github.com/eAlexandrohin/tuickly/errs"
	"github.com/eAlexandrohin/tuickly/utils"
	"github.com/eAlexandrohin/tuickly/ux/streamlist"
	"github.com/eAlexandrohin/tuickly/vars"
	helix "github.com/nicklaw5/helix"
)

type LiveMsg []list.Item

func Live(ctx *ctx.Ctx) tea.Cmd {
	return func() tea.Msg {
		live, err := utils.Paginate(func(after string) ([]helix.Stream, string, error) {
			params := &helix.FollowedStreamsParams{
				UserID: ctx.Auth.User.ID,
				First:  100,
				After:  after,
			}

			resp, err := vars.Client.GetFollowedStream(params)
			if err != nil {
				return nil, "", err
			}

			return resp.Data.Streams, resp.Data.Pagination.Cursor, nil
		})
		if err != nil {
			return errs.ErrorMsg{Err: err}
		}

		// log.Println(live)

		streamItems := utils.ConvertToItems(live, func(s helix.Stream) list.Item {
			return streamlist.Item{
				UserID:      s.UserID,
				UserLogin:   s.UserLogin,
				UserName:    s.UserName,
				GameName:    s.GameName,
				Title:       s.Title,
				ViewerCount: s.ViewerCount,
				StartedAt:   s.StartedAt,
			}
		})

		return LiveMsg(streamItems)
	}
}
