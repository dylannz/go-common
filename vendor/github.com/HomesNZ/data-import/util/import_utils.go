package util
import (
	"encoding/json"
)

type UploadResult struct {
	TotalRecords       int
	ExistingRecords    int
	RecordsToBeCreated int
	NewRecordsCreated  int
}


func ToJSON(rec interface{}) ([]byte, error) {
	b, err := json.MarshalIndent(rec, "", "    ")
	if err != nil {
		return nil, err
	}
	return b, err
}

