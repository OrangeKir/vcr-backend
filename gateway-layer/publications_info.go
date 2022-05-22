package gateway_layer

import (
	publications_info_controllers "Backend/controllers/publications_info"
	publications_info_models "Backend/models/publications_info"
	"encoding/json"
	"github.com/jackc/pgx/v4"
	"go.uber.org/zap"
	"net/http"
)

func AddPublicationsInfo(conn *pgx.Conn) {
	http.HandleFunc("/create-publishing-house", func(w http.ResponseWriter, r *http.Request) {
		var request publications_info_models.CreatePublishingHouseRequest
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

		err := publications_info_controllers.CreatePublishingHouse(conn, request)

		if err != nil {
			zap.L().Error(err.Error())
			w.WriteHeader(500)
			return
		}
		w.WriteHeader(200)
	})

	http.HandleFunc("/get-publishing-houses-info", func(w http.ResponseWriter, r *http.Request) {

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

		response, err := publications_info_controllers.GetPublishingHousesInfo(conn)

		if err != nil {
			zap.L().Error(err.Error())
			w.WriteHeader(500)
			return
		}

		json.NewEncoder(w).Encode(response)
		w.WriteHeader(200)
	})

	http.HandleFunc("/create-users-publication", func(w http.ResponseWriter, r *http.Request) {
		var request publications_info_models.CreateUsersPublicationRequest
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

		err := publications_info_controllers.CreateUsersPublication(conn, request)

		if err != nil {
			zap.L().Error(err.Error())
			w.WriteHeader(500)
			return
		}
		w.WriteHeader(200)
	})

	http.HandleFunc("/get-user-publications", func(w http.ResponseWriter, r *http.Request) {
		var request publications_info_models.GetUserPublicationsInfoByLoginRequest
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

		response, err := publications_info_controllers.GetUserPublications(conn, request)

		if err != nil {
			zap.L().Error(err.Error())
			w.WriteHeader(500)
			return
		}

		json.NewEncoder(w).Encode(response)
		w.WriteHeader(200)
	})
}
