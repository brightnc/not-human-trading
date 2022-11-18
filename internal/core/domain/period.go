package domain

type Period string

const (
	// Min1 - 1 Minute time period
	Min1 Period = "60"
	// Min3 - 3 Minute time period
	Min3 Period = "3m"
	// Min5 - 5 Minute time period
	Min5 Period = "300"
	// Min15 - 15 Minute time period
	Min15 Period = "900"
	// Min30 - 30 Minute time period
	Min30 Period = "1800"
	// Min60 - 60 Minute time period
	Min60 Period = "3600"
	// Hour2 - 2 hour time period
	Hour2 Period = "2h"
	// Hour4 - 4 hour time period
	Hour4 Period = "4h"
	// Hour6 - 6 hour time period
	Hour6 Period = "6h"
	// Hour8 - 8 hour time period
	Hour8 Period = "8h"
	// Hour12 - 12 hour time period
	Hour12 Period = "12h"
	// Daily time period
	Daily Period = "d"
	// Day3 - 3 day time period
	Day3 Period = "3d"
	// Weekly time period
	Weekly Period = "w"
	// Monthly time period
	Monthly Period = "m"
)
