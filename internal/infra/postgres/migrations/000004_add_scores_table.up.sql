CREATE TABLE scores (
    id UUID PRIMARY KEY,
    project_id UUID NOT NULL,
    judge_id UUID NOT NULL,
    score INTEGER NOT NULL,  
    comment TEXT,
    created_at TIMESTAMPTZ,
    UNIQUE(project_id, judge_id) 
);