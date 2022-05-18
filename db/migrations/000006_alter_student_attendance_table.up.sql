ALTER TABLE student_attendance
    ADD CONSTRAINT student_date_unique UNIQUE (student_id, date)
