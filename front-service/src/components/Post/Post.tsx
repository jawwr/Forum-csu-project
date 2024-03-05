import React from 'react';

import * as SC from './style';

export const Post = ({data} : {data: any}) => {
  return (
    <SC.Wrapper>
      {data}
    </SC.Wrapper>
  );
};
