package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

const file = "inputs.txt"

type Folder struct {
	Name       string
	Size       int
	SubFolders []*Folder
	Parent     *Folder
}

func (folder *Folder) SumOfSizes() int {
	totalSize := 0

	if folder.Size <= 100000 {
		totalSize += folder.Size
	}

	for _, subfolder := range folder.SubFolders {
		totalSize += subfolder.SumOfSizes()
	}

	return totalSize
}

func main() {
	// Open the file
	file, _ := os.Open(file)
	scanner := bufio.NewScanner(file)
	rootFolder := &Folder{Name: "/", Size: 0, SubFolders: nil, Parent: nil}
	var currentFolder *Folder
	sizes := make(map[string]int)

	for scanner.Scan() {
		line := scanner.Text()
		lineSplit := strings.Split(line, " ")
		if lineSplit[1] == "cd" {
			if lineSplit[2] == "/" {
				currentFolder = rootFolder
			} else if lineSplit[2] == ".." {
				currentFolder.Parent.Size += currentFolder.Size
				currentFolder = currentFolder.Parent
				sizes = UpdateSize(sizes, currentFolder.Name, currentFolder.Size)
			} else {
				newFolder := Folder{
					Name:       lineSplit[2],
					Size:       0,
					SubFolders: nil,
					Parent:     currentFolder,
				}
				currentFolder.SubFolders = append(currentFolder.SubFolders, &newFolder)
				currentFolder = &newFolder
			}
		} else if lineSplit[0] != "dir" && lineSplit[1] != "ls" {
			filesize, _ := strconv.Atoi(lineSplit[0])
			currentFolder.Size += filesize
		}
	}

	//If not in root folder
	for currentFolder.Name != "/" {
		currentFolder.Parent.Size += currentFolder.Size
		currentFolder = currentFolder.Parent
		sizes = UpdateSize(sizes, currentFolder.Name, currentFolder.Size)
	}

	log.Println("FolderSize of root folder:", rootFolder.Size)
	log.Println("Sum of sizes smaller than 100000:", rootFolder.SumOfSizes())
	neededSpace := 30000000 - (70000000 - rootFolder.Size)
	log.Println("Needed space:", neededSpace)
	log.Println("Foldersize to delete:", part2(neededSpace, sizes))
}

func part2(neededSpace int, sizes map[string]int) int {
	var smallest int
	for _, size := range sizes {
		if size >= neededSpace && (smallest == 0 || size < smallest) {
			smallest = size
		}
	}
	return smallest
}

func UpdateSize(sizes map[string]int, newName string, newSize int) map[string]int {
	foundSize, found := sizes[newName]
	if !found || foundSize < newSize {
		sizes[newName] = newSize
	}
	return sizes
}
