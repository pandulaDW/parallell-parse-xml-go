package main

func main() {
	rrModel := GliefModel{prefix: "rr", category: "Relationship", recordsPerRoutine: 2000}
	concurrentProcessing(rrModel, "XMLFileRead")
	// concurrentProcessing("rr", "Relationship")
}
