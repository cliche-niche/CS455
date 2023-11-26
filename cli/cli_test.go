package cli

import (
	"testing"

	"github.com/cliche-niche/CS455/blob"
	"github.com/stretchr/testify/assert"
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
