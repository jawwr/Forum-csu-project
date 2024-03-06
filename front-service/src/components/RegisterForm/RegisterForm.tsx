import React from 'react';

import {Input} from "../Input";
import {Button} from "../Button";

import * as SC from './style';

export const RegisterForm = () => {
  return (
    <SC.Wrapper>
      <Input title="Логин" placeholder="Плакать сюда" />
      <Input title="Пароль" placeholder="Плакать сюда" />
      <Input title="Пароль" placeholder="Плакать сюда" />
      <Button>
        Зарегаться
      </Button>
    </SC.Wrapper>
  );
};
