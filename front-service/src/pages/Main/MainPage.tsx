import React from 'react';

import {MainLayout} from "../../components/Layout";
import {PostList} from "../../components/PostList";

export const MainPage = () => {
  return (
    <MainLayout>
      <PostList />
    </MainLayout>
  );
};
