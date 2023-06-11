package helpers

import (
	"os"

	"github.com/hebecoding/digital-dash-commons/utils"
	"github.com/pkg/errors"
)

func ReadInJSONTestDataFile(logger utils.LoggerInterface, path string) (*os.File, error) {
	// read in test data from file
	logger.Info("Reading in test data from file")
	file, err := os.Open(path)
	if err != nil {
		logger.Error("error opening file in path: ", path)
		return nil, errors.Wrap(err, "error opening test data file")
	}
	logger.Info("Successfully read in test data from file")
	return file, nil
}
