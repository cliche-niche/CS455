// Demo code for the TreeView primitive.
package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// Show a navigable tree view of the current directory.
func main() {
	app := tview.NewApplication()

	rootDir := "./util"
	root := tview.NewTreeNode(rootDir).
		SetColor(tcell.ColorYellow)
	tree := tview.NewTreeView().
		SetRoot(root).
		SetCurrentNode(root)

	// A helper function which adds the files and directories of the given path
	// to the given target node.
	add := func(target *tview.TreeNode, path string) {
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

	// Add the current directory to the root node.
	add(root, rootDir)

	// If a directory was selected, open it.
	tree.SetSelectedFunc(func(node *tview.TreeNode) {
		if node.GetColor() == tcell.ColorGreen {
			reference := node.GetReference()
			if reference == nil {
				return // Selecting the root node does nothing.
			}
			children := node.GetChildren()
			if len(children) == 0 {
				// Load and show files in this directory.
				path := reference.(string)
				add(node, path)
			} else {
				// Collapse if visible, expand if collapsed.
				node.SetExpanded(!node.IsExpanded())
			}
		}
	})


	textArea := tview.NewTextArea().
		SetPlaceholder("Enter text here...")
	textArea.SetTitle("Text Area").SetBorder(true)
	helpInfo := tview.NewTextView().
		SetText(" Press F1 for help, press Ctrl-C to exit")
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
	updateInfos()

	mainView := tview.NewGrid().
		SetRows(0, 1).
		SetColumns(-2, -3, -3, -1).
		AddItem(tree, 0, 0, 1, 1, 0, 0, true).
		AddItem(textArea, 0, 1, 1, 3, 0, 0, false).
		AddItem(helpInfo, 1, 0, 1, 2, 0, 0, false).
		AddItem(position, 1, 2, 1, 2, 0, 0, false)

	pages.AddPage("main", mainView, true, true)

	if err := app.SetRoot(pages, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
