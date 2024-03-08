import React from 'react';
import axios from "axios";

import {Box} from "../Wrapper";

import {UserModel} from "types/entities";

import {useToken} from "helpers/hooks";

import * as SC from './style';

export interface UserListProps {
  users: UserModel[];
}

export const UserList = ({users}: UserListProps) => {
  const [token] = useToken();

  const handleButtonClick = (id: number) => {
    axios.post(`http://localhost:8083/api/users/${id}/subscribe`, {}, {
      headers: {
        'Authorization': `Bearer ${token}`,
      }
    })
      .then(res => {
        if (res.status === 200) {
          users.filter(u => u.id !== id);
        } else {
          alert('Не получилось подписаться!');
        }
      })
      .catch(e => {
        alert('Ошибка!' + e);
      })
  }

  return (
    <SC.Wrapper>
      {users.map(u => (
        <Box>
          <h6>
            {u.login}
          </h6>
          <button onClick={() => handleButtonClick(u.id)}>
            Подписаться
          </button>
        </Box>
      ))}
    </SC.Wrapper>
  );
};
