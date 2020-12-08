package main

// func test() {
// 	start := time.Now()
// 	bufferedReader, err := createBufferedReader("data/20201202-gleif-concatenated-file-lei2.xml.5fc7579cab4ee.zip")
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	content, _ := ioutil.ReadAll(bufferedReader)
// 	records := LEIData{}

// 	if err := xml.Unmarshal(content, &records); err != nil {
// 		log.Println(err)
// 	}

// 	// manually run the gc to clean up memory
// 	content = nil
// 	runtime.GC()

// 	zipWriter, zipFile := createGzipWriter("file.zip", "testfile.csv")
// 	bufferedWriter := bufio.NewWriter(zipWriter)

// 	for i, record := range records.LEIRecords {
// 		recordStr := fmt.Sprintf("%v -> %v\n", record.LEI, record.Entity.LegalName)
// 		bufferedWriter.WriteString(recordStr)

// 		if i%50 == 0 {
// 			bufferedWriter.Flush()
// 		}
// 	}

// 	// write any remaining data to the underlying writer
// 	bufferedWriter.Flush()

// 	zipFile.Close()
// 	zipWriter.Close()

// 	fmt.Println(len(records.LEIRecords), " records were parsed")
// 	fmt.Println("time taken: ", time.Since(start))
// }
