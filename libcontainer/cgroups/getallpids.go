package cgroups

import (
	"io/fs"
	"path/filepath"
)

// GetAllPids returns all pids from the cgroup identified by path, and all its
// sub-cgroups.
func GetAllPids(path string) ([]int, error) {
	var pids []int
	/*通过func遍历path下所有文件及目录*/
	err := filepath.WalkDir(path, func(p string, d fs.DirEntry, iErr error) error {
		if iErr != nil {
			return iErr
		}
		if !d.IsDir() {
			/*跳过非目录*/
			return nil
		}
		cPids, err := readProcsFile(p)
		if err != nil {
			return err
		}
		pids = append(pids, cPids...)
		return nil
	})
	return pids, err
}
