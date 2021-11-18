-- +goose Up
CREATE TABLE cars
(
    id SERIAL NOT NULL CONSTRAINT cars_pk PRIMARY KEY,
    model      VARCHAR(50) NOT NULL,
    color      VARCHAR(15) NOT NULL,
    deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE TABLE car_orders
(
    id SERIAL NOT NULL PRIMARY KEY,
    deleted_at TIMESTAMP WITH TIME ZONE,
    start_at TIMESTAMP WITH TIME ZONE NOT NULL,
    end_at TIMESTAMP WITH TIME ZONE NOT NULL,
    car_id INTEGER NOT NULL REFERENCES cars (id)

);

-- +goose Down
DROP TABLE cars;
DROP TABLE car_orders;