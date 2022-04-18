package tries

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInMemoryFileSystem(t *testing.T) {
	t.Parallel()

	imfs := NewInMemoryFileSystem()
	ls, err := imfs.LS("/")
	require.NoError(t, err)
	assert.Equal(t, []string{}, ls)

	ls, err = imfs.LS("/does/not/exist")
	require.Error(t, err)

	fileCContent := "i am file c"
	require.NoError(t, imfs.AddFile("/a/b/c", fileCContent))

	ls, err = imfs.LS("/a/b")
	require.NoError(t, err)
	assert.Equal(t, []string{"c"}, ls)

	content, err := imfs.ReadContentFromFile("/a/b/c")
	require.NoError(t, err)
	assert.Equal(t, fileCContent, content)

	_, err = imfs.ReadContentFromFile("a/b/c/d")
	require.Error(t, err)

	anotherFileContent := "some content"
	require.NoError(t, imfs.AddFile("/a/b/another", anotherFileContent))

	content, err = imfs.ReadContentFromFile("/a/b/another")
	require.NoError(t, err)
	assert.Equal(t, anotherFileContent, content)

	moreAnotherFileContent := " and more content"
	require.NoError(t, imfs.AppendToFile("/a/b/another", moreAnotherFileContent))

	content, err = imfs.ReadContentFromFile("/a/b/another")
	require.NoError(t, err)
	assert.Equal(t, anotherFileContent+moreAnotherFileContent, content)

	ls, err = imfs.LS("/a/b")
	require.NoError(t, err)
	assert.Equal(t, []string{"another", "c"}, ls)

	err = imfs.Mkdir("/a/b/d")
	require.NoError(t, err)

	ls, err = imfs.LS("/a/b")
	require.NoError(t, err)
	assert.Equal(t, []string{"another", "c", "d"}, ls)

	err = imfs.Mkdir("/a/b/c")
	require.Error(t, err)

	ls, err = imfs.LS("/a/b/another")
	require.NoError(t, err)
	assert.Equal(t, []string{"another"}, ls)

	content, err = imfs.ReadContentFromFile("/b/doesnotexist")
	require.Error(t, err)

	content, err = imfs.ReadContentFromFile("/")
	require.Error(t, err)

	err = imfs.AppendToFile("/", "")
	require.Error(t, err)
}

func TestSplitPath(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		inputPath      string
		expectedOutput []string
	}{
		{
			name:           "empty-path",
			inputPath:      "",
			expectedOutput: []string{},
		},
		{
			name:           "root-path",
			inputPath:      "/",
			expectedOutput: []string{},
		},
		{
			name:           "multiple-path",
			inputPath:      "/a/b/c/d/e",
			expectedOutput: []string{"a", "b", "c", "d", "e"},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			res := splitPath(tt.inputPath)

			assert.Equal(t, tt.expectedOutput, res)
		})
	}
}

func Print(f *FSNode, currentPath string) {
	for k, v := range f.directory {
		fmt.Println(fmt.Sprintf("%s/%s -> IsDir: %v", currentPath, k, v.isDirectory))
		Print(v, currentPath+"/"+k)
	}
}
