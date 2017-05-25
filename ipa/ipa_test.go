package ipa

import (
	"testing"

	"os"
	"path/filepath"

	"github.com/bitrise-io/go-utils/fileutil"
	"github.com/bitrise-io/go-utils/pathutil"
	"github.com/stretchr/testify/require"
)

func TestFindFileInPayloadDir(t *testing.T) {
	t.Log("app name == ipa name")
	{
		tmpDir, err := pathutil.NormalizedOSTempDirPath("__ipa__")
		require.NoError(t, err)

		payloadDir := filepath.Join(tmpDir, "Payload")
		appDir := filepath.Join(payloadDir, "test.app")
		require.NoError(t, os.MkdirAll(appDir, 0700))

		infoPlistPth := filepath.Join(appDir, "Info.plist")
		require.NoError(t, fileutil.WriteStringToFile(infoPlistPth, ""))

		pth, err := findFileInPayloadDir(payloadDir, "test", "Info.plist")
		require.NoError(t, err)
		require.Equal(t, infoPlistPth, pth)
	}

	t.Log("app name != ipa name")
	{
		tmpDir, err := pathutil.NormalizedOSTempDirPath("__ipa__")
		require.NoError(t, err)

		payloadDir := filepath.Join(tmpDir, "Payload")
		appDir := filepath.Join(payloadDir, "test.app")
		require.NoError(t, os.MkdirAll(appDir, 0700))

		infoPlistPth := filepath.Join(appDir, "Info.plist")
		require.NoError(t, fileutil.WriteStringToFile(infoPlistPth, ""))

		pth, err := findFileInPayloadDir(payloadDir, "not_test", "Info.plist")
		require.NoError(t, err)
		require.Equal(t, infoPlistPth, pth)
	}
}