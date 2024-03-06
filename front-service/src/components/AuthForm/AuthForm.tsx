import React from 'react';

import {Input} from "../Input";
import {Button} from "../Button";

import * as SC from './style';

export const AuthForm = () => {
  return (
    <SC.Wrapper>
      <Input title="Логин" placeholder="Введите логин" />
      <Input title="Пароль" placeholder="Введите пароль" />
      <Button>
        Войти
      </Button>
    </SC.Wrapper>
  );
};
