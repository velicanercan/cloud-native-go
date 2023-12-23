package health

import (
	"net/http"
)

// Read godoc
//
//	@summary		Health API
//	@description	Health API
//	@tags			health
//	@success		200
//	@router			/.../health [get]
func Read(w http.ResponseWriter, r *http.Request) {}
