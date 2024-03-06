import React from 'react';

import {MainPage, AuthPage} from "pages";
import {BrowserRouter, Route, Routes} from "react-router-dom";
import {RegisterPage} from "./pages/Register";

function App() {
  const isAuth = true;

  if (!isAuth) {
    return (
      <AuthPage />
    )
  }
  return (
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<MainPage />} />
          <Route path="/register" element={<RegisterPage />} />
          <Route path="/login" element={<AuthPage />}/>
        </Routes>
      </BrowserRouter>
  );
}

export default App;
