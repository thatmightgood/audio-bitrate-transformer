package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	err := filepath.Walk("./input", func(path string, info os.FileInfo, err error) error {
		name := info.Name()
		cmd := exec.Command("powershell", "./ffmpeg.exe", "-i", "'./input/"+name+"'", "-b", "128k", "'./out/"+name+"'")
		if err := cmd.Run(); err != nil {
			fmt.Println("Error: ", err)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

}
