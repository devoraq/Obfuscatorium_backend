-- 1. Удаление триггеров (обязательно указываем правильные схемы)
DROP TRIGGER IF EXISTS update_contests_updated_at ON contest_service.contests;
DROP TRIGGER IF EXISTS update_users_updated_at ON user_service.users;

-- 2. Удаление функции
DROP FUNCTION IF EXISTS update_updated_at_column();

-- 3. Удаление Junction-таблиц (сначала их, так как они ссылаются на основные)
DROP TABLE IF EXISTS contest_service.contests_teams;
DROP TABLE IF EXISTS contest_service.contest_participants;
DROP TABLE IF EXISTS team_service.team_members; -- Исправлена схема

-- 4. Удаление основных таблиц
DROP TABLE IF EXISTS contest_service.contests;
DROP TABLE IF EXISTS team_service.teams;        -- Исправлена схема
DROP TABLE IF EXISTS user_service.users;

-- 5. Удаление схем
DROP SCHEMA IF EXISTS contest_service CASCADE;
DROP SCHEMA IF EXISTS user_service CASCADE;
DROP SCHEMA IF EXISTS team_service CASCADE;     -- Добавлена схема