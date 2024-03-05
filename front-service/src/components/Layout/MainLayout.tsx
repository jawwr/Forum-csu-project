import React, {PropsWithChildren} from 'react';

import * as Styles from './styles';


export const MainLayout = (props: PropsWithChildren) => {
  return (
    <Styles.Wrapper>
      <Styles.Container>
        {props.children}
      </Styles.Container>
    </Styles.Wrapper>
  );
};
