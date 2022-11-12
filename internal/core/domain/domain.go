package domain

/*
	|--------------------------------------------------------------------------
	| Structure/Model and the Application Rules
	|--------------------------------------------------------------------------
	|
	| Here you can define struct to use in your application and you may define
	| some application rule.
	|
*/

type BusinessLogicRequest struct {
	Name string `validate:"gt=10"`
}
