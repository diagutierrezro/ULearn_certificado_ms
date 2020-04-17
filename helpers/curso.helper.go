package helpers

func ObtenerCurso(cur int) map[string]interface{} {
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
		"id_curso":        cur,
		"nombre":          "GO",
		"categoria":       "backend",
		"duracion":        "35 horas",
		"nombre_profesor": "pepito",
		"calificacion":    "5",
	}
	return respuestaTemp
}
