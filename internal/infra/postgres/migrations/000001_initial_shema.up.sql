CREATE SCHEMA IF NOT EXISTS user_service;
CREATE SCHEMA IF NOT EXISTS contest_service;
CREATE SCHEMA IF NOT EXISTS team_service;

-- Таблица users (пользователи) — принадлежит user_service
CREATE TABLE user_service.users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email TEXT UNIQUE,
    username TEXT NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    avatar TEXT,
    bio TEXT,
    role TEXT NOT NULL DEFAULT 'user',
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Таблица teams (команды) — принадлежит team_service
CREATE TABLE team_service.teams (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    description TEXT,
    type TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT chk_teams_type CHECK (type IN ('public', 'private'))
);

-- Таблица contests — принадлежит contest_service
CREATE TABLE contest_service.contests (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    description TEXT,
    status TEXT NOT NULL,
    type TEXT NOT NULL,
    start_date TIMESTAMPTZ,           
    end_date TIMESTAMPTZ,             
    registration_start TIMESTAMPTZ,   
    registration_end TIMESTAMPTZ,     
    max_participants INTEGER,         
    max_teams INTEGER,                
    min_team_size INTEGER,            
    max_team_size INTEGER,            
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT chk_contests_status CHECK (status IN ('draft', 'active', 'finished', 'cancelled')),
    CONSTRAINT chk_contests_type CHECK (type IN ('individual', 'team'))
);

-- Junction таблица team_members (связь пользователей и команд с ролями) — внутри team_service
CREATE TABLE team_service.team_members (
    user_id UUID NOT NULL,
    team_id UUID NOT NULL,
    role TEXT NOT NULL,
    PRIMARY KEY (user_id, team_id),
    CONSTRAINT fk_team_members_team FOREIGN KEY (team_id) REFERENCES team_service.teams(id) ON DELETE CASCADE,
    CONSTRAINT chk_team_members_role CHECK (role IN ('captain', 'member'))
);

-- Junction таблица contest_participants (связь пользователей и конкурсов)
CREATE TABLE contest_service.contest_participants (
    user_id UUID NOT NULL,
    contest_id UUID NOT NULL,
    joined_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, contest_id),
    CONSTRAINT fk_contest_participants_contest FOREIGN KEY (contest_id) REFERENCES contest_service.contests(id) ON DELETE CASCADE
);

-- Junction таблица contests_teams (связь команд и конкурсов)
CREATE TABLE contest_service.contests_teams (
    team_id UUID NOT NULL,
    contest_id UUID NOT NULL,
    joined_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (team_id, contest_id),
    CONSTRAINT fk_contests_teams_contest FOREIGN KEY (contest_id) REFERENCES contest_service.contests(id) ON DELETE CASCADE
);

-- Индексы для таблицы users
CREATE INDEX idx_users_created_at ON user_service.users(created_at);

-- Индексы для таблицы teams
CREATE INDEX idx_teams_type ON team_service.teams(type);
CREATE INDEX idx_teams_created_at ON team_service.teams(created_at);

CREATE INDEX idx_contests_status ON contest_service.contests(status);
CREATE INDEX idx_contests_type ON contest_service.contests(type);
CREATE INDEX idx_contests_created_at ON contest_service.contests(created_at);

-- Индексы для junction таблиц
-- Для team_members
CREATE INDEX idx_team_members_user_id ON team_service.team_members(user_id);
CREATE INDEX idx_team_members_team_id ON team_service.team_members(team_id);
CREATE INDEX idx_team_members_role ON team_service.team_members(role);

-- Для contest_participants
CREATE INDEX idx_contest_participants_joined_at ON contest_service.contest_participants(joined_at);

-- Для contests_teams
CREATE INDEX idx_contests_teams_joined_at ON contest_service.contests_teams(joined_at);

-- Функция для автоматического обновления updated_at
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language 'plpgsql';

-- Триггеры для автоматического обновления updated_at
CREATE TRIGGER update_users_updated_at BEFORE UPDATE ON user_service.users
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_contests_updated_at BEFORE UPDATE ON contest_service.contests
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
