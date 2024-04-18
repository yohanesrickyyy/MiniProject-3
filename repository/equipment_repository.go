package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"miniproject-3/helper"
	"miniproject-3/model/domain"
)

type EquipmentRepository interface {
	Save(context.Context, *sql.Tx, domain.Equipment) error
	Update(context.Context, *sql.Tx, domain.Equipment) error
	Delete(context.Context, *sql.Tx, int) error
	FindById(context.Context, *sql.Tx, int) (*domain.Equipment, error)
	FindAll(context.Context, *sql.Tx) ([]domain.Equipment, error)
}

type EquipmentRepositoryImpl struct{}

func (repository *EquipmentRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, equipment domain.Equipment) error {
	SQL := `INSERT INTO equipment (Name, Availability, RentalCosts, Category) VALUES (?, ?, ?, ?) RETURNING ID`
	row := tx.QueryRowContext(ctx, SQL, equipment.Name, equipment.Availability, equipment.RentalCosts, equipment.Category)
	var id int
	err := row.Scan(&id)
	helper.PanicIfError(err)
	equipment.ID = uint(id)
	return nil
}

func (repository *EquipmentRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, equipment domain.Equipment) error {
	SQL := `UPDATE equipment SET Name=?, Availability=?, RentalCosts=?, Category=? WHERE ID=?`
	_, err := tx.ExecContext(ctx, SQL, equipment.Name, equipment.Availability, equipment.RentalCosts, equipment.Category, equipment.ID)
	helper.PanicIfError(err)
	return nil
}

func (repository *EquipmentRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id int) error {
	SQL := `DELETE FROM equipment WHERE ID=?`
	_, err := tx.ExecContext(ctx, SQL, id)
	helper.PanicIfError(err)
	return nil
}

func (repository *EquipmentRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (*domain.Equipment, error) {
	SQL := `SELECT ID, Name, Availability, RentalCosts, Category FROM equipment WHERE ID=?`
	row := tx.QueryRowContext(ctx, SQL, id)
	var equipment domain.Equipment
	err := row.Scan(&equipment.ID, &equipment.Name, &equipment.Availability, &equipment.RentalCosts, &equipment.Category)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("error getting equipment: %v", err)
	}
	return &equipment, nil
}

func (repository *EquipmentRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]domain.Equipment, error) {
	SQL := `SELECT ID, Name, Availability, RentalCosts, Category FROM equipment`
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	var equipments []domain.Equipment
	for rows.Next() {
		var equipment domain.Equipment
		err := rows.Scan(&equipment.ID, &equipment.Name, &equipment.Availability, &equipment.RentalCosts, &equipment.Category)
		if err != nil {
			return nil, fmt.Errorf("error scanning equipment row: %v", err)
		}
		equipments = append(equipments, equipment)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating equipment rows: %v", err)
	}
	return equipments, nil
}
