package main

import (
	"fmt"
	"flag"
	"log"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/cliche-niche/CS455/blob"
	"github.com/cliche-niche/CS455/util"
)

const autoSaveInterval = 300

func autoSave(b *blob.Blob, textArea *tview.TextArea) {
	for {
		b.SetContents(textArea.GetText())
		go func() {
			err := b.SaveBlob()
			if err != nil {
				panic(err)
			}
		}()
		time.Sleep(autoSaveInterval * time.Second)
	}
}

func runApp(b *blob.Blob) {
	app := tview.NewApplication()

	textArea := tview.NewTextArea().
		SetPlaceholder("Enter text here...")
	textArea.SetTitle("Text Area").SetBorder(true)
	helpInfo := tview.NewTextView().
		SetText(" Press Ctrl-S to save, press Ctrl-C to exit")

	position := tview.NewTextView().
		SetDynamicColors(true).
		SetTextAlign(tview.AlignRight)
	pages := tview.NewPages()

	updateInfos := func() {
		fromRow, fromColumn, toRow, toColumn := textArea.GetCursor()
		if fromRow == toRow && fromColumn == toColumn {
			position.SetText(fmt.Sprintf("Row: [yellow]%d[white], Column: [yellow]%d ", fromRow, fromColumn))
		} else {
			position.SetText(fmt.Sprintf("[red]From[white] Row: [yellow]%d[white], Column: [yellow]%d[white] - [red]To[white] Row: [yellow]%d[white], To Column: [yellow]%d ", fromRow, fromColumn, toRow, toColumn))
		}
	}

	textArea.SetMovedFunc(updateInfos)
	textArea.SetText(b.GetContents(), true)
	updateInfos()

	mainView := tview.NewGrid().
		SetRows(0, 1).
		AddItem(textArea, 0, 0, 1, 2, 0, 0, true).
		AddItem(helpInfo, 1, 0, 1, 1, 0, 0, false).
		AddItem(position, 1, 1, 1, 1, 0, 0, false)

	modal := tview.NewModal().
				SetText("You have unsaved changes. Do you want to save them?").
				AddButtons([]string {"Yes", "No", "Cancel"}).
				SetDoneFunc(func(buttonIndex int, buttonLabel string) {
						if buttonIndex == 0 {
							b.SetContents(textArea.GetText())
							err := b.SaveBlob()
							if err != nil {
								panic(err);
							}
							app.Stop()
						} else if buttonIndex == 1 {
							app.Stop()
						} else {
							pages.SwitchToPage("main")
						}
				})

	pages.AddPage("main", mainView, true, true).
		AddPage("exit", modal, false, false)

	go autoSave(b, textArea)

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyCtrlS {
			b.SetContents(textArea.GetText());
			go func() {
				err := b.SaveBlob()
				if err != nil {
					panic(err);
				}
			}()
			return nil
		} else if event.Key() == tcell.KeyCtrlC {
			if b.GetContents() == textArea.GetText() {
				return event
			} else {
				pages.ShowPage("exit")
				return nil
			}
		}
		return event
	})

	if err := app.SetRoot(pages,
		true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}


func main() {
	// Define command-line flags
	location := flag.String("location", ".", "Location to store the blob file")

	// Parse command-line arguments
	flag.Parse()

	// name and blobtype are obtained from location string
	name, blobtype, path := util.UnrollLocation(*location)

	// Initialize the Blob
	var b blob.Blob

	err := b.InitBlob(name, blobtype, path)
	if err != nil {
		log.Fatalf("Failed to initialize the blob: %v", err)
	}

	runApp(&b)

	defer func() {
		if err := b.CloseBlob(); err != nil {
			log.Fatalf("Failed to close the blob: %v", err)
		}
	}()
}
