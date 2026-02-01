package repositorios

import (
	"api/src/models"
	"database/sql"
)

type usuarios struct {
	db *sql.DB
}

func NovoRepositorioDeUsuarios(db *sql.DB) *usuarios{
	return &usuarios{db}
}

//insere usuario no banco de dados
func (repo usuarios) Criar(usuario models.Usuario) (uint64, error){

	//criar statement mysql
	statement, erro := repo.db.Prepare(
		"Insert into usuarios (nome, nick, email, senha) values (?, ?, ?, ?)")
	if erro != nil {
		return 0, erro
	}

	defer statement.Close()

	//execuar o statement
	insercao, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if erro != nil {
		return 0, erro
	}

	//acompanhar id do usuario
	idInserido, erro := insercao.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(idInserido), nil 
}
