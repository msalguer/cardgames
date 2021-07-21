package cardgames

import "net/http"

//************************** API HOME PAGE *******************************************
func HomePage(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte(Loadreadme()))
}
