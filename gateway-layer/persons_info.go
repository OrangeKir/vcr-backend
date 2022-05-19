package gateway_layer

import (
	persons_info_controllers "Backend/controllers/persons_info"
	persons_info_models "Backend/models/persons_info"
	"encoding/json"
	"github.com/jackc/pgx/v4"
	"go.uber.org/zap"
	"net/http"
)

func AddPersonInfo(conn *pgx.Conn) {
	http.HandleFunc("/create-person", func(w http.ResponseWriter, r *http.Request) {
		var request persons_info_models.CreatePersonRequest
		json.NewDecoder(r.Body).Decode(&request)

		//isValidToken, _, role := helpers.ValidateToken(request.Token)
		//
		//if !isValidToken {
		//	w.WriteHeader(401)
		//	return
		//}
		//
		//if !(role == types.Administrator) {
		//	w.WriteHeader(403)
		//	return
		//}

		err := persons_info_controllers.CreatePerson(conn, request)

		if err != nil {
			zap.L().Error(err.Error())
			w.WriteHeader(500)
			return
		}
		w.WriteHeader(200)
	})

	http.HandleFunc("/update-person-info", func(w http.ResponseWriter, r *http.Request) {
		var request persons_info_models.UpdatePersonInfoRequest
		json.NewDecoder(r.Body).Decode(&request)

		//isValidToken, login, role := helpers.ValidateToken(request.Token)
		//
		//if !isValidToken {
		//	w.WriteHeader(401)
		//	return
		//}
		//
		//if !(role == types.Administrator || login != request.Login) {
		//	w.WriteHeader(403)
		//	return
		//}

		err := persons_info_controllers.UpdatePersonInfo(conn, request)

		if err != nil {
			zap.L().Error(err.Error())
			w.WriteHeader(500)
			return
		}
		w.WriteHeader(200)
	})

	http.HandleFunc("/get-person-info", func(w http.ResponseWriter, r *http.Request) {
		var request persons_info_models.GetPersonInfoRequest
		json.NewDecoder(r.Body).Decode(&request)

		response, err := persons_info_controllers.GetPersonInfo(conn, request)

		if err != nil {
			zap.L().Error(err.Error())
			w.WriteHeader(500)
			return
		}

		json.NewEncoder(w).Encode(response)
		w.WriteHeader(200)
	})
}
