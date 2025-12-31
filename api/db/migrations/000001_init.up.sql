CREATE TABLE Users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    username TEXT UNIQUE NOT NULL,
    created_at TIMESTAMPZ DEFAULT NOW()
    deleted_at TIMESTAMPZ
);

CREATE TABLE Topics (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    topicname TEXT UNIQUE NOT NULL,
    description TEXT,
    created_at TIMESTAMPZ DEFAULT NOW(),
    author_id UUID NOT NULL,
    FOREIGN KEY (author_id) REFERENCES Users(id)
);

CREATE TABLE Posts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title TEXT NOT NULL,
    description TEXT,
    created_at TIMESTAMPZ DEFAULT NOW(),
    topic_id UUID NOT NULL,
    FOREIGN KEY (topic_id) REFERENCES Topics(id),
    author_id UUID NOT NULL,
    FOREIGN KEY (author_id) REFERENCES Users(id)
);

CREATE TABLE Comments (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    content TEXT NOT NULL,
    created_at TIMESTAMPZ DEFAULT NOW(),
    post_id UUID NOT NULL,
    FOREIGN KEY (post_id) REFERENCES Posts(id),
    author_id UUID NOT NULL,
    FOREIGN KEY (author_id) REFERENCES Users(id)
);