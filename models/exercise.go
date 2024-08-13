package models

type Exercise struct {
	Id          string             `bson:"_id,omitempty"`
	Name        string             `bson:"name"`
	Description string             `bson:"description"`
	Parameters  ExerciseParameters `bson:"parameters"`
}

type ExerciseParameters struct {
	Time   bool
	Reps   bool
	Weight bool
}
