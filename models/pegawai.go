package models

import (
	"net/http"

	"github.com/kuma-coffee/go-rest-api/db"
)

type Pegawai struct {
	Id      int    `json:"id"`
	Nama    string `json:"nama"`
	Alamat  string `json:"alamat"`
	Telepon string `json:"telepon"`
}

func FetchAllPegawai() (Response, error) {
	var (
		obj    Pegawai
		arrObj []Pegawai
	)

	conn := db.CreateConn()

	sqlStatement := `select * from pegawai`

	rows, err := conn.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return Response{}, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.Nama, &obj.Alamat, &obj.Telepon)
		if err != nil {
			return Response{}, err
		}

		arrObj = append(arrObj, obj)
	}

	return Response{
		Status:  http.StatusOK,
		Message: "Success Fetch Data",
		Data:    arrObj,
	}, nil
}
