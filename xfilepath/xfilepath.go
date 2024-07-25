package xfilepath

import (
	"io/fs"
	"os"
	"path/filepath"
)

func WalkRenameFile(dir string, renameFn func(filePath string, info fs.FileInfo) string) error {
	return filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if newName := renameFn(path, info); newName != "" {
			return os.Rename(path, newName)
		}
		return nil
	})
}

//filepath.Walk(dir, func(filePath string, info fs.FileInfo, err error) error {
//	if strings.HasSuffix(filePath, ".go") {
//		filename := info.Name()
//		if strings.HasPrefix(filename, "g") {
//			fp := filepath.Dir(filePath)
//			newName := path.Join(fp, "x"+filename[1:])
//			return os.Rename(filePath, newName)
//		}
//	}
//	return nil
//})
