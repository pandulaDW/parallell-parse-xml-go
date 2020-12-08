package main

// GliefModel includes information needed for Glief file processing
// Based on the processing stage, (testFile, testRar, testAWS) different
// properties will be used
type GliefModel struct {
	prefix            string
	category          string
	recordsPerRoutine int
	xmlFileName       string
	archiveFileName   string
	GZipFileName      string
	csvFileName       string
	url               string
}

// GliefFileOps includes operations that should be performed for Glief files
type GliefFileOps interface {
	create() GliefModel
	unmarshal(content *string) *string
}

// GliefRelationship implement GliefFileOps methods
type GliefRelationship struct {
	GliefModel
}

func create() {

}
