package main

func main() {
	repexModel := createReportingExceptionModel()
	repexModel.xmlFileName = "data/20210216-gleif-concatenated-file-repex.xml"
	repexModel.csvFileName = "data/repexCSV.csv"
	concurrentProcessing(*repexModel, XMLFileRead, CSVFileWrite)
}
