package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

func (g *Game) Draw(screen *ebiten.Image) {
	if g.timer > 0 {
		screen.DrawImage(g.splash, nil)
	} else {

		screen.Fill(bg)

		screen.DrawImage(level.board, level.boardOpts)

		for _, p := range level.pieces {
			screen.DrawImage(p.image, p.opts)
		}

		text.Draw(screen, level.label, *regular, (w-bs)/2, (h-bs)/2-10, fg)
		if len(state) == 1 {
			text.Draw(screen, fmt.Sprintf("moves remaining: %d", level.moves), *regular, (w-bs)/2, bs+(h-bs)/2+24, fg)
		} else {
			text.Draw(screen, fmt.Sprintf("moves remaining: %d\npress 'esc' to undo", level.moves), *regular, (w-bs)/2, bs+(h-bs)/2+24, fg)
		}
	}

}
