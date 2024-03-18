package repository

import (
	"database/sql"
	"strings"

	pruebagrpc "github.com/ticaso/pruebagrpc/api"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetAspirante(id int64) (*pruebagrpc.Aspirante, error) {
	var aspirante pruebagrpc.Aspirante
	var lenguajes string
	err := r.db.QueryRow("SELECT id, nombre, apellido, imagen_principal, imagen_secundaria, carrera, lenguajes_programacion FROM aspirante WHERE id = ?", id).
		Scan(&aspirante.Id, &aspirante.Nombre, &aspirante.Apellido, &aspirante.ImagenPrincipal, &aspirante.ImagenSecundaria, &aspirante.Carrera, &lenguajes)
	if err != nil {
		return nil, err
	}

	aspirante.LenguajesProgramacion = strings.Split(lenguajes, ",")
	return &aspirante, nil
}

func (r *Repository) GetAspirantes() ([]*pruebagrpc.Aspirante, error) {
	rows, err := r.db.Query("SELECT id, nombre, apellido, imagen_principal, imagen_secundaria, carrera, lenguajes_programacion FROM aspirante")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var aspirantes []*pruebagrpc.Aspirante
	for rows.Next() {
		var a pruebagrpc.Aspirante
		var lenguajes string
		if err := rows.Scan(&a.Id, &a.Nombre, &a.Apellido, &a.ImagenPrincipal, &a.ImagenSecundaria, &a.Carrera, &lenguajes); err != nil {
			return nil, err
		}
		a.LenguajesProgramacion = strings.Split(lenguajes, ",")
		aspirantes = append(aspirantes, &a)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return aspirantes, nil
}
