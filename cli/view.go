package cli

import (
	"fmt"
	"os"
	"log"
	"path/filepath"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/cliche-niche/CS455/blob"
	"github.com/cliche-niche/CS455/util"
)

func addNode(target *tview.TreeNode, path string) {
	files, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		node := tview.NewTreeNode(file.Name()).
			SetReference(filepath.Join(path, file.Name())).
			SetSelectable(true)
		if file.IsDir() {
			node.SetColor(tcell.ColorGreen)
		}
		target.AddChild(node)
	}
}

type View struct {
	textArea 	*tview.TextArea
	treeView	*tview.TreeView
	updateInfos	func()
	mainView	*tview.Grid
	nextB		*blob.Blob
}

func (mv *View) InitView(b *blob.Blob) {
	mv.textArea = tview.NewTextArea().
		SetPlaceholder("Enter text here...").SetClipboard(nil, nil)
	mv.textArea.SetTitle("Text Area").SetBorder(true)
	helpInfo := tview.NewTextView().
		SetText(" Press F1 for help, Ctrl-S to save, Ctrl-C to exit")

	position := tview.NewTextView().
		SetDynamicColors(true).
		SetTextAlign(tview.AlignRight)

	mv.updateInfos = func() {
		fromRow, fromColumn, toRow, toColumn := mv.textArea.GetCursor()
		if fromRow == toRow && fromColumn == toColumn {
			position.SetText(fmt.Sprintf("Row: [yellow]%d[white], Column: [yellow]%d ", fromRow, fromColumn))
		} else {
			position.SetText(fmt.Sprintf("[red]From[white] Row: [yellow]%d[white], Column: [yellow]%d[white] - [red]To[white] Row: [yellow]%d[white], To Column: [yellow]%d ", fromRow, fromColumn, toRow, toColumn))
		}
	}

	mv.textArea.SetMovedFunc(mv.updateInfos)
	mv.textArea.SetText(b.GetContents(), true)
	mv.updateInfos()

	mv.mainView = tview.NewGrid().
		SetRows(0, 1).
		AddItem(mv.textArea, 0, 0, 1, 2, 0, 0, true).
		AddItem(helpInfo, 1, 0, 1, 1, 0, 0, false).
		AddItem(position, 1, 1, 1, 1, 0, 0, false)
	
}

func (mv *View) InitViewDir(rootDir string, cli *Cli) {

	root := tview.NewTreeNode(rootDir).
		SetColor(tcell.ColorYellow)
	mv.treeView = tview.NewTreeView().
		SetRoot(root).
		SetCurrentNode(root)

	addNode(root, rootDir)
	mv.TreeSelectFunc(cli)

	mv.nextB = nil
	mv.textArea = nil
	helpInfo := tview.NewTextView().
		SetText(" Press F1 for help, Ctrl-S to save, Ctrl-C to exit")
	position := tview.NewTextView().
		SetDynamicColors(true).
		SetTextAlign(tview.AlignRight)
	
	mv.updateInfos = func() {
		fromRow, fromColumn, toRow, toColumn := 0, 0, 0, 0
		if mv.textArea != nil {
			fromRow, fromColumn, toRow, toColumn = mv.textArea.GetCursor()
		}

		if fromRow == toRow && fromColumn == toColumn {
			position.SetText(fmt.Sprintf("Row: [yellow]%d[white], Column: [yellow]%d ", fromRow, fromColumn))
		} else {
			position.SetText(fmt.Sprintf("[red]From[white] Row: [yellow]%d[white], Column: [yellow]%d[white] - [red]To[white] Row: [yellow]%d[white], To Column: [yellow]%d ", fromRow, fromColumn, toRow, toColumn))
		}
	}
	mv.updateInfos()

	mv.mainView = tview.NewGrid().
		SetRows(0, 1).
		SetColumns(-2, -3, -3, -1).
		AddItem(mv.treeView, 0, 0, 1, 1, 0, 0, true).
		AddItem(helpInfo, 1, 0, 1, 2, 0, 0, false).
		AddItem(position, 1, 2, 1, 2, 0, 0, false)
}

func (mv *View) NewTextArea(b *blob.Blob) {
	mv.textArea = tview.NewTextArea().
	SetPlaceholder("Enter text here...")
	mv.textArea.SetTitle("Text Area").SetBorder(true)
	
	mv.textArea.SetMovedFunc(mv.updateInfos)
	mv.textArea.SetText(b.GetContents(), true)
	mv.updateInfos()
	
	mv.mainView.AddItem(mv.textArea, 0, 1, 1, 3, 0, 0, false)
}

func (mv *View) TreeSelectFunc(cli *Cli) {
	mv.treeView.SetSelectedFunc(
		func(node *tview.TreeNode) {
			if node.GetColor() == tcell.ColorGreen {
				reference := node.GetReference()
				children := node.GetChildren()

				if len(children) == 0 {
					// Load and show files in this directory.
					path := reference.(string)
					addNode(node, path)
				} else {
					// Collapse if visible, expand if collapsed.
					node.SetExpanded(!node.IsExpanded())
				}
			} else {
				reference := node.GetReference()
				
				if reference == nil {
					return // Selecting the root node does nothing.
				}
				
				name, blobtype, path := util.UnrollLocation(reference.(string))
				var b blob.Blob
				err := b.InitBlob(name, blobtype, path)
				if err != nil {
					log.Fatalf("Failed to initialize the blob: %v", err)
				}

				mv.nextB = &b
				
				if mv.textArea == nil {
					mv.NewTextArea(&b)
					cli.b = &b
					
				} else {
					if cli.b.GetContents() != mv.textArea.GetText() {
						cli.pages.ShowPage("fileChange")
					} else {
						mv.mainView.RemoveItem(mv.textArea)
						mv.NewTextArea(&b)
						cli.b = &b
					}

				}
					
			}
	})
}

