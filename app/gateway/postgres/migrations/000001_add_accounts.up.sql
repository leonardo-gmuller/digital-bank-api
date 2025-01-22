begin;
create table if not exists Account (
    id bigint generated always as identity primary key,
    name text not null,
    cpf text not null,
    secret text not null,
    balance int not null DEFAULT 0,
    created_at timestamptz not null default current_timestamp,
    update_at timestamptz not null default current_timestamp,
    deleted_at timestamptz
);

create table if not exists Transfer (
    id bigint generated always as identity primary key,
    account_origin_id bigint not null,
    account_destination_id bigint not null,
    amount int not null DEFAULT 0,
    created_at timestamptz not null default current_timestamp,
    update_at timestamptz not null default current_timestamp,
    deleted_at timestamptz
);
commit;