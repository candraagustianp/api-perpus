package model

type Book struct {
	ID       uint   `gorm:"primarykey"`
	Judul    string `json:"judul" gorm:"size:255"`
	Penulis  string `json:"penulis" gorm:"size:50"`
	Penerbit string `json:"penerbit" gorm:"size:50"`
	Tahun    int    `json:"tahun"`
	Isbn     string `json:"isbn" gorm:"size:16"`
	IdJenis  int    `json:"id_jenis"`
}

type Jenis struct {
	ID    uint   `gorm:"primarykey"`
	Jenis string `json:"jenis" gorm:"size:30"`
	Books []Book `gorm:"foreignKey:IdJenis"`
}
