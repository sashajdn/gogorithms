package tries

import (
	"fmt"
	"sort"
	"strings"
)

var (
	ErrDirectoryDoesNotExist = fmt.Errorf("directory does not exist")
	ErrFileDoesNotExist      = fmt.Errorf("file does not exist")
	ErrDirectoryIsNotAFile   = fmt.Errorf("directory is not a file")
	ErrDirectoryIsAFile      = fmt.Errorf("directory is a file")
	ErrFileIsNotADirectory   = fmt.Errorf("file is not a directory")
	ErrInvalidFilePath       = fmt.Errorf("invalid file path")
	ErrAlreadyExists         = fmt.Errorf("already exists")
)

func NewInMemoryFileSystem() *InMemoryFileSystem {
	return &InMemoryFileSystem{
		root: &FSNode{
			directory:   map[string]*FSNode{},
			isDirectory: true,
		},
	}
}

type InMemoryFileSystem struct {
	root *FSNode
}

type FSNode struct {
	directory   map[string]*FSNode
	isDirectory bool
	content     string
}

func (i *InMemoryFileSystem) LS(path string) ([]string, error) {
	splits := splitPath(path)

	var current = i.root
	for _, split := range splits {
		var ok bool
		current, ok = current.directory[split]
		if !ok {
			return nil, fmt.Errorf("%w: %s", ErrDirectoryDoesNotExist, split)
		}
	}

	if !current.isDirectory {
		return splits[len(splits)-1:], nil
	}

	var directory = make([]string, 0, len(current.directory))
	for k := range current.directory {
		directory = append(directory, k)
	}

	sort.Strings(directory)
	return directory, nil
}

func (i *InMemoryFileSystem) ReadContentFromFile(filepath string) (string, error) {
	splits := splitPath(filepath)

	var current = i.root
	for _, split := range splits {
		if !current.isDirectory {
			return "", fmt.Errorf("%s: %w", split, ErrFileIsNotADirectory)
		}

		var ok bool
		current, ok = current.directory[split]
		if !ok {
			return "", fmt.Errorf("%s: %w", split, ErrFileDoesNotExist)
		}
	}

	if current.isDirectory {
		return "", fmt.Errorf("%s: %w", filepath, ErrDirectoryIsNotAFile)
	}

	return current.content, nil
}

func (i *InMemoryFileSystem) AppendToFile(filepath, content string) error {
	splits := splitPath(filepath)

	var current = i.root
	for _, split := range splits {
		if !current.isDirectory {
			return fmt.Errorf("%s: %w", split, ErrFileIsNotADirectory)
		}

		var ok bool
		current, ok = current.directory[split]
		if !ok {
			return fmt.Errorf("%s: %w", split, ErrFileDoesNotExist)
		}
	}

	if current.isDirectory {
		return fmt.Errorf("%s: %w", filepath, ErrDirectoryIsNotAFile)
	}

	current.content += content
	return nil
}

func (i *InMemoryFileSystem) AddFile(filepath, content string) error {
	splits := splitPath(filepath)
	if len(splits) == 0 {
		return fmt.Errorf("%v: %w", filepath, ErrInvalidFilePath)
	}

	head, tail := splits[:len(splits)-1], splits[len(splits)-1]

	var current = i.root
	for _, split := range head {
		if _, ok := current.directory[split]; !ok {
			current.directory[split] = &FSNode{
				directory:   map[string]*FSNode{},
				isDirectory: true,
			}
		}

		current, _ = current.directory[split]
	}

	if _, ok := current.directory[tail]; ok {
		return fmt.Errorf("file %s: %w", filepath, ErrAlreadyExists)
	}

	current.directory[tail] = &FSNode{
		directory:   map[string]*FSNode{},
		isDirectory: false,
		content:     content,
	}

	return nil
}

func (i *InMemoryFileSystem) Mkdir(directoryPath string) error {
	splits := splitPath(directoryPath)

	var current = i.root
	for _, split := range splits {
		if _, ok := current.directory[split]; !ok {
			current.directory[split] = &FSNode{
				directory:   map[string]*FSNode{},
				isDirectory: true,
			}
		}

		current, _ = current.directory[split]
		if !current.isDirectory {
			return fmt.Errorf("directory %s is a file: %w", split, ErrDirectoryIsAFile)
		}
	}

	return nil
}

func splitPath(path string) []string {
	if path == "" || path == "/" {
		return []string{}
	}

	var splits []string
	for _, split := range strings.Split(path, "/") {
		if split == "" {
			continue
		}

		splits = append(splits, split)
	}

	return splits
}
