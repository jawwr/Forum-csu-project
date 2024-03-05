import React from 'react';

import {Post} from "../Post";

import * as SC from './style';

const POSTS = [
  '1', '2', '3'
]

export const PostList = () => {
  return (
    <SC.Wrapper>
      {POSTS.map(p => (
        <Post key={p} data={p} />
      ))}
    </SC.Wrapper>
  );
};
