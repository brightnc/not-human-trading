package port

/*
	|--------------------------------------------------------------------------
	| Application Port
	|--------------------------------------------------------------------------
	|
	| Here you can define an interface which will allow foreign actors to
	| communicate with the Application
	|
*/

type Repository interface {
	SomeFunction() error
}
