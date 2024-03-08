import React from 'react';

import {Input} from "../Input";
import {Button} from "../Button";

import * as SC from './style';

export const RegisterForm = () => {
  return (
    <SC.Wrapper>
      <Input
        title="Логин"
        placeholder="Введите логин"
        type="text"
      />
      <Input
        title="Пароль"
        placeholder="Введите пароль"
        type="password"
      />
      <Input
        title="Пароль"
        placeholder="Повторите пароль"
        type="password"
      />
      <Button>
        Зарегаться
      </Button>
    </SC.Wrapper>
  );
};
