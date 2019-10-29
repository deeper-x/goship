package ldb

import (
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestArrivalsRegister(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		log.Println(err)
	}

	defer db.Close()

	expectedRows := sqlmock.NewRows([]string{"id_control_unit_data", "ship_description", "ts_last_ship_activity", "ship_current_activities.description", "anchorage_points.description", "type_acronym", "iso3", "gross_tonnage", "ships.length", "ships.width", "agencies.description", "shipped_goods_data.shipped_goods_row", "data_previsione_arrivo_nave.ts_mooring_time", "data_arrivo_in_rada.ts_readiness"})

	mock.ExpectQuery(`SELECT id_control_unit_data AS id_trip,`).WithArgs(26, 26, 26, 26, "28", "2019-01-10 01:00", "2019-02-10 01:00").WillReturnRows(expectedRows)
	// idArrivalPrevision, idArrivalPrevision, idArrivalPrevision, idArrivalPrevision, idPortinformer, start, stop

	mockDB := NewRepository(db)
	mockDB.GetArrivalsRegister("28", 26, "2019-01-10 01:00", "2019-02-10 01:00")

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("testing error: %s", err)
	}
}

func TestMooredRegister(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		log.Println(err)
	}

	defer db.Close()

	expectedRows := sqlmock.NewRows([]string{"id_control_unit_data", "ship_description", "ts_last_ship_activity", "ship_current_activities.description", "anchorage_points.description", "type_acronym", "iso3", "gross_tonnage", "ships.length", "ships.width", "agencies.description", "shipped_goods_data.shipped_goods_row", "data_previsione_arrivo_nave.ts_mooring_time", "data_arrivo_in_rada.ts_readiness"})

	mock.ExpectQuery(`SELECT id_control_unit_data, ship_description, type_description, ts_fine_ormeggio,`).WithArgs("28", "2019-01-10 01:00", "2019-02-10 01:00", "28", "2019-01-10 01:00", "2019-02-10 01:00", "28", "2019-01-10 01:00", "2019-02-10 01:00").WillReturnRows(expectedRows)

	mockDB := NewRepository(db)
	mockDB.GetMooredRegister("28", "2019-01-10 01:00", "2019-02-10 01:00")
	//idPortinformer, start, stop, idPortinformer, start, stop, idPortinformer, start, stop

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("testing error: %s", err)
	}
}

func TestAnchoredRegister(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		log.Println(err)
	}

	defer db.Close()

	expectedRows := sqlmock.NewRows([]string{"id_control_unit_data", "ship_description", "ts_last_ship_activity", "ship_current_activities.description", "anchorage_points.description", "type_acronym", "iso3", "gross_tonnage", "ships.length", "ships.width", "agencies.description", "shipped_goods_data.shipped_goods_row", "data_previsione_arrivo_nave.ts_mooring_time", "data_arrivo_in_rada.ts_readiness"})

	mock.ExpectQuery(`SELECT id_control_unit_data, ship_description, type_description, ts_anchor_drop,`).WithArgs("28", "2019-01-10 01:00", "2019-02-10 01:00", "28", "2019-01-10 01:00", "2019-02-10 01:00", "28", "2019-01-10 01:00", "2019-02-10 01:00").WillReturnRows(expectedRows)

	mockDB := NewRepository(db)
	mockDB.GetRoadsteadRegister("28", "2019-01-10 01:00", "2019-02-10 01:00")
	//idPortinformer, start, stop, idPortinformer, start, stop, idPortinformer, start, stop

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("testing error: %s", err)
	}
}

func TestDeparturesRegister(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		log.Println(err)
	}

	defer db.Close()

	expectedRows := sqlmock.NewRows([]string{"id_control_unit_data", "ship_description", "ts_last_ship_activity", "ship_current_activities.description", "anchorage_points.description", "type_acronym", "iso3", "gross_tonnage", "ships.length", "ships.width", "agencies.description", "shipped_goods_data.shipped_goods_row", "data_previsione_arrivo_nave.ts_mooring_time", "data_arrivo_in_rada.ts_readiness"})

	mock.ExpectQuery(`SELECT id_control_unit_data AS id_trip,`).WithArgs(10, "28", "2019-01-10 01:00", "2019-02-10 01:00").WillReturnRows(expectedRows)

	mockDB := NewRepository(db)
	mockDB.GetDeparturesRegister("28", 10, "2019-01-10 01:00", "2019-02-10 01:00")
	//idPortinformer, start, stop, idPortinformer, start, stop, idPortinformer, start, stop

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("testing error: %s", err)
	}
}

func TestShiftingsRegister(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		log.Println(err)
	}

	defer db.Close()

	expectedRows := sqlmock.NewRows([]string{"id_control_unit_data", "ship_description", "ts_last_ship_activity", "ship_current_activities.description", "anchorage_points.description", "type_acronym", "iso3", "gross_tonnage", "ships.length", "ships.width", "agencies.description", "shipped_goods_data.shipped_goods_row", "data_previsione_arrivo_nave.ts_mooring_time", "data_arrivo_in_rada.ts_readiness"})

	mock.ExpectQuery(`SELECT id_control_unit_data, ts_main_event_field_val,`).WithArgs("28", "2019-01-10 01:00", "2019-02-10 01:00").WillReturnRows(expectedRows)

	mockDB := NewRepository(db)
	mockDB.GetShiftingsRegister("28", "2019-01-10 01:00", "2019-02-10 01:00")

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("testing error: %s", err)
	}
}

func TestShippedGoodsRegister(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		log.Println(err)
	}

	defer db.Close()

	expectedRows := sqlmock.NewRows([]string{"id_control_unit_data", "ship_description", "ts_last_ship_activity", "ship_current_activities.description", "anchorage_points.description", "type_acronym", "iso3", "gross_tonnage", "ships.length", "ships.width", "agencies.description", "shipped_goods_data.shipped_goods_row", "data_previsione_arrivo_nave.ts_mooring_time", "data_arrivo_in_rada.ts_readiness"})

	mock.ExpectQuery(`SELECT id_control_unit_data AS id_trip,`).WithArgs("28", "2019-01-10 01:00", "2019-02-10 01:00").WillReturnRows(expectedRows)

	mockDB := NewRepository(db)
	mockDB.GetShippedGoodsRegister("28", "2019-01-10 01:00", "2019-02-10 01:00")

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("testing error: %s", err)
	}
}

func TestTrafficListRegister(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		log.Println(err)
	}

	defer db.Close()

	expectedRows := sqlmock.NewRows([]string{"id_control_unit_data", "ship_description", "ts_last_ship_activity", "ship_current_activities.description", "anchorage_points.description", "type_acronym", "iso3", "gross_tonnage", "ships.length", "ships.width", "agencies.description", "shipped_goods_data.shipped_goods_row", "data_previsione_arrivo_nave.ts_mooring_time", "data_arrivo_in_rada.ts_readiness"})

	mock.ExpectQuery(`SELECT control_unit_data.id_control_unit_data AS id_trip,`).WithArgs("28", "2019-01-10 01:00", "2019-02-10 01:00").WillReturnRows(expectedRows)

	mockDB := NewRepository(db)
	mockDB.GetRegisterTrafficList("28", "2019-01-10 01:00", "2019-02-10 01:00")

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("testing error: %s", err)
	}
}
