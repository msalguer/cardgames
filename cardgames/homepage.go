package cardgames

import "net/http"

//************************** API HOME PAGE *******************************************
//func HomePage(c echo.Context) error {
func HomePage(res http.ResponseWriter, req *http.Request) {
	//return c.String(http.StatusOK, Readme)
	res.Write([]byte(Loadreadme()))
	//res.Write([]byte(Readme))
}
