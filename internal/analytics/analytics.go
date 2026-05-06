package analytics

import (
	"sort"

	"gointel/internal/models"
)

func GetLargestFiles(files []models.FileInfo, limit int) []models.FileInfo {

	sort.Slice(files, func(i, j int) bool {
		return files[i].Size > files[j].Size
	})

	if len(files) < limit {
		return files
	}

	return files[:limit]
}