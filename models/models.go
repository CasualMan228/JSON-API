package models

type Plane struct {
	Id             int     `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	Regnum         string  `json:"regnum" gorm:"column:regnum"`
	PricePerWeek   float64 `json:"pricePerWeek" gorm:"column:pricePerWeek"` //USD
	CountryOfReg   string  `json:"countryOfReg" gorm:"column:countryOfReg"`
	TotalFly       int     `json:"totalFly" gorm:"column:totalFly"` //KM
	Name           string  `json:"name" gorm:"column:name"`
	Description    string  `json:"description" gorm:"column:description"`
	YearOfCreation int     `json:"yearOfCreation" gorm:"column:yearOfCreation"`
	TypeId         int     `json:"typeId" gorm:"column:typeId"`
	CategoryId     int     `json:"categoryId" gorm:"column:categoryId"`
}
type Type struct {
	Id   int    `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	Name string `json:"name" gorm:"column:name"`
}

type Category struct {
	Id   int    `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	Name string `json:"name" gorm:"column:name"`
}

type Photo struct {
	Id      int    `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	Photo   []byte `json:"photo" gorm:"column:photo"` //base64 (https://base64.guru/converter/decode/image)
	PlaneId int    `json:"planeId" gorm:"column:planeId"`
}

type User struct {
	Id     int    `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	Name   string `json:"name" gorm:"column:name"`
	Pass   string `json:"pass" gorm:"column:pass"`
	RoleId int    `json:"roleId" gorm:"column:roleId"`
}

type Role struct {
	Id          int    `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	Name        string `json:"name" gorm:"column:name"`
	Description string `json:"description" gorm:"column:description"`
}

type Bill struct {
	Id         int     `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	TotalPrice float64 `json:"totalPrice" gorm:"column:totalPrice"`
	Weeks      int     `json:"weeks" gorm:"column:weeks"`
	UserId     int     `json:"userId" gorm:"column:userId"`
	PlaneId    int     `json:"planeId" gorm:"column:planeId"`
}

//чтобы GORM нашел нужный столбец нужно ему явно подсказывать, т.к. Camel нотацию он преобразовывает в Snake
//JSON автоматом кодирует бинарную информацию в формат BASE64
//кодировать изображение в BASE64 -> https://online-photo-converter.com/ru/image-to-base64
//old -> INSERT INTO dbo.Photos (photo) SELECT CAST('' AS XML).value('xs:base64Binary("")', 'VARBINARY(MAX)'); -> ssms преобразовать BASE64 в Binary Data

//new -> INSERT INTO dbo.Photos (photo, planeId)
//SELECT CAST('' AS XML).value('xs:base64Binary("")', 'VARBINARY(MAX)'), 1;

//DBCC CHECKIDENT ('dbo.TABLENAME', RESEED, 999); -> сброс PK ID
