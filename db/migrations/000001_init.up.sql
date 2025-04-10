create table if not exists public.users(
    id       uuid primary key,
    username text not null,
    email    text unique not null,
    token    text not null
);

create table if not exists public.lessons(
    id          uuid primary key,
    section_id  uuid not null,
    name        text not null,
    description text not null,
    videos      text[],
    files       text[],
    exercises   text[],
    comments    text[],
    rating      bigint not null default 0
    -- constraint fk_section
    --     foreign key(section) references public.sections(id) on delete cascade
);