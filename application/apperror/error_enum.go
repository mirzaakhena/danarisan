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
	PesertaSudahJoinUndangan       ErrorType = "ER1000 peserta sudah join undangan"
	PesertaSudahMenolakUndangan    ErrorType = "ER1000 peserta sudah menolak undangan"
	PesertaSudahDiundang           ErrorType = "ER1000 peserta sudah diundang"
	ArisanTidakDitemukan           ErrorType = "ER1000 arisan tidak ditemukan"
	PesertaTidakDitemukan          ErrorType = "ER1000 peserta tidak ditemukan"
	TagihanTidakDitemukan          ErrorType = "ER1000 tagihan tidak ditemukan"
	UndianTidakDitemukan           ErrorType = "ER1000 undian tidak ditemukan"
	PesertaSudahMenjadiAdmin       ErrorType = "ER1000 peserta sudah menjadi admin"
	PesertaBukanAdmin              ErrorType = "ER1000 peserta bukan admin"
	BalanceUnderZeroError          ErrorType = "ER1000 balance under zero error"
)
