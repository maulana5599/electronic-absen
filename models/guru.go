package models

type Guru struct {
	Id         int
	Nip        string
	NamaGuru   string
	TahunMasuk string
}

type RequestGuru struct {
	Nip        string
	NamaGuru   string
	TahunMasuk string
}

func (Guru) TableName() string {
	return "guru"
}
