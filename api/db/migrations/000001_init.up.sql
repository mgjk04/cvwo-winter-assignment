CREATE TABLE Users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    username TEXT UNIQUE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE TABLE Topics (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    topicname TEXT UNIQUE NOT NULL,
    description TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    author_id UUID NOT NULL,
    FOREIGN KEY (author_id) REFERENCES Users(id)
);

CREATE TABLE Posts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title TEXT NOT NULL,
    description TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    topic_id UUID NOT NULL,
    FOREIGN KEY (topic_id) REFERENCES Topics(id) ON DELETE CASCADE,
    author_id UUID NOT NULL,
    FOREIGN KEY (author_id) REFERENCES Users(id)
);

CREATE TABLE Comments (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    content TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    post_id UUID NOT NULL,
    FOREIGN KEY (post_id) REFERENCES Posts(id) ON DELETE CASCADE,
    author_id UUID NOT NULL,
    FOREIGN KEY (author_id) REFERENCES Users(id)
);