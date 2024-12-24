package main

import (
	"os"

	"github.com/Jordation/j2s"
	"github.com/sirupsen/logrus"
)

func main() {
	if len(os.Args) != 4 {
		logrus.Fatal("Usage: j2s <root struct name> <json source file> <output file>")
	}

	rootName := os.Args[1]
	srcFileName := os.Args[2]
	outputFileName := os.Args[3]

	outputFile, err := os.OpenFile(outputFileName, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		logrus.WithError(err).Fatal("failed to get output file")
	}
	defer outputFile.Close()

	srcFile, err := os.OpenFile(srcFileName, 0, 0644)
	if err != nil {
		logrus.WithError(err).Fatal("failed to get output file")
	}
	defer srcFile.Close()

	result, err := j2s.NewTranslator(rootName, srcFile).Translate()
	if err != nil {
		logrus.WithError(err).Fatal("failed to build sets")
	}

	types := j2s.BuildTypeRepresentations(result.Sets, j2s.DefaultRenamer)
	if _, err := types.WriteTo(outputFile); err != nil {
		logrus.WithError(err).Fatal("failed to write to file")
	}

}
