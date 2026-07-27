package code

import (
	"fmt"
	"os"
)

func GetPathSize(path string, human bool) (string, error) {
	fileInfo, err := os.Lstat(path)
	if err != nil {
		return err.Error(), nil
	}
	if !fileInfo.Mode().IsRegular() {
		return fmt.Sprintf("Путь '%s' не указывает на файл", path), nil
	}

	return formatSize(fileInfo.Size(), human), nil
}

func formatSize(size int64, human bool) string {
	if !human {
		return fmt.Sprintf("%dB", size)
	}
	units := []string{"B", "KB", "MB", "GB", "TB"}
	i := 0
	for size >= 1024 && i < len(units)-1 {
		size /= 1024
		i++
	}
	return fmt.Sprintf("%d%s", size, units[i])
}
