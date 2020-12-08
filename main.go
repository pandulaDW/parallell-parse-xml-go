package main

func main() {
	rrModel := GliefModel{prefix: "rr", category: "Relationship", recordsPerRoutine: 1000}
	rrModel.archiveFileName = "data/20201203-gleif-concatenated-file-rr.xml.5fc8c1302bde7.zip"
	rrModel.GZipFileName = "data/file.zip"
	rrModel.csvFileName = "rr.csv"
	concurrentProcessing(rrModel, ZipFileRead, CSVFileWrite)
}
