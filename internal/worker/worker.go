package worker

import (
	"log"
	"os"

	"gointel/internal/checksum"
	"gointel/internal/models"
)

func StartWorker(
	id int,
	jobs <-chan string,
	results chan<- models.FileInfo,
) {

	for path := range jobs {

		log.Printf("[Worker %d] Processing: %s\n", id, path)

		info, err := os.Stat(path)
		if err != nil {
			log.Println(err)
			continue
		}

		hash, err := checksum.GenerateSHA256(path)
		if err != nil {
			log.Println(err)
			continue
		}

		results <- models.FileInfo{
			Path:     path,
			Size:     info.Size(),
			Checksum: hash,
		}
	}
}