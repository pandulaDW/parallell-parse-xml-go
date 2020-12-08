package main

func main() {
	// rrModel := GliefModel{prefix: "rr", category: "Relationship", recordsPerRoutine: 1000}
	leiModel := createLEIModel()
	leiModel.zipFileName = "data/20201202-gleif-concatenated-file-lei2.xml.5fc7579cab4ee.zip"
	concurrentProcessing(*leiModel, ZipFileRead, CSVFileWrite)
}
