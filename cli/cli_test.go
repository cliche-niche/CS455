package cli

import (
	"testing"
	"strings"

	"github.com/cliche-niche/CS455/blob"
	"github.com/stretchr/testify/assert"
	"github.com/gdamore/tcell/v2"
)

func TestInitCli(t *testing.T) {
	var cli Cli
	var b blob.Blob

	cli.InitCli(&b, "", false)

	assert.NotNil(t, cli.app)
	assert.NotNil(t, cli.pages)
	assert.NotNil(t, cli.b)
	assert.NotNil(t, cli.view)
	assert.NotNil(t, cli.withDir)
	assert.NotNil(t, cli.toggle)
}

func TestAddExitModal(t *testing.T) {
	var cli Cli
	var b blob.Blob

	cli.InitCli(&b, "", false)
	cli.AddExitModal()

	assert.True(t, cli.pages.HasPage("exit"))
}

func TestAppInputCapture(t *testing.T) {
	var cli Cli
	var b blob.Blob

	cli.InitCli(&b, "", false)
	cli.AppInputCapture()

	assert.NotNil(t, cli.app.GetInputCapture())
}

func TestFileChangeModal(t *testing.T) {
	var cli Cli
	var b blob.Blob

	cli.InitCli(&b, "", false)
	cli.FileChangeModal()

	assert.True(t, cli.pages.HasPage("fileChange"))
}

func TestAddHelp(t *testing.T) {
	var cli Cli
	var b blob.Blob

	cli.InitCli(&b, "", false)
	cli.AddHelp()

	assert.True(t, cli.pages.HasPage("help"))
}

func TestSelectStart(t *testing.T){
	// When cursor is at start
	var cli Cli
	var b blob.Blob
	text := "abc12\ndef34"

	b.InitBlob("", "", "")
	b.EditBlob(text)
	cli.InitCli(&b, "", false)
	
	cli.view.textArea.SetText(text, true)
	
	l, r := selectLineText(&cli)
	s := text[l : r]

	assert.True(t,  s == text[ : strings.Index(text, "\n")])
}

func TestSelectEnd(t *testing.T){
	// When cursor is at end
	var cli Cli
	var b blob.Blob
	text := "abc12\ndef34"

	b.InitBlob("", "", "")
	b.EditBlob(text)
	cli.InitCli(&b, "", false)
	
	cli.view.textArea.SetText(text, false)
	
	l, r := selectLineText(&cli)
	s := text[l : r]

	assert.True(t,  s == text[strings.Index(text, "\n") + 1 : ])
}