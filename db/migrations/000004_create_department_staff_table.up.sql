CREATE table IF NOT EXISTS department_staff(
       id SERIAL PRIMARY KEY ,
       department_id integer REFERENCES departments(id),
       staff_id integer references staff(id)
)
