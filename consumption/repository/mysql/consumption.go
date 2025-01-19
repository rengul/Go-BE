package mysql

import (
	"context"
	"database/sql"
	"re-home/models"
	mod "re-home/models"
)

type ConsumptionRepository struct {
	db *sql.DB
}

func NewConsumptionRepository(db *sql.DB) *ConsumptionRepository {
	return &ConsumptionRepository{
		db: db,
	}
}

func (r ConsumptionRepository) GetAllConsumption(ctx context.Context, userId string) ([]*mod.Consumption, error) {
	query := `SELECT 
	DATE(lastupdate) AS lastupdate,
	MAX(hotwater) - MIN(hotwater) AS hotwater,
	MAX(coldwater) - MIN(coldwater) AS coldwater,
	MAX(heating) - MIN(heating) AS heating,
	MAX(cooling) - MIN(cooling) AS cooling
	FROM 
	consumption
WHERE 
	MONTH(lastupdate) = MONTH(CURRENT_DATE()) 
	AND YEAR(lastupdate) = YEAR(CURRENT_DATE())    
GROUP BY 
	DATE(lastupdate)
ORDER BY 
	lastupdate;`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var consumptions []*mod.Consumption
	for rows.Next() {
		var consumption mod.Consumption
		err := rows.Scan(&consumption.LastUpdate, &consumption.HotWater, &consumption.ColdWater, &consumption.Heating, &consumption.Cooling)
		if err != nil {
			return nil, err
		}
		consumptions = append(consumptions, &consumption)
	}

	return consumptions, nil
}

func (r ConsumptionRepository) GetHotWater(ctx context.Context, userId string) ([]*mod.Consumption, error) {
	query := `SELECT 
	DATE(lastupdate) AS lastupdate,
	MAX(hotwater) - MIN(hotwater) AS hotwater	
FROM 
	consumption
WHERE 
	MONTH(lastupdate) = MONTH(CURRENT_DATE()) 
	AND YEAR(lastupdate) = YEAR(CURRENT_DATE())    
GROUP BY 
	DATE(lastupdate)
ORDER BY 
	lastupdate;`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var consumptions []*mod.Consumption
	for rows.Next() {
		var consumption mod.Consumption
		err := rows.Scan(&consumption.LastUpdate, &consumption.HotWater)
		if err != nil {
			return nil, err
		}
		consumptions = append(consumptions, &consumption)
	}

	return consumptions, nil
}

func (r ConsumptionRepository) GetColdWater(ctx context.Context, userId string) ([]*mod.Consumption, error) {
	query := `SELECT 
	DATE(lastupdate) AS lastupdate,
	MAX(coldwater) - MIN(coldwater) AS coldwater	
FROM 
	consumption
WHERE 
	MONTH(lastupdate) = MONTH(CURRENT_DATE()) 
	AND YEAR(lastupdate) = YEAR(CURRENT_DATE())    
GROUP BY 
	DATE(lastupdate)
ORDER BY 
	lastupdate;`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var consumptions []*mod.Consumption
	for rows.Next() {
		var consumption mod.Consumption
		err := rows.Scan(&consumption.LastUpdate, &consumption.ColdWater)
		if err != nil {
			return nil, err
		}
		consumptions = append(consumptions, &consumption)
	}

	return consumptions, nil
}

func (r ConsumptionRepository) GetHeating(ctx context.Context, userId string) ([]*models.Consumption, error) {

	query := `SELECT 
	DATE(lastupdate) AS lastupdate,
	MAX(heating) - MIN(heating) AS heating	
FROM 
	consumption
WHERE 
	MONTH(lastupdate) = MONTH(CURRENT_DATE()) 
	AND YEAR(lastupdate) = YEAR(CURRENT_DATE())    
GROUP BY 
	DATE(lastupdate)
ORDER BY 
	lastupdate;`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var consumptions []*mod.Consumption
	for rows.Next() {
		var consumption mod.Consumption
		err := rows.Scan(&consumption.LastUpdate, &consumption.Heating)
		if err != nil {
			return nil, err
		}
		consumptions = append(consumptions, &consumption)
	}

	return consumptions, nil
}
