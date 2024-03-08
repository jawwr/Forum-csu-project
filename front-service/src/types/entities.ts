export interface UserCredModel {
  login: string;
  password: string;
}

export interface UserModel {
  id: number;
  login: string;
}

export interface PostModel {
  id: number;
  text: string;
  title: string;
  user: UserModel;
}

export interface CreatePostModel extends Omit<PostModel, 'id' | 'user'> {}
