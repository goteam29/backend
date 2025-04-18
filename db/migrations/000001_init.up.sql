create table if not exists public.users(
    id       uuid primary key,
    username text not null,
    email    text unique not null,
    token    text not null
);

create table if not exists public.classes(
    number   integer primary key,
    subject_ids uuid[]
);

create table if not exists public.subjects(
    id           uuid primary key,
    name         text not null,
    class_number integer not null references public.classes(number) on delete cascade,
    section_ids     uuid[]
);

create table if not exists public.sections(
    id          uuid primary key,
    subject_id  uuid not null references public.subjects(id) on delete cascade,
    name        text not null,
    description text not null,
    lesson_ids     uuid[]
);

create table if not exists public.lessons(
    id          uuid primary key,
    section_id  uuid not null references public.sections(id) on delete cascade,
    name        text not null,
    description text not null,
    videos      text[],
    files       text[],
    exercises   text[],
    comments    text[],
    rating      bigint not null default 0
);