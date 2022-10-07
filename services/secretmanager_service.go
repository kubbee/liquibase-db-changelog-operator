package services

func GetUsernameAndPasswordFromSecreManager() (map[string]string, error) {

	/*
	 * TODO: get db username and password from secret manager
	 */

	dbAuthoriozation := make(map[string]string)

	dbAuthoriozation["username"] = "teste"
	dbAuthoriozation["password"] = "teste"

	return dbAuthoriozation, nil

}
