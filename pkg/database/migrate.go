package database

func MigrateDB() error {
	err := DB.AutoMigrate(
		&Controller{},
		&Certificate{},
		&ControllerHold{},
		&Evaluation{},
		&RatingChange{},
		&ControllerVisit{},
		&ControllerRole{},
		&ControllerRosterRequest{},
		&APIUser{},
		&APIToken{},
		&APITokenLog{},
	)
	if err != nil {
		return err
	}
	return nil
}
