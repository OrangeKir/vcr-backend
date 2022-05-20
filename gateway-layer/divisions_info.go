package gateway_layer

import (
	division_info_controllers "Backend/controllers/divisions_info"
	division_info_models "Backend/models/divisions_info"
	"encoding/json"
	"github.com/jackc/pgx/v4"
	"go.uber.org/zap"
	"net/http"
)

func AddDivisionInfo(conn *pgx.Conn) {
	http.HandleFunc("/create-division", func(w http.ResponseWriter, r *http.Request) {
		var request division_info_models.CreateDivisionRequest
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

		err := division_info_controllers.CreateDivision(conn, request)

		if err != nil {
			zap.L().Error(err.Error())
			w.WriteHeader(500)
			return
		}
		w.WriteHeader(200)
	})

	http.HandleFunc("/update-division-info", func(w http.ResponseWriter, r *http.Request) {
		var request division_info_models.UpdateDivisionInfoRequest
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

		err := division_info_controllers.UpdateDivisionInfo(conn, request)

		if err != nil {
			zap.L().Error(err.Error())
			w.WriteHeader(500)
			return
		}
		w.WriteHeader(200)
	})

	http.HandleFunc("/get-division-info", func(w http.ResponseWriter, r *http.Request) {
		var request division_info_models.GetDivisionInfoRequest
		json.NewDecoder(r.Body).Decode(&request)

		response, err := division_info_controllers.GetDivisionInfo(conn, request)

		if err != nil {
			zap.L().Error(err.Error())
			w.WriteHeader(500)
			return
		}

		json.NewEncoder(w).Encode(response)
		w.WriteHeader(200)
	})
}
