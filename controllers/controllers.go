package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"personalidades/database"
	"personalidades/models"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func Health(w http.ResponseWriter, r *http.Request) {
	respondJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

func respondError(w http.ResponseWriter, status int, message string) {
	respondJSON(w, status, map[string]string{"erro": message})
}

func validarPersonalidade(p *models.Personalidade) string {
	if strings.TrimSpace(p.Nome) == "" {
		return "O campo 'nome' é obrigatório"
	}
	return ""
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	db := database.DB
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")

	if pageStr != "" || limitStr != "" {
		page := 1
		limit := 10

		if pageStr != "" {
			var err error
			page, err = strconv.Atoi(pageStr)
			if err != nil || page < 1 {
				respondError(w, http.StatusBadRequest, "Parâmetro 'page' inválido")
				return
			}
		}
		if limitStr != "" {
			var err error
			limit, err = strconv.Atoi(limitStr)
			if err != nil || limit < 1 {
				respondError(w, http.StatusBadRequest, "Parâmetro 'limit' inválido")
				return
			}
		}

		var total int64
		if err := database.DB.Model(&models.Personalidade{}).Count(&total).Error; err != nil {
			respondError(w, http.StatusInternalServerError, "Erro ao contar personalidades")
			return
		}
		w.Header().Set("X-Total-Count", strconv.FormatInt(total, 10))
		db = db.Offset((page - 1) * limit).Limit(limit)
	}

	var p []models.Personalidade
	if err := db.Find(&p).Error; err != nil {
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
	if msg := validarPersonalidade(&novaPersonalidade); msg != "" {
		respondError(w, http.StatusBadRequest, msg)
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
	if msg := validarPersonalidade(&personalidade); msg != "" {
		respondError(w, http.StatusBadRequest, msg)
		return
	}
	if err := database.DB.Save(&personalidade).Error; err != nil {
		respondError(w, http.StatusInternalServerError, "Erro ao atualizar personalidade")
		return
	}
	respondJSON(w, http.StatusOK, personalidade)
}
