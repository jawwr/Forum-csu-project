import React, {ComponentPropsWithoutRef, PropsWithoutRef} from 'react';

import * as SC from './style';

export interface InputProps extends ComponentPropsWithoutRef<'input'> {
  title: string;
}

export const Input = ({title, ...rest}: InputProps) => {
  return (
    <SC.Wrapper>
      <SC.Label>
        {title}
      </SC.Label>
      <SC.Input
        {...rest}
      />
    </SC.Wrapper>
  );
};
