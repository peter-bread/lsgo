package fileinfo

import (
	"fmt"
	"os"
	"os/user"
	"strings"
	"syscall"
)

// type FileInfo struct{}

// TODO: put all this in my own struct so it can be passed around more easily
// TODO: clean up code too, add error handling

func getInfo(filepath string) (entry string) {
	fileInfo, err := os.Stat(filepath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	name := fileInfo.Name()
	mode := fileInfo.Mode()
	size := fileInfo.Size()
	modTime := fileInfo.ModTime().Format("_2 Jan 15:04")

	var UID int
	var GID int
	var hardlinks int
	var owner *user.User
	var group *user.Group

	if stat, ok := fileInfo.Sys().(*syscall.Stat_t); ok {
		UID = int(stat.Uid)
		GID = int(stat.Gid)
		owner, _ = user.LookupId(fmt.Sprint(UID))
		group, _ = user.LookupGroupId(fmt.Sprint(GID))
		hardlinks = int(stat.Nlink)
	}

	entry = fmt.Sprintf("%-11s %2d %10s %6s %4s %12s %s\n",
		mode, hardlinks, owner.Username, group.Name, fmt.Sprintf("%8d", size), modTime, name)

	// entry = fmt.Sprintf("%s %d %s  %s  %d %s %s\n", mode, hardlinks, owner.Username, group.Name, size, modTime, name)
	return
}

func ls(filepath string) {
	if !strings.HasSuffix(filepath, string(os.PathSeparator)) {
		filepath += string(os.PathSeparator)
	}

	files, _ := os.ReadDir(filepath)

	for _, file := range files {
		entry := getInfo(filepath + file.Name())
		fmt.Print(entry)
	}
}

func Exec() {
	filepath := "/Users/petersheehan/Developer/peter-bread/lsgo"
	ls(filepath)
}
