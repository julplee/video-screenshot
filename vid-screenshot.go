package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	var files []string

	root := "./"
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() || filepath.Ext(path) != ".mp4" {
			return nil
		}

		files = append(files, path)
		return nil
	})

	if err != nil {
		panic(err)
	}

	for _, file := range files {
		createScreenshot(file)
	}
}

func createScreenshot(inputVideoFilepath string) {

	// use ffmpeg to generate a jpg file
	width := 640
	height := 360
	cmd := exec.Command("ffmpeg.exe", "-i", inputVideoFilepath, "-ss", "00:01:45.000", "-vframes", "1", "-s", fmt.Sprintf("%dx%d", width, height), "-f", "singlejpeg", "-")

	var buffer bytes.Buffer
	cmd.Stdout = &buffer
	if cmd.Run() != nil {
		panic("could not generate frame")
	}

	// write the whole body at once
	err := ioutil.WriteFile(filepath.Base(inputVideoFilepath)+".jpg", buffer.Bytes(), 0644)
	check(err)
}
