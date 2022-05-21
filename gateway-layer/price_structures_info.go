package gateway_layer

import (
	price_structure_controllers "Backend/controllers/price_structures_info"
	price_structure_models "Backend/models/price_structures_info"
	"encoding/json"
	"github.com/jackc/pgx/v4"
	"go.uber.org/zap"
	"net/http"
)

func AddPriceStructureInfo(conn *pgx.Conn) {
	http.HandleFunc("/create-price-structure", func(w http.ResponseWriter, r *http.Request) {
		var request price_structure_models.CreatePriceStructureRequest
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

		err := price_structure_controllers.CreatePriceStructure(conn, request)

		if err != nil {
			zap.L().Error(err.Error())
			w.WriteHeader(500)
			return
		}
		w.WriteHeader(200)
	})

	http.HandleFunc("/update-price-structure-info", func(w http.ResponseWriter, r *http.Request) {
		var request price_structure_models.UpdatePriceStructureInfoRequest
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

		err := price_structure_controllers.UpdatePriceStructureInfo(conn, request)

		if err != nil {
			zap.L().Error(err.Error())
			w.WriteHeader(500)
			return
		}
		w.WriteHeader(200)
	})

	http.HandleFunc("/get-price-structure-info", func(w http.ResponseWriter, r *http.Request) {
		var request price_structure_models.GetPriceStructureInfoRequest
		json.NewDecoder(r.Body).Decode(&request)

		response, err := price_structure_controllers.GetPriceStructureInfo(conn, request)

		if err != nil {
			zap.L().Error(err.Error())
			w.WriteHeader(500)
			return
		}

		json.NewEncoder(w).Encode(response)
		w.WriteHeader(200)
	})
}
