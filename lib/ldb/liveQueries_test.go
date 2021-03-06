package ldb

import (
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestAllRoadstead(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		log.Println(err)
	}

	defer db.Close()

	expectedRows := sqlmock.NewRows([]string{"agency", "anchorage_point", "anchoring_time", "current_activity", "gross_tonnage", "id_trip", "iso3", "length", "ship", "ship_type", "shipped_goods", "ts_planned_mooring", "ts_readiness", "width"})

	mock.ExpectQuery(`SELECT`).WithArgs("28", "28").WillReturnRows(expectedRows)

	mockDB := NewRepository(db)
	mockDB.GetAllRoadstead("28")

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("testing error: %s", err)
	}
}

func TestAllActiveTrips(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		log.Println(err)
	}

	defer db.Close()

	expectedRows := sqlmock.NewRows([]string{"id_control_unit_data", "ship_name", "ship_type", "length", "width", "gross_tonnage", "net_tonnage", "details"})

	mock.ExpectQuery(`SELECT`).WithArgs("28").WillReturnRows(expectedRows)

	mockDB := NewRepository(db)
	mockDB.GetActiveTrips("28")

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("testing error: %s", err)
	}
}

func TestAllMoored(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		log.Println(err)
	}

	defer db.Close()

	expectedRows := sqlmock.NewRows([]string{"id_control_unit_data", "ship_description", "ts_last_ship_activity", "ship_current_activities.description", "anchorage_points.description", "type_acronym", "iso3", "gross_tonnage", "ships.length", "ships.width", "agencies.description", "shipped_goods_data.shipped_goods_row", "data_previsione_arrivo_nave.ts_mooring_time", "data_arrivo_in_rada.ts_readiness"})

	mock.ExpectQuery(`SELECT`).WithArgs("28", "28").WillReturnRows(expectedRows)

	mockDB := NewRepository(db)
	mockDB.GetAllMoored("28")

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("testing error: %s", err)
	}
}

func TestArrivalsToday(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		log.Println(err)
	}

	defer db.Close()

	expectedRows := sqlmock.NewRows([]string{"id_control_unit_data", "ship_description", "ts_last_ship_activity", "ship_current_activities.description", "anchorage_points.description", "type_acronym", "iso3", "gross_tonnage", "ships.length", "ships.width", "agencies.description", "shipped_goods_data.shipped_goods_row", "data_previsione_arrivo_nave.ts_mooring_time", "data_arrivo_in_rada.ts_readiness"})

	mock.ExpectQuery(`SELECT`).WithArgs(10, 10, 10, 10, "28").WillReturnRows(expectedRows)

	mockDB := NewRepository(db)
	mockDB.GetTodayArrivals("28", 12)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("testing error: %s", err)
	}
}

func TestDeparturesToday(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		log.Println(err)
	}

	defer db.Close()

	expectedRows := sqlmock.NewRows([]string{"id_control_unit_data", "ship_description", "ts_last_ship_activity", "ship_current_activities.description", "anchorage_points.description", "type_acronym", "iso3", "gross_tonnage", "ships.length", "ships.width", "agencies.description", "shipped_goods_data.shipped_goods_row", "data_previsione_arrivo_nave.ts_mooring_time", "data_arrivo_in_rada.ts_readiness"})

	mock.ExpectQuery(`SELECT`).WithArgs(26, "28").WillReturnRows(expectedRows)

	mockDB := NewRepository(db)
	mockDB.GetTodayDepartures("28", 26)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("testing error: %s", err)
	}
}

func TestArrivalPrevisionsToday(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		log.Println(err)
	}

	defer db.Close()

	expectedRows := sqlmock.NewRows([]string{"id_control_unit_data", "ship_description", "ts_last_ship_activity", "ship_current_activities.description", "anchorage_points.description", "type_acronym", "iso3", "gross_tonnage", "ships.length", "ships.width", "agencies.description", "shipped_goods_data.shipped_goods_row", "data_previsione_arrivo_nave.ts_mooring_time", "data_arrivo_in_rada.ts_readiness"})

	mock.ExpectQuery(`SELECT`).WithArgs("28").WillReturnRows(expectedRows)

	mockDB := NewRepository(db)
	mockDB.GetArrivalPrevisions("28")

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("testing error: %s", err)
	}
}

func TestShippedGoods(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		log.Println(err)
	}

	defer db.Close()

	expectedRows := sqlmock.NewRows([]string{"id_control_unit_data", "ship_description", "ts_last_ship_activity", "ship_current_activities.description", "anchorage_points.description", "type_acronym", "iso3", "gross_tonnage", "ships.length", "ships.width", "agencies.description", "shipped_goods_data.shipped_goods_row", "data_previsione_arrivo_nave.ts_mooring_time", "data_arrivo_in_rada.ts_readiness"})

	mock.ExpectQuery(`SELECT`).WithArgs("28").WillReturnRows(expectedRows)

	mockDB := NewRepository(db)
	mockDB.GetTodayShippedGoods("28")

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("testing error: %s", err)
	}
}

func TestTrafficList(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		log.Println(err)
	}

	defer db.Close()

	expectedRows := sqlmock.NewRows([]string{"id_control_unit_data", "ship_description", "ts_last_ship_activity", "ship_current_activities.description", "anchorage_points.description", "type_acronym", "iso3", "gross_tonnage", "ships.length", "ships.width", "agencies.description", "shipped_goods_data.shipped_goods_row", "data_previsione_arrivo_nave.ts_mooring_time", "data_arrivo_in_rada.ts_readiness"})

	mock.ExpectQuery(`SELECT`).WithArgs("28").WillReturnRows(expectedRows)

	mockDB := NewRepository(db)
	mockDB.GetTodayTrafficList("28")

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("testing error: %s", err)
	}
}

func TestShiftingPrevisionsToday(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		log.Println(err)
	}

	defer db.Close()

	expectedRows := sqlmock.NewRows([]string{"id_control_unit_data", "ship_description", "ts_last_ship_activity", "ship_current_activities.description", "anchorage_points.description", "type_acronym", "iso3", "gross_tonnage", "ships.length", "ships.width", "agencies.description", "shipped_goods_data.shipped_goods_row", "data_previsione_arrivo_nave.ts_mooring_time", "data_arrivo_in_rada.ts_readiness"})

	mock.ExpectQuery(`SELECT`).WithArgs("28").WillReturnRows(expectedRows)

	mockDB := NewRepository(db)
	mockDB.GetShiftingPrevisions("28")

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("testing error: %s", err)
	}
}

func TestDeparturePrevisionsToday(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		log.Println(err)
	}

	defer db.Close()

	expectedRows := sqlmock.NewRows([]string{"id_control_unit_data", "ship_description", "ts_last_ship_activity", "ship_current_activities.description", "anchorage_points.description", "type_acronym", "iso3", "gross_tonnage", "ships.length", "ships.width", "agencies.description", "shipped_goods_data.shipped_goods_row", "data_previsione_arrivo_nave.ts_mooring_time", "data_arrivo_in_rada.ts_readiness"})

	mock.ExpectQuery(`SELECT`).WithArgs("28").WillReturnRows(expectedRows)

	mockDB := NewRepository(db)
	mockDB.GetDeparturePrevisions("28")

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("testing error: %s", err)
	}
}

func TestShiftingsToday(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		log.Println(err)
	}

	defer db.Close()

	expectedRows := sqlmock.NewRows([]string{"id_control_unit_data", "ship_description", "ts_last_ship_activity", "ship_current_activities.description", "anchorage_points.description", "type_acronym", "iso3", "gross_tonnage", "ships.length", "ships.width", "agencies.description", "shipped_goods_data.shipped_goods_row", "data_previsione_arrivo_nave.ts_mooring_time", "data_arrivo_in_rada.ts_readiness"})

	mock.ExpectQuery(`SELECT`).WithArgs("28").WillReturnRows(expectedRows)

	mockDB := NewRepository(db)
	mockDB.GetTodayShiftings("28")

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("testing error: %s", err)
	}
}
