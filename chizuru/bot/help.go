package bot

import (
	"Chizuru-GO/chizuru/nekosAPI"
	"fmt"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"strings"
)

var NekosHelpStr = fmt.Sprintf(`
<b>This is the help for the nekos.life package</b>

This bot supports a variety of endpoints, use them with the <code>/nekos</code> cmd.

Example: <code>/nekos neko</code>.

Supported endpoints:
%s
`, nekosAPI.GetEndpointsHTML())

var BotHelpSTR = `
<b>Hi there, I'm %s, I fetch lewds, just for you!</b>

Currently I have these modules, 
do /help modname to know about em.
➨ <code>nekos</code>
➨ <code>e621</code>

`

var e621HelpStr = `
<b>This is the help for the e621 package</b>.

This makes me send furry posts fetched from <code>e621.net</code>

To use, send <code>/e621 tag</code>
Example: <code>/e621 fox</code>`

func getHelp(b *gotgbot.Bot, ctx *ext.Context) (err error) {
	if ctx.EffectiveChat.Type == "private" {
		args := ctx.Args()
		if len(args) == 1 {
			_, err = ctx.EffectiveMessage.Reply(b, fmt.Sprintf(BotHelpSTR, b.User.FirstName), &gotgbot.SendMessageOpts{ParseMode: "html"})
			return err
		} else if strings.ToLower(args[1]) == "nekos" {
			_, err = ctx.EffectiveMessage.Reply(b, NekosHelpStr, &gotgbot.SendMessageOpts{ParseMode: "html"})
			return err
		} else if strings.ToLower(args[1]) == "e621" {
			_, err = ctx.EffectiveMessage.Reply(b, e621HelpStr, &gotgbot.SendMessageOpts{ParseMode: "html"})
			return err
		}
	} else {
		_, err = ctx.EffectiveMessage.Reply(b, "Command restricted to PM.", &gotgbot.SendMessageOpts{ParseMode: "html"})
		return err
	}
	return err
}
