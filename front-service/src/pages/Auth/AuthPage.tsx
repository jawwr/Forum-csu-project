import React from 'react';

import {MainLayout} from "components/Layout";
import {AuthForm} from "../../components/AuthForm";

export const AuthPage = () => {
  return (
    <MainLayout>
      <AuthForm />
    </MainLayout>
  );
};
