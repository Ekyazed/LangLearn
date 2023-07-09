package main

import "time"

var (
	APPR1 = func() *time.Time {
		h := time.Now().Add(time.Hour * 4)
		return &h
	}
	APPR2 = func() *time.Time {
		h := time.Now().Add(time.Hour * 8)
		return &h
	}
	APPR3 = func() *time.Time {
		h := time.Now().Add(time.Hour * 24)
		return &h
	}
	APPR4 = func() *time.Time {
		h := time.Now().Add(time.Hour * 24 * 2)
		return &h
	}
	CONN1 = func() *time.Time {
		h := time.Now().Add(time.Hour * 24 * 7)
		return &h
	}
	CONN2 = func() *time.Time {
		h := time.Now().Add(time.Hour * 24 * 7 * 2)
		return &h
	}
	MTRE1 = func() *time.Time {
		h := time.Now().Add(time.Hour * 24 * 7 * 4)
		return &h
	}
	MTRE2 = func() *time.Time {
		h := time.Now().Add(time.Hour * 24 * 7 * 4 * 2)
		return &h
	}
	EXP1 = func() *time.Time {
		h := time.Now().Add(time.Hour * 24 * 7 * 4 * 4)
		return &h
	}
	EXP2 = func() *time.Time {
		h := time.Now().Add(time.Hour * 24 * 7 * 4 * 12)
		return &h
	}
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
