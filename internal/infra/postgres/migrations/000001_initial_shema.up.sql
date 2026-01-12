-- Таблица trophies (трофеи)
CREATE TABLE trophies (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    description TEXT,
    image TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Таблица users (пользователи)
CREATE TABLE users (
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

-- Таблица teams (команды)
CREATE TABLE teams (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    description TEXT,
    type TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT chk_teams_type CHECK (type IN ('public', 'private'))
);

-- Таблица contests
CREATE TABLE contests (
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

-- Junction таблица team_members (связь пользователей и команд с ролями)
CREATE TABLE team_members (
    user_id UUID NOT NULL,
    team_id UUID NOT NULL,
    role TEXT NOT NULL,
    PRIMARY KEY (user_id, team_id),
    CONSTRAINT fk_team_members_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT fk_team_members_team FOREIGN KEY (team_id) REFERENCES teams(id) ON DELETE CASCADE,
    CONSTRAINT chk_team_members_role CHECK (role IN ('captain', 'member'))
);

-- Junction таблица users_trophies (связь пользователей и трофеев)
CREATE TABLE users_trophies (
    user_id UUID NOT NULL,
    trophy_id UUID NOT NULL,
    received_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, trophy_id),
    CONSTRAINT fk_users_trophies_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT fk_users_trophies_trophy FOREIGN KEY (trophy_id) REFERENCES trophies(id) ON DELETE CASCADE
);

-- Junction таблица contest_participants (связь пользователей и конкурсов)
CREATE TABLE contest_participants (
    user_id UUID NOT NULL,
    contest_id UUID NOT NULL,
    joined_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, contest_id),
    CONSTRAINT fk_contest_participants_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT fk_contest_participants_contest FOREIGN KEY (contest_id) REFERENCES contests(id) ON DELETE CASCADE
);

-- Junction таблица contests_teams (связь команд и конкурсов)
CREATE TABLE contests_teams (
    team_id UUID NOT NULL,
    contest_id UUID NOT NULL,
    joined_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (team_id, contest_id),
    CONSTRAINT fk_contests_teams_team FOREIGN KEY (team_id) REFERENCES teams(id) ON DELETE CASCADE,
    CONSTRAINT fk_contests_teams_contest FOREIGN KEY (contest_id) REFERENCES contests(id) ON DELETE CASCADE
);

-- Индексы для таблицы users
CREATE INDEX idx_users_created_at ON users(created_at);

-- Индексы для таблицы teams
CREATE INDEX idx_teams_type ON teams(type);
CREATE INDEX idx_teams_created_at ON teams(created_at);

-- Индексы для таблицы contests
CREATE INDEX idx_contests_status ON contests(status);
CREATE INDEX idx_contests_type ON contests(type);
CREATE INDEX idx_contests_created_at ON contests(created_at);

-- Индексы для таблицы trophies
CREATE INDEX idx_trophies_created_at ON trophies(created_at);

-- Индексы для junction таблиц
-- Для team_members
CREATE INDEX idx_team_members_user_id ON team_members(user_id);
CREATE INDEX idx_team_members_team_id ON team_members(team_id);
CREATE INDEX idx_team_members_role ON team_members(role);

-- Для users_trophies
CREATE INDEX idx_users_trophies_received_at ON users_trophies(received_at);

-- Для contest_participants
CREATE INDEX idx_contest_participants_joined_at ON contest_participants(joined_at);

-- Для contests_teams
CREATE INDEX idx_contests_teams_joined_at ON contests_teams(joined_at);

-- Функция для автоматического обновления updated_at
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language 'plpgsql';

-- Триггеры для автоматического обновления updated_at
CREATE TRIGGER update_users_updated_at BEFORE UPDATE ON users
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_contests_updated_at BEFORE UPDATE ON contests
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
