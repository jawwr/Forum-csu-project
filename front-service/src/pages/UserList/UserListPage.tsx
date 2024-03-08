import React, {useEffect, useState} from 'react';
import axios from "axios";

import {UserModel} from "types/entities";

import {MainLayout} from "components/Layout";
import {UserList} from "components/UserList";

import {useToken} from "helpers/hooks";

export const UserListPage = () => {
  const [token] = useToken();

  const [users, setUsers] = useState<UserModel[]>([]);

  useEffect(() => {
    axios.get('http://localhost:8083/api/users', {
      headers: {
        'Authorization': `Bearer ${token}`,
      }
    })
      .then(res => {
        setUsers(res.data.posts);
      })
  }, []);

  return (
    <MainLayout>
      <UserList users={users}/>
    </MainLayout>
  );
};
