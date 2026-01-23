export interface dataState {
  page: number;
  limit: number;
}

export interface topic {
  id: string,
  topicname: string,
  description: string,
  created_at: string 
  author_id: string
}
export interface topicsList {
  topics: topic[]
}