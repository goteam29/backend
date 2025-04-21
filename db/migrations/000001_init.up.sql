create table if not exists public.users(
    id       uuid primary key,
    username text not null,
    email    text unique not null,
    token    text not null
);

create table if not exists public.classes(
    id          uuid primary key,
    number      integer unique not null
);

create table if not exists public.subjects(
    id          uuid primary key,
    name        text unique not null
);

create table if not exists public.classes_subjects(
    class_id    uuid not null references public.classes(id) on delete cascade,
    subject_id  uuid not null references public.subjects(id) on delete cascade,
    primary key (class_id, subject_id)
);  

create table if not exists public.sections(
    id          uuid primary key,
    subject_id  uuid not null references public.subjects(id) on delete cascade,
    name        text not null,
    description text not null,
    unique      (subject_id, name)
);

create table if not exists public.lessons(
    id           uuid primary key,
    section_id   uuid not null references public.sections(id) on delete cascade,
    name         text not null,
    description  text not null,
    rating       int not null default 0,
    unique       (section_id, name)
);

create table if not exists public.videos(
    id          uuid primary key,
    lesson_id   uuid not null references public.lessons(id) on delete cascade,
    url         text not null
);

create table if not exists public.files(
    id          uuid primary key,
    lesson_id   uuid not null references public.lessons(id) on delete cascade,
    url         text not null
);

create table if not exists public.exercises(
    id          uuid primary key,
    lesson_id   uuid not null references public.lessons(id) on delete cascade,
    name        text not null,
    description text not null,
    answer      text not null,
    unique      (lesson_id, name)
);

create table if not exists public.comments(
    id          uuid primary key,
    lesson_id   uuid not null references public.lessons(id) on delete cascade,
    user_id     uuid not null references public.users(id) on delete cascade,
    text        text not null,
    created_at  timestamp not null default now(),
    rating      int not null default 0
);
