CREATE TABLE cities (
    id UUID PRIMARY KEY,
    name TEXT CHECK (CHAR_LENGTH(name) > 3 AND CHAR_LENGTH(name) <= 30),
    created_at TIMESTAMP DEFAULT NOW()
);


create table customers (
    id uuid primary key,
    full_name text,
    phone text unique,
    email text unique,
    created_at timestamp default now()
);

create table drivers (
    id uuid primary key ,
    full_name text,
    phone text unique,
    from_city_id uuid references cities(id),
    to_city_id uuid references cities(id),
    created_at timestamp default now()
);

create table cars (
    id uuid primary key ,
    model varchar(30),
    brand varchar(30),
    number varchar(30) unique,
    status boolean default true,
    driver_id uuid references drivers(id),
    created_at timestamp default now()
);

create table trips (
    id uuid primary key,
    trip_number_id varchar(5) unique,
    from_city_id uuid references cities(id),
    to_city_id uuid references cities(id),
    driver_id uuid references drivers(id),
    price int default 0 check (price >= 0),
    created_at timestamp default now()
);

create table trip_customers (
    id uuid primary key,
    trip_id uuid references trips(id),
    customer_id uuid references customers(id),
    created_at timestamp default now()
);