package db

import (
	"context"
	"fmt"
)

func (dbq *PostgreSQLDatabaseQueries) UnsafeListAllApplicationStates(ctx context.Context, applicationStates *[]ApplicationState) error {

	if err := validateUnsafeQueryParamsNoPK(dbq); err != nil {
		return err
	}

	if err := dbq.dbConnection.Model(applicationStates).Context(ctx).Select(); err != nil {
		return err
	}

	return nil
}

func (dbq *PostgreSQLDatabaseQueries) DeleteApplicationStateById(ctx context.Context, id string) (int, error) {

	if err := validateQueryParams(id, dbq); err != nil {
		return 0, err
	}

	result := &ApplicationState{
		Applicationstate_application_id: id,
	}

	deleteResult, err := dbq.dbConnection.Model(result).WherePK().Context(ctx).Delete()
	if err != nil {
		return 0, fmt.Errorf("error on deleting application state: %v", err)
	}

	return deleteResult.RowsAffected(), nil
}

func (dbq *PostgreSQLDatabaseQueries) CreateApplicationState(ctx context.Context, obj *ApplicationState) error {

	if err := validateQueryParamsEntity(obj, dbq); err != nil {
		return err
	}

	if err := isEmptyValues("CreateApplicationState",
		"Applicationstate_application_id", obj.Applicationstate_application_id,
		"Health", obj.Health,
		"Sync_Status", obj.Sync_Status); err != nil {
		return err
	}

	if err := validateFieldLength(obj); err != nil {
		return err
	}

	// Inserting ApplicationState object
	result, err := dbq.dbConnection.Model(obj).Context(ctx).Insert()
	if err != nil {
		return fmt.Errorf("error on inserting application %v", err)
	}

	if result.RowsAffected() != 1 {
		return fmt.Errorf("unexpected number of rows affected: %d", result.RowsAffected())
	}
	return nil
}

func (dbq *PostgreSQLDatabaseQueries) UpdateApplicationState(ctx context.Context, obj *ApplicationState) error {

	if err := validateQueryParamsEntity(obj, dbq); err != nil {
		return err
	}

	if err := isEmptyValues("UpdateApplicationState",
		"Applicationstate_application_id", obj.Applicationstate_application_id,
		"Health", obj.Health,
		"Sync_Status", obj.Sync_Status); err != nil {
		return err
	}

	if err := validateFieldLength(obj); err != nil {
		return err
	}

	result, err := dbq.dbConnection.Model(obj).Context(ctx).
		Where("Applicationstate_application_id = ?", obj.Applicationstate_application_id).Update()
	if err != nil {
		return fmt.Errorf("error on updating application %v", err)
	}

	if result.RowsAffected() != 1 {
		return fmt.Errorf("unexpected number of rows affected: %d", result.RowsAffected())
	}

	return nil
}

func (dbq *PostgreSQLDatabaseQueries) GetApplicationStateById(ctx context.Context, obj *ApplicationState) error {

	if err := validateQueryParamsEntity(obj, dbq); err != nil {
		return err
	}

	if isEmpty(obj.Applicationstate_application_id) {
		return fmt.Errorf("applicationstate_application_id is nil")
	}

	var results []ApplicationState

	if err := dbq.dbConnection.Model(&results).
		Where("Applicationstate_application_id = ?", obj.Applicationstate_application_id).
		Context(ctx).
		Select(); err != nil {

		return fmt.Errorf("error on retrieving ApplicationState row: %v", err)
	}

	if len(results) == 0 {
		return NewResultNotFoundError(fmt.Sprintf("ApplicationState row '%s'", obj.Applicationstate_application_id))
	}

	if len(results) > 1 {
		return fmt.Errorf("multiple results found on retrieving ApplicationState row: %v", obj.Applicationstate_application_id)
	}

	*obj = results[0]

	return nil
}
