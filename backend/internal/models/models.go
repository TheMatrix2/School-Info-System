package models

import (
	"time"
	
	"github.com/satori/go.uuid"
	"github.com/uptrace/bun"
)

type Role struct { // user's role (admin, student, teacher etc.)
	bun.BaseModel `bun:"table:roles"`	
	
	ID		uuid.UUID	`bun:",pk,type:uuid"`
	Name	string		`bun:",unique,notnull"`
}

type User struct { // user of system
	bun.BaseModel `bun:"table:users"`
	
	ID			uuid.UUID	`bun:",pk,type:uuid"`
	Surname 	string		`bun:",notnull"`
	Name		string		`bun:",notnull"`
	Patronymic 	string		`bun:",nullzero"`
	Email		string		`bun:",unique,notnull"`
	Phone		string		`bun:",unique,notnull"`
	
	RoleID		uuid.UUID	`bun:"type:uuid,notnull"`
	Role		*Role		`bun:"rel:belongs-to,join:role_id"`
	
	Deleted		bool		`bun:",notnull"`
	
	CreatedAt	time.Time
	UpdatedAt 	time.Time
	DeletedAt	time.Time
}

type Application struct { // application for admission to school
	bun.BaseModel `bun:"table:applications"`
	
	ID			uuid.UUID	`bun:",pk,type:uuid"`
	
	ApplicantID	uuid.UUID	`bun:"type:uuid,notnull"`
	Applicant	*User		`bun:"rel:belongs-to,join:applicant_id=id"`
	
	Status		string		`bun:",notnull"`
	
	CreatedAt	time.Time
	UpdatedAt 	time.Time
}

type Course struct { // learning course
	bun.BaseModel `bun:"table:courses"`

	ID          uuid.UUID `bun:",pk,type:uuid"`
	Name        string    `bun:",notnull"`
	Description string    `bun:",nullzero"`

	TeacherID 	uuid.UUID `bun:"type:uuid,notnull"`
	Teacher   	*User     `bun:"rel:belongs-to,join:teacher_id=id"`
}

type Group struct { // learning group
	bun.BaseModel `bun:"table:groups"`

	ID       uuid.UUID 	`bun:",pk,type:uuid"`
	Name     string    	`bun:",notnull"`
	CourseID uuid.UUID 	`bun:"type:uuid,notnull"`

	Course	 *Course 	`bun:"rel:belongs-to,join:course_id=id"`
}

type Lesson struct { // lesson in schedule
	bun.BaseModel `bun:"table:lessons"`

	ID      uuid.UUID	`bun:",pk,type:uuid"`
	
	GroupID uuid.UUID 	`bun:"type:uuid,notnull"`
	Group   *Group    	`bun:"rel:belongs-to,join:group_id=id"`
	
	StartAt time.Time 	`bun:",notnull"`
	EndAt   time.Time 	`bun:",notnull"`
	Room	int64		
}

type Visit struct { // student's visit at lesson
	bun.BaseModel `bun:"table:visits"`

	ID       	uuid.UUID	`bun:",pk,type:uuid"`
	LessonID 	uuid.UUID	`bun:"type:uuid,notnull"`
	StudentID 	uuid.UUID	`bun:"type:uuid,notnull"`

	Lesson  	*Lesson		`bun:"rel:belongs-to,join:lesson_id=id"`
	Student 	*User		`bun:"rel:belongs-to,join:student_id=id"`

	Present 	bool		`bun:",notnull"`
}