package main

import "time"

var (
	APPR1 = func() time.Time { return time.Now().Add(time.Hour * 4) }
	APPR2 = func() time.Time { return time.Now().Add(time.Hour * 8) }
	APPR3 = func() time.Time { return time.Now().AddDate(0, 0, 1) }
	APPR4 = func() time.Time { return time.Now().AddDate(0, 0, 2) }
	CONN1 = func() time.Time { return time.Now().AddDate(0, 0, 7) }
	CONN2 = func() time.Time { return time.Now().AddDate(0, 0, 14) }
	MTRE1 = func() time.Time { return time.Now().AddDate(0, 1, 0) }
	MTRE2 = func() time.Time { return time.Now().AddDate(0, 2, 0) }
	EXP1  = func() time.Time { return time.Now().AddDate(0, 4, 0) }
	EXP2  = func() time.Time { return time.Now().AddDate(1, 0, 0) }
)

const (
	FRESH          = "FRESH"
	APPRENTISSAGE1 = "APPRENTISSAGE1"
	APPRENTISSAGE2 = "APPRENTISSAGE2"
	APPRENTISSAGE3 = "APPRENTISSAGE3"
	APPRENTISSAGE4 = "APPRENTISSAGE4"
	CONNAISSEUR1   = "CONNAISSEUR1"
	CONNAISSEUR2   = "CONNAISSEUR2"
	MAITRE1        = "MAITRE1"
	MAITRE2        = "MAITRE2"
	EXPERT1        = "EXPERT1"
	EXPERT2        = "EXPERT2"
	BURNED         = "BURNED"
)
