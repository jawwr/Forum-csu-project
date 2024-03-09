import React, {useState} from 'react';
import axios from "axios";

import {useToken} from "helpers/hooks";

import {CreatePostModel} from "types/entities";

import {Input} from "../Input";

export const CreatePost = () => {
  const [token] = useToken();

  const [title, setTitle] = useState('');
  const [text, setText] = useState('');

  const handleButtonClick = () => {
    const post: CreatePostModel = {
      text,
      title,
    };

    axios.post('http://localhost:8082/api/post', post, {
      headers: {
        'Authorization': `Bearer ${token}`,
      }
    })
  }

  return (
    <form>
      <Input
        type="text"
        title="Название"
        onClick={(e) => setTitle(e.currentTarget.value)}
      />
      <Input
        type="text"
        title="Содержание"
        onClick={(e) => setText(e.currentTarget.value)}
      />
      <button
        onClick={handleButtonClick}
        disabled={title.length < 2 || text.length < 2}
      >
        Создать новый пост
      </button>
    </form>
  );
};

export default CreatePost;
