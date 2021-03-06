// AUTOGENERATED - DO NOT EDIT
// Generated by go-windows.

// Package remotedifferentialcompression implements the Windows.Win32.RemoteDifferentialCompression namespace.
package remotedifferentialcompression

type RdcLibrary struct {
}

type RdcGeneratorParameters struct {
}

type RdcGeneratorFilterMaxParameters struct {
}

type RdcGenerator struct {
}

type RdcFileReader struct {
}

type RdcSignatureReader struct {
}

type RdcComparator struct {
}

type SimilarityReportProgress struct {
}

type SimilarityTableDumpState struct {
}

type SimilarityTraitsTable struct {
}

type SimilarityFileIdTable struct {
}

type Similarity struct {
}

type RdcSimilarityGenerator struct {
}

type FindSimilarResults struct {
}

type SimilarityTraitsMapping struct {
}

type SimilarityTraitsMappedView struct {
}

type RdcNeed struct {
	M_BLOCKTYPE   int
	M_FILEOFFSET  int
	M_BLOCKLENGTH int
}

type RdcBufferPointer struct {
	M_SIZE int
	M_USED int
	M_DATA int
}

type RdcNeedPointer struct {
	M_SIZE int
	M_USED int
	M_DATA int
}

type RdcSignature struct {
	M_SIGNATURE   int
	M_BLOCKLENGTH int
}

type RdcSignaturePointer struct {
	M_SIZE int
	M_USED int
	M_DATA int
}

type SimilarityMappedViewInfo struct {
	M_DATA   int
	M_LENGTH int
}

type SimilarityData struct {
	M_DATA int
}

type FindSimilarFileIndexResults struct {
	M_FILEINDEX  int
	M_MATCHCOUNT int
}

type SimilarityDumpData struct {
	M_FILEINDEX int
	M_DATA      int
}

type SimilarityFileId struct {
	M_FILEID int
}
