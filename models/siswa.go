package models

type Siswa struct {
	IdSiswa   int
	NIS       int
	NamaSiswa string
}

type SiswaDetail struct {
	IdSiswa   int
	NIS       int
	NamaSiswa string
	Abensi    Absensi `gorm:"foreignKey:nis;references:nis"`
}

func (Siswa) TableName() string {
	return "siswa"
}

func (SiswaDetail) TableName() string {
	return "siswa"
}
