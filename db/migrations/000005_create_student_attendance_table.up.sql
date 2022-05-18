CREATE table IF NOT EXISTS student_attendance(
    id SERIAL PRIMARY KEY ,
    student_id integer REFERENCES students(id),
    login_time time ,
    logout_time time,
    date date
)
