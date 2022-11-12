package repository

/*
	|--------------------------------------------------------------------------
	| The Repository Adaptor
	|--------------------------------------------------------------------------
	|
	| An Adapter will initiate the interaction with the Application through
	| a Port, using specific technology that means you can choose
	| any technology you want for your application or business logic.
	|
*/

type Postgres struct{}

func NewPostgres() *Postgres {
	return &Postgres{}
}

func (r *Postgres) SomeFunction() error { return nil }
