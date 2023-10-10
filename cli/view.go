package cli

import (
	"fmt"

	"github.com/rivo/tview"
	"github.com/cliche-niche/CS455/blob"
)

type View struct {
	textArea 	*tview.TextArea
	mainView	*tview.Grid
}

func (mv *View) InitView(b *blob.Blob) {
	mv.textArea = tview.NewTextArea().
		SetPlaceholder("Enter text here...")
	mv.textArea.SetTitle("Text Area").SetBorder(true)
	helpInfo := tview.NewTextView().
		SetText(" Press Ctrl-S to save, press Ctrl-C to exit")

	position := tview.NewTextView().
		SetDynamicColors(true).
		SetTextAlign(tview.AlignRight)

	updateInfos := func() {
		fromRow, fromColumn, toRow, toColumn := mv.textArea.GetCursor()
		if fromRow == toRow && fromColumn == toColumn {
			position.SetText(fmt.Sprintf("Row: [yellow]%d[white], Column: [yellow]%d ", fromRow, fromColumn))
		} else {
			position.SetText(fmt.Sprintf("[red]From[white] Row: [yellow]%d[white], Column: [yellow]%d[white] - [red]To[white] Row: [yellow]%d[white], To Column: [yellow]%d ", fromRow, fromColumn, toRow, toColumn))
		}
	}

	mv.textArea.SetMovedFunc(updateInfos)
	mv.textArea.SetText(b.GetContents(), true)
	updateInfos()

	mv.mainView = tview.NewGrid().
		SetRows(0, 1).
		AddItem(mv.textArea, 0, 0, 1, 2, 0, 0, true).
		AddItem(helpInfo, 1, 0, 1, 1, 0, 0, false).
		AddItem(position, 1, 1, 1, 1, 0, 0, false)
	
}