import React, {useEffect, useState} from 'react';
import axios from "axios";

import {PostModel} from "types/entities";

import {MainLayout} from "components/Layout";
import {PostList} from "components/PostList";
import {CreatePost} from "components/CreatePost";

import {useToken} from "helpers/hooks";

export const MainPage = () => {
  const [token] = useToken();

  const [posts, setPosts] = useState<PostModel[]>([]);

  useEffect(() => {
    axios.get('http://localhost:8082/api/post', {
      headers: {
        'Authorization': `Bearer ${token}`,
      }
    })
      .then(res => {
        setPosts(res.data.posts);
      })
  }, []);

  return (
    <MainLayout>
      <CreatePost />
      <PostList posts={posts} />
    </MainLayout>
  );
};
