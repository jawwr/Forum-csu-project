import React from 'react';
import {Route, Routes} from "react-router-dom";

import {MainPage, AuthPage, RegisterPage, UserListPage} from "pages";

import {useToken} from "helpers/hooks";

function App() {
  const [token] = useToken();

  const isAuth = !!token;

  if (!isAuth) {
    return (
      <AuthPage/>
    )
  }
  return (
    <Routes>
      <Route path="/" element={<MainPage/>}/>
      <Route path="/register" element={<RegisterPage/>}/>
      <Route path="/login" element={<AuthPage/>}/>
      <Route path="/users" element={<UserListPage/>}/>
    </Routes>
  );
}

export default App;

// TODO endpoints:
//  user-service:8083/api/auth (POST), { UserCred }
//  user-service:8083/api/users (GET),
//  user-service:8083/api/users/{id} (GET),
//  user-service:8083/api/users/{id}/subscribe (POST),
//  post-service:8082/api/post (POST), { Post }
//  post-service:8082/api/post (GET),
//  post-service:8082/api/post/{id} (GET)
