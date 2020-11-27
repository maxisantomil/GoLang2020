package chat

import (
	"github.com/jmoiron/sqlx"
	"github.com/maxisantomil/GoLang2020.git/internal/config"
)

// Vino ...
type Vino struct {
	ID     int64
	Name   string
	tipo   string
	año    int
	precio int
}

// Service ...
type Service interface {
	AddVino(Vino) (int64, error)
	FindByID(int64) *Vino
	FindAll() []*Vino
	UpdateVino(Vino, int64)
	DeleteVino(int64) *Vino
}

// devuelve algo que es privado
type service struct {
	db   *sqlx.DB
	conf *config.Config
}

// New ...
func New(db *sqlx.DB, c *config.Config) (Service, error) {
	return service{db, c}, nil
}

// agregar ...
func (s service) AddVino(v Vino) (int64, error) {
	query := `INSERT INTO Vino (Name, tipo, año, precio) VALUES (?,?,?,?)`
	return s.db.MustExec(query, v.Name, v.tipo, v.año, v.precio).LastInsertId()
}

// modificar ...
func (s service) UpdateVino(v Vino, id int64) {
	query := `UPDATE Vino SET Name = ?, tipo = ?, año = ?, precio = ? WHERE id = ?`
	s.db.MustExec(query, v.Name, v.tipo, v.año, v.precio, id)
}

// busca por id ...
func (s service) FindByID(ID int64) *Vino {
	vino := &Vino{}
	query := `SELECT * FROM Vino WHERE id = ?`
	err := s.db.Get(vino, query, ID)
	if err != nil {
		return nil
	}
	return vino
}

func (s service) FindAll() []*Vino {
	var list []*Vino
	if err := s.db.Select(&list, "SELECT * FROM vino"); err != nil {
		panic(err)
	}
	return list
}

// eliminar ...
func (s service) DeleteVino(id int64) *Vino {
	vino := s.FindByID(id)
	query := `DELETE FROM vino WHERE id = ?`
	s.db.MustExec(query, id).RowsAffected()
	return vino
}
