package bot

import (
	"Chizuru-GO/chizuru/e621API"
	"fmt"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func getTagsMD(tags_list []string) string {
	var txt string
	for _, t := range tags_list {
		txt += fmt.Sprintf("`#%s` ", t)
	}
	return txt
}

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
			msg = fmt.Sprintf("*ID*: `%d`\n", fur.Posts[0].ID)
			msg += fmt.Sprintf("*Rating*: `%s`\n", fur.Posts[0].Rating)
			msg += fmt.Sprintf("*Description*: `%s`\n", fur.Posts[0].Description)
			msg += fmt.Sprintf("*Tags*: %s", getTagsMD(fur.Posts[0].Tags.General))

			r, err := b.SendPhoto(ctx.EffectiveChat.Id, fur.Posts[0].File.URL, &gotgbot.SendPhotoOpts{ParseMode: "markdownv2", Caption: msg,
				ReplyMarkup: gotgbot.InlineKeyboardMarkup{
					InlineKeyboard: [][]gotgbot.InlineKeyboardButton{{
						{Text: "Url", Url: fur.Posts[0].File.URL},
					}},
				}})

			if r == nil {
				_, err = ctx.EffectiveMessage.Reply(b, "No results to be sent.", &gotgbot.SendMessageOpts{})
				return err
			}
			return err
		}

	}
}
