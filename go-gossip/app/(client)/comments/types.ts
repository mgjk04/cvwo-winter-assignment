export enum verb {
  Create = 'Create',
  Edit =  'Edit',
}

export interface comment {
  id: string;
  content: string;
  description: string;
  created_at: string;
  post_id: string;
  author_id: string;
  authorname: string;
}