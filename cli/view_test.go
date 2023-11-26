package cli

import (
	"testing"

	"github.com/cliche-niche/CS455/blob"
	"github.com/stretchr/testify/assert"
)

func TestInitView(t *testing.T) {
	var mv View
	var b blob.Blob

	b.EditBlob("abc");
	mv.InitView(&b);

	assert.NotNil(t, mv.textArea);
	assert.True(t, "abc" == mv.textArea.GetText());
	assert.NotNil(t, mv.updateInfos);
	assert.NotNil(t, mv.mainView);
}

func TestInitViewDir(t *testing.T) {
	mv := &View{}
	cli := &Cli{}
	rootDir := "../"

	mv.InitViewDir(rootDir, cli)

	assert.Nil(t, mv.textArea);
	assert.NotNil(t, mv.updateInfos);
	assert.NotNil(t, mv.mainView);
	assert.NotNil(t, mv.treeView);
	assert.Nil(t, mv.nextB);
}

func TestNewTextArea(t *testing.T) {
	var mv View
	var b blob.Blob

	b.EditBlob("abc")
	mv.InitView(&b);

	assert.NotNil(t, mv.textArea);
	assert.True(t, "abc" == mv.textArea.GetText());

	var nb blob.Blob

	nb.EditBlob("bcd")
	mv.NewTextArea(&nb);

	assert.NotNil(t, mv.textArea);
	assert.True(t, "bcd" == mv.textArea.GetText());
}