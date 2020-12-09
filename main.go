package main

func main() {
	leiModel := createLEIModel()
	leiModel.zipFileName = "data/20201202-gleif-concatenated-file-lei2.xml.5fc7579cab4ee.zip"
	concurrentProcessing(*leiModel, XMLWriteAndRead, CSVFileWrite)
}
