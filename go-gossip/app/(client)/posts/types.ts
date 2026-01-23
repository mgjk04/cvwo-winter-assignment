export enum verb {
  Create = 'Create',
  Edit =  'Edit',
}

export interface post {
  id: string;
  title: string;
  description: string;
  created_at: string;
  topic_id: string;
  author_id: string;
  authorname: string;
}