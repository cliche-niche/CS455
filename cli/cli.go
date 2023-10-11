package cli

import (
	"time"
	"log"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/cliche-niche/CS455/blob"
)

const autoSaveInterval = 300

func autoSave(b *blob.Blob, textArea *tview.TextArea) {
	for {
		b.SetContents(textArea.GetText())
		err := b.SaveBlob()
		if err != nil {
			panic(err)
		}
		time.Sleep(autoSaveInterval * time.Second)
	}
}

type Cli struct  {
	app 	*tview.Application 
	view 	*View
	b 		*blob.Blob
	pages	*tview.Pages
	withDir	bool
}

func (cli *Cli) InitCli(b *blob.Blob, dirPath string, withDir bool) {
	cli.app = tview.NewApplication()
	cli.pages = tview.NewPages()
	cli.b = b
	cli.withDir = withDir

	var view View
	if withDir {
		view.InitViewDir(dirPath, cli)
	} else {
		view.InitView(b)
	}
	cli.view = &view

	cli.pages.AddPage("main", cli.view.mainView, true, true)
}

func (cli *Cli) AddExitModal() {
	modal := tview.NewModal().
			SetText("You have unsaved changes. Do you want to save them?").
			AddButtons([]string {"Yes", "No"}).
			SetDoneFunc(
				func(buttonIndex int, buttonLabel string) {
					if buttonIndex == 0 {
						cli.b.SetContents(cli.view.textArea.GetText())
						err := cli.b.SaveBlob()
						if err != nil {
							panic(err);
						}
					}
					cli.app.Stop()
					defer func() {
						if err := cli.b.CloseBlob(); err != nil {
							log.Fatalf("Failed to close the blob: %v", err)
						}
					}()
				})

	cli.pages.AddPage("exit", modal, false, false)
}

func (cli *Cli) AppInputCapture() {
	cli.app.SetInputCapture(
		func(event *tcell.EventKey) *tcell.EventKey {
			if event.Key() == tcell.KeyCtrlS {
				cli.b.SetContents(cli.view.textArea.GetText());
				go func() {
					err := cli.b.SaveBlob()
					if err != nil {
						panic(err);
					}
				}()
				return nil
			} else if event.Key() == tcell.KeyCtrlC {
				if cli.b == nil {
					return event
				}
				if cli.b.GetContents() == cli.view.textArea.GetText() {
					return event
				} else {
					cli.pages.ShowPage("exit")
					return nil
				}
			} else if event.Key() == tcell.KeyF1 {
				cli.pages.ShowPage("help")
				return nil
			}
			return event
		})
}

func (cli *Cli) FileChangeModal() {
	modal := tview.NewModal().
			SetText("You have unsaved changes. Do you want to save them?").
			AddButtons([]string {"Yes", "No"}).
			SetDoneFunc(
				func(buttonIndex int, buttonLabel string) {

					if buttonIndex == 0 {
						cli.b.SetContents(cli.view.textArea.GetText())
						err := cli.b.SaveBlob()
						if err != nil {
							panic(err);
						}
					}
					if err := cli.b.CloseBlob(); err != nil {
						log.Fatalf("Failed to close the blob: %v", err)
					}

					cli.pages.SwitchToPage("main")
					cli.view.mainView.RemoveItem(cli.view.textArea)
					cli.view.NewTextArea(cli.view.nextB)
					cli.b = cli.view.nextB
				})

	cli.pages.AddPage("fileChange", modal, true, false)
}

func (cli *Cli) AddHelp() {
	help1 := tview.NewTextView().
		SetDynamicColors(true).
		SetText(`[green]Navigation

[yellow]Left arrow[white]: Move left.
[yellow]Right arrow[white]: Move right.
[yellow]Down arrow[white]: Move down.
[yellow]Up arrow[white]: Move up.
[yellow]Ctrl-A, Home[white]: Move to the beginning of the current line.
[yellow]Ctrl-E, End[white]: Move to the end of the current line.
[yellow]Ctrl-F, page down[white]: Move down by one page.
[yellow]Ctrl-B, page up[white]: Move up by one page.
[yellow]Alt-Up arrow[white]: Scroll the page up.
[yellow]Alt-Down arrow[white]: Scroll the page down.
[yellow]Alt-Left arrow[white]: Scroll the page to the left.
[yellow]Alt-Right arrow[white]:  Scroll the page to the right.
[yellow]Alt-B, Ctrl-Left arrow[white]: Move back by one word.
[yellow]Alt-F, Ctrl-Right arrow[white]: Move forward by one word.

[blue]Press Enter for more help, press Escape to return.`)
	help2 := tview.NewTextView().
		SetDynamicColors(true).
		SetText(`[green]Editing[white]

Type to enter text.
[yellow]Ctrl-H, Backspace[white]: Delete the left character.
[yellow]Ctrl-D, Delete[white]: Delete the right character.
[yellow]Ctrl-K[white]: Delete until the end of the line.
[yellow]Ctrl-W[white]: Delete the rest of the word.
[yellow]Ctrl-U[white]: Delete the current line.

[blue]Press Enter for more help, press Escape to return.`)
	help3 := tview.NewTextView().
		SetDynamicColors(true).
		SetText(`[green]Selecting Text[white]

Move while holding Shift or drag the mouse.
Double-click to select a word.
[yellow]Ctrl-L[white] to select entire text.

[green]Clipboard

[yellow]Ctrl-Q[white]: Copy.
[yellow]Ctrl-X[white]: Cut.
[yellow]Ctrl-V[white]: Paste.
		
[green]Undo

[yellow]Ctrl-Z[white]: Undo.
[yellow]Ctrl-Y[white]: Redo.

[blue]Press Enter for more help, press Escape to return.`)

help := tview.NewFrame(help1).
		SetBorders(1, 1, 0, 0, 2, 2)
	help.SetBorder(true).
		SetTitle("Help").
		SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			if event.Key() == tcell.KeyEscape {
				cli.pages.SwitchToPage("main")
				return nil
			} else if event.Key() == tcell.KeyEnter {
				switch {
				case help.GetPrimitive() == help1:
					help.SetPrimitive(help2)
				case help.GetPrimitive() == help2:
					help.SetPrimitive(help3)
				case help.GetPrimitive() == help3:
					help.SetPrimitive(help1)
				}
				return nil
			}
			return event
		})

	cli.pages.AddPage("help", tview.NewGrid().
			SetColumns(0, 64, 0).
			SetRows(0, 22, 0).
			AddItem(help, 1, 1, 1, 1, 0, 0, true), true, false)
}

func (cli *Cli) RunApp() error{

	// go autoSave(cli.b, cli.view.textArea)

	err := cli.app.SetRoot(cli.pages,true).
			EnableMouse(true).Run()
	
	return err
}