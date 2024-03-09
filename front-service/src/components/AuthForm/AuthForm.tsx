import React, {useState} from 'react';
import {Link, redirect} from "react-router-dom";
import axios from "axios";

import {Input} from "../Input";
import {Button} from "../Button";

import {useToken} from "helpers/hooks";

import {UserCredModel} from "types/entities";

import * as SC from './style';

export const AuthForm = () => {
  const [_, setToken] = useToken();

  const [login, setLogin] = useState('');
  const [pass, setPass] = useState('');

  const handleEnterClick = () => {
    const userCred: UserCredModel = {
      login,
      password: pass,
    }

    axios.post('http://localhost:8083/api/auth', userCred)
      .then(res => {
        if (res.status === 200) {
          console.log(res.data.token);

          setToken(res.data.token);
          redirect('/');
        } else {
          alert('Авторизация неудачная!');
        }
      })
      .catch((e) => {
        alert('ОШИБКА: ' + e.target.value);
      })
  }

  return (
    <SC.Wrapper>
      <Input
        title="Логин"
        placeholder="Введите логин"
        type="text"
        onChange={(e) => {
          setLogin(e.target.value);
        }}
      />
      <Input
        title="Пароль"
        placeholder="Введите пароль"
        type="password"
        onChange={(e) => {
          setPass(e.target.value);
        }}
      />
      <Button
        onClick={handleEnterClick}
        disabled={login.length < 2 || pass.length < 2}
      >
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
