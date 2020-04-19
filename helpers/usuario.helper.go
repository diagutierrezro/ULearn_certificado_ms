package helpers

func ObtenerUsuario(usu int) map[string]interface{} {
	/*	var (
		url_crud  string
		respuesta map[string]interface{}
	)*/

	/*	if _, err := request.getJson(URLCrud,respuesta); error = nil{
			return respuesta,  nil
		}else {
			nil, err
		}	*/
	respuestaTemp := map[string]interface{}{
		"user_id":  usu,
		"username": "diagutierrezro",
		"password": "12345",
		"names":    "Diego Alejandro",
		"surnames": "Gutierrez Rojas",
	}
	return respuestaTemp
}
