# GLEIF XML File Parser
The Global Legal Entity Identifier Foundation (GLEIF) publishes three separate ‘Concatenated Files’ daily. The Concatenated Files include Legal Entity Identifiers (LEIs) and related LEI reference data published by the LEI issuing organizations.

The GLEIF Concatenated Files, which are available for download, include specific information on LEI records based on the relevant Common Data File (CDF) format. It contains Level 1 data, i.e. LEI records and related reference data that provide information on who is who based on the LEI-CDF format.

## Problem Statement
- Data files are updated daily.
- Files are concatenated daily as XML files and they are available to download as compressed zip files.
- One of the files is about 3.8GB of size and parsing it using a language like Python takes a long time.
- XML is an extremely difficult format to ingest into databases or do analysis on. 

## Solution
- Just by running the executable file provided, it will download the latest zip files, unzip it, parse it and will save to the current working directory as CSV files.
- The solution is written entirely in GO and due to parallel processing of data, in a machine with 8 logical threads, it will process and write the files in around 3 minutes.