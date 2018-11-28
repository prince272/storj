// Copyright (C) 2018 Storj Labs, Inc.
// See LICENSE for copying information.

package satellitedb

import (
	"context"

	"github.com/zeebo/errs"

	"github.com/skyrings/skyring-common/tools/uuid"
	"storj.io/storj/pkg/satellite"
	"storj.io/storj/pkg/satellite/satellitedb/dbx"
)

// implementation of Companies interface repository using spacemonkeygo/dbx orm
type companies struct {
	db *dbx.DB
}

// Get is a method for querying company from the database by user id
func (companies *companies) GetByUserID(ctx context.Context, userID uuid.UUID) (*satellite.Company, error) {
	company, err := companies.db.Get_Company_By_UserId(ctx, dbx.Company_UserId(userID[:]))
	if err != nil {
		return nil, err
	}

	return companyFromDBX(company)
}

// Insert is a method for inserting company into the database
func (companies *companies) Insert(ctx context.Context, company *satellite.Company) (*satellite.Company, error) {
	createdCompany, err := companies.db.Create_Company(
		ctx,
		dbx.Company_UserId(company.UserID[:]),
		dbx.Company_Name(company.Name),
		dbx.Company_Address(company.Address),
		dbx.Company_Country(company.Country),
		dbx.Company_City(company.City),
		dbx.Company_State(company.State),
		dbx.Company_PostalCode(company.PostalCode))

	if err != nil {
		return nil, err
	}

	return companyFromDBX(createdCompany)
}

// Delete is a method for deleting company by Id from the database.
func (companies *companies) Delete(ctx context.Context, userID uuid.UUID) error {
	_, err := companies.db.Delete_Company_By_UserId(ctx, dbx.Company_UserId(userID[:]))

	return err
}

// Update is a method for updating company entity
func (companies *companies) Update(ctx context.Context, company *satellite.Company) error {
	_, err := companies.db.Update_Company_By_UserId(
		ctx,
		dbx.Company_UserId(company.UserID[:]),
		getCompanyUpdateFields(company))

	return err
}

// companyFromDBX is used for creating Company entity from autogenerated dbx.Company struct
func companyFromDBX(company *dbx.Company) (*satellite.Company, error) {
	if company == nil {
		return nil, errs.New("company parameter is nil")
	}

	userID, err := bytesToUUID(company.UserId)
	if err != nil {
		return nil, err
	}

	return &satellite.Company{
		UserID:     userID,
		Name:       company.Name,
		Address:    company.Address,
		Country:    company.Country,
		City:       company.City,
		State:      company.State,
		PostalCode: company.PostalCode,
		CreatedAt:  company.CreatedAt,
	}, nil
}

// getCompanyUpdateFields is used to generate company update fields
func getCompanyUpdateFields(company *satellite.Company) dbx.Company_Update_Fields {
	return dbx.Company_Update_Fields{
		Name:       dbx.Company_Name(company.Name),
		Address:    dbx.Company_Address(company.Address),
		Country:    dbx.Company_Country(company.Country),
		City:       dbx.Company_City(company.City),
		State:      dbx.Company_State(company.State),
		PostalCode: dbx.Company_PostalCode(company.PostalCode),
	}
}
