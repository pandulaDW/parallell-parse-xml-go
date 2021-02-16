package main

func main() {
	repexModel := createReportingExceptionModel()
	repexModel.xmlFileName = "data/20210216-gleif-concatenated-file-repex.xml"
	concurrentProcessing(*repexModel, XMLFileRead, CSVFileWrite)
}
