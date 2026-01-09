-- Откат миграции: удаление всех объектов в обратном порядке

-- Удаление триггеров
DROP TRIGGER IF EXISTS update_contests_updated_at ON contests;
DROP TRIGGER IF EXISTS update_users_updated_at ON users;

-- Удаление функции
DROP FUNCTION IF EXISTS update_updated_at_column();

-- Удаление junction таблиц (сначала удаляем зависимости)
DROP TABLE IF EXISTS contests_teams;
DROP TABLE IF EXISTS contest_participants;
DROP TABLE IF EXISTS users_trophies;
DROP TABLE IF EXISTS team_members;

-- Удаление основных таблиц
DROP TABLE IF EXISTS contests;
DROP TABLE IF EXISTS teams;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS trophies;
