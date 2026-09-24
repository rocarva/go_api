package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"personalidades/database"
	"personalidades/models"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

func respondError(w http.ResponseWriter, status int, message string) {
	respondJSON(w, status, map[string]string{"erro": message})
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	var p []models.Personalidade
	if err := database.DB.Find(&p).Error; err != nil {
		respondError(w, http.StatusInternalServerError, "Erro ao buscar personalidades")
		return
	}
	respondJSON(w, http.StatusOK, p)
}

func RetornaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var personalidade models.Personalidade
	result := database.DB.First(&personalidade, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		respondError(w, http.StatusNotFound, "Personalidade não encontrada")
		return
	}
	if result.Error != nil {
		respondError(w, http.StatusInternalServerError, "Erro ao buscar personalidade")
		return
	}
	respondJSON(w, http.StatusOK, personalidade)
}

func CriaUmaNovaPersonalidade(w http.ResponseWriter, r *http.Request) {
	var novaPersonalidade models.Personalidade
	if err := json.NewDecoder(r.Body).Decode(&novaPersonalidade); err != nil {
		respondError(w, http.StatusBadRequest, "Corpo da requisição inválido")
		return
	}
	if err := database.DB.Create(&novaPersonalidade).Error; err != nil {
		respondError(w, http.StatusInternalServerError, "Erro ao criar personalidade")
		return
	}
	respondJSON(w, http.StatusCreated, novaPersonalidade)
}

func DeletaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var personalidade models.Personalidade
	result := database.DB.First(&personalidade, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		respondError(w, http.StatusNotFound, "Personalidade não encontrada")
		return
	}
	if result.Error != nil {
		respondError(w, http.StatusInternalServerError, "Erro ao buscar personalidade")
		return
	}
	if err := database.DB.Delete(&personalidade).Error; err != nil {
		respondError(w, http.StatusInternalServerError, "Erro ao deletar personalidade")
		return
	}
	respondJSON(w, http.StatusOK, personalidade)
}

func EditaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var personalidade models.Personalidade
	result := database.DB.First(&personalidade, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		respondError(w, http.StatusNotFound, "Personalidade não encontrada")
		return
	}
	if result.Error != nil {
		respondError(w, http.StatusInternalServerError, "Erro ao buscar personalidade")
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&personalidade); err != nil {
		respondError(w, http.StatusBadRequest, "Corpo da requisição inválido")
		return
	}
	if err := database.DB.Save(&personalidade).Error; err != nil {
		respondError(w, http.StatusInternalServerError, "Erro ao atualizar personalidade")
		return
	}
	respondJSON(w, http.StatusOK, personalidade)
}
