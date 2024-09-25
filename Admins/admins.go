// данная папка хранит данные про админов, в реальной жизни - это БД,
//можно запушить миграции

package admins

type Admin struct {
	ID         int    `json:"id"`
	Department string `json:"departametn"`
}

func InitializeAdmins() []Admin {
	result := make([]Admin, 0)
	a1 := Admin{
		ID:         1,
		Department: "Support",
	}

	a2 := Admin{
		ID:         2,
		Department: "IT",
	}

	a3 := Admin{
		ID:         3,
		Department: "Billing",
	}

	// Append the instances to the result slice
	result = append(result, a1, a2, a3)

	return result
}
