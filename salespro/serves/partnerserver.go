package serves

import (
	"context"
	"fmt"

	"github.com/getfarook/go_grpc_crud_server/prodpb"
	"github.com/getfarook/go_grpc_crud_server/salespro/dbcon"
)

type Pserver struct{}

func (*Pserver) AddPartner(ctx context.Context, req *prodpb.Partner) (*prodpb.PartnerId, error) {
	fmt.Printf("Add Product function was invoked with : %v \n", req)
	name := req.GetName()
	openBal := req.GetOpenBal()
	currBal := req.GetCurrBal()
	contactNo := req.GetContactNo()
	details := req.GetDetails()

	db := dbcon.DbHandle()
	defer db.Close()

	sqlStatement := "INSERT INTO partners (name, open_bal, curr_bal, contact_no, details) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	var id int32 = 0

	err := db.QueryRow(sqlStatement, name, openBal, currBal, contactNo, details).Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Println("New Partner ID is:", id)

	resp := &prodpb.PartnerId{

		Id: id,
	}

	return resp, nil
}

func (*Pserver) DeletePartner(ctx context.Context, req *prodpb.PartnerId) (*prodpb.DelPartnerResponse, error) {
	id := req.GetId()

	db := dbcon.DbHandle()
	defer db.Close()

	sqlStatement := "DELETE FROM partners WHERE id = $1;"

	res, err1 := db.Exec(sqlStatement, id)
	if err1 != nil {
		panic(err1)
	}

	count, err2 := res.RowsAffected()
	if err2 != nil {
		panic(err2)
	}

	fmt.Print("\n No. of Partner Rows deleted were: ")
	fmt.Println(count)

	resp := &prodpb.DelPartnerResponse{}
	return resp, nil

}

func (*Pserver) UpdatePartner(ctx context.Context, req *prodpb.Partner) (*prodpb.DelPartnerResponse, error) {
	fmt.Printf("UpdatePartner function is invoked")

	id := req.GetId()
	name := req.GetName()
	openBal := req.GetOpenBal()
	currBal := req.GetCurrBal()
	contactNo := req.GetContactNo()
	details := req.GetDetails()

	db := dbcon.DbHandle()
	defer db.Close()

	sqlStatement := "UPDATE partners SET name = $2, open_bal = $3, curr_bal = $4, contact_no = $5, details = $6  WHERE id = $1;"

	res, err1 := db.Exec(sqlStatement, id, name, openBal, currBal, contactNo, details)
	if err1 != nil {
		panic(err1)
	}

	count, err2 := res.RowsAffected()
	if err2 != nil {
		panic(err2)
	}

	fmt.Print("\n No of Partner rows updated: ")
	fmt.Println(count)

	resp := &prodpb.DelPartnerResponse{}
	return resp, nil

}

func (*Pserver) GetPartners(ctx context.Context, req *prodpb.GetPartnerRequest) (*prodpb.Partners, error) {
	fmt.Printf("GetPartners function is invoked")

	db := dbcon.DbHandle()
	defer db.Close()

	rows, err := db.Query("SELECT id, name, open_bal, curr_bal, contact_no, details FROM partners ")
	if err != nil {
		// handle this error better than this
		panic(err)
	}
	defer rows.Close()

	loadedPartners := []*prodpb.Partner{}

	for rows.Next() {
		var id int32
		var name string
		var openBal float64
		var currBal float64
		var contactNo string
		var details string

		err1 := rows.Scan(&id, &name, &openBal, &currBal, &contactNo, &details)
		if err1 != nil {
			panic(err1)
		}

		newPartner := &prodpb.Partner{
			Id:        id,
			Name:      name,
			OpenBal:   openBal,
			CurrBal:   currBal,
			ContactNo: contactNo,
			Details:   details,
		}

		loadedPartners = append(loadedPartners, newPartner)

	}
	partners := &prodpb.Partners{
		MyPartners: loadedPartners,
	}

	err2 := rows.Err()
	if err2 != nil {
		panic(err2)
	}
	return partners, nil
}
