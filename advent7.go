package main

import (
	"bufio"
	"log"
	"math"
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

func (f *Folder) CalculateTotalSize() int {
	totalSize := f.Size

	for _, subFolder := range f.SubFolders {
		totalSize += subFolder.CalculateTotalSize()
	}
	return totalSize
}

func (f *Folder) TraverseAndCalculateSizes() []int {
	var sizes []int
	sizes = append(sizes, f.CalculateTotalSize())
	for _, subFolder := range f.SubFolders {
		subSizes := subFolder.TraverseAndCalculateSizes()
		sizes = append(sizes, subSizes...)
	}
	return sizes
}

func main() {
	// Open the file
	file, _ := os.Open(file)
	scanner := bufio.NewScanner(file)
	rootFolder := &Folder{Name: "/", Size: 0, SubFolders: nil, Parent: nil}
	var currentFolder *Folder

	for scanner.Scan() {
		line := scanner.Text()
		lineSplit := strings.Split(line, " ")
		if lineSplit[1] == "cd" {
			if lineSplit[2] == "/" {
				currentFolder = rootFolder
			} else if lineSplit[2] == ".." {
				currentFolder = currentFolder.Parent
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

	slices := rootFolder.TraverseAndCalculateSizes()
	log.Println("FolderSize of root folder:", rootFolder.CalculateTotalSize())
	log.Println("Sum of sizes smaller than 100000:", part1(slices))
	neededSpace := 30000000 - (70000000 - rootFolder.CalculateTotalSize())
	log.Println("Needed space:", neededSpace)
	log.Println("All folders", part2(neededSpace, slices))
}

func part1(slices []int) int {
	totalSize := 0

	for _, size := range slices {
		if size <= 100000 {
			totalSize += size
		}
	}
	return totalSize
}

func part2(requiredSpace int, slices []int) int {
	smallest := math.MaxInt

	for _, value := range slices {
		if value > requiredSpace && value < smallest {
			smallest = value
		}
	}

	return smallest
}
