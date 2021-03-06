package gateway_layer

import (
	conferences_info_controllers "Backend/controllers/conferences_info"
	"Backend/helpers"
	conferences_info_models "Backend/models/conferences_info"
	"Backend/types"
	"encoding/json"
	"github.com/jackc/pgx/v4"
	"go.uber.org/zap"
	"net/http"
)

func AddConferencesInfo(conn *pgx.Conn) {
	http.HandleFunc("/create-conference", func(w http.ResponseWriter, r *http.Request) {
		var request conferences_info_models.CreateConferenceRequest
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

		err := conferences_info_controllers.CreateConference(conn, request)

		if err != nil {
			zap.L().Error(err.Error())
			w.WriteHeader(500)
			helpers.SetupResponse(&w)
			return
		}
		w.WriteHeader(200)

		helpers.SetupResponse(&w)
	})

	http.HandleFunc("/get-conferences", func(w http.ResponseWriter, r *http.Request) {
		response, err := conferences_info_controllers.GetConferencesInfo(conn)

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

	http.HandleFunc("/create-user-conference-topic", func(w http.ResponseWriter, r *http.Request) {
		var request conferences_info_models.CreateUserConferenceTopicRequest
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

		err := conferences_info_controllers.CreateUserConferenceTopic(conn, request)

		if err != nil {
			zap.L().Error(err.Error())
			w.WriteHeader(500)
			helpers.SetupResponse(&w)
			return
		}
		w.WriteHeader(200)

		helpers.SetupResponse(&w)
	})

	http.HandleFunc("/get-user-conferences", func(w http.ResponseWriter, r *http.Request) {
		var request conferences_info_models.GetUserConferencesInfoRequest
		json.NewDecoder(r.Body).Decode(&request)

		response, err := conferences_info_controllers.GetUserConferencesInfo(conn, request)

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
