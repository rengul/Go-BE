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

func (r ConsumptionRepository) Get(ctx context.Context, user *models.User) ([]*models.Consumption, error) {
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
		err := rows.Scan(&consumption.Heating, &consumption.LastUpdate)
		if err != nil {
			return nil, err
		}
		consumptions = append(consumptions, &consumption)
	}

	return consumptions, nil
}
