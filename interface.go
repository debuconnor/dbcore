package dbcore

type Dml interface {
	Select() error
	Insert() error
	Update() error
	Delete() error
}

type Ddl interface {
	CreateTable() error
	AlterTable() error
	DropTable() error
}
