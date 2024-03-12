package vehicle

import (
	"net/http"
	"strconv"
	"fmt"
	"github.com/CecileTalec/vehicle-server/storage"
	"go.uber.org/zap"
	"../vehiclestore/store.go"
)

type DeleteHandler struct {
	store  storage.Store
	logger *zap.Logger
}

func NewDeleteHandler(store storage.Store, logger *zap.Logger) *DeleteHandler {
	return &DeleteHandler{
		store:  store,
		logger: logger.With(zap.String("handler", "delete_vehicles")),
	}
}

func (d *DeleteHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	http.Error(rw, "Not Implemented", http.StatusInternalServerError)
	id := r.PathValue("id")
	id_num, err := strconv.ParseInt(id, 10, 64);
	if err == nil {
		fmt.Printf("%T, %v\n", id_num, err) //cas erreur
	}
	reponse := d.store.Delete(r.Context, id_num)
	if !reponse{
		fmt.Println("ERREUR 404")
	}
	rw.WriteHeader(http.StatusNoContent)

}