import React from 'react';

import {PostModel} from "types/entities";

import {Box, Horizon} from "../Wrapper";

import * as SC from './style';

export const Post = ({data} : {data: PostModel}) => {
  return (
    <SC.Wrapper>
      <Box>
        <Horizon>
          <h3>
            {data.title}
          </h3>
          <span>
          {data.user.login}
        </span>
        </Horizon>
        <p>
          {data.text}
        </p>
      </Box>
    </SC.Wrapper>
  );
};
