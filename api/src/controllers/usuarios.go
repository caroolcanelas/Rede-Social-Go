package controllers

import (
	"api/src/banco"
	"api/src/models"
	"api/src/repositorios"
	"api/src/respostas"
	"encoding/json"

	"io"
	"net/http"
)

//Criar novo usuario
func CriarUsuario(w http.ResponseWriter, r *http.Request){
	corpoRequisicao, erro := io.ReadAll(r.Body) //pegar o corpo da req
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}								

	var usuario models.Usuario // variavel pra criar novo usuario

	//converter body para struct
	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	//conectar ao banco
	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	usuario.ID, erro = repositorio.Criar(usuario)

	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}


	respostas.JSON(w, http.StatusCreated, usuario)
}

//buscar todos usuarios
func BuscarUsuarios(w http.ResponseWriter, r *http.Request){
	
}

//buscar um usuario
func BuscarUsuario(w http.ResponseWriter, r *http.Request){
	
}

//atualizar dados de usuario
func AtualizarUsuario(w http.ResponseWriter, r *http.Request){
	
}

//deletar um usuario
func DeletarUsuario(w http.ResponseWriter, r *http.Request){
	
}