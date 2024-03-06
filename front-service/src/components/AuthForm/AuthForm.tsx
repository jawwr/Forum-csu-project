import React from 'react';
import {Link} from "react-router-dom";

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
      <span>
        {'Нет аккаунта? ну не знаю иди поплачься '}
        <Link to="/register">
          кому-нибудь
        </Link>
      </span>
    </SC.Wrapper>
  );
};
