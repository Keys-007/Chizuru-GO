package bot

import (
	"Chizuru-GO/chizuru/e621API"
	"fmt"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func getE621(b *gotgbot.Bot, ctx *ext.Context) (err error) {
	args := ctx.Args()
	if len(args) == 1 {
		_, err = ctx.EffectiveMessage.Reply(b, "Provide a tag", &gotgbot.SendMessageOpts{})
		return err
	} else {
		fur, err := e621API.DoRequest(args[1])
		var msg string
		if err != nil {
			msg = err.Error()
			_, err = ctx.EffectiveMessage.Reply(b, msg, &gotgbot.SendMessageOpts{})
			return err
		} else {
			msg = fmt.Sprintf("<b>ID</b>: <code>%d</code>\n", fur.Posts[0].ID)
			msg += fmt.Sprintf("<b>Rating</b>: <code>%s</code>\n", fur.Posts[0].Rating)
			msg += fmt.Sprintf("<b>Description</b>: <code>%s<code>\n", fur.Posts[0].Description)
			msg += fmt.Sprintf("<b>Tags</b>: <code>%s</code>\n", fur.Posts[0].Tags.General)

			_, err = b.SendPhoto(ctx.EffectiveChat.Id, fur.Posts[0].File.URL, &gotgbot.SendPhotoOpts{ParseMode: "html", Caption: msg})
			return err
		}

	}
}
