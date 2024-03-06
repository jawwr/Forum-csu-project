import React from 'react';

import {MainPage, AuthPage} from "pages";

function App() {
  const isAuth = false;

  if (!isAuth) {
    return (
      <AuthPage />
    )
  }
  return (
      <MainPage />
  );
}

export default App;
