package main

import (
	"flag"
	"io"
	"os"
	"strings"

	"github.com/Jordation/j2s"
	"github.com/sirupsen/logrus"
)

var (
	rootTypeName, srcPath, outPath, lang string
)

func init() {
	flag.StringVar(&rootTypeName, "tn", "Unnamed", "what the (t)ype(n)ame for the top level datatype should be called")

	flag.StringVar(&srcPath, "i", "", "the input json or xml file to build the type from")

	flag.StringVar(&outPath, "o", "output.go", "the output file for the generated types")

	flag.StringVar(&lang, "l", "go", "the output (l)anguages you want to generate types for")
}

func main() {
	flag.Parse()

	if srcPath == "" {
		flag.PrintDefaults()
		logrus.Fatalf("j2s !h %d", len(os.Args))
	}

	srcFile, err := os.OpenFile(srcPath, 0, 0644)
	if err != nil {
		logrus.WithError(err).
			Fatalf("failed to get src file @ %s", srcPath)
	}

	setsOfTypes, err := j2s.
		NewTranslator(rootTypeName, srcFile).
		Translate()

	if err != nil {
		logrus.WithError(err).Fatal("failed to build sets")
	}

	srcFile.Close()

	typeWriters := make(map[string]io.WriterTo, 2)
	for _, t := range strings.Split(lang, ",") {
		switch t {
		case "go":
			typeWriters["go"] = j2s.NewGoTypeWriter(setsOfTypes)
		case "ts":
			typeWriters["ts"] = j2s.NewTsTypeWriter(setsOfTypes)
		}
	}

	for ext, typeWriter := range typeWriters {
		path := getFilePath(outPath, ext)

		outputFile, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0644)
		if err != nil {
			logrus.WithError(err).
				Fatalf("failed to get output file @ %s", path)
		}

		if _, err := typeWriter.WriteTo(outputFile); err != nil {
			logrus.WithError(err).
				Fatal("failed to write to output file")
		}

		outputFile.Close()
	}
}

func getFilePath(name, wantExt string) string {
	if strings.HasSuffix(name, wantExt) {
		return name
	} else {
		return strings.Split(name, ".")[0] + "." + wantExt
	}
}
