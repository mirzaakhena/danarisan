package apperror

const (
	FailUnmarshalResponseBodyError ErrorType = "ER1000 Fail to unmarshal response body"  // used by controller
	UnrecognizedEnum               ErrorType = "ER1001 %s is not recognized %s enum"     // used by enum
	DatabaseNotFoundInContextError ErrorType = "ER1002 Database is not found in context" // used by repoimpl
	ArisanSudahDimulai             ErrorType = "ER1000 arisan sudah dimulai"
	ArisanSudahSelesai             ErrorType = "ER1000 arisan sudah selesai"
	PesertaArisanMasihKurang       ErrorType = "ER1000 peserta arisan masih kurang"
	NominalHarusLebihBesarDariNol  ErrorType = "ER1000 nominal tidak boleh nol"
	SemuaPesertaSudahMenang        ErrorType = "ER1000 semua peserta sudah menang"
	ArisanSudahBerakhir            ErrorType = "ER1000 arisan sudah selesai"
)
