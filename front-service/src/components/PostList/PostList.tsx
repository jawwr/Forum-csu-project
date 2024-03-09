import React from 'react';

import {Post} from "../Post";

import {PostModel} from "types/entities";

import * as SC from './style';

export interface PostListProps {
  posts: PostModel[];
}

export const PostList = ({posts}: PostListProps) => {
  return (
    <SC.Wrapper>
      {posts.map(p => (
        <Post key={p.id} data={p} />
      ))}
    </SC.Wrapper>
  );
};
