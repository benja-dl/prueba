package repositories

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/luccasniccolas/monitor/data"
	"github.com/luccasniccolas/monitor/database"
	"github.com/luccasniccolas/monitor/utils"
)

func IsMailRegistered(email string) (bool, error) {
	var count int64
	err := database.DB.QueryRow(`SELECT COUNT(*) FROM users WHERE email = $1`, email).Scan(&count)

	if err != nil {
		return false, err
	}
	// Si el contador es mayor a 1, el mail ya esta registrado
	return count > 0, nil
}

func GetUserByEmail(email string) (*data.User, error) {
	var user data.User

	row := database.DB.QueryRow("SELECT id, email, first_name, last_name FROM users WHERE email=$1", email)
	err := row.Scan(&user.ID, &user.Email, &user.FirstName, &user.LastName)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("usuario no encontrado")
		}
		return nil, err
	}

	return &user, nil
}

func Login(email, password string) (*data.User, error) {
	user, err := GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	if !utils.VerifyHashData(user.Password, password) {
		return nil, errors.New("por favor revisar las credenciales")
	}

	return user, nil

}

func RegisterUser(user *data.User) error {
	if !utils.IsValidEmail(user.Email) {
		return errors.New("correo invalido")
	}

	isRegistered, _ := IsMailRegistered(user.Email)
	if isRegistered {
		return errors.New("el correo ya esta asociado a una cuenta")
	}

	hashPassword, err := utils.HashData(user.Password)
	if err != nil {
		return errors.New("error al hashear la contraseña")
	}

	// Falta agregar al usuario
	_, err = database.DB.Exec(`INSERT INTO users (email, password, first_name, last_name) 
			VALUES ($1, $2, $3, $4)`, user.Email, hashPassword, user.FirstName, user.LastName)

	if err != nil {
		return err
	}

	return nil
}
