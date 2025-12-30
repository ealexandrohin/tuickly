// Package cmds defines Bubble Tea commands
package cmds

import (
	"time"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	helix "github.com/nicklaw5/helix/v2"

	"github.com/ealexandrohin/tuickly/consts"
	"github.com/ealexandrohin/tuickly/ctx"
	"github.com/ealexandrohin/tuickly/errs"
	"github.com/ealexandrohin/tuickly/msgs"
	"github.com/ealexandrohin/tuickly/utils"
	"github.com/ealexandrohin/tuickly/ux/items/stream"
)

// ClockTick is a Bubble Tea command that sends a [msgs.ClockTick] message
// every second.
func ClockTick() tea.Cmd {
	// return tea.Tick(1*time.Minute, func(t time.Time) tea.Msg {
	return tea.Tick(1*time.Second, func(t time.Time) tea.Msg {
		return msgs.ClockTick(t)
	})
}

// Live is a Bubble Tea command that fetches [helix.Stream] followed by the user.
// It uses the [utils.Paginate] to retrieve all pages of followed streams.
// For each stream, it generates a preview image.
// On success, it returns a [msgs.LiveMsg] containing a slice of [list.Item]
// representing the live streams. On error, it returns an [errs.ErrorMsg].
func Live(ctx *ctx.Ctx) tea.Cmd {
	return func() tea.Msg {
		live, err := utils.Paginate(func(after string) ([]helix.Stream, string, error) {
			params := &helix.FollowedStreamsParams{
				UserID: ctx.Auth.User.ID,
				First:  100,
				After:  after,
			}

			resp, err := consts.Client.GetFollowedStream(params)
			if err != nil {
				return nil, "", err
			}

			return resp.Data.Streams, resp.Data.Pagination.Cursor, nil
		})
		if err != nil {
			return errs.ErrorMsg{Err: err}
		}

		// log.Printf("%+v", live)

		streamItems := make([]list.Item, len(live))
		for i, s := range live {
			preview, err := utils.GetStreamPreview(s, ctx.Styles.Sizes.StreamList.Inner.Width)
			if err != nil {
				return errs.ErrorMsg{Err: err}
			}

			streamItems[i] = stream.Item{
				UserID:       s.UserID,
				UserLogin:    s.UserLogin,
				UserName:     s.UserName,
				GameName:     s.GameName,
				Title:        s.Title,
				ViewerCount:  s.ViewerCount,
				StartedAt:    s.StartedAt,
				ThumbnailURL: s.ThumbnailURL,
				Preview:      preview,
			}
		}

		return msgs.LiveMsg(streamItems)
	}
}
