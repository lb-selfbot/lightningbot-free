package events

import (
	"github.com/lb-selfbot/lightningbot-free/utils"
	"github.com/LightningDev1/dgc"
)

func (e *_Events) OnCommand(following dgc.ExecutionHandler) dgc.ExecutionHandler {
	return func(ctx *dgc.Ctx) {
		utils.Console.NewConsoleMessage().
			SetTitle("Command Used").
			SetColor("#C1EFFF").
			AddInfo("Command", ctx.Command.Name).
			Show()

		following(ctx)
	}
}
