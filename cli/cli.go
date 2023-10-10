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

func (cli *Cli) RunApp() error{

	// go autoSave(cli.b, cli.view.textArea)

	err := cli.app.SetRoot(cli.pages,true).
			EnableMouse(true).Run()
	
	return err
}