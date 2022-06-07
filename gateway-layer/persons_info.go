package gateway_layer

import (
	persons_info_controllers "Backend/controllers/persons_info"
	"Backend/helpers"
	persons_info_models "Backend/models/persons_info"
	"Backend/types"
	"encoding/json"
	"github.com/jackc/pgx/v4"
	"go.uber.org/zap"
	"net/http"
)

func AddPersonInfo(conn *pgx.Conn) {
	http.HandleFunc("/create-person", func(w http.ResponseWriter, r *http.Request) {
		var request persons_info_models.CreatePersonRequest
		json.NewDecoder(r.Body).Decode(&request)

		isValidToken, _, role := helpers.ValidateToken(request.Token)

		if !isValidToken {
			w.WriteHeader(500)
			helpers.SetupResponse(&w)
			return
		}

		if !(role == types.Administrator) {
			w.WriteHeader(500)
			helpers.SetupResponse(&w)
			return
		}

		err := persons_info_controllers.CreatePerson(conn, request)

		if err != nil {
			zap.L().Error(err.Error())
			w.WriteHeader(500)
			helpers.SetupResponse(&w)
			return
		}
		w.WriteHeader(200)

		helpers.SetupResponse(&w)
	})

	http.HandleFunc("/update-person-info", func(w http.ResponseWriter, r *http.Request) {
		var request persons_info_models.UpdatePersonInfoRequest
		json.NewDecoder(r.Body).Decode(&request)

		isValidToken, login, role := helpers.ValidateToken(request.Token)

		if !isValidToken {
			w.WriteHeader(500)
			helpers.SetupResponse(&w)
			return
		}

		if !(role == types.Administrator || login != request.Login) {
			w.WriteHeader(500)
			helpers.SetupResponse(&w)
			return
		}

		err := persons_info_controllers.UpdatePersonInfo(conn, request)

		if err != nil {
			zap.L().Error(err.Error())
			w.WriteHeader(500)
			helpers.SetupResponse(&w)
			return
		}
		w.WriteHeader(200)

		helpers.SetupResponse(&w)
	})

	http.HandleFunc("/get-person-info", func(w http.ResponseWriter, r *http.Request) {
		var request persons_info_models.GetPersonInfoRequest
		json.NewDecoder(r.Body).Decode(&request)

		response, err := persons_info_controllers.GetPersonInfo(conn, request)

		if err != nil {
			zap.L().Error(err.Error())
			w.WriteHeader(500)
			helpers.SetupResponse(&w)
			return
		}

		json.NewEncoder(w).Encode(response)
		w.WriteHeader(200)

		helpers.SetupResponse(&w)
	})

	http.HandleFunc("/get-persons", func(w http.ResponseWriter, r *http.Request) {
		response, err := persons_info_controllers.GetPersons(conn)

		if err != nil {
			zap.L().Error(err.Error())
			w.WriteHeader(500)
			helpers.SetupResponse(&w)
			return
		}

		json.NewEncoder(w).Encode(response)
		w.WriteHeader(200)

		helpers.SetupResponse(&w)
	})
}
